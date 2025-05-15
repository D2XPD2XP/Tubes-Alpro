package main

import "fmt"

type data_rekening struct {
	nama_depan, nama_belakang, nama_bank, cabang_bank, username, password string
	tanggal_lahir, tahun_lahir, bulan_lahir, no_rekening, saldo           int
}

type history_rekening struct {
	namadepan_history, namabelakang_history, nama_bank, nd_pengirim, nb_pengirim                      string
	no_rekening, total_transaksi, tanggal_transaksi, bulan_transaksi, tahun_transaksi, norek_pengirim int
}

type data_nasabah struct {
	namadepan, namabelakang, bank_nasabah string
	norek, NIK                            int
}

const MaxRekening int = 50
const MaxHistory int = 1000
const MaxNasabah int = 50

type tabRekening [MaxRekening]data_rekening
type tabHistory [MaxHistory]history_rekening
type tabNasabah [MaxNasabah]data_nasabah

func main() {
	var rekening tabRekening
	var history tabHistory
	var nasabah tabNasabah
	var nRekening, nPenerima, nNasabah int
	var input_menu int
	for {
		bank_interface()
		fmt.Scan(&input_menu)
		for input_menu <= 0 || input_menu > 4 {
			fmt.Println("Silahkan Pilih Menu Yang Tersedia!")
			fmt.Scan(&input_menu)
		}
		if input_menu == 1 {
			if nRekening < MaxRekening {
				simpan_rekening(&nasabah, &rekening, &nRekening, nNasabah)
			} else {
				fmt.Println("Anda Sudah Mencapai Batas Penambahan Rekening :(")
			}
		}
		if input_menu == 2 {
			transfer_saldo(&rekening, &history, nRekening, &nPenerima)
		}
		if input_menu == 3 {
			dataRekening(&rekening, &history, &nRekening, &nPenerima)
		}
		if input_menu == 4 {
			input_nasabah(&nasabah, &nNasabah)
		}
	}
}

func bank_interface() {
	fmt.Println("======== BANK MANAGER =======")
	fmt.Println("=============================")
	fmt.Println("|| 1. Simpan Rekening      ||")
	fmt.Println("=============================")
	fmt.Println("=============================")
	fmt.Println("|| 2. Transfer             ||")
	fmt.Println("=============================")
	fmt.Println("|| 3. Data Rekening        ||")
	fmt.Println("=============================")
	fmt.Println("=============================")
	fmt.Println("|| 4. Data Nasabah         ||")
	fmt.Println("=============================")
	fmt.Println("SILAHKAN PILIH MENU : ")
}

func input_nasabah(nasabah *tabNasabah, ns *int) {
	var i int
	var nd_valid, nb_valid bool
	i = 0 + *ns
	fmt.Println("======= DATA NASABAH =======")
	fmt.Print("Nama Depan: ")
	fmt.Scan(&nasabah[i].namadepan)
	nd_valid = koreksi_string(nasabah[i].namadepan)
	for nd_valid == false {
		fmt.Println("===== NAMA TIDAK VALID =====")
		fmt.Print("Nama Depan: ")
		fmt.Scan(&nasabah[i].namadepan)
		nd_valid = koreksi_string(nasabah[i].namadepan)
	}
	fmt.Print("Nama Belakang: ")
	fmt.Scan(&nasabah[i].namabelakang)
	nb_valid = koreksi_string(nasabah[i].namabelakang)
	for nb_valid == false {
		fmt.Println("===== NAMA TIDAK VALID =====")
		fmt.Print("Nama Depan: ")
		fmt.Scan(&nasabah[i].namabelakang)
		nb_valid = koreksi_string(nasabah[i].namabelakang)
	}
	if nasabah[i].namabelakang == "-" {
		nasabah[i].namabelakang = " "
	}
	fmt.Print("NIK: ")
	fmt.Scan(&nasabah[i].NIK)
	for koreksi_number(nasabah[i].NIK) == false {
		fmt.Println("===== NIK TIDAK VALID =====")
		fmt.Print("NIK: ")
		fmt.Scan(&nasabah[i].NIK)
		koreksi_number(nasabah[i].NIK)
	}
	fmt.Print("Bank: ")
	fmt.Scan(&nasabah[i].bank_nasabah)
	for koreksi_string(nasabah[i].bank_nasabah) == false {
		fmt.Println("==== NAMA BANK TIDAK VALID ====")
		fmt.Print("Bank: ")
		fmt.Scan(&nasabah[i].bank_nasabah)
		koreksi_string(nasabah[i].bank_nasabah)
	}
	fmt.Print("No Rekening: ")
	fmt.Scan(&nasabah[i].norek)
	if norek_ada(nasabah, nasabah[i].norek, ns) == true {
		fmt.Println("==== NO REKENING SUDAH ADA ====")
		fmt.Print("No Rekening: ")
	}
	for norek_ada(nasabah, nasabah[i].norek, ns) == true {
		fmt.Scan(&nasabah[i].norek)
	}
	for koreksi_number(nasabah[i].norek) == false {
		fmt.Println("==== NO REKENING TIDAK VALID ====")
		fmt.Print("No Rekening: ")
		fmt.Scan(&nasabah[i].norek)
		koreksi_number(nasabah[i].norek)
	}
	*ns++
	fmt.Println("======== DATA TERSIMPAN =======")
}

