package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main(){
	for l := len(os.Args)-1;l > 0 ; l --{
		mac := os.Args[l]
		hw , err := net.ParseMAC(mac)
		if err != nil{
			panic(err)
		}
		macArray :=  strings.Split(hw.String(),":")
		for i := range macArray {
			val := macArray[i]
			n, err := strconv.ParseUint(val, 16, 32)
			if err != nil{
				log.Println(err.Error())
				break
			}
			if n != 0&& i != 0{
				n2 := uint32(n)
				fmt.Print(n2)
			}

		}
		fmt.Print("\n")
	}
}
