package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"flag"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main(){
	var uredjaj = flag.String("i", "eth0", "this is a way to read from a specific interface")
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
	
	nesto := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range nesto.Packets(){
	
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil{
			ipBetter, _ := ipLayer.(*layers.IPv4)
			fmt.Println("IP(Src): ", ipBetter.SrcIP,", IP(Dst):", ipBetter.DstIP )
		}//ip ispisivac
		
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil{
			tcpBetter, _:=tcpLayer.(*layers.TCP)
			fmt.Println("TCP(Src): ", tcpBetter.SrcPort, ", TCP(Dst): ", tcpBetter.DstPort)
		}
	}
}