func norek_ada(nasabah *tabNasabah, input int, ns *int) bool {
	for i := 0; i < *ns; i++ {
		if input == nasabah[i].norek {
			return true
		}
	}
	return false
}
func koreksi_string(input string) bool {
	for j := 0; j < len(input); j++ {
		if input[0] == '-' {
			return true
		}
		if input[j] >= '!' && input[j] <= '@' {
			return false
		}
	}
	return true
}

func koreksi_number(input int) bool {
	if input == 0 {
		return false
	}
	return true
}

func simpan_rekening(nasabah *tabNasabah, rekening *tabRekening, n *int, ns int) {
	var input_menu string
	var nama_sesuai, tl_sesuai bool
	var i int
	if ns == 0 {
		fmt.Println("=========== TIDAK ADA DATA NASABAH ============")
		fmt.Println("SILAHKAN TAMBAHKAN DATA NASABAH TERLEBIH DAHULU")
		fmt.Println("===============================================")
	} else {
		i = 0 + *n
		fmt.Println("===== SILAHKAN ISI DATA BERIKUT =====")
		fmt.Print("Nama Depan: ")
		fmt.Scan(&rekening[i].nama_depan)
		fmt.Print("Nama Belakang: ")
		fmt.Scan(&rekening[i].nama_belakang)
		if rekening[i].nama_belakang == "-" {
			rekening[i].nama_belakang = " "
		}
		nama_sesuai = check_nama(nasabah, rekening, ns, i)
		for nama_sesuai == false {
			fmt.Println("== NAMA NASABAH TIDAK DITEMUKAN ==")
			fmt.Print("Nama Depan: ")
			fmt.Scan(&rekening[i].nama_depan)
			fmt.Print("Nama Belakang: ")
			fmt.Scan(&rekening[i].nama_belakang)
			if rekening[i].nama_belakang == "-" {
				rekening[i].nama_belakang = " "
			}
			nama_sesuai = check_nama(nasabah, rekening, ns, i)
		}
		fmt.Print("Tanggal Lahir (DD-MM-YYYY): ")
		fmt.Scan(&rekening[i].tanggal_lahir, &rekening[i].bulan_lahir, &rekening[i].tahun_lahir)
		tl_sesuai = check_tanggallahir(rekening, i)
		if tl_sesuai == false {
			fmt.Println("=== TANGGAL LAHIR TIDAK VALID ===")
			fmt.Print("Tanggal Lahir (DD-MM-YYYY): ")
		}
		for tl_sesuai == false {
			fmt.Scan(&rekening[i].tanggal_lahir, &rekening[i].bulan_lahir, &rekening[i].tahun_lahir)
			tl_sesuai = check_tanggallahir(rekening, i)
		}
		fmt.Print("Bank: ")
		fmt.Scan(&rekening[i].nama_bank)
		for check_bank(nasabah, rekening, ns, i) == false {
			fmt.Println("===== BANK TIDAK SESUAI =====")
			fmt.Print("Bank: ")
			fmt.Scan(&rekening[i].nama_bank)
			check_bank(nasabah, rekening, ns, i)
		}
		fmt.Print("Cabang Bank: ")
		fmt.Scan(&rekening[i].cabang_bank)
		for koreksi_string(rekening[i].cabang_bank) == false {
			fmt.Println("===== CABANG BANK TIDAK VALID =====")
			fmt.Print("Cabang Bank: ")
			fmt.Scan(&rekening[i].cabang_bank)
			koreksi_string(rekening[i].cabang_bank)
		}
		fmt.Print("No Rekening: ")
		fmt.Scan(&rekening[i].no_rekening)
		for check_norek(nasabah, rekening, ns, i) == false {
			fmt.Println("=== NO REKENING TIDAK DITEMUKAN ===")
			fmt.Print("No Rekening: ")
			fmt.Scan(&rekening[i].no_rekening)
		}
		fmt.Print("Saldo: ")
		fmt.Scan(&rekening[i].saldo)
		for koreksi_number(rekening[i].saldo) == false {
			fmt.Println("===== SALDO TIDAK VALID =====")
			fmt.Print("Saldo: ")
			fmt.Scan(&rekening[i].saldo)
		}
		fmt.Println("=======  USERNAME & PASSWORD =======")
		fmt.Println("Buat Username & Password tanpa Menggunakan Spasi")
		fmt.Print("Username: ")
		fmt.Scan(&rekening[i].username)
		fmt.Print("Password: ")
		fmt.Scan(&rekening[i].password)
		fmt.Println("===== REKENING BERHASIL DIBUAT =====")
		*n++
		fmt.Print("TAMBAH REKENING LAGI (Y/G): ")
		fmt.Scan(&input_menu)
		for input_menu != "Y" && input_menu != "G" {
			fmt.Println("Pilih (Y/G): ")
			fmt.Scan(&input_menu)
		}
		if input_menu == "Y" {
			simpan_rekening(nasabah, rekening, n, ns)
		}
		if input_menu == "G" {

		}
	}
}

