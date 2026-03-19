package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {

	var a, b, c, d int

	fmt.Print("masukkan nilai a : ")
	fmt.Scan(&a)

	fmt.Print("masukkan nilai b : ")
	fmt.Scan(&b)

	fmt.Print("masukkan nilai c : ")
	fmt.Scan(&c)

	fmt.Print("masukkan nilai d : ")
	fmt.Scan(&d)

	fmt.Println("hasil permutasi", a, "dan", c, "adalah :", permutasi(a, c))
	fmt.Println("hasil kombinasi", a, "dan", c, "adalah :", kombinasi(a, c))

	fmt.Println("hasil permutasi", b, "dan", d, "adalah :", permutasi(b, d))
	fmt.Println("hasil kombinasi", b, "dan", d, "adalah :", kombinasi(b, d))

}