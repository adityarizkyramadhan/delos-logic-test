package main

import (
	"fmt"
	"os"

	soaldua "github.com/adityarizkyramadhan/delos-logic-test/soal_dua"
	soalsatu "github.com/adityarizkyramadhan/delos-logic-test/soal_satu"
	soaltiga "github.com/adityarizkyramadhan/delos-logic-test/soal_tiga"
)

func main() {
	for {
		var input int8
		fmt.Printf("%s\n%s\n%s\n%s\n", "1. Library Fine Calculator", "2. Sour Candy", "3. Sum Sub Array", "4. Exit")
		fmt.Print("Pilih soal yang dijalankan (1-4) :")
		fmt.Scan(&input)
		switch input {
		case 1:
			fmt.Println("===== Fine Calculator =====")
			var d1, m1, y1, d2, m2, y2 int
			fmt.Print("d1 : ")
			fmt.Scan(&d1)
			fmt.Print("m1 : ")
			fmt.Scan(&m1)
			fmt.Print("y1 : ")
			fmt.Scan(&y1)
			fmt.Print("d2 : ")
			fmt.Scan(&d2)
			fmt.Print("m2 : ")
			fmt.Scan(&m2)
			fmt.Print("y2 : ")
			fmt.Scan(&y2)
			fmt.Printf("Result : %d\n\n", soalsatu.FineCalculator(d1, m1, y1, d2, m2, y2))
			fmt.Println("======= Selesai =======")
		case 2:
			fmt.Println("===== Sour Candy =====")
			var student, candies, firstStudent int
			fmt.Print("Jumlah siswa : ")
			fmt.Scan(&student)
			fmt.Print("Jumlah permen : ")
			fmt.Scan(&candies)
			fmt.Print("Dimulai dari : ")
			fmt.Scan(&firstStudent)
			fmt.Printf("Result : %d\n\n", soaldua.SourCandyDecision(student, candies, firstStudent))
			fmt.Println("======= Selesai =======")
		case 3:
			fmt.Println("===== Sum Sub Array =====")
			var dataLen int
			var arr []int
			fmt.Print("Panjang array : ")
			fmt.Scan(&dataLen)
			for i := 0; i < dataLen; i++ {
				var temp int
				fmt.Printf("Masukkan data ke-%d : ", i+1)
				fmt.Scan(&temp)
				arr = append(arr, temp)
			}
			fmt.Printf("Result : %s\n\n", soaltiga.SumOfElements(arr))
			fmt.Println("======= Selesai =======")
		case 4:
			fmt.Println("========= Exit =========")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
