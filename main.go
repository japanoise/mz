package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type word uint16
type dword uint32

func main() {
	var file *os.File
	var filename string
	if len(os.Args) > 1 {
		var err error
		file, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		filename = os.Args[1]
		defer file.Close()
	} else {
		file = os.Stdin
		filename = "<stdin>"
	}

	var mz [2]byte
	binary.Read(file, binary.LittleEndian, &mz)
	if !(mz[0] == 'M' && mz[1] == 'Z') {
		fmt.Fprintf(os.Stderr, "Not an MZ executable: %s\n", filename)
		os.Exit(1)
	}
	fmt.Printf("%s: MZ header OK!\n", filename)
	var w word
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Bytes in last page:                 0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Number of pages (inc last):         0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Number of relocation entries:       0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Header size (paragraphs):           0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Min. Memory allocated (paragraphs): 0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Max. Memory allocated (paragraphs): 0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Initial Stack Segment:              0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Initial Stack Pointer:              0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Checksum (0 for none):              0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Initial Instruction Pointer:        0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Initial Code Segment:               0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Offset of relocation table:         0x%04x\n", w)
	binary.Read(file, binary.LittleEndian, &w)
	fmt.Printf("Overlay number:                     0x%04x\n", w)
}
