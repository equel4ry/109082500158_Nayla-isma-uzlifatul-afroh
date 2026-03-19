package main

import "fmt"

func hitungPersegi(sisi int) {
	fmt.Println()
	fmt.Println("Masukkan sisi :", sisi)
	fmt.Println("Luas persegi :", sisi*sisi)
	fmt.Println("Keliling persegi :", 4*sisi)
}

func hitungPersegiPanjang(panjang int, lebar int) {
	fmt.Println()
	fmt.Println("Masukkan panjang :", panjang)
	fmt.Println("Masukkan lebar :", lebar)
	fmt.Println("Luas persegi panjang :", panjang*lebar)
	fmt.Println("Keliling persegi panjang :", 2*(panjang+lebar))
}

func hitungLingkaran(r float64) {
	fmt.Println()
	fmt.Println("Masukkan jari-jari :", r)
	fmt.Println("Luas lingkaran :", 3.14*r*r)
	fmt.Println("Keliling lingkaran :", 2*3.14*r)
}

func main() {
	var pilih int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		hitungPersegi(1057)
	case 2:
		hitungPersegiPanjang(106, 88)
	case 3:
		hitungLingkaran(13.87)
	}
}