func check_nama(nasabah *tabNasabah, rekening *tabRekening, ns, i int) bool {
	var found bool
	var j int
	found = false
	for j < ns && found == false {
		if rekening[i].nama_depan == nasabah[j].namadepan && rekening[i].nama_belakang == nasabah[j].namabelakang {
			found = true
		}
		j++
	}
	return found
}

func check_tanggallahir(rekening *tabRekening, i int) bool {
	var check bool
	if rekening[i].tanggal_lahir <= 0 || rekening[i].tanggal_lahir > 31 || rekening[i].bulan_lahir <= 0 || rekening[i].bulan_lahir > 12 || rekening[i].tahun_lahir <= 0 {
		check = false
	} else {
		check = true
	}
	return check
}

func check_bank(nasabah *tabNasabah, rekening *tabRekening, ns, i int) bool {
	var found bool
	var j int
	found = false
	for j < ns && found == false {
		if rekening[i].nama_depan == nasabah[j].namadepan && rekening[i].nama_belakang == nasabah[j].namabelakang && rekening[i].nama_bank == nasabah[j].bank_nasabah {
			found = true
		}
		j++
	}
	return found
}

func check_norek(nasabah *tabNasabah, rekening *tabRekening, ns, i int) bool {
	var found bool
	var j int
	found = false
	for j < ns && found == false {
		if rekening[i].nama_depan == nasabah[j].namadepan && rekening[i].nama_belakang == nasabah[j].namabelakang && rekening[i].no_rekening == nasabah[j].norek {
			found = true
		}
		j++
	}
	return found
}

