package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var seed int64
	// use little endian byte order
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}
