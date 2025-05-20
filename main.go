package main

import "fmt"

type admin struct {
	Nama     string
	password string
}

type peminjam struct {
	nama   string
	jumlah int
	tenor  int
}

func daftarAdmin(data *[10]admin, i int) {
	fmt.Println("====== Buat Akun ======")
	fmt.Print("Nama: ")
	fmt.Scan(&data[i].Nama)
	fmt.Print("Password: ")
	fmt.Scan(&data[i].password)
	fmt.Println("=======================")
}

func loginAdmin(data [10]admin, Nama string, Password string) bool {
	for _, isi := range data {
		if Nama == isi.Nama && Password == isi.password {
			return true
		}
	}
	return false
}

func tambahData(data *[1000]peminjam, nama string, jumlah, tenor, i int) {
	data[i].nama = nama
	data[i].jumlah = jumlah
	data[i].tenor = tenor
}

func ubahData(data *[1000]peminjam, nama string, pilihan int) {
	var indeks, pinjamanPengganti, tenorPengganti int
	var namaPengganti string
	for i, isi := range data {
		if nama == isi.nama {
			indeks = i
		}
	}
	switch pilihan {
	case 1:
		fmt.Print("Nama pengganti: ")
		fmt.Scan(&namaPengganti)
		data[indeks].nama = namaPengganti
	case 2:
		fmt.Print("Jumlah pinjmanan pengganti: ")
		fmt.Scan(&pinjamanPengganti)
		data[indeks].jumlah = pinjamanPengganti
	case 3:
		fmt.Print("Tenor pengganti: ")
		fmt.Scan(&tenorPengganti)
		data[indeks].tenor = tenorPengganti
	}
}

func main() {
	var pilihan, pilihanUbah int
	var dataAdmin [10]admin
	var login admin
	var tambah, ubah peminjam
	var dataPeminjam [1000]peminjam
	iLogin, iTambah := 0, 0
	cek, cekMenu := true, true

	fmt.Println("===============")
	fmt.Println("1. login")
	fmt.Println("2. Buat akun")
	fmt.Println("3. Keluar")
	fmt.Println("===============")
	fmt.Println()
	for cek {
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("====Login======")
			fmt.Print("Nama: ")
			fmt.Scan(&login.Nama)
			fmt.Print("Password: ")
			fmt.Scan(&login.password)
			fmt.Println("===============")
			cekLogin := loginAdmin(dataAdmin, login.Nama, login.password)
			if cekLogin == true {
				cek = false
			}
		case 2:
			daftarAdmin(&dataAdmin, iLogin)
			iLogin++
		case 3:
			cek = false
			cekMenu = false
		}
		fmt.Println()
	}
	fmt.Println("================")
	fmt.Println("MENU")
	fmt.Println("1. MENAMBAHKAN")
	fmt.Println("2. MENGUBAH")
	fmt.Println("3. MENGHAPUS")
	fmt.Println("4. MENCARI PEMINJAM")
	fmt.Println("5. MENGURUTKAN")
	fmt.Println("6. MENAMPILKAN LAPORAN KESELERUHAN")
	fmt.Println("7. KELUAR")
	fmt.Println("================")
	for cekMenu {
		fmt.Print("Pilihan Menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("Tambahkan Data")
			fmt.Println("Nama: ")
			fmt.Scan(&tambah.nama)
			fmt.Println("Jumlah Pinjaman: ")
			fmt.Scan(&tambah.jumlah)
			fmt.Println("Tenor: ")
			fmt.Scan(&tambah.tenor)
			tambahData(&dataPeminjam, tambah.nama, tambah.jumlah, tambah.tenor, iTambah)
			iTambah++
		case 2:
			fmt.Println("Ubah Data")
			fmt.Print("Nama: ")
			fmt.Scan(&ubah.nama)
			fmt.Println("Data yang ingin diubah:")
			fmt.Println("1. Nama")
			fmt.Println("2. Jumlah pinjaman")
			fmt.Println("3. Tenor")
			fmt.Scan(&pilihanUbah)
			ubahData(&dataPeminjam, ubah.nama, pilihanUbah)
		case 3:
			fmt.Println("pilihan 3")
		case 4:
			fmt.Println("pilihan 4")
		case 5:
			fmt.Println("pilihan 5")
		case 6:
			fmt.Println("pilihan 6")
			fmt.Println(dataPeminjam)
		case 7:
			cekMenu = false
		}
	}
	fmt.Println("Terimakasih telah menggunakan layanan kami")
	fmt.Println(dataAdmin)
	fmt.Println(dataPeminjam)
}
