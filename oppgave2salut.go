package main

import "fmt"

func main() {
	str := "\x53\x61\x6C\x75\x74\x2C\x20\xE7\x61\x20\x76\x61" +
		"\x20\xB0\x2D\x29\x20\x80\x35\x30"

	fmt.Println(str)
	//for i := 0; i < len(str); i++ {
	//fmt.Printf("%c", str[i])
	//}
}

// Vi skal skrive "Salut, " og litt mer
