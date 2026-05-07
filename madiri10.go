package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var w int
	fmt.Print("Masukkan jumlah produk: ")
	fmt.Scanln(&w)

	if w == 0 {
		return
	}

	produk := make([]struct {
		nama   string
		jumlah int
	}, w)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < w; i++ {
		fmt.Printf("Produk ke-%d\n", i+1)
		fmt.Print("Nama produk: ")
		scanner.Scan()
		produk[i].nama = strings.TrimSpace(scanner.Text())

		for {
			fmt.Print("Jumlah produk: ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			jumlah, err := strconv.Atoi(input)
			if err != nil {
				produk[i].nama += " " + input
			} else {
				produk[i].jumlah = jumlah
				break
			}
		}
	}

	maxIndex, minIndex := 0, 0
	total := 0

	for i := 0; i < w; i++ {
		total += produk[i].jumlah
		if produk[i].jumlah > produk[maxIndex].jumlah {
			maxIndex = i
		}
		if produk[i].jumlah < produk[minIndex].jumlah {
			minIndex = i
		}
	}

	fmt.Println("===== HASIL ANALISIS PENJUALAN =====")
	fmt.Printf("Produk Terlaris    : %s\n", produk[maxIndex].nama)
	fmt.Printf("Penjualan Maksimum : %d\n", produk[maxIndex].jumlah)
	fmt.Printf("Produk Terendah    : %s\n", produk[minIndex].nama)
	fmt.Printf("Penjualan Minimum  : %d\n", produk[minIndex].jumlah)
	fmt.Printf("Selisih Max-Min    : %d\n", produk[maxIndex].jumlah-produk[minIndex].jumlah)
	fmt.Printf("Rata-rata Penjualan: %.0f\n", float64(total)/float64(w))
}