func transfer_saldo(rekening *tabRekening, history *tabHistory, nS int, nP *int) {
	var input_norek, index_rekening, input_transfer, norek_penerima, tanggal, bulan, tahun int
	var input_username, input_password, input_option, namadepan, namabelakang, bank_penerima string
	var rekening_ada, status_transaksi bool
	fmt.Print("Silahkan Masukan No Rekening Anda: ")
	fmt.Scan(&input_norek)
	for i := 0; i < nS; i++ {
		if input_norek == rekening[i].no_rekening {
			rekening_ada = true
			index_rekening = i
		}
	}
	if rekening_ada == true {
		fmt.Print("Username: ")
		fmt.Scan(&input_username)
		for input_username != rekening[index_rekening].username {
			fmt.Println("USERNAME SALAH SILAHKAN COBA LAGI")
			fmt.Print("Username: ")
			fmt.Scan(&input_username)
		}
		fmt.Print("Password: ")
		fmt.Scan(&input_password)
		for input_password != rekening[index_rekening].password {
			fmt.Println("PASSWORD SALAH SILAHKAN COBA LAGI")
			fmt.Print("Password: ")
			fmt.Scan(&input_password)
		}
		fmt.Println("====== LOGIN BERHASIL ======")
		fmt.Print("Bank tujuan: ")
		fmt.Scan(&bank_penerima)
		for koreksi_string(bank_penerima) == false {
			fmt.Println("===== BANK TIDAK VALID =====")
			fmt.Print("Bank tujuan: ")
			fmt.Scan(&bank_penerima)
			koreksi_string(bank_penerima)
		}
		fmt.Print("No Rekening: ")
		fmt.Scan(&norek_penerima)
		for koreksi_number(norek_penerima) == false {
			fmt.Println("==== NO REKENING TIDAK VALID ====")
			fmt.Print("No Rekening: ")
			fmt.Scan(&norek_penerima)
			koreksi_number(norek_penerima)
		}
		fmt.Print("Nama Depan Penerima: ")
		fmt.Scan(&namadepan)
		for koreksi_string(namadepan) == false {
			fmt.Println("==== NAMA TIDAK VALID ====")
			fmt.Print("Nama Depan Penerima: ")
			fmt.Scan(&namadepan)
			koreksi_string(namadepan)
		}
		fmt.Print("Nama Belakang Penerima: ")
		fmt.Scan(&namabelakang)
		for koreksi_string(namabelakang) == false {
			fmt.Println("==== NAMA TIDAK VALID ====")
			fmt.Print("Nama Belakang Penerima: ")
			fmt.Scan(&namabelakang)
			koreksi_string(namabelakang)
		}
		fmt.Print("Masukan Nominal Transfer: ")
		fmt.Scan(&input_transfer)
		for input_transfer < 50000 {
			fmt.Println("Jumlah Minimal Transfer Adalah Rp.50.000")
			fmt.Print("Masukan Nominal Transfer: ")
			fmt.Scan(&input_transfer)
		}
		fmt.Print("Tanggal Transaksi (DD-MM-YYYY): ")
		fmt.Scan(&tanggal, &bulan, &tahun)
		for koreksi_tanggaltransaksi(tanggal, bulan, tahun) == false {
			fmt.Println("=== TANGGAL TRANSAKSI TIDAK VALID ===")
			fmt.Print("Tanggal Transaksi(DD-MM-YYYY): ")
			fmt.Scan(&tanggal, &bulan, &tahun)
		}
		if rekening[index_rekening].saldo > input_transfer {
			status_transaksi = true
		} else {
			status_transaksi = false
		}
		if status_transaksi == true {
			fmt.Println("====== PENGIRIM ======")
			fmt.Print("Nama: ")
			fmt.Print(rekening[index_rekening].nama_depan, " ", rekening[index_rekening].nama_belakang)
			fmt.Println(" ")
			fmt.Print("No Rekening: ")
			fmt.Print(rekening[index_rekening].no_rekening)
			fmt.Println(" ")
			fmt.Print("Bank: ")
			fmt.Print(rekening[index_rekening].nama_bank)
			fmt.Println(" ")
			fmt.Println("====== PENERIMA ======")
			fmt.Print("Nama: ")
			fmt.Print(namadepan, " ", namabelakang)
			fmt.Println(" ")
			fmt.Print("No Rekening: ")
			fmt.Print(norek_penerima)
			fmt.Println(" ")
			fmt.Print("Bank: ")
			fmt.Print(bank_penerima)
			fmt.Println(" ")
			fmt.Println("======================")
			fmt.Print("Tanggal: ")
			if bulan < 10 && tanggal < 10 {
				fmt.Print("0", tanggal, " ", "0", bulan, " ", tahun)
			} else if tanggal < 10 && tahun < 10 {
				fmt.Print("0", tanggal, " ", bulan, " ", "0", tahun)
			} else if tahun < 10 && bulan < 10 {
				fmt.Print(tanggal, " ", "0", bulan, " ", "0", tahun)
			} else if tanggal < 10 && bulan < 10 && tahun < 10 {
				fmt.Print("0", tanggal, " ", "0", bulan, " ", "0", tahun)
			} else if tanggal < 10 {
				fmt.Print("0", tanggal, " ", bulan, " ", tahun)
			} else if bulan < 10 {
				fmt.Print(tanggal, " ", "0", bulan, " ", tahun)
			} else if tahun < 10 {
				fmt.Print(tanggal, " ", bulan, " ", "0", tahun)
			} else {
				fmt.Print(tanggal, " ", bulan, " ", tahun)
			}
			fmt.Println(" ")
			fmt.Println("Lanjutkan Transaksi? (Y/G)")
			fmt.Scanln(&input_option)
			for input_option != "Y" && input_option != "G" {
				fmt.Print("Pilih (Y/G): ")
				fmt.Scan(&input_option)
			}
			if input_option == "Y" {
				fmt.Println("===== TRANSAKSI BERHASIL =====")
				rekening[index_rekening].saldo -= input_transfer
				history[*nP].nama_bank = bank_penerima
				history[*nP].no_rekening = norek_penerima
				history[*nP].namadepan_history = namadepan
				history[*nP].namabelakang_history = namabelakang
				history[*nP].tanggal_transaksi = tanggal
				history[*nP].bulan_transaksi = bulan
				history[*nP].tahun_transaksi = tahun
				history[*nP].total_transaksi = input_transfer
				history[*nP].norek_pengirim = rekening[index_rekening].no_rekening
				fmt.Print("Sisa Saldo Anda: ")
				fmt.Print("Rp. ", rekening[index_rekening].saldo)
				fmt.Println(" ")
				*nP++
			}
			if input_option == "G" {
				fmt.Println("======= TRANSAKSI GAGAL =======")
			}
		}
		if status_transaksi == false {
			fmt.Println("====== SALDO ANDA TIDAK CUKUP ======")
		}
	}
	if rekening_ada == false {
		fmt.Println("Rekening Tidak Tersedia")
		if input_option != "Y" && input_option != "G" {
			fmt.Println("Coba Lagi (Y/G): ")
		}
		for input_option != "Y" && input_option != "G" {
			fmt.Scanln(&input_option)
		}
		if input_option == "Y" {
			transfer_saldo(rekening, history, nS, nP)
		}
		if input_option == "G" {

		}
	}
}

