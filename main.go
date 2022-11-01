package main

import "fmt"

func rentangAngka(a, b int32) {
	for i := a; i <= b; i++ {
		fmt.Println(i)
	}
}

func ganjil(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

func main() {
	//06.md- Rentang angka
	fmt.Println("06.md- Rentang Angka")
	rentangAngka(4, 8)

	//07.md- angka ganjil 1-100
	fmt.Println("Angka Ganjil 1-100")
	ganjil(100)

	//10.md - array
	fmt.Println("Tambah item ARRAY")
	var stuff []string = []string{"Meja", "Buku", "Topi", "Baju", "Kayu"}
	//add handuk ke awal
	stuff = append([]string{"Handuk"}, stuff...)
	//add Celana ke akhir
	stuff = append(stuff, "Celana")
	fmt.Println(stuff)

	//11.md - Hapus item Array
	fmt.Println("Tambah item ARRAY")
	buah := []string{"Jeruk", "Pepaya", "Jambu", "Anggur", "Melon"}
	//hapus jambu
	buah = append(buah[:2], buah[3:]...)
	fmt.Println(buah)


	//14.md - Hapus item Array
	fmt.Println("Filter MAP Array")
	orang := []map[string]interface{}{
		{"id": 1, "name": "Udin", "age": 12},
		{"id": 2, "name": "Wati", "age": 51},
		{"id": 3, "name": "Budi", "age": 34},
		{"id": 4, "name": "Agus", "age": 16},
		{"id": 5, "name": "Sari", "age": 19},
		{"id": 6, "name": "Ririn", "age": 21},
	}
	for _, v := range orang {
		if v["age"].(int) < 21 {
			fmt.Printf("id: %v, name: %v, Age: %v \n", v["id"], v["name"], v["age"])
		}
	}
}
