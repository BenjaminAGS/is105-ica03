package main

import (
	"ascii"
	"log"
)

func main() {
	str := "\x00\x01\x02\x20"
	for i := 0; i < len(str); i++ {
		for j := 0; i < len(ascii.ASCII); j++ {
			if str[i] != ascii.ASCII[j] {
				log.Fatal()
			}
		}
	}

}
