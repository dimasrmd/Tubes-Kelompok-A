package main

import (
	"fmt"
)

type admin struct {
	Nama, password string
}

type peminjam struct {
	nama                   string
	jumlah, tenor          int
	status                 string
	bunga, bayaranPerbulan float64
	// Terbayar, Dalam Tenor, Lewat Tenor
}

var dataAdmin [10]admin
var dataPeminjam [1000]peminjam

func daftarAdmin(data *[10]admin, i int) {
	fmt.Println()
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

func tambahData(data *[1000]peminjam, nama string, jumlah, tenor int, status string, i int) {
	data[i].nama = nama
	data[i].jumlah = jumlah
	data[i].tenor = tenor
	data[i].status = status
	hitungBunga(i)
	hitungBayaranperbulan(i)
}
func hitungBunga(indeks int) {
	if dataPeminjam[indeks].status == "Lunas" || dataPeminjam[indeks].status == "DalamTenor" {
		dataPeminjam[indeks].bunga = float64(dataPeminjam[indeks].jumlah) * 0.02
	} else if dataPeminjam[indeks].status == "LewatTenor" {
		dataPeminjam[indeks].bunga = float64(dataPeminjam[indeks].jumlah) * 0.03
	}
}

func hitungBayaranperbulan(indeks int) {
	dataPeminjam[indeks].bayaranPerbulan = (float64(dataPeminjam[indeks].jumlah) + dataPeminjam[indeks].bunga) / float64(dataPeminjam[indeks].tenor)
}

func ubahData(data *[1000]peminjam, indeks int, pilihan string) {
	var ubah peminjam
	cek := true
	for cek {
		switch pilihan {
		case "1":
			fmt.Print("Nama pengganti: ")
			fmt.Scan(&ubah.nama)
			data[indeks].nama = ubah.nama
			cek = false
		case "2":
			fmt.Print("Jumlah pinjmanan pengganti: ")
			fmt.Scan(&ubah.jumlah)
			data[indeks].jumlah = ubah.jumlah
			cek = false
		case "3":
			fmt.Print("Tenor pengganti: ")
			fmt.Scan(&ubah.tenor)
			data[indeks].tenor = ubah.tenor
			cek = false
		case "4":
			fmt.Print("Status pengganti (Lunas/DalamTenor/LewatTenor): ")
			fmt.Scan(&ubah.status)
			data[indeks].status = ubah.status
			cek = false
		case "5":
			fmt.Println("SYSTEM: Kembali")
			cek = false
		default:
			fmt.Println("SYSTEM: INPUT TIDAK VALID!")
			fmt.Println("---------------------------")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)
		}

	}
	hitungBunga(indeks)
	hitungBayaranperbulan(indeks)
}

func showData(data [1000]peminjam, iTambah int) {
	if iTambah >= 1 {
		for i := 0; i < iTambah; i++ {
			fmt.Printf("> %s dengan jumlah pinjaman Rp%d dan tenor %d bulan, cicilan perbulan Rp%.2f, status: %s\n", data[i].nama, data[i].jumlah, data[i].tenor, data[i].bayaranPerbulan, data[i].status)
		}
	} else {
		fmt.Println("SYSTEM: DATA KOSONG!")
	}
}

func hapusData(data *[1000]peminjam, iTambah, pilih int) {
	for i := pilih; i < iTambah; i++ {
		data[i] = data[i+1]
	}

}

func ascending(jumIndeks int, pilihan string) {
	// Selection Sort
	switch pilihan {
	case "1": // Urut keatas berdasar jumlah
		for i := 0; i < jumIndeks-1; i++ {
			minIdx := i
			for j := i + 1; j < jumIndeks; j++ {
				if dataPeminjam[j].jumlah < dataPeminjam[minIdx].jumlah {
					minIdx = j
				}
			}
			dataPeminjam[i], dataPeminjam[minIdx] = dataPeminjam[minIdx], dataPeminjam[i]
		}
	case "2": // Urut keatas berdasar tenor
		for i := 0; i < jumIndeks-1; i++ {
			minIdx := i
			for j := i + 1; j < jumIndeks; j++ {
				if dataPeminjam[j].tenor < dataPeminjam[minIdx].tenor {
					minIdx = j
				}
			}
			dataPeminjam[i], dataPeminjam[minIdx] = dataPeminjam[minIdx], dataPeminjam[i]
		}
	}
}

func descending(jumIndeks int, pilihan string) {
	// Insertion Sort
	switch pilihan {
	case "1": // Urut kebawah berdasar jumlah
		for i := 1; i < jumIndeks; i++ {
			mark := dataPeminjam[i]
			j := i - 1
			for j >= 0 && dataPeminjam[j].jumlah < mark.jumlah {
				dataPeminjam[j+1] = dataPeminjam[j]
				j--
			}
			dataPeminjam[j+1] = mark
		}
	case "2": // Urut kebawah berdasar tenor
		for i := 1; i < jumIndeks; i++ {
			mark := dataPeminjam[i]
			j := i - 1
			for j >= 0 && dataPeminjam[j].tenor < mark.tenor {
				dataPeminjam[j+1] = dataPeminjam[j]
				j--
			}
			dataPeminjam[j+1] = mark
		}
	}

}
func main() {
	var inputData int
	var login admin
	var tambah, ubah, hapus peminjam
	var cekMenu bool
	var pilihan1, pilihan2 string
	iLogin, iTambah := 0, 0
	cek := true
	for cek {
		fmt.Println("===== MENU =====")
		fmt.Println("1. login")
		fmt.Println("2. Buat akun")
		fmt.Println("3. Keluar")
		fmt.Println("================")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan1)
		switch pilihan1 {
		case "1":
			fmt.Println()
			fmt.Println("====Login======")
			fmt.Print("Nama: ")
			fmt.Scan(&login.Nama)
			fmt.Print("Password: ")
			fmt.Scan(&login.password)
			fmt.Println("===============")
			cekLogin := loginAdmin(dataAdmin, login.Nama, login.password)
			if cekLogin == true {
				cek = false
				cekMenu = true
			} else {
				fmt.Println("SYSTEM: Data tidak ditemukan")
			}
		case "2":
			daftarAdmin(&dataAdmin, iLogin)
			iLogin++
			cekMenu = false
		case "3":
			cek = false
			cekMenu = false
		default:
			fmt.Println("SYSTEM: Input tidak valid!")
		}
		fmt.Println()
		for cekMenu {
			fmt.Println("Selamat datang yang terhormat", login.Nama)
			fmt.Println("================ MENU ================")
			fmt.Println("1. MENAMBAHKAN")
			fmt.Println("2. MENGUBAH")
			fmt.Println("3. MENGHAPUS")
			fmt.Println("4. MENCARI PEMINJAM")
			fmt.Println("5. MENGURUTKAN")
			fmt.Println("6. MENAMPILKAN LAPORAN KESELERUHAN")
			fmt.Println("7. KEMBALI")
			fmt.Println("======================================")
			fmt.Print("Pilihan Menu: ")
			fmt.Scan(&pilihan2)
			switch pilihan2 {
			case "1":
				fmt.Println("\n=================== INPUT DATA ===================")
				fmt.Print("Berapa data yang ditambahkan (0 untuk kembali): ")
				fmt.Scan(&inputData)
				for inputData >= 1 {
					for i := 1; i <= inputData; i++ {
						fmt.Println("----------------------------------------------")
						fmt.Print("Nama: ")
						fmt.Scan(&tambah.nama)
						fmt.Print("Jumlah Pinjaman: ")
						fmt.Scan(&tambah.jumlah)
						fmt.Print("Tenor: ")
						fmt.Scan(&tambah.tenor)
						fmt.Print("Status (Lunas/DalamTenor/LewatTenor): ")
						fmt.Scan(&tambah.status)
						tambahData(&dataPeminjam, tambah.nama, tambah.jumlah, tambah.tenor, tambah.status, iTambah)
						iTambah++
					}
				}
				fmt.Println("==================================================")
			case "2":
				var indeksUbah, indeksUbah1, duplikatUbah int
				var konfirmasiUbah string
				var cekUbah, lanjutUbah, konfirmasiUbah2 bool = true, true, true
				var arrayCekUbah [1000]int
				fmt.Println("\n================ UBAH DATA ================")
				for cekUbah {
					fmt.Print("Nama tujuan: ")
					fmt.Scan(&ubah.nama)
					for _, isi := range dataPeminjam {
						if isi.nama == ubah.nama {
							duplikatUbah++ // 2
						}
					}
					if duplikatUbah > 1 {
						fmt.Println("SYSTEM: Data duplikat, pilih salah satu:")
						for i, isi := range dataPeminjam {
							if isi.nama == ubah.nama {
								fmt.Println(">", isi, "pilih", i)
								arrayCekUbah[indeksUbah1] = i
								indeksUbah1++
							}
						}
						cekArrayUbah := true
						indeksArrayUbah := -1
						for cekArrayUbah {
							fmt.Print("Pilihan ubah: ")
							fmt.Scan(&indeksArrayUbah)
							for _, isi := range arrayCekUbah {
								if isi == indeksArrayUbah {
									cekUbah = false
									cekArrayUbah = false
									indeksUbah = indeksArrayUbah
								}
							}
							if cekArrayUbah {
								fmt.Println("SYSTEM: Input tidak valid!")
								fmt.Println("---------------------------")
							}
						}
					} else if duplikatUbah == 1 {
						for i, isi := range dataPeminjam {
							if isi.nama == ubah.nama {
								indeksUbah = i
							}
						}
						cekUbah = false
					} else {
						fmt.Println("SYSTEM: DATA TIDAK ADA ATAU TIDAK DITEMUKAN")
						for konfirmasiUbah2 {
							fmt.Print("Keluar (YA/TIDAK): ")
							fmt.Scan(&konfirmasiUbah)
							if konfirmasiUbah == "YA" {
								cekUbah, lanjutUbah, konfirmasiUbah2 = false, false, false
							} else if konfirmasiUbah == "TIDAK" {
								fmt.Println("---------------------------")
								cekUbah, lanjutUbah, konfirmasiUbah2 = true, true, false
							} else {
								konfirmasiUbah2 = true
							}
						}
						konfirmasiUbah2 = true
					}
				}
				var pilihanUbah string
				if lanjutUbah {
					fmt.Println("---------------------------")
					fmt.Println("Data yang ingin diubah:")
					fmt.Println("1. Nama")
					fmt.Println("2. Jumlah pinjaman")
					fmt.Println("3. Tenor")
					fmt.Println("4. Status")
					fmt.Println("5. Kembali")
					fmt.Print("Pilihan: ")
					fmt.Scan(&pilihanUbah)
					ubahData(&dataPeminjam, indeksUbah, pilihanUbah)
				}
				fmt.Println("===========================================")
			case "3":
				var cekHapus bool = true
				var konfirmasi string
				var duplikatHapus, indeksHapus int
				fmt.Println("\n============ HAPUS DATA ============")
				for cekHapus {
					cekHapus1 := true
					fmt.Print("Nama tujuan (KETIK 0 UNTUK KELUAR): ")
					fmt.Scan(&hapus.nama)
					if hapus.nama == "0" {
						cekHapus1, cekHapus = false, false
					}
					if cekHapus1 {
						for _, isi := range dataPeminjam {
							if isi.nama == hapus.nama {
								duplikatHapus++ // 2
							}
						}
						if duplikatHapus > 1 {
							fmt.Println("Data duplikat, pilih salah satu:")
							for i, isi := range dataPeminjam {
								if isi.nama == hapus.nama {
									fmt.Println(">", isi, "pilih", i)
								}
							}
							fmt.Print("Pilihan ubah: ")
							fmt.Scan(&indeksHapus)
							fmt.Print("Ketik 'KONFIRMASI' untuk lanjut: ")
							fmt.Scan(&konfirmasi)
							cekHapus = false
						} else if duplikatHapus == 1 {
							for i, isi := range dataPeminjam {
								if isi.nama == hapus.nama {
									indeksHapus = i
								}
							}
							fmt.Print("Ketik 'KONFIRMASI' untuk lanjut: ")
							fmt.Scan(&konfirmasi)
							cekHapus = false
						} else {
							fmt.Println("DATA TIDAK ADA ATAU TIDAK DITEMUKAN")
							fmt.Println("------------------------------------")
						}
						if konfirmasi == "KONFIRMASI" {
							hapusData(&dataPeminjam, iTambah, indeksHapus)
							iTambah--
						}
					}
				}
				fmt.Println("====================================")
			case "4":
				// Pake sequential search karena berdasarkan nama
				fmt.Println("\n=============== MENCARI DATA ===============")
				var cekData_4 bool = false
				var cekNama_4 string
				jumData_4 := 0
				fmt.Print("Nama: ")
				fmt.Scan(&cekNama_4)
				for i := 0; i < iTambah; i++ {
					if dataPeminjam[i].nama == cekNama_4 {
						fmt.Println(">", dataPeminjam[i])
						cekData_4 = true
						jumData_4++
					}
				}
				if !cekData_4 {
					fmt.Println("SYSTEM: DATA TIDAK DITEMUKAN!")
				} else {
					fmt.Printf("SYSTEM: TERDAPAT %d BUAH DATA\n", jumData_4)
				}
				fmt.Println("============================================")
			case "5":
				var pilihan1_5, pilihan_5 string
				cekCase5 := true
				fmt.Println("\n=============== MENGURUTKAN ===============")
				for cekCase5 && iTambah >= 1 {
					cek1Case5 := true
					fmt.Println("1. Jumlah Pinjaman")
					fmt.Println("2. Tenor")
					fmt.Println("3. Keluar")
					fmt.Print("Pilihan: ")
					fmt.Scan(&pilihan_5)
					switch pilihan_5 {
					case "1":
						fmt.Println("---------------------------")
						for cek1Case5 {
							fmt.Println("Ingin diurutkan secara:")
							fmt.Println("1. Ascending")
							fmt.Println("2. Descending")
							fmt.Print("Pilihan: ")
							fmt.Scan(&pilihan1_5)
							switch pilihan1_5 {
							case "1":
								ascending(iTambah, pilihan_5)
								cek1Case5 = false
								fmt.Println("SYSTEM: DATA BERHASIL DIURUTKAN BERDASARKAN JUMLAH PINJAMAN SECARA ASCENDING!")
							case "2":
								descending(iTambah, pilihan_5)
								cek1Case5 = false
								fmt.Println("SYSTEM: DATA BERHASIL DIURUTKAN BERDASARKAN JUMLAH PINJAMAN SECARA DESCENDING!")
							default:
								fmt.Println("SYSTEM: Masukkan tidak valid!")
								fmt.Println("---------------------------")
							}
						}
						cekCase5 = false
					case "2":
						fmt.Println("---------------------------")
						for cek1Case5 {
							fmt.Println("Ingin diurutkan secara:")
							fmt.Println("1. Ascending")
							fmt.Println("2. Descending")
							fmt.Print("Pilihan: ")
							fmt.Scan(&pilihan1_5)
							switch pilihan1_5 {
							case "1":
								ascending(iTambah, pilihan_5)
								cek1Case5 = false
								fmt.Println("SYSTEM: DATA BERHASIL DIURUTKAN BERDASARKAN TENOR PINJAMAN SECARA ASCENDING!")
							case "2":
								descending(iTambah, pilihan_5)
								cek1Case5 = false
								fmt.Println("SYSTEM: DATA BERHASIL DIURUTKAN BERDASARKAN TENOR SECARA DESCENDING!")
							default:
								fmt.Println("SYSTEM: Masukkan tidak valid!")
								fmt.Println("---------------------------")
							}
						}
						cekCase5 = false
					case "3":
						cekCase5 = false
					default:
						fmt.Println("SYSTEM: Masukkan tidak valid!")
						fmt.Println("---------------------------")
					}
				}
				if iTambah == 0 {
					fmt.Println("SYSTEM: DATA BELUM TERISI!")
				}
				fmt.Println("===========================================")
			case "6":
				fmt.Println("\n=============== MENAMPILKAN DATA =============== ")
				showData(dataPeminjam, iTambah)
				fmt.Println("================================================ ")
			case "7":
				cek = true
				cekMenu = false
			default:
				fmt.Println("SYSTEM: Input tidak valid!")
			}
			fmt.Println()
		}
	}
	fmt.Println("SYSTEM: Terimakasih telah menggunakan layanan kami!🙏😊")
	fmt.Println("===============================")
	fmt.Println("|Build by:                    |")
	fmt.Println("|Dimas Ramadhani/103112400065 |")
	fmt.Println("|Ryan Akeyla                  |")
	fmt.Println("===============================")
}
