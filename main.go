package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"flag"
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
	
}