func koreksi_tanggaltransaksi(tanggal, bulan, tahun int) bool {
	if tanggal <= 0 || tanggal > 31 || bulan <= 0 || bulan > 12 || tahun == 0 {
		return false
	}
	return true
}

func dataRekening(rekening *tabRekening, history *tabHistory, nS, nP *int) {
	var input_norek, idx int
	var input_password, input_username, input_option string
	var rekening_ada bool

	fmt.Print("Silahkan Masukan No Rekening Anda: ")
	fmt.Scan(&input_norek)

	for i := 0; i < *nS; i++ {
		if input_norek == rekening[i].no_rekening {
			rekening_ada = true
			idx = i
		}
	}

	if rekening_ada {
		fmt.Print("Username: ")
		fmt.Scan(&input_username)

		for input_username != rekening[idx].username {
			fmt.Println("USERNAME SALAH SILAHKAN COBA LAGI")
			fmt.Print("Username: ")
			fmt.Scan(&input_username)
		}

		fmt.Print("Password: ")
		fmt.Scan(&input_password)

		for input_password != rekening[idx].password {
			fmt.Println("PASSWORD SALAH SILAHKAN COBA LAGI")
			fmt.Print("Password: ")
			fmt.Scan(&input_password)
		}

		fmt.Println("====== LOGIN BERHASIL ======")

		fmt.Println("======= DATA REKENING ANDA =======")
		fmt.Println("Nama Depan:", rekening[idx].nama_depan)
		fmt.Println("Nama Belakang:", rekening[idx].nama_belakang)
		fmt.Println("Tanggal Lahir:", rekening[idx].tanggal_lahir, "/", rekening[idx].bulan_lahir, "/", rekening[idx].tahun_lahir)
		fmt.Println("Bank:", rekening[idx].nama_bank)
		fmt.Println("Cabang Bank:", rekening[idx].cabang_bank)
		fmt.Println("No Rekening:", rekening[idx].no_rekening)
		fmt.Println("Saldo:", rekening[idx].saldo)
		fmt.Println("===================================")

		fmt.Println("Menu:")
		fmt.Println("1. Edit Data Rekening")
		fmt.Println("2. History Transaksi")
		fmt.Println("3. Hapus Rekening")
		fmt.Println("4. Kembali ke Main Menu")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&input_option)

		for input_option != "4" {
			if input_option == "1" {
				edit_transfer(rekening, history, idx)
				dataRekening(rekening, history, nS, nP)
			}
			if input_option == "2" {
				history_transaksi(rekening, history, nP, idx)
				input_option = "4"
			}
			if input_option == "3" {
				delete_rekening(rekening, history, nS, nP, idx)
				input_option = "4"
			}
		}

	} else {
		fmt.Println("Rekening Tidak Tersedia")
		fmt.Print("Coba Lagi (Y/G): ")
		fmt.Scan(&input_option)
		for input_option != "Y" && input_option != "G" {
			fmt.Print("Pilih (Y/G): ")
			fmt.Scan(&input_option)
		}
		if input_option == "Y" {
			dataRekening(rekening, history, nS, nP)
		} else {

		}
	}
}

