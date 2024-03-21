package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	dataMahasiswa := make(map[string]Mahasiswa)

	dataMahasiswa["4522210002"] = Mahasiswa{"Dimas Satrio Parikesit", "4522210002", "Teknik Informatika"}
	dataMahasiswa["4522210016"] = Mahasiswa{"Muhammad Sulthan Zharfan", "4522210016", "Teknik Informatika"}
	dataMahasiswa["4523210018"] = Mahasiswa{"Jimly Asidiq Anwar", "4522210016", "Teknik Informatika"}
	dataMahasiswa["4523210033"] = Mahasiswa{"Faathir Akbar Nugroho", "4522210016", "Teknik Informatika"}

	for npm, mahasiswa := range dataMahasiswa {
		fmt.Printf("NPM: %s, Nama: %s, Jurusan: %s\n", npm, mahasiswa.Nama, mahasiswa.Jurusan)
	}

	cariNPM := "4522210002"
	mahasiswa, found := dataMahasiswa[cariNPM]
	if found {
		fmt.Printf("\nData Mahasiswa dengan NPM %s ditemukan:\n", cariNPM)
		fmt.Printf("Nama: %s, Jurusan: %s\n", mahasiswa.Nama, mahasiswa.Jurusan)
	} else {
		fmt.Printf("\nData Mahasiswa dengan NPM %s tidak ditemukan.\n", cariNPM)
	}
}
