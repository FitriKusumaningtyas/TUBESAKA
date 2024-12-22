package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Barang struct {
	Nama  string
	Harga int
	Stok  int
}

// Binary search iterative
func binarySearchIterative(barangList []Barang, target string) (*Barang, bool) {
	left, right := 0, len(barangList)-1
	target = strings.ToLower(target)
	for left <= right {
		mid := left + (right-left)/2
		if strings.ToLower(barangList[mid].Nama) == target {
			return &barangList[mid], true
		} else if strings.ToLower(barangList[mid].Nama) < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nil, false
}

// Binary search recursive
func binarySearchRecursive(barangList []Barang, target string, left, right int) (*Barang, bool) {
	if left > right {
		return nil, false
	}
	mid := left + (right-left)/2
	target = strings.ToLower(target)
	if strings.ToLower(barangList[mid].Nama) == target {
		return &barangList[mid], true
	} else if strings.ToLower(barangList[mid].Nama) < target {
		return binarySearchRecursive(barangList, target, mid+1, right)
	} else {
		return binarySearchRecursive(barangList, target, left, mid-1)
	}
}

// Measure execution time
func measureTime(fn func([]Barang, string) (*Barang, bool), barangList []Barang, target string, iterations int) float64 {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn(barangList, target)
	}
	return time.Since(start).Seconds()
}

func main() {
	// Daftar barang di toko
	barangList := []Barang{
		{"Minyak", 16000, 50},
		{"Gula", 12000, 30},
		{"Indomie", 4500, 100},
		{"Lee Mineral", 3000, 45},
		{"Beras", 10000, 25},
		{"Tepung", 8000, 20},
		{"Kopi", 15000, 40},
		{"Susu", 12000, 35},
		{"Telur", 23000, 60},
		{"Keju", 45000, 15},
		{"Mentega", 20000, 10},
		{"Teh", 7500, 50},
		{"Kecap", 9000, 25},
		{"Saos", 11000, 30},
		{"Garpu", 5000, 20},
		{"Piring", 10000, 10},
		{"Pisau", 15000, 5},
		{"Sabun Cuci", 6000, 35},
		{"Sikat Gigi", 5000, 25},
		{"Shampoo", 18000, 15},
	}

	// Mengurutkan daftar barang berdasarkan nama
	sort.Slice(barangList, func(i, j int) bool {
		return barangList[i].Nama < barangList[j].Nama
	})

	// Menampilkan daftar barang
	fmt.Println("\n========== Daftar Barang ==========")
	fmt.Println("No\tNama Barang\tHarga\t\tStok")
	for i, barang := range barangList {
		fmt.Printf("%2d. %-15s    Rp %7d        %3d\n", i+1, barang.Nama, barang.Harga, barang.Stok)
	}

	// Meminta input user untuk memilih barang
	var pilihan int
	fmt.Print("\nMasukkan nomor barang yang ingin dicari: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(barangList) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	target := barangList[pilihan-1].Nama
	fmt.Printf("\nMencari barang: %s\n", target)

	// Jumlah iterasi
	iterationsList := []int{10000, 50000, 100000, 150000}

	fmt.Println("\n========== Perbandingan Waktu Eksekusi ==========")
	fmt.Printf("%-12s %-20s %-20s\n", "Iterations", "Iterative Time (s)", "Recursive Time (s)")
	fmt.Println(strings.Repeat("-", 55))
	for _, iterations := range iterationsList {
		iterativeTime := measureTime(binarySearchIterative, barangList, target, iterations)
		recursiveTime := measureTime(func(barangList []Barang, target string) (*Barang, bool) {
			return binarySearchRecursive(barangList, target, 0, len(barangList)-1)
		}, barangList, target, iterations)

		fmt.Printf("%-12d %-20.4f %-20.4f\n", iterations, iterativeTime, recursiveTime)
	}
}
