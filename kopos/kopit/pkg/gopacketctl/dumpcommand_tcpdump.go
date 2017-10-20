// Inspired by:
//   https://github.com/google/gopacket/blob/master/dumpcommand/tcpdump.go

// Package dumpcommand implements a run function for pfdump and pcapdump
// with many similar flags/features to tcpdump.  This code is split out seperate
// from data sources (pcap/pfring) so it can be used by both.
package gopacketctl

// package dumpcommand

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/ip4defrag"
	"github.com/google/gopacket/layers" // pulls in all layers decoders
)

var (
	print       = flag.Bool("print", true, "Print out packets, if false only prints out statistics")
	maxcount    = flag.Int("c", -1, "Only grab this many packets, then exit")
	decoder     = flag.String("decoder", "Ethernet", "Name of the decoder to use")
	dump        = flag.Bool("X", true, "If true, dump very verbose info on each packet")
	statsevery  = flag.Int("stats", 1000, "Output statistics every N packets")
	printErrors = flag.Bool("errors", false, "Print out packet dumps of decode errors, useful for checking decoders against live traffic")
	lazy        = flag.Bool("lazy", false, "If true, do lazy decoding")
	defrag      = flag.Bool("defrag", false, "If true, do IPv4 defrag")
)

func Run_dumpcommand(src gopacket.PacketDataSource) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[gopacketctl] ")
	log.SetOutput(os.Stderr)

	if !flag.Parsed() {
		// log.Fatalln("Run called without flags.Parse() being called")
		log.Println("Run called without flags.Parse() being called")
		return fmt.Errorf("Configuration required")
	}
	var dec gopacket.Decoder
	var ok bool
	if dec, ok = gopacket.DecodersByLayerName[*decoder]; !ok {
		// log.Fatalln("No decoder named", *decoder)
		log.Println("No decoder named", *decoder)
		return fmt.Errorf("No decoder named %v", *decoder)
	}
	source := gopacket.NewPacketSource(src, dec)
	source.Lazy = *lazy
	source.NoCopy = true
	source.DecodeStreamsAsDatagrams = true
	fmt.Fprintln(os.Stderr, "Starting to read packets")
	count := 0
	bytes := int64(0)
	start := time.Now()
	errors := 0
	truncated := 0
	layertypes := map[gopacket.LayerType]int{}
	defragger := ip4defrag.NewIPv4Defragmenter()

	for packet := range source.Packets() {
		count++
		bytes += int64(len(packet.Data()))

		// defrag the IPv4 packet if required
		if *defrag {
			ip4Layer := packet.Layer(layers.LayerTypeIPv4)
			if ip4Layer == nil {
				continue
			}
			ip4 := ip4Layer.(*layers.IPv4)
			l := ip4.Length

			newip4, err := defragger.DefragIPv4(ip4)
			if err != nil {
				// log.Fatalln("Error while de-fragmenting", err)
				log.Println("Error while de-fragmenting", err)
				return fmt.Errorf("Error while de-fragmenting, %v", err)
			} else if newip4 == nil {
				continue // packet fragment, we don't have whole packet yet.
			}
			if newip4.Length != l {
				fmt.Printf("Decoding re-assembled packet: %s\n", newip4.NextLayerType())
				pb, ok := packet.(gopacket.PacketBuilder)
				if !ok {
					// panic("Not a PacketBuilder")
					log.Println("Not a PacketBuilder")
					return fmt.Errorf("Not a PacketBuilder")
				}
				nextDecoder := newip4.NextLayerType()
				nextDecoder.Decode(newip4.Payload, pb)
			}
		}

		if *dump {
			// fmt.Println(packet.Dump())
			d := packet.Dump()
			if err := verbose(d); err != nil {
				log.Printf("Failed to parse: %v", err)
			}
		} else if *print {
			fmt.Println(packet)
			if err := stats(packet.String()); err != nil {
				log.Printf("Failed to parse: %v", err)
			}
		} else {
			log.Printf("print=%b, maxcount=%d, decoder=%s, dump=%b, statsevery=%d, printErrors=%b, lazy=%b, defrag=%b", *print, *maxcount, *decoder, *dump, *statsevery, *printErrors, *lazy, *defrag)
		}
		if !*lazy || *print || *dump { // if we've already decoded all layers...
			for _, layer := range packet.Layers() {
				layertypes[layer.LayerType()]++
			}
			if packet.Metadata().Truncated {
				truncated++
			}
			if errLayer := packet.ErrorLayer(); errLayer != nil {
				errors++
				if *printErrors {
					fmt.Println("Error:", errLayer.Error())
					fmt.Println("--- Packet ---")
					fmt.Println(packet.Dump())
				}
			}
		}
		done := *maxcount > 0 && count >= *maxcount
		if count%*statsevery == 0 || done {
			fmt.Fprintf(os.Stderr, "Processed %v packets (%v bytes) in %v, %v errors and %v truncated packets\n", count, bytes, time.Since(start), errors, truncated)
			if len(layertypes) > 0 {
				fmt.Fprintf(os.Stderr, "Layer types seen: %+v\n", layertypes)
			}
		}
		if done {
			break
		}

		time.Sleep(time.Second)
	}

	return nil
}

func Run_dumponce(src gopacket.PacketDataSource, timeout time.Duration) ([]string, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[gopacketctl] ")
	log.SetOutput(os.Stderr)

	if !flag.Parsed() {
		// log.Fatalln("Run called without flags.Parse() being called")
		log.Println("Run called without flags.Parse() being called")
		return []string{}, fmt.Errorf("Configuration required")
	}
	var dec gopacket.Decoder
	var ok bool
	if dec, ok = gopacket.DecodersByLayerName[*decoder]; !ok {
		// log.Fatalln("No decoder named", *decoder)
		log.Println("No decoder named", *decoder)
		return []string{}, fmt.Errorf("No decoder named %v", *decoder)
	}
	source := gopacket.NewPacketSource(src, dec)
	source.Lazy = *lazy
	source.NoCopy = true
	source.DecodeStreamsAsDatagrams = true
	fmt.Fprintln(os.Stderr, "Starting to read packets")
	count := 0
	bytes := int64(0)
	start := time.Now()
	errors := 0
	truncated := 0
	layertypes := map[gopacket.LayerType]int{}
	defragger := ip4defrag.NewIPv4Defragmenter()

	result := make([]string, 0)
	for packet := range source.Packets() {
		count++
		bytes += int64(len(packet.Data()))

		// defrag the IPv4 packet if required
		if *defrag {
			ip4Layer := packet.Layer(layers.LayerTypeIPv4)
			if ip4Layer == nil {
				continue
			}
			ip4 := ip4Layer.(*layers.IPv4)
			l := ip4.Length

			newip4, err := defragger.DefragIPv4(ip4)
			if err != nil {
				// log.Fatalln("Error while de-fragmenting", err)
				log.Println("Error while de-fragmenting", err)
				return []string{}, fmt.Errorf("Error while de-fragmenting, %v", err)
			} else if newip4 == nil {
				continue // packet fragment, we don't have whole packet yet.
			}
			if newip4.Length != l {
				fmt.Printf("Decoding re-assembled packet: %s\n", newip4.NextLayerType())
				pb, ok := packet.(gopacket.PacketBuilder)
				if !ok {
					// panic("Not a PacketBuilder")
					log.Println("Not a PacketBuilder")
					return []string{}, fmt.Errorf("Not a PacketBuilder")
				}
				nextDecoder := newip4.NextLayerType()
				nextDecoder.Decode(newip4.Payload, pb)
			}
		}

		if *dump {
			// fmt.Println(packet.Dump())
			d := packet.Dump()
			result = append(result, d)
		} else if *print {
			fmt.Println(packet)
			result = append(result, packet.String())
		}
		if !*lazy || *print || *dump { // if we've already decoded all layers...
			for _, layer := range packet.Layers() {
				layertypes[layer.LayerType()]++
			}
			if packet.Metadata().Truncated {
				truncated++
			}
			if errLayer := packet.ErrorLayer(); errLayer != nil {
				errors++
				if *printErrors {
					fmt.Println("Error:", errLayer.Error())
					fmt.Println("--- Packet ---")
					fmt.Println(packet.Dump())
				}
			}
		}
		done := *maxcount > 0 && count >= *maxcount
		if count%*statsevery == 0 || done {
			fmt.Fprintf(os.Stderr, "Processed %v packets (%v bytes) in %v, %v errors and %v truncated packets\n", count, bytes, time.Since(start), errors, truncated)
			if len(layertypes) > 0 {
				fmt.Fprintf(os.Stderr, "Layer types seen: %+v\n", layertypes)
			}
		}
		if done {
			break
		}

		if time.Since(start) > timeout {
			break
		}
	}

	return result, nil
}
