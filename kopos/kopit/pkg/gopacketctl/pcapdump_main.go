// Inspired by
//
// https://github.com/google/gopacket/blob/master/examples/pcapdump/main.go

// The pcapdump binary implements a tcpdump-like command line tool with gopacket
// using pcap as a backend data collection mechanism.
package gopacketctl

// package main

import (
	// "bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	// "github.com/google/gopacket/dumpcommand"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/pcap"
)

func Pcapdump(ifacename string) error {
	iface = &ifacename
	*fname = ""
	*snaplen = 65536
	*tstype = ""
	*promisc = true

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[gopacketctl] ")
	log.SetOutput(os.Stderr)

	return pcapdump_main()
}

func PcapdumpOnce(ifacename string, timeout time.Duration) ([]string, error) {
	iface = &ifacename
	*fname = ""
	*snaplen = 65536
	*tstype = ""
	*promisc = true

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[gopacketctl] ")
	log.SetOutput(os.Stderr)

	defer util.Run()()
	var handle *pcap.Handle
	var err error

	// This is a little complicated because we want to allow all possible options
	// for creating the packet capture handle... instead of all this you can
	// just call pcap.OpenLive if you want a simple handle.
	inactive, err := pcap.NewInactiveHandle(*iface)
	if err != nil {
		// log.Fatalf("could not create: %v", err)
		log.Println("could not create:", err)
		return []string{}, err
	}
	defer inactive.CleanUp()
	if err = inactive.SetSnapLen(*snaplen); err != nil {
		// log.Fatalf("could not set snap length: %v", err)
		log.Println("could not set snap length:", err)
		return []string{}, err
	} else if err = inactive.SetPromisc(*promisc); err != nil {
		// log.Fatalf("could not set promisc mode: %v", err)
		log.Println("could not set promisc mode:", err)
		return []string{}, err
	} else if err = inactive.SetTimeout(time.Second); err != nil {
		// log.Fatalf("could not set timeout: %v", err)
		log.Println("could not set timeout:", err)
		return []string{}, err
	}
	if *tstype != "" {
		if t, err := pcap.TimestampSourceFromString(*tstype); err != nil {
			// log.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
			log.Println("Supported timestamp types:", inactive.SupportedTimestamps())
			return []string{}, err
		} else if err := inactive.SetTimestampSource(t); err != nil {
			// log.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
			log.Println("Supported timestamp types:", inactive.SupportedTimestamps())
		}
	}
	if handle, err = inactive.Activate(); err != nil {
		// log.Fatal("PCAP Activate error:", err)
		log.Println("PCAP Activate error:", err)
	}
	defer handle.Close()

	return Run_dumponce(handle, timeout)
}

var iface = flag.String("i", "eth0", "Interface to read packets from")
var fname = flag.String("r", "", "Filename to read from, overrides -i")
var snaplen = flag.Int("s", 65536, "Snap length (number of bytes max to read per packet")
var tstype = flag.String("timestamp_type", "", "Type of timestamps to use")
var promisc = flag.Bool("promisc", true, "Set promiscuous mode")

func pcapdump_main() error {
	defer util.Run()()
	var handle *pcap.Handle
	var err error
	if *fname != "" {
		if handle, err = pcap.OpenOffline(*fname); err != nil {
			log.Fatal("PCAP OpenOffline error:", err)
		}
	} else {
		// This is a little complicated because we want to allow all possible options
		// for creating the packet capture handle... instead of all this you can
		// just call pcap.OpenLive if you want a simple handle.
		inactive, err := pcap.NewInactiveHandle(*iface)
		if err != nil {
			// log.Fatalf("could not create: %v", err)
			log.Println("could not create:", err)
			return err
		}
		defer inactive.CleanUp()
		if err = inactive.SetSnapLen(*snaplen); err != nil {
			// log.Fatalf("could not set snap length: %v", err)
			log.Println("could not set snap length:", err)
			return err
		} else if err = inactive.SetPromisc(*promisc); err != nil {
			// log.Fatalf("could not set promisc mode: %v", err)
			log.Println("could not set promisc mode:", err)
			return err
		} else if err = inactive.SetTimeout(time.Second); err != nil {
			// log.Fatalf("could not set timeout: %v", err)
			log.Println("could not set timeout:", err)
			return err
		}
		if *tstype != "" {
			if t, err := pcap.TimestampSourceFromString(*tstype); err != nil {
				// log.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
				log.Println("Supported timestamp types:", inactive.SupportedTimestamps())
				return err
			} else if err := inactive.SetTimestampSource(t); err != nil {
				// log.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
				log.Println("Supported timestamp types:", inactive.SupportedTimestamps())
			}
		}
		if handle, err = inactive.Activate(); err != nil {
			// log.Fatal("PCAP Activate error:", err)
			log.Println("PCAP Activate error:", err)
		}
		defer handle.Close()
	}
	if len(flag.Args()) > 0 {
		bpffilter := strings.Join(flag.Args(), " ")
		fmt.Fprintf(os.Stderr, "Using BPF filter %q\n", bpffilter)
		if err = handle.SetBPFFilter(bpffilter); err != nil {
			log.Fatal("BPF filter error:", err)
		}
	}
	// dumpcommand.Run(handle)
	return Run_dumpcommand(handle)
}
