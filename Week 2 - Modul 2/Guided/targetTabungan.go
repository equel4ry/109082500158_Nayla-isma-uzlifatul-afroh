package main
import"fmt"

func main(){
	var target, tabungan, total, hari int

	fmt.Print("masukan target uang yang anda ingin dicapai: ")
	fmt.Scanln(&target)

	total= 0
	hari= 0
	for total < target {
		hari++
		fmt.Printf("masukan nominal tabungan hari ke-%d; ", hari)
		fmt.Scanln(&tabungan)
		total = total + tabungan
	}

	//output setelah loop selesai (target tercapai)
	fmt.Printf("selamat! target tercapai dalam %d hari.\n", hari)
	fmt.Printf("total tabungan anda terkumpul: Rp%d\n", total)
}