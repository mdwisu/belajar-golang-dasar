// slice

package main

import "fmt"

func main()  {
	months := [...]string{1: "Januari", 2: "Februari", 3: "Maret", 4: "April", 5: "Mei", 6: "Juni", 7: "Juli", 8: "Agustus", 9: "September", 10: "Oktober", 11: "November", 12: "Desember"}

	var slice1 = months[4:7]
	fmt.Println(slice1)

	slice2 := months[3:]
	fmt.Println(slice2)

	slice3 := months[:8]
	fmt.Println(slice3)

	var slice4 []string = months[:]
	fmt.Println(slice4)

	// function slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Lagi"
	daySlice1[1] = "Minggu Lagi"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice baru dari awal sekali
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Dwi"
	// newSlice[2] = "Susanto" // error, harusnya menggunakan append
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Susanto")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "budi"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// hati hati membuat slice dan array
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}