func edit_transfer(rekening *tabRekening, history *tabHistory, idx int) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println("===== EDIT REKENING =====")
		fmt.Println("1. Edit Nama")
		fmt.Println("2. Edit Tanggal Lahir")
		fmt.Println("3. Kembali Ke Menu Sebelumnya")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Print("Nama Depan Baru: ")
			fmt.Scan(&rekening[idx].nama_depan)
			fmt.Print("Nama Belakang Baru: ")
			fmt.Scan(&rekening[idx].nama_belakang)
			fmt.Println("Informasi rekening berhasil diubah!")
		}
		if pilihan == 2 {
			fmt.Print("Tanggal Lahir Baru (DD/MM/YYYY): ")
			fmt.Scan(&rekening[idx].tanggal_lahir, &rekening[idx].bulan_lahir, &rekening[idx].tahun_lahir)
			fmt.Println("Informasi rekening berhasil diubah!")
		}
		if pilihan != 1 && pilihan != 2 && pilihan != 3 {
			fmt.Println("Pilihan tidak valid.")
			fmt.Print("Coba Lagi: ")
			fmt.Scan(&pilihan)
		}
	}
}

func history_transaksi(rekening *tabRekening, history *tabHistory, nP *int, idx int) {
	var option int
	fmt.Println("===== HISTORY TRANSFER =====")
	sort_history(history, nP)
	for i := 0; i < *nP; i++ {
		if rekening[idx].no_rekening == history[i].norek_pengirim {
			fmt.Printf("Nama Penerima: %s %s\n", history[i].namadepan_history, history[i].namabelakang_history)
			fmt.Printf("Nominal: Rp.%d\n", history[i].total_transaksi)
			fmt.Printf("Tanggal: %d/%d/%d\n", history[i].tanggal_transaksi, history[i].bulan_transaksi, history[i].tahun_transaksi)
			fmt.Printf("Bank Penerima: %s\n", history[i].nama_bank)
			fmt.Printf("No Rekening Penerima: %d\n", history[i].no_rekening)
			fmt.Println("===========================")
		}
	}
	fmt.Println("1. Hapus History")
	fmt.Println("2. Kembali Ke Main Menu")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&option)
	for option != 2 {
		if option == 1 {
			delete_history(rekening, history, nP, idx)
		}
		if option != 1 {
			fmt.Println("=== Silahkan Pilih Menu Yang Tersedia ===")
			fmt.Print("Pilih Menu: ")
			fmt.Scan(&option)
		}
	}
}

func delete_rekening(rekening *tabRekening, history *tabHistory, nR, nP *int, idx int) {
	fmt.Println("======== DELETE REKENING ========")
	for j := 0; j < *nP; j++ {
		if rekening[idx].no_rekening == history[j].norek_pengirim {
			for k := j; k < *nP-1; k++ {
				history[k] = history[k+1]
			}
			*nP--
		}
	}
	for i := idx; i < *nR-1; i++ {
		rekening[i] = rekening[i+1]
	}
	*nR--
	fmt.Println("===== REKENING BERHASIL DIHAPUS =====")
}

func delete_history(rekening *tabRekening, history *tabHistory, nP *int, idx int) {
	for j := 0; j < *nP; j++ {
		if rekening[idx].no_rekening == history[j].norek_pengirim {
			for k := j; k < *nP-1; k++ {
				history[k] = history[k+1]
			}
			*nP--
		}
	}
}

func sort_history(history *tabHistory, nP *int) {
	var pass, i, idx int
	var temp history_rekening
	pass = 1
	for pass < *nP {
		idx = pass - 1
		i = pass
		for i < *nP {
			if history[idx].tahun_transaksi == history[i].tahun_transaksi && history[idx].bulan_transaksi == history[i].bulan_transaksi && history[idx].tanggal_transaksi < history[i].tanggal_transaksi {
				idx = i
			}
			if history[idx].tahun_transaksi == history[i].tahun_transaksi && history[idx].bulan_transaksi < history[i].bulan_transaksi {
				idx = i
			}
			if history[idx].tahun_transaksi < history[i].tahun_transaksi {
				idx = i
			}
			i++
		}
		temp = history[pass-1]
		history[pass-1] = history[idx]
		history[idx] = temp
		pass++
	}
}
