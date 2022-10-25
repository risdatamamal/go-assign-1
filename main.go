package main

import (
	"fmt"
	"os"
	"strconv"
)

type siswa struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	xstr := os.Args[1]

	x, _ := strconv.Atoi(xstr)

	var dbSiswa = []siswa{
		{Name: "Tamam", Alamat: "Haji Kodir", Pekerjaan: "Mahasiswa", Alasan: "Ingin jadi backend developer"},
		{Name: "Budi", Alamat: "Palem Merah", Pekerjaan: "Mahasiswa", Alasan: "Supaya jadi programmer handal"},
		{Name: "Asep", Alamat: "Palem X", Pekerjaan: "Mahasiswa", Alasan: "Mau jadi backend engineer"},
		{Name: "Daffa", Alamat: "Palem V", Pekerjaan: "Mahasiswa", Alasan: "Modal buat cari kerja"},
		{Name: "Tole", Alamat: "Matra", Pekerjaan: "PNS", Alasan: "Memperdalam ilmu coding"},
		{Name: "Rizky", Alamat: "Sumpah Pemuda", Pekerjaan: "Mahasiswa", Alasan: "Dapetin skill baru"},
		{Name: "Mamat", Alamat: "Halim", Pekerjaan: "Pengangguran", Alasan: "Nyari kesibukan"},
		{Name: "Aco", Alamat: "Jakarta barat", Pekerjaan: "Mahasiswa", Alasan: "Diajak teman"},
		{Name: "Raihan", Alamat: "Tangerang", Pekerjaan: "Mahasiswa", Alasan: "Disuruh orang tua"},
		{Name: "Parija", Alamat: "BSD", Pekerjaan: "Freelancer", Alasan: "Dapetin sertifikat"},
	}

	if x > len(dbSiswa) || x < 1 {
		showError()
	} else {
		tampilData(dbSiswa[x-1])
	}
}

func tampilData(siswa siswa) {
	fmt.Println("Nama:", siswa.Name)
	fmt.Println("Alamat:", siswa.Alamat)
	fmt.Println("Pekerjaan: ", siswa.Pekerjaan)
	fmt.Println("Alasan:", siswa.Alasan)
}

func showError() {
	fmt.Println("No Absen Tidak Ditemukan")
}
