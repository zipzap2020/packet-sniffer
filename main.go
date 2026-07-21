package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"flag"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"encoding/hex"
)
func isPrintableASCII(payload []byte) bool{

	for _, bpay := range payload{
		switch{
			case bpay == 9 || bpay == 10 || bpay == 13:
			case bpay >= 32 && bpay <= 126:
			default:
			return false
			}
		}
		return true
}


func main(){
	var uredjaj = flag.String("i", "eth0", "this is a way to read from a specific interface")
	var filter = flag.String("f", "", "this is a way to filter your output")
	flag.Parse() //ovaj deo je da bi podrazumevane stvari u flag mogle da se zamene u buduce. ne diraj!!

//	devices, err := pcap.FindAllDevs()
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	for _, device := range devices {
//		fmt.Print("ime: ", device.Name, "\n")
//		fmt.Print("opis: ", device.Description, "\n")
//	}	//ovaj deo ispisuje sve na mreznoj kartici
	handle, err := pcap.OpenLive(*uredjaj , 65535, true, pcap.BlockForever)
	if err != nil{
		fmt.Println("error with opening this specific card", err)
		return
	}	//ovaj deo cita sa specificne mrezne kartice
	defer handle.Close() //ovo samo zatvara program kako nebi bio u vecnom loop-u
	
	handle.SetBPFFilter(*filter)
	
	nesto := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range nesto.Packets(){
	
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil{
			ipBetter, _ := ipLayer.(*layers.IPv4)
			fmt.Println("IP(Src): ", ipBetter.SrcIP,", IP(Dst):", ipBetter.DstIP)
		}//ip ispisivac
		
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil{
			tcpBetter, _:=tcpLayer.(*layers.TCP)
			fmt.Println("TCP(Src): ", tcpBetter.SrcPort, ", TCP(Dst): ", tcpBetter.DstPort)
		}
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer != nil{
		udpBetter, _ := udpLayer.(*layers.UDP)
			fmt.Println("UDP(Src): ", udpBetter.SrcPort, ", UDP(Dst): ", udpBetter.DstPort)
		}
		app := packet.ApplicationLayer()
		if app != nil{
			payload := app.Payload()
			if isPrintableASCII(payload){
				fmt.Println(string(payload))
			}else{
				fmt.Println("payload: \n" + hex.Dump(payload))

			}
		}
		dns := packet.Layer(layers.LayerTypeDNS)
		if dns != nil{
			dnsBetter, _ := dns.(*layers.DNS)
			for _, dnsQuestion := range dnsBetter.Questions{
				fmt.Println("DNS: ", string(dnsQuestion.Name))
			}
		}
		
		fmt.Println(" ")
	}
}

