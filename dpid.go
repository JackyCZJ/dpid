package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	for l := len(os.Args)-1;l > 0 ; l --{
		mac := os.Args[l]
		macArray :=  strings.Split(mac,":")
		if len(macArray) < 6{
			log.Println("InVail mac address")
			continue
		}
		for i := range macArray {
			val := macArray[i]
			n, err := strconv.ParseUint(val, 16, 32)
			if err != nil{
				log.Println(err.Error())
				break
			}
			n2 := uint32(n)
			fmt.Print(n2)
		}
		fmt.Print("\n")
	}
}
