package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("PROGRAM KONVERSI SUHU")

	// Meminta input suhu Celsius
	fmt.Print("Masukkan suhu dalam Celsius: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Validasi input harus berupa angka
	celsius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error: Input harus berupa angka!")
		return
	}

	// Validasi suhu tidak boleh di bawah nol absolut
	if celsius < -273.15 {
		fmt.Println("Error: Suhu tidak boleh di bawah -273.15 Celsius!")
		return
	}

	// Proses konversi suhu
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15

	// Menampilkan hasil
	fmt.Println("HASIL KONVERSI")
	fmt.Printf("Celsius    : %.2f °C\n", celsius)
	fmt.Printf("Fahrenheit : %.2f °F\n", fahrenheit)
	fmt.Printf("Kelvin     : %.2f K\n", kelvin)
	fmt.Println("Konversi suhu berhasil!")
}
