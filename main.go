package main

import (
	"fmt"
)

// func main() {
// 	var input string
// 	fmt.Print("Masukkan sesuatu: ")
// 	fmt.Scan(&input)
// 	fmt.Println("Anda memasukkan:", input)
// }

func main() {
	// var input string
	// fmt.Print("Masukkan sesuatu: ")
	// fmt.Scan(&input)
	// fmt.Println("Anda memasukkan:", input)

	// Inisialisasi slice
	slice := []int{1, 2, 3, 4}

	fmt.Println(slice[:2])
	fmt.Println(slice[2+1:])

	// Misalnya kita ingin menghapus elemen di indeks 2
	indexToRemove := 2

	// Menghapus elemen dari slice
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)

	fmt.Println(slice)

}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"runtime"
// )

// // This function clears the terminal screen.
// func ClearScreen() {
// 	if runtime.GOOS == "windows" {
// 		cmd := exec.Command("cmd", "/c", "cls")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	} else {
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	}
// }

// func main() {
// 	// Call the ClearScreen function to clear the terminal screen.
// 	var input int
// 	fmt.Println("Menu")
// 	fmt.Println("1. Home")
// 	fmt.Println("2. Profile")
// 	fmt.Print("Masukkan nomor menu: ")
// 	fmt.Scan(&input)

// 	if input == 1 {
// 		ClearScreen()
// 		fmt.Println("ini halaman home")
// 	} else if input == 2 {
// 		ClearScreen()
// 		fmt.Println("ini halaman profile")
// 	} else {
// 		ClearScreen()
// 		fmt.Println("halaman yang kamu pilih tidak tersedia")
// 	}

// }
