package gopacketctl

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/ip4defrag"
	"github.com/google/gopacket/layers" // pulls in all layers decoders

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util"
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
	logger      = util.Logger
)

func Run(src gopacket.PacketDataSource) error {
	if !flag.Parsed() {
		log.Fatalln("Run called without flags.Parse() being called")
	}
	var dec gopacket.Decoder
	var ok bool
	if dec, ok = gopacket.DecodersByLayerName[*decoder]; !ok {
		// log.Fatalln("No decoder named", *decoder)
		logger.Println("No decoder named", *decoder)
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
				logger.Println("Error while de-fragmenting", err)
				return fmt.Errorf("Error while de-fragmenting, %v", err)
			} else if newip4 == nil {
				continue // packet fragment, we don't have whole packet yet.
			}
			if newip4.Length != l {
				fmt.Printf("Decoding re-assembled packet: %s\n", newip4.NextLayerType())
				pb, ok := packet.(gopacket.PacketBuilder)
				if !ok {
					// panic("Not a PacketBuilder")
					logger.Println("Not a PacketBuilder")
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
				logger.Printf("Failed to parse: %v", err)
			}
		} else if *print {
			fmt.Println(packet)
			if err := stats(packet.String()); err != nil {
				logger.Printf("Failed to parse: %v", err)
			}
		} else {
			logger.Printf("print=%b, maxcount=%d, decoder=%s, dump=%b, statsevery=%d, printErrors=%b, lazy=%b, defrag=%b", *print, *maxcount, *decoder, *dump, *statsevery, *printErrors, *lazy, *defrag)
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

/*
PACKET: 78 bytes, wire length 78 cap length 78 @ 2017-07-02 14:02:17.556352 -0700 PDT
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..64..] SrcMAC=ac:bc:32:7c:3c:37 DstMAC=00:19:5b:27:b6:9e EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..44..] Version=4 IHL=5 TOS=0 Length=64 Id=953 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=24393 SrcIP=192.168.0.5 DstIP=74.125.204.139 Options=[] Padding=[]}
- Layer 3 (44 bytes) = TCP	{Contents=[..44..] Payload=[] SrcPort=62907 DstPort=443(https) Seq=2708963281 Ack=0 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65535 Checksum=59464 Urgent=0 Options=[..9..] Padding=[]}

PACKET: 1392 bytes, wire length 1392 cap length 1392 @ 2017-07-02 14:02:17.895049 -0700 PDT
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..1378..] SrcMAC=ac:bc:32:7c:3c:37 DstMAC=00:19:5b:27:b6:9e EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..1358..] Version=4 IHL=5 TOS=0 Length=1378 Id=13471 Flags= FragOffset=0 TTL=64 Protocol=UDP Checksum=33264 SrcIP=192.168.0.5 DstIP=64.233.189.101 Options=[] Padding=[]}
- Layer 3 (08 bytes) = UDP	{Contents=[..8..] Payload=[..1350..] SrcPort=58208 DstPort=443(https) Length=1358 Checksum=18777}
- Layer 4 (1350 bytes) = Payload	1350 byte(s)

PACKET: 60 bytes, wire length 60 cap length 60 @ 2017-07-09 09:29:32.705818 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..46..] SrcMAC=52:54:00:12:35:03 DstMAC=08:00:27:46:54:e7 EthernetType=ARP Length=0}
- Layer 2 (28 bytes) = ARP	{Contents=[..28..] Payload=[..18..] AddrType=Ethernet Protocol=IPv4 HwAddressSize=6 ProtAddressSize=4 Operation=2 SourceHwAddress=[..6..] SourceProtAddress=[10, 0, 2, 3] DstHwAddress=[..6..] DstProtAddress=[10, 0, 2, 15]}
- Layer 3 (18 bytes) = Payload	18 byte(s)

PACKET: 89 bytes, wire length 89 cap length 89 @ 2017-07-09 09:44:48.86587 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..75..] SrcMAC=08:00:27:46:54:e7 DstMAC=52:54:00:12:35:03 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..55..] Version=4 IHL=5 TOS=0 Length=75 Id=18166 Flags=DF FragOffset=0 TTL=64 Protocol=UDP Checksum=56218 SrcIP=10.0.2.15 DstIP=10.0.2.3 Options=[] Padding=[]}
- Layer 3 (08 bytes) = UDP	{Contents=[..8..] Payload=[..47..] SrcPort=41241 DstPort=53(domain) Length=55 Checksum=6234}
- Layer 4 (47 bytes) = DNS	{Contents=[..47..] Payload=[] ID=61304 QR=false OpCode=Query AA=false TC=false RD=true RA=false Z=0 ResponseCode=No Error QDCount=1 ANCount=0 NSCount=0 ARCount=0 Questions=[{Name=[..29..] Type=AAAA Class=IN}] Answers=[] Authorities=[] Additionals=[]}

PACKET: 164 bytes, wire length 164 cap length 164 @ 2017-07-09 09:44:48.986902 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..150..] SrcMAC=52:54:00:12:35:02 DstMAC=08:00:27:46:54:e7 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..130..] Version=4 IHL=5 TOS=0 Length=150 Id=20838 Flags= FragOffset=0 TTL=64 Protocol=UDP Checksum=4320 SrcIP=10.0.2.3 DstIP=10.0.2.15 Options=[] Padding=[]}
- Layer 3 (08 bytes) = UDP	{Contents=[..8..] Payload=[..122..] SrcPort=53(domain) DstPort=41241 Length=130 Checksum=25342}
- Layer 4 (122 bytes) = DNS	{Contents=[..122..] Payload=[] ID=61304 QR=true OpCode=Query AA=false TC=false RD=true RA=true Z=0 ResponseCode=Non-Existent Domain QDCount=1 ANCount=0 NSCount=1 ARCount=0 Questions=[{Name=[..29..] Type=AAAA Class=IN}] Answers=[] Authorities=[{Name=[] Type=SOA Class=IN TTL=86377 DataLength=64 Data=[..64..] IP=<nil> NS=[] CNAME=[] PTR=[] TXTs=[] SOA={ MName=[..18..] RName=[..22..] Serial=2017070900 Refresh=1800 Retry=900 Expire=604800 Minimum=86400} SRV={ Priority=0 Weight=0 Port=0 Name=[]} MX={ Preference=0 Name=[]} TXT=[]}] Additionals=[]}

PACKET: 76 bytes, wire length 76 cap length 76 @ 2017-07-09 09:44:54.188587 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..62..] SrcMAC=08:00:27:46:54:e7 DstMAC=52:54:00:12:35:03 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..42..] Version=4 IHL=5 TOS=0 Length=62 Id=19912 Flags=DF FragOffset=0 TTL=64 Protocol=UDP Checksum=54485 SrcIP=10.0.2.15 DstIP=10.0.2.3 Options=[] Padding=[]}
- Layer 3 (08 bytes) = UDP	{Contents=[..8..] Payload=[..34..] SrcPort=43356 DstPort=53(domain) Length=42 Checksum=6221}
- Layer 4 (34 bytes) = DNS	{Contents=[..34..] Payload=[] ID=43680 QR=false OpCode=Query AA=false TC=false RD=true RA=false Z=0 ResponseCode=No Error QDCount=1 ANCount=0 NSCount=0 ARCount=0 Questions=[{Name=[..16..] Type=A Class=IN}] Answers=[] Authorities=[] Additionals=[]}

PACKET: 76 bytes, wire length 76 cap length 76 @ 2017-07-09 09:44:54.188623 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..62..] SrcMAC=08:00:27:46:54:e7 DstMAC=52:54:00:12:35:03 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..42..] Version=4 IHL=5 TOS=0 Length=62 Id=19913 Flags=DF FragOffset=0 TTL=64 Protocol=UDP Checksum=54484 SrcIP=10.0.2.15 DstIP=10.0.2.3 Options=[] Padding=[]}
- Layer 3 (08 bytes) = UDP	{Contents=[..8..] Payload=[..34..] SrcPort=43356 DstPort=53(domain) Length=42 Checksum=6221}
- Layer 4 (34 bytes) = DNS	{Contents=[..34..] Payload=[] ID=10654 QR=false OpCode=Query AA=false TC=false RD=true RA=false Z=0 ResponseCode=No Error QDCount=1 ANCount=0 NSCount=0 ARCount=0 Questions=[{Name=[..16..] Type=AAAA Class=IN}] Answers=[] Authorities=[] Additionals=[]}

PACKET: 104 bytes, wire length 104 cap length 104 @ 2017-07-09 09:44:54.307345 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..90..] SrcMAC=52:54:00:12:35:02 DstMAC=08:00:27:46:54:e7 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..70..] Version=4 IHL=5 TOS=0 Length=90 Id=20845 Flags= FragOffset=0 TTL=64 Protocol=UDP Checksum=4373 SrcIP=10.0.2.3 DstIP=10.0.2.15 Options=[] Padding=[]}
- Layer 3 (08 bytes) = UDP	{Contents=[..8..] Payload=[..62..] SrcPort=53(domain) DstPort=43356 Length=70 Checksum=18924}
- Layer 4 (62 bytes) = DNS	{Contents=[..62..] Payload=[] ID=10654 QR=true OpCode=Query AA=false TC=false RD=true RA=true Z=0 ResponseCode=No Error QDCount=1 ANCount=1 NSCount=0 ARCount=0 Questions=[{Name=[..16..] Type=AAAA Class=IN}] Answers=[{Name=[..16..] Type=AAAA Class=IN TTL=299 DataLength=16 Data=[..16..] IP=2404:6800:4004:817::2003 NS=[] CNAME=[] PTR=[] TXTs=[] SOA={ MName=[] RName=[] Serial=0 Refresh=0 Retry=0 Expire=0 Minimum=0} SRV={ Priority=0 Weight=0 Port=0 Name=[]} MX={ Preference=0 Name=[]} TXT=[]}] Authorities=[] Additionals=[]}

PACKET: 92 bytes, wire length 92 cap length 92 @ 2017-07-09 09:44:54.307428 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..78..] SrcMAC=52:54:00:12:35:02 DstMAC=08:00:27:46:54:e7 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..58..] Version=4 IHL=5 TOS=0 Length=78 Id=20846 Flags= FragOffset=0 TTL=64 Protocol=UDP Checksum=4384 SrcIP=10.0.2.3 DstIP=10.0.2.15 Options=[] Padding=[]}
- Layer 3 (08 bytes) = UDP	{Contents=[..8..] Payload=[..50..] SrcPort=53(domain) DstPort=43356 Length=58 Checksum=7272}
- Layer 4 (50 bytes) = DNS	{Contents=[..50..] Payload=[] ID=43680 QR=true OpCode=Query AA=false TC=false RD=true RA=true Z=0 ResponseCode=No Error QDCount=1 ANCount=1 NSCount=0 ARCount=0 Questions=[{Name=[..16..] Type=A Class=IN}] Answers=[{Name=[..16..] Type=A Class=IN TTL=299 DataLength=4 Data=[216, 58, 200, 195] IP=216.58.200.195 NS=[] CNAME=[] PTR=[] TXTs=[] SOA={ MName=[] RName=[] Serial=0 Refresh=0 Retry=0 Expire=0 Minimum=0} SRV={ Priority=0 Weight=0 Port=0 Name=[]} MX={ Preference=0 Name=[]} TXT=[]}] Authorities=[] Additionals=[]}
*/
func stats(s string) error {
	var t, l1_type, l1_srcmac, l1_dstmac, l1_ethertype string
	var l2_protocol, l2_ipv4flags, l2_ipv4flagoffset, l2_ipv4protocol, l2_ipv4srcip, l2_ipv4dstip string
	var l3_protocol, l3_ipv4srcport, l3_ipv4dstport, l3_tcpseq, l3_udplength, l3_ICMPv4_typecode string
	var l4_payload_name, l4_payload_value string

	r := bufio.NewReader(strings.NewReader(s))
	for {
		line, err := r.ReadString(byte('\n'))
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		switch {
		case len(line) == 0:
			t = ""
			l1_type = ""
			l1_srcmac = ""
			l1_dstmac = ""
			l1_ethertype = ""
			l2_protocol = ""
			l2_ipv4flags = ""
			l2_ipv4flagoffset = ""
			l2_ipv4protocol = ""
			l2_ipv4srcip = ""
			l2_ipv4dstip = ""
			l3_protocol = ""
			l3_ipv4srcport = ""
			l3_ipv4dstport = ""
			l3_tcpseq = ""
			l3_ICMPv4_typecode = ""
			l3_udplength = ""
			l4_payload_name = ""
			l4_payload_value = ""
			break
		case strings.HasPrefix(line, "PACKET:"):
			i := strings.LastIndex(line, " @ ")
			t = line[i+3 : len(line)-1]
			fmt.Println("time:" + t)
		case strings.HasPrefix(line, "- Layer 1"):
			i := strings.Index(line, " = ") + 3
			j := i + strings.Index(line[i:], string(byte(9)))
			l1_type = line[i:j]
			fmt.Println("decoder:" + l1_type)
			if "Ethernet" == l1_type { // https://en.wikipedia.org/wiki/EtherType
				j++
				i = j + strings.Index(line[j:], "SrcMAC=") + 7
				j = i + 6*2 + 5
				l1_srcmac = line[i:j]
				i = j + 1 + 7
				j = i + 6*2 + 5
				l1_dstmac = line[i:j]
				j = i + 1 + 7
				i = j + strings.Index(line[j:], "EthernetType=") + 13
				j = i + strings.Index(line[i:], " ")
				l1_ethertype = line[i:j]
				fmt.Println("src_mac:"+l1_srcmac, "dst_mac:"+l1_dstmac, "ethernet_type:"+l1_ethertype)
			} else {
				logger.Print("...:", line[j:])
			}
		case strings.HasPrefix(line, "- Layer 2"):
			i := strings.Index(line, " = ") + 3
			j := i + strings.Index(line[i:], string(byte(9)))
			l2_protocol = line[i:j]
			fmt.Println("protocol:" + l2_protocol)
			if "IPv4" == l2_protocol {
				i = j + strings.Index(line[j:], "Flags=") + 6 // https://en.wikipedia.org/wiki/IPv4#Flags
				j = i + strings.Index(line[i:], " ")
				l2_ipv4flags = line[i:j]
				i = j + strings.Index(line[j:], "FragOffset=") + 11
				j = i + strings.Index(line[i:], " ")
				l2_ipv4flagoffset = line[i:j]
				i = j + strings.Index(line[j:], "Protocol=") + 9 // https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
				j = i + strings.Index(line[i:], " ")
				l2_ipv4protocol = line[i:j]
				i = j + strings.Index(line[j:], "SrcIP=") + 6
				j = i + strings.Index(line[i:], " ")
				l2_ipv4srcip = line[i:j]
				i = j + 1 + 6
				j = i + strings.Index(line[i:], " ")
				l2_ipv4dstip = line[i:j]
				fmt.Println("flags:"+l2_ipv4flags, "flag_offset:"+l2_ipv4flagoffset, "protocol:"+l2_ipv4protocol, "src_ip:"+l2_ipv4srcip, "dst_ip:"+l2_ipv4dstip)
			} else {
				logger.Print("...:", line[j:])
			}
		case strings.HasPrefix(line, "- Layer 3"):
			if "IPv4" == l2_protocol || "TCP" == l2_ipv4protocol || "UDP" == l2_ipv4protocol {
				i := strings.Index(line, " = ") + 3
				j := i + strings.Index(line[i:], "\t")
				l3_protocol = line[i:j]
				i = j + strings.Index(line[j:], "SrcPort=") + 8
				j = i + strings.Index(line[i:], " ")
				l3_ipv4srcport = line[i:j]
				i = j + strings.Index(line[j:], "DstPort=") + 8
				j = i + strings.Index(line[i:], " ")
				l3_ipv4dstport = line[i:j]
				fmt.Println("protocol:" + l3_protocol)
				fmt.Print("src_port:"+l3_ipv4srcport, " ", "dst_port:"+l3_ipv4dstport, " ")
				if "UDP" == l3_protocol { // https://en.wikipedia.org/wiki/User_Datagram_Protocol#Packet_structure
					i = j + strings.Index(line[j:], "Length=") + 7
					j = i + strings.Index(line[i:], " ")
					l3_udplength = line[i:j]
					fmt.Println("length:" + l3_udplength)
				} else if "TCP" == l3_protocol { //https://en.wikipedia.org/wiki/Transmission_Control_Protocol#TCP_segment_structure
					i = j + strings.Index(line[j:], "Seq=") + 4
					j = i + strings.Index(line[i:], " ")
					l3_tcpseq = line[i:j]
					fmt.Println("sequence:" + l3_tcpseq)
				} else if "ICMPv4" == l3_protocol {
					i = j + strings.Index(line[j:], "TypeCode=") + 9
					j = i + strings.Index(line[i:], " ")
					l3_ICMPv4_typecode = line[i:j]
					fmt.Println("type_code:" + l3_ICMPv4_typecode)
				} else {
					logger.Print("...:", line[j:])
				}
			} else if "ARP" == l2_protocol {
				logger.Print("...:", line)
			} else {
				logger.Print("...:", line)
			}
		case strings.HasPrefix(line, "- Layer 4"):
			i := strings.Index(line, " = ") + 3
			j := i + strings.Index(line[i:], "\t")
			l4_payload_name = line[i:j]
			if "DNS" == l4_payload_name {
				l4_payload_value = line[j+1 : len(line)-1]
				fmt.Println("application:" + l4_payload_name)
				fmt.Println("payload:", l4_payload_value)
			} else if "Payload" == l4_payload_name {
				l4_payload_value = line[j+1 : len(line)-1]
				fmt.Println("payload:", l4_payload_value)
			} else {
				logger.Print("...:", line)
			}
		default:
			logger.Print("...:", line)
		}
	}
	fmt.Println()
	return nil
}

/*
-- FULL PACKET DATA (78 bytes) ------------------------------------
00000000  00 19 5b 27 b6 9e ac bc  32 7c 3c 37 08 00 45 00  |..['....2|<7..E.|
00000010  00 40 93 62 40 00 40 06  cf c5 c0 a8 00 05 4a 7d  |.@.b@.@.......J}|
00000020  cc 65 f3 5e 01 bb a6 38  a5 d9 00 00 00 00 b0 02  |.e.^...8........|
00000030  ff ff 8a d8 00 00 02 04  05 b4 01 03 03 05 01 01  |................|
00000040  08 0a 1a c7 78 a1 00 00  00 00 04 02 00 00        |....x.........|
--- Layer 1 ---
Ethernet	{Contents=[..14..] Payload=[..64..] SrcMAC=ac:bc:32:7c:3c:37 DstMAC=00:19:5b:27:b6:9e EthernetType=IPv4 Length=0}
00000000  00 19 5b 27 b6 9e ac bc  32 7c 3c 37 08 00        |..['....2|<7..|
--- Layer 2 ---
IPv4	{Contents=[..20..] Payload=[..44..] Version=4 IHL=5 TOS=0 Length=64 Id=37730 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=53189 SrcIP=192.168.0.5 DstIP=74.125.204.101 Options=[] Padding=[]}
00000000  45 00 00 40 93 62 40 00  40 06 cf c5 c0 a8 00 05  |E..@.b@.@.......|
00000010  4a 7d cc 65                                       |J}.e|
--- Layer 3 ---
TCP	{Contents=[..44..] Payload=[] SrcPort=62302 DstPort=443(https) Seq=2788730329 Ack=0 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65535 Checksum=35544 Urgent=0 Options=[..9..] Padding=[]}
00000000  f3 5e 01 bb a6 38 a5 d9  00 00 00 00 b0 02 ff ff  |.^...8..........|
00000010  8a d8 00 00 02 04 05 b4  01 03 03 05 01 01 08 0a  |................|
00000020  1a c7 78 a1 00 00 00 00  04 02 00 00              |..x.........|

-- FULL PACKET DATA (227 bytes) ------------------------------------
00000000  ac bc 32 7c 3c 37 00 19  5b 27 b6 9e 08 00 45 00  |..2|<7..['....E.|
00000010  00 d5 00 00 00 00 98 11  7c 86 72 72 72 72 c0 a8  |........|.rrrr..|
00000020  00 05 00 35 fa 10 00 c1  08 96 b4 8f 81 80 00 01  |...5............|
00000030  00 08 00 00 00 00 08 63  6c 69 65 6e 74 73 36 06  |.......clients6.|
00000040  67 6f 6f 67 6c 65 03 63  6f 6d 00 00 01 00 01 c0  |google.com......|
00000050  0c 00 05 00 01 00 00 00  37 00 0c 07 63 6c 69 65  |........7...clie|
00000060  6e 74 73 01 6c c0 15 c0  31 00 05 00 01 00 00 01  |nts.l...1.......|
00000070  a4 00 10 0d 63 6c 69 65  6e 74 73 2d 63 68 69 6e  |....clients-chin|
00000080  61 c0 39 c0 49 00 01 00  01 00 00 00 62 00 04 40  |a.9.I.......b..@|
00000090  e9 bd 8a c0 49 00 01 00  01 00 00 00 62 00 04 40  |....I.......b..@|
000000a0  e9 bd 66 c0 49 00 01 00  01 00 00 00 62 00 04 40  |..f.I.......b..@|
000000b0  e9 bd 64 c0 49 00 01 00  01 00 00 00 62 00 04 40  |..d.I.......b..@|
000000c0  e9 bd 71 c0 49 00 01 00  01 00 00 00 62 00 04 40  |..q.I.......b..@|
000000d0  e9 bd 8b c0 49 00 01 00  01 00 00 00 62 00 04 40  |....I.......b..@|
000000e0  e9 bd 65                                          |..e|
--- Layer 1 ---
Ethernet	{Contents=[..14..] Payload=[..213..] SrcMAC=00:19:5b:27:b6:9e DstMAC=ac:bc:32:7c:3c:37 EthernetType=IPv4 Length=0}
00000000  ac bc 32 7c 3c 37 00 19  5b 27 b6 9e 08 00        |..2|<7..['....|
--- Layer 2 ---
IPv4	{Contents=[..20..] Payload=[..193..] Version=4 IHL=5 TOS=0 Length=213 Id=0 Flags= FragOffset=0 TTL=152 Protocol=UDP Checksum=31878 SrcIP=114.114.114.114 DstIP=192.168.0.5 Options=[] Padding=[]}
00000000  45 00 00 d5 00 00 00 00  98 11 7c 86 72 72 72 72  |E.........|.rrrr|
00000010  c0 a8 00 05                                       |....|
--- Layer 3 ---
UDP	{Contents=[..8..] Payload=[..185..] SrcPort=53(domain) DstPort=64016 Length=193 Checksum=2198}
00000000  00 35 fa 10 00 c1 08 96                           |.5......|
--- Layer 4 ---
DNS	{Contents=[..185..] Payload=[] ID=46223 QR=true OpCode=Query AA=false TC=false RD=true RA=true Z=0 ResponseCode=No Error QDCount=1 ANCount=8 NSCount=0 ARCount=0 Questions=[{Name=[..19..] Type=A Class=IN}] Answers=[..8..] Authorities=[] Additionals=[]}
00000000  b4 8f 81 80 00 01 00 08  00 00 00 00 08 63 6c 69  |.............cli|
00000010  65 6e 74 73 36 06 67 6f  6f 67 6c 65 03 63 6f 6d  |ents6.google.com|
00000020  00 00 01 00 01 c0 0c 00  05 00 01 00 00 00 37 00  |..............7.|
00000030  0c 07 63 6c 69 65 6e 74  73 01 6c c0 15 c0 31 00  |..clients.l...1.|
00000040  05 00 01 00 00 01 a4 00  10 0d 63 6c 69 65 6e 74  |..........client|
00000050  73 2d 63 68 69 6e 61 c0  39 c0 49 00 01 00 01 00  |s-china.9.I.....|
00000060  00 00 62 00 04 40 e9 bd  8a c0 49 00 01 00 01 00  |..b..@....I.....|
00000070  00 00 62 00 04 40 e9 bd  66 c0 49 00 01 00 01 00  |..b..@..f.I.....|
00000080  00 00 62 00 04 40 e9 bd  64 c0 49 00 01 00 01 00  |..b..@..d.I.....|
00000090  00 00 62 00 04 40 e9 bd  71 c0 49 00 01 00 01 00  |..b..@..q.I.....|
000000a0  00 00 62 00 04 40 e9 bd  8b c0 49 00 01 00 01 00  |..b..@....I.....|
000000b0  00 00 62 00 04 40 e9 bd  65                       |..b..@..e|


-- FULL PACKET DATA (42 bytes) ------------------------------------
00000000  ff ff ff ff ff ff 0a 00  27 00 00 00 08 06 00 01  |........'.......|
00000010  08 00 06 04 00 01 0a 00  27 00 00 00 ac 11 04 01  |........'.......|
00000020  00 00 00 00 00 00 ac 11  04 c8                    |..........|
--- Layer 1 ---
Ethernet	{Contents=[..14..] Payload=[..28..] SrcMAC=0a:00:27:00:00:00 DstMAC=ff:ff:ff:ff:ff:ff EthernetType=ARP Length=0}
00000000  ff ff ff ff ff ff 0a 00  27 00 00 00 08 06        |........'.....|
--- Layer 2 ---
ARP	{Contents=[..28..] Payload=[] AddrType=Ethernet Protocol=IPv4 HwAddressSize=6 ProtAddressSize=4 Operation=1 SourceHwAddress=[..6..] SourceProtAddress=[172, 17, 4, 1] DstHwAddress=[..6..] DstProtAddress=[172, 17, 4, 200]}
00000000  00 01 08 00 06 04 00 01  0a 00 27 00 00 00 ac 11  |..........'.....|
00000010  04 01 00 00 00 00 00 00  ac 11 04 c8              |............|

-- FULL PACKET DATA (60 bytes) ------------------------------------
00000000  0a 00 27 00 00 00 08 00  27 9a ac 74 08 06 00 01  |..'.....'..t....|
00000010  08 00 06 04 00 02 08 00  27 9a ac 74 ac 11 04 c8  |........'..t....|
00000020  0a 00 27 00 00 00 ac 11  04 01 00 00 00 00 00 00  |..'.............|
00000030  00 00 00 00 00 00 00 00  00 00 00 00              |............|
--- Layer 1 ---
Ethernet	{Contents=[..14..] Payload=[..46..] SrcMAC=08:00:27:9a:ac:74 DstMAC=0a:00:27:00:00:00 EthernetType=ARP Length=0}
00000000  0a 00 27 00 00 00 08 00  27 9a ac 74 08 06        |..'.....'..t..|
--- Layer 2 ---
ARP	{Contents=[..28..] Payload=[..18..] AddrType=Ethernet Protocol=IPv4 HwAddressSize=6 ProtAddressSize=4 Operation=2 SourceHwAddress=[..6..] SourceProtAddress=[172, 17, 4, 200] DstHwAddress=[..6..] DstProtAddress=[172, 17, 4, 1]}
00000000  00 01 08 00 06 04 00 02  08 00 27 9a ac 74 ac 11  |..........'..t..|
00000010  04 c8 0a 00 27 00 00 00  ac 11 04 01              |....'.......|
--- Layer 3 ---
Payload	18 byte(s)
00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000010  00 00                                             |..|


-- FULL PACKET DATA (60 bytes) ------------------------------------
00000000  0a 00 27 00 00 00 08 00  27 9a ac 74 08 06 00 01  |..'.....'..t....|
00000010  08 00 06 04 00 01 08 00  27 9a ac 74 ac 11 04 c8  |........'..t....|
00000020  00 00 00 00 00 00 ac 11  04 01 00 00 00 00 00 00  |................|
00000030  00 00 00 00 00 00 00 00  00 00 00 00              |............|
--- Layer 1 ---
Ethernet	{Contents=[..14..] Payload=[..46..] SrcMAC=08:00:27:9a:ac:74 DstMAC=0a:00:27:00:00:00 EthernetType=ARP Length=0}
00000000  0a 00 27 00 00 00 08 00  27 9a ac 74 08 06        |..'.....'..t..|
--- Layer 2 ---
ARP	{Contents=[..28..] Payload=[..18..] AddrType=Ethernet Protocol=IPv4 HwAddressSize=6 ProtAddressSize=4 Operation=1 SourceHwAddress=[..6..] SourceProtAddress=[172, 17, 4, 200] DstHwAddress=[..6..] DstProtAddress=[172, 17, 4, 1]}
00000000  00 01 08 00 06 04 00 01  08 00 27 9a ac 74 ac 11  |..........'..t..|
00000010  04 c8 00 00 00 00 00 00  ac 11 04 01              |............|
--- Layer 3 ---
Payload	18 byte(s)
00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000010  00 00                                             |..|

-- FULL PACKET DATA (42 bytes) ------------------------------------
00000000  08 00 27 9a ac 74 0a 00  27 00 00 00 08 06 00 01  |..'..t..'.......|
00000010  08 00 06 04 00 02 0a 00  27 00 00 00 ac 11 04 01  |........'.......|
00000020  08 00 27 9a ac 74 ac 11  04 c8                    |..'..t....|
--- Layer 1 ---
Ethernet	{Contents=[..14..] Payload=[..28..] SrcMAC=0a:00:27:00:00:00 DstMAC=08:00:27:9a:ac:74 EthernetType=ARP Length=0}
00000000  08 00 27 9a ac 74 0a 00  27 00 00 00 08 06        |..'..t..'.....|
--- Layer 2 ---
ARP	{Contents=[..28..] Payload=[] AddrType=Ethernet Protocol=IPv4 HwAddressSize=6 ProtAddressSize=4 Operation=2 SourceHwAddress=[..6..] SourceProtAddress=[172, 17, 4, 1] DstHwAddress=[..6..] DstProtAddress=[172, 17, 4, 200]}
00000000  00 01 08 00 06 04 00 02  0a 00 27 00 00 00 ac 11  |..........'.....|
00000010  04 01 08 00 27 9a ac 74  ac 11 04 c8              |....'..t....|


-- FULL PACKET DATA (98 bytes) ------------------------------------
00000000  08 00 27 5a 1a a4 0a 00  27 00 00 00 08 00 45 00  |..'Z....'.....E.|
00000010  00 54 87 bb 00 00 40 01  92 98 ac 11 04 01 ac 11  |.T....@.........|
00000020  04 32 08 00 65 44 d0 68  00 0a 59 61 e0 6f 00 0d  |.2..eD.h..Ya.o..|
00000030  9d 67 08 09 0a 0b 0c 0d  0e 0f 10 11 12 13 14 15  |.g..............|
00000040  16 17 18 19 1a 1b 1c 1d  1e 1f 20 21 22 23 24 25  |.......... !"#$%|
00000050  26 27 28 29 2a 2b 2c 2d  2e 2f 30 31 32 33 34 35  |&'()*+,-./012345|
00000060  36 37                                             |67|
--- Layer 1 ---
Ethernet	{Contents=[..14..] Payload=[..84..] SrcMAC=0a:00:27:00:00:00 DstMAC=08:00:27:5a:1a:a4 EthernetType=IPv4 Length=0}
00000000  08 00 27 5a 1a a4 0a 00  27 00 00 00 08 00        |..'Z....'.....|
--- Layer 2 ---
IPv4	{Contents=[..20..] Payload=[..64..] Version=4 IHL=5 TOS=0 Length=84 Id=34747 Flags= FragOffset=0 TTL=64 Protocol=ICMPv4 Checksum=37528 SrcIP=172.17.4.1 DstIP=172.17.4.50 Options=[] Padding=[]}
00000000  45 00 00 54 87 bb 00 00  40 01 92 98 ac 11 04 01  |E..T....@.......|
00000010  ac 11 04 32                                       |...2|
--- Layer 3 ---
ICMPv4	{Contents=[..8..] Payload=[..56..] TypeCode=EchoRequest Checksum=25924 Id=53352 Seq=10}
00000000  08 00 65 44 d0 68 00 0a                           |..eD.h..|
--- Layer 4 ---
Payload	56 byte(s)
00000000  59 61 e0 6f 00 0d 9d 67  08 09 0a 0b 0c 0d 0e 0f  |Ya.o...g........|
00000010  10 11 12 13 14 15 16 17  18 19 1a 1b 1c 1d 1e 1f  |................|
00000020  20 21 22 23 24 25 26 27  28 29 2a 2b 2c 2d 2e 2f  | !"#$%&'()*+,-./|
00000030  30 31 32 33 34 35 36 37                           |01234567|

*/
func verbose(v string) error {
	var t, l0_PDU, l1_FrameType, l1_SrcMAC, l1_DstMAC, l1_EtherType string
	var l2_PacketType, l2_IPv4Flags, l2_IPv4FlagOffset, l2_IPv4Protocol, l2_IPv4SrcIP, l2_IPv4DstIP string
	var l3_SegmentDatagram, l3_IPv4SrcPort, l3_IPv4DstPort, l3_IPv4TCP_Seq, l3_IPv4UDP_Length, l3_ICMPv4_TypeCode string
	var l4_DataType, l4_PayloadStats, l4_DNS_ResponseCode string

	r := bufio.NewReader(strings.NewReader(v))
	for {
		line, err := r.ReadString(byte('\n'))
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		switch {
		case len(line) == 0:
			t = time.Now().String()
			l0_PDU = ""
			logger.Println(t)
			break
		case strings.HasPrefix(line, "-- FULL PACKET DATA"):
			i := strings.Index(line[0:], "(") + 1
			j := i + strings.Index(line[i:], " byte")
			if c, err := strconv.Atoi(line[i:j]); err != nil {
				return fmt.Errorf("Failed to read full packet data: %s", err.Error())
			} else {
				b := new(bytes.Buffer)
				for row := 0; row < c/16+1; row++ {
					if s, err := r.ReadString(byte('\n')); err != nil {
						return fmt.Errorf("Failed to read full packet data: %s", err.Error())
					} else {
						if _, err := b.WriteString(s); err != nil {
							return fmt.Errorf("Failed to read full packet data: %s", err.Error())
						} else {
							l0_PDU = b.String()
						}
					}
				}
			}
			logger.Printf("\n%s", l0_PDU)
		case strings.HasPrefix(line, "--- Layer 1 ---"): // Ethernet	{Contents=[..14..] Payload=[..115..] SrcMAC=ac:bc:32:7c:3c:37 DstMAC=00:19:5b:27:b6:9e EthernetType=IPv4 Length=0}
			logger.Print(line)
			var err error
			c := 1
			for row := 0; row < c/16+2; row++ {
				line, err = r.ReadString('\n')
				if err != nil {
					return fmt.Errorf("Failed to read data link layer: %s", err.Error())
				}
				if 0 != row {
					logger.Print(line)
					continue
				}
				j := strings.Index(line, string('\t'))
				l1_FrameType = line[:j]
				i := j + strings.Index(line[j:], "{Contents=[") + 11
				j = i + strings.Index(line[i:], "]")
				if c, err = strconv.Atoi(strings.Replace(line[i:j], ".", "", -1)); err != nil {
					return fmt.Errorf("Failed to read data link layer: %s", err.Error())
				}
				switch {
				case "Ethernet" == l1_FrameType:
					i = j + strings.Index(line[j:], "SrcMAC=") + 7
					j = i + 6*2 + 5
					l1_SrcMAC = line[i:j]
					i = j + 1 + 7
					j = i + 6*2 + 5
					l1_DstMAC = line[i:j]
					j = i + 1 + 7
					i = j + strings.Index(line[j:], "EthernetType=") + 13
					j = i + strings.Index(line[i:], " ")
					l1_EtherType = line[i:j]
					logger.Printf("{ethernet:{src_mac:%s,dst_mac:%s,ether_type:%s}}\n", l1_SrcMAC, l1_DstMAC, l1_EtherType)
				default:
					logger.Print(line)
				}
			}
		case strings.HasPrefix(line, "--- Layer 2 ---"): // IPv4	{Contents=[..20..] Payload=[..64..] Version=4 IHL=5 TOS=0 Length=84 Id=34747 Flags= FragOffset=0 TTL=64 Protocol=ICMPv4 Checksum=37528 SrcIP=172.17.4.1 DstIP=172.17.4.50 Options=[] Padding=[]}
			logger.Print(line)
			var err error
			c := 1
			for row := 0; row < c/16+2; row++ {
				line, err = r.ReadString('\n')
				if err != nil {
					return fmt.Errorf("Failed to read network layer: %s", err.Error())
				}
				if 0 != row {
					logger.Print(line)
					continue
				}
				j := strings.Index(line, string('\t'))
				l2_PacketType = line[:j]
				i := j + strings.Index(line[j:], "{Contents=[") + 11
				j = i + strings.Index(line[i:], "]")
				if c, err = strconv.Atoi(strings.Replace(line[i:j], ".", "", -1)); err != nil {
					return fmt.Errorf("Failed to read network layer: %s", err.Error())
				}
				switch {
				case "IPv4" == l2_PacketType:
					i = j + strings.Index(line[j:], "Flags=") + 6 // https://en.wikipedia.org/wiki/IPv4#Flags
					j = i + strings.Index(line[i:], " ")
					l2_IPv4Flags = line[i:j]
					i = j + strings.Index(line[j:], "FragOffset=") + 11
					j = i + strings.Index(line[i:], " ")
					l2_IPv4FlagOffset = line[i:j]
					i = j + strings.Index(line[j:], "Protocol=") + 9 // https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
					j = i + strings.Index(line[i:], " ")
					l2_IPv4Protocol = line[i:j]
					i = j + strings.Index(line[j:], "SrcIP=") + 6
					j = i + strings.Index(line[i:], " ")
					l2_IPv4SrcIP = line[i:j]
					i = j + 1 + 6
					j = i + strings.Index(line[i:], " ")
					l2_IPv4DstIP = line[i:j]
					logger.Printf("{ipv4:{flags:%s,flag_offset:%s,protocol:%s,src_ip:%s,dst_ip:%s}}\n", l2_IPv4Flags, l2_IPv4FlagOffset, l2_IPv4Protocol, l2_IPv4SrcIP, l2_IPv4DstIP)
				default:
					logger.Print(line)
				}
			}
		case strings.HasPrefix(line, "--- Layer 3 ---"): // TCP	{Contents=[..44..] Payload=[] SrcPort=62302 DstPort=443(https) Seq=2788730329 Ack=0 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65535 Checksum=35544 Urgent=0 Options=[..9..] Padding=[]}
			logger.Print(line)
			var err error
			c := 1
			for row := 0; row < c/16+1; row++ {
				line, err = r.ReadString('\n')
				if err != nil {
					return fmt.Errorf("Failed to read transport layer: %s", err.Error())
				}
				if 0 != row {
					logger.Print(line)
					continue
				}
				j := strings.Index(line, string('\t'))
				l3_SegmentDatagram = line[:j]
				if "Payload" != l3_SegmentDatagram {
					i := j + strings.Index(line[j:], "{Contents=[") + 11
					j = i + strings.Index(line[i:], "]")
					if c, err = strconv.Atoi(strings.Replace(line[i:j], ".", "", -1)); err != nil {
						return fmt.Errorf("Failed to read transport layer: %s", err.Error())
					}
					switch {
					case "TCP" == l3_SegmentDatagram || "UDP" == l3_SegmentDatagram:
						i = j + strings.Index(line[j:], "SrcPort=") + 8
						j = i + strings.Index(line[i:], " ")
						l3_IPv4SrcPort = line[i:j]
						i = j + strings.Index(line[j:], "DstPort=") + 8
						j = i + strings.Index(line[i:], " ")
						l3_IPv4DstPort = line[i:j]
						if "UDP" == l3_SegmentDatagram { // https://en.wikipedia.org/wiki/User_Datagram_Protocol#Packet_structure
							i = j + strings.Index(line[j:], "Length=") + 7
							j = i + strings.Index(line[i:], " ")
							l3_IPv4UDP_Length = line[i:j]
							logger.Printf("{udp:{src_port:%s,dst_port:%s,length:%s}}\n", l3_IPv4SrcPort, l3_IPv4DstPort, l3_IPv4UDP_Length)
						} else if "TCP" == l3_SegmentDatagram { //https://en.wikipedia.org/wiki/Transmission_Control_Protocol#TCP_segment_structure
							i = j + strings.Index(line[j:], "Seq=") + 4
							j = i + strings.Index(line[i:], " ")
							l3_IPv4TCP_Seq = line[i:j]
							logger.Printf("{tcp:{src_port:%s,dst_port:%s,sequence:%s}}\n", l3_IPv4SrcPort, l3_IPv4DstPort, l3_IPv4TCP_Seq)
						} else if "ICMPv4" == l3_SegmentDatagram {
							i = j + strings.Index(line[j:], "TypeCode=") + 9
							j = i + strings.Index(line[i:], " ")
							l3_ICMPv4_TypeCode = line[i:j]
							logger.Printf("{icmpv4:{src_port:%s,dst_port:%s,sequence:%s}}\n", l3_IPv4SrcPort, l3_IPv4DstPort, l3_ICMPv4_TypeCode)
						} else {
							logger.Print(line)
						}
					case "ARP" == l3_SegmentDatagram:
						logger.Print(line)
					default:
						logger.Print(line)
					}
				} else {
					logger.Print(line)
				}
			}
		case strings.HasPrefix(line, "--- Layer 4 ---"): // DNS	{Contents=[..185..] Payload=[] ID=46223 QR=true OpCode=Query AA=false TC=false RD=true RA=true Z=0 ResponseCode=No Error QDCount=1 ANCount=8 NSCount=0 ARCount=0 Questions=[{Name=[..19..] Type=A Class=IN}] Answers=[..8..] Authorities=[] Additionals=[]}
			logger.Print(line)
			var err error
			c := 1
			for row := 0; row < c/16+1; row++ {
				line, err = r.ReadString('\n')
				if err != nil {
					return fmt.Errorf("Failed to read application layer: %s", err.Error())
				}
				if 0 != row {
					logger.Print(line)
					continue
				}
				j := strings.Index(line, string(' '))
				l4_DataType = line[:j]
				switch {
				case "Payload" == l4_DataType:
					i := j + 1
					j = len(line) - 1
					l4_PayloadStats = line[i:j]
					j = strings.Index(l4_PayloadStats, " ")
					if c, err = strconv.Atoi(l4_PayloadStats[:j]); err != nil {
						return fmt.Errorf("Failed to read application layer: %s", err.Error())
					}
					logger.Printf("{payload:{stats:%s}}\f", l4_PayloadStats)
				case "DNS" == l4_DataType:
					i := j + strings.Index(line[j:], "{Contents=[") + 11
					j = i + strings.Index(line[i:], "]")
					if c, err = strconv.Atoi(strings.Replace(line[i:j], ".", "", -1)); err != nil {
						return fmt.Errorf("Failed to read application layer: %s", err.Error())
					}
					i = j + strings.Index(line[j:], "ResponseCode=") + 13
					j = i + strings.Index(line[i:], "QDCount=") - 1
					l4_DNS_ResponseCode = line[i:j]
					fmt.Printf("{payload:{application:%s,response_code:%s}}\n"+l4_DataType, l4_DNS_ResponseCode)
				case "NTP" == l4_DataType:
					logger.Print(line)
				default:
					logger.Print(line)
				}
			}
		default:
			logger.Print(line)
		}
	}
	fmt.Println()
	return nil
}
