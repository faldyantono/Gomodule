package modulego

import "fmt"

//function sebagai parameter
type Censored func(string) string

func SayHelloCensored(name string, censored Censored) {
	nameCensored := censored(name)
	fmt.Println("Selamat Datang", nameCensored, "--Portal Resmi SAMSAT Kota Depok--")
}
func FilterSensor(name string string) string {
	if name == "babi" {
		return "*a*i"
	} else if name == "anjing" {
		return "an****"
	} else {
		return name
	}
}

//function return value
func Pengguna(name string) string {
	if name == "" {
		return "Mohon isi identitas pemilik kendaraan anda dengan huruf."
	} else {
		return "BPKB atas nama                      			: " + name
	}
}

//struct
type Pajak struct {
	BanyakKendaraan, Gaji, Tunjangan int
}

func (n Pajak) Perkalian() int {
	return ((n.Gaji + n.Tunjangan) * n.BanyakKendaraan) / 100
}

//interface
type HitungPajak interface {
	Perkalian() int
}

//struct method
type Denda struct {
	MasaHabisSIM, MasaHabisSTNK, MasaHabisKTP int
}

func (s Denda) Perkalian2() int {
	return ((s.MasaHabisKTP * 20000) + (s.MasaHabisSTNK * 200000) + (s.MasaHabisSIM * 40000))
}

//interface
type HitungDenda interface {
	Perkalian2() int
}
