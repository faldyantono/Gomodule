package main

import (
	"fmt"
)

func main() {
	//output dari function parameter
	SayHelloCensored("Haris", FilterSensor)
	//output dari return value
	fmt.Println(Pengguna("Haris"))
  //struct
  var x1 karyawan
  x1.gaji = 4500000
  x1.pekerjaan = "Tukang sapu"
  x1.umur = 13
  fmt.Println("Gaji anda                            :", x1.gaji)
  fmt.Println("Pekerjaan                            :", x1.pekerjaan)
  fmt.Println("Umur                                 :", x1.umur)
  //anonymous function
  func(i int) {
    fmt.Println("Jumlah kendaraan anda adalah         :", i, "unit")
  }(4)
  //struct method
  var s1 kelengkapan
  s1.KTP = 31748530911
  s1.SIM = "Berlaku Hingga 2024"
  s1.STNK = "Habis"
  s1.DENDA = x1.gaji * 0.01
  s1.administrasi()
	periodepajak, bobotpajak := BobotPajak(25, 125)
  //anonymous struct
  vehicle := struct {
    brand string
    breed string
    class string
    cc int
  }{
    brand: "Honda",
    breed: "Supra-X",
    class: "Bebek",
    cc: 125,
  }
  fmt.Println("Kendaraan anda adalah                :", vehicle)
	//output dari multiple return value
	fmt.Println("Sisa Pajak Anda Adalah               :", periodepajak, "Bulan")
  fmt.Println("-------------------------KETERANGAN----------------------------")
	fmt.Println("Pajak yang harus dibayar hingga tahun berikutnya : ", bobotpajak)
	if periodepajak > 12 {
		fmt.Println("Anda menerima subsidi sebesar 30%")
	} else if periodepajak > 6 {
		fmt.Println("Anda menerima subsidi sebesar 10%")
	} else if periodepajak > 3 {
		fmt.Println("Anda tidak menerima subsidi")
	} else if periodepajak < 3 {
		fmt.Println("Anda terkena pajak progresif")
	} else {
		fmt.Println("Kendaraan anda kami sita")
	}


}

//function sebagai parameter
type Censored func(string) string

func SayHelloCensored(name string, censored Censored) {
	nameCensored := censored(name)
	fmt.Println("Selamat Datang", nameCensored, "--Portal Resmi SAMSAT Kota Depok--")
}
func FilterSensor(name string) string {
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
	if name == " " {
		return "Mohon isi identitas pemilik kendaraan anda dengan huruf."
	} else {
		return "BPKB atas nama                       : " + name
	}
}

//function multiple return value
func BobotPajak(periodepajak int, ccmesin int) (int, int) {
	Periodepajak := 25
	bobotpajak := ((ccmesin * 1) - (periodepajak * 4)) * 10000
	return Periodepajak, bobotpajak
}

//struct
type karyawan struct {
	gaji      float64
	pekerjaan string
	umur      int
}

//struct method
type kelengkapan struct {
  SIM, STNK string
	KTP int
  DENDA float64
}

func (identitas kelengkapan) administrasi() {
	fmt.Println("Nomor Induk KTP                      :", identitas.KTP)
	fmt.Println("Masa Berlaku SIM                     :", identitas.SIM)
	fmt.Println("Masa Berlaku STNK                    :", identitas.STNK)
	fmt.Println("ANDA MEMILIKI DENDA SEBESAR          :", identitas.DENDA)
}
