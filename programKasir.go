package main

import (
	"fmt"
)

func main() {
	// Menu Makanan
	fmt.Println("|=============================|")
	fmt.Println("|Program Menu Kasir Sederhana |")
	fmt.Println("|=============================|")
	fmt.Println("|    Pilihan Menu Makanan     |")
	fmt.Println("|=============================|")
	fmt.Println("|  1. Nasi Goreng   Rp.12000  |")
	fmt.Println("|  2. Ayam Geprek   Rp.12000  |")
	fmt.Println("|  3. Mie Goreng    Rp.5000   |")
	fmt.Println("|  4. Ketoprak      Rp.7000   |")
	fmt.Println("|  5. Ikan Bakar    Rp.20000  |")
	fmt.Println("|=============================|")

	var pilihanMenu int
	fmt.Print("Pilih Menu (1/2/3/4/5): ")
	fmt.Scan(&pilihanMenu)

	var jumlahPorsi int
	fmt.Print("Berapa Porsi: ")
	fmt.Scan(&jumlahPorsi)

	var hargaMakanan int
	var namaMakanan string

	switch pilihanMenu {
	case 1:
		hargaMakanan = 12000
		fmt.Printf("Nasi Goreng %d Porsi : Rp.%d\n", jumlahPorsi, hargaMakanan*jumlahPorsi)
		namaMakanan = "Nasi Goreng"
	case 2:
		hargaMakanan = 12000
		fmt.Printf("Ayam Geprek %d Porsi : Rp.%d\n", jumlahPorsi, hargaMakanan*jumlahPorsi)
		namaMakanan = "Ayam Geprek"
	case 3:
		hargaMakanan = 5000
		fmt.Printf("Mie Goreng %d Porsi : Rp.%d\n", jumlahPorsi, hargaMakanan*jumlahPorsi)
		namaMakanan = "Mie Goreng"
	case 4:
		hargaMakanan = 7000
		fmt.Printf("Ketoprak %d Porsi : Rp.%d\n", jumlahPorsi, hargaMakanan*jumlahPorsi)
		namaMakanan = "Ketoprak"
	case 5:
		hargaMakanan = 20000
		fmt.Printf("Ikan Bakar %d Porsi : Rp.%d\n", jumlahPorsi, hargaMakanan*jumlahPorsi)
		namaMakanan = "Ikan Bakar"
	default:
		fmt.Println("Pilihan Menu Tidak Tersedia")
		return
	}

	// Menu Minuman
	fmt.Println("|=============================|")
	fmt.Println("|    Pilihan Menu Minuman     |")
	fmt.Println("|=============================|")
	fmt.Println("|  1. Es Teh         Rp.3000  |")
	fmt.Println("|  2. Es Jeruk       Rp.4000  |")
	fmt.Println("|  3. Soda Gembira   Rp.5000  |")
	fmt.Println("|  4. Kopi Susu      Rp.5000  |")
	fmt.Println("|=============================|")

	var pilihanMinum int
	fmt.Print("Pilih Menu (1/2/3/4): ")
	fmt.Scan(&pilihanMinum)

	var jumlahGelas int
	fmt.Print("Berapa Gelas: ")
	fmt.Scan(&jumlahGelas)

	var hargaMinuman int
	var namaMinuman string

	switch pilihanMinum {
	case 1:
		hargaMinuman = 3000
		fmt.Printf("Es Teh %d Gelas : Rp.%d\n", jumlahGelas, hargaMinuman*jumlahGelas)
		namaMinuman = "Es Teh"
	case 2:
		hargaMinuman = 4000
		fmt.Printf("Es Jeruk %d Gelas : Rp.%d\n", jumlahGelas, hargaMinuman*jumlahGelas)
		namaMinuman = "Es Jeruk"
	case 3:
		hargaMinuman = 5000
		fmt.Printf("Soda Gembira %d Gelas : Rp.%d\n", jumlahGelas, hargaMinuman*jumlahGelas)
		namaMinuman = "Soda Gembira"
	case 4:
		hargaMinuman = 5000
		fmt.Printf("Kopi Susu %d Gelas : Rp.%d\n", jumlahGelas, hargaMinuman*jumlahGelas)
		namaMinuman = "Kopi Susu"
	default:
		fmt.Println("Pilihan Menu Tidak Tersedia")
		return
	}

	// Daftar Pembelian
	fmt.Println("|=============================|")
	fmt.Println("|       Daftar Pembelian      |")
	fmt.Println("|=============================|")
	fmt.Printf("|  Makanan  : %s\n", namaMakanan)
	fmt.Printf("|  Porsi    : %d\n", jumlahPorsi)
	fmt.Printf("|  Minuman  : %s\n", namaMinuman)
	fmt.Printf("|  Gelas    : %d\n", jumlahGelas)
	fmt.Println("|=============================|")

	// Total Pembayaran
	totalHarga := (hargaMakanan * jumlahPorsi) + (hargaMinuman * jumlahGelas)
	fmt.Println("|  Total Yang Harus Di Bayar  |")
	fmt.Println("|=============================|")
	fmt.Printf("|  Total    : Rp.%d\n", totalHarga)
	fmt.Println("|=============================|")

	// Pembayaran dan Kembalian
	var uangDibayar int
	fmt.Print("Uang Pembayaran: Rp.")
	fmt.Scan(&uangDibayar)

	if uangDibayar < totalHarga {
		fmt.Println("Uang Tidak Mencukupi!")
	} else {
		kembalian := uangDibayar - totalHarga
		fmt.Printf("Uang Kembalian: Rp.%d\n", kembalian)
	}
}
