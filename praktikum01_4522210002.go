package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Masukkan nama Anda: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan usia Anda: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)

	isValid := true
	for _, char := range ageStr {
		if char < '0' || char > '9' {
			isValid = false
			break
		}
	}

	if !isValid {
		fmt.Println("Input usia tidak valid. Masukkan angka yang benar.")
		return
	}

	age := 0
	for _, char := range ageStr {
		age = age*10 + int(char-'0')
	}

	var ageCategory string
	if age < 18 {
		ageCategory = "anak-anak"
	} else {
		ageCategory = "dewasa"
	}

	fmt.Printf("Selamat datang, %s! Anda termasuk kategori %s.\n", name, ageCategory)
}
