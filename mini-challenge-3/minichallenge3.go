package minichallenge3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	index     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func (e Employee) person() {
	fmt.Println("Index:", e.index)
	fmt.Println("Nama:", e.nama)
	fmt.Println("Alamat:", e.alamat)
	fmt.Println("Pekerjaan:", e.pekerjaan)
	fmt.Println("Alasan:", e.alasan)
}

func MiniChallenge3() {
	// Mini Challenge 3 -- Struct, Slice, Map, Function
	employee1 := Employee{
		index:     0,
		nama:      "Dadang",
		alamat:    "Bekasi",
		pekerjaan: "Kuli",
		alasan:    "Alasan dadang",
	}

	employee2 := Employee{
		index:     1,
		nama:      "Ujang",
		alamat:    "Jakarta",
		pekerjaan: "Kuli juga",
		alasan:    "Alasan ujang",
	}

	sliceOfData := []Employee{employee1, employee2}

	if len(os.Args) < 2 {
		fmt.Println("Masukan argumen!")
		return
	}

	args := os.Args[1]

	if index, err := strconv.Atoi(args); err == nil {
		if index >= 0 && index < len(sliceOfData) {
			fmt.Println("Data dari index: ")
			sliceOfData[index].person()
		} else {
			fmt.Println("Index tidak valid")
		}

	} else {
		for _, j := range sliceOfData {
			if args == strings.ToLower(j.nama) {
				fmt.Println("Data dari nama: ")
				j.person()
			}
		}
	}
}
