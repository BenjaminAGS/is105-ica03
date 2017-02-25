package main

import "fmt"

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

func main() {
	IterateOverASCIIStringLiteral()
	GreetingASCII()
}

// Funksjon tar en streng med kun ASCII tegn og lager en utskrift på
// følgende format:

// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]

// Eksempel (bruk denne funksjonen i en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
// ...

func IterateOverASCIIStringLiteral() {

	fmt.Println()
	for i := 0; i < len(ascii); i++ {
		fmt.Printf("%X ", ascii[i])
		fmt.Printf("%c ", ascii[i])
		fmt.Printf("%b\n", ascii[i])

	}
}

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen greetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
// Gjelder oppgave 1b
func GreetingASCII() {
	str := "\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29"
	fmt.Println()
	fmt.Println(str)
}
