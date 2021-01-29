package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: " + `
	dpid [your mac address] [your next mac address] [...]`)
		os.Exit(0)
	}
	for l := 1; l < len(os.Args); l++ {
		mac := os.Args[l]
		hw, err := net.ParseMAC(mac)
		if err != nil {
			fmt.Println(mac, " : ", err)
			continue
		}
		macStr := strings.Replace(hw.String(),":","",-1)
		dpid , err := strconv.ParseInt(macStr,16,64)
		if err != nil{
			fmt.Println(mac, " : ", err)
			continue
		}
		fmt.Println(dpid)
	}
}
