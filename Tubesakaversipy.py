import time
import matplotlib.pyplot as plt

class Barang:
    def _init_(self, nama, harga, stok):
        self.nama = nama
        self.harga = harga
        self.stok = stok

    def _repr_(self):
        return f"{self.nama} - Rp {self.harga} - Stok: {self.stok}"

# Daftar barang di toko
barang_list = [
    Barang("Minyak", 16000, 50),
    Barang("Gula", 12000, 30),
    Barang("Indomie", 4500, 100),
    Barang("Lee Mineral", 3000, 45),
    Barang("Beras", 10000, 25),
    Barang("Tepung", 8000, 20),
    Barang("Kopi", 15000, 40),
    Barang("Susu", 12000, 35),
    Barang("Telur", 23000, 60),
    Barang("Keju", 45000, 15),
    Barang("Mentega", 20000, 10),
    Barang("Teh", 7500, 50),
    Barang("Kecap", 9000, 25),
    Barang("Saos", 11000, 30),
    Barang("Garpu", 5000, 20),
    Barang("Piring", 10000, 10),
    Barang("Pisau", 15000, 5),
    Barang("Sabun Cuci", 6000, 35),
    Barang("Sikat Gigi", 5000, 25),
    Barang("Shampoo", 18000, 15),
]

# Mengurutkan daftar barang berdasarkan nama
barang_list.sort(key=lambda b: b.nama)

def binary_search_iteratif(barang_list, target):
    left, right = 0, len(barang_list) - 1
    target = target.lower()
    while left <= right:
        mid = left + (right - left) // 2
        if barang_list[mid].nama.lower() == target:
            return barang_list[mid], True
        elif barang_list[mid].nama.lower() < target:
            left = mid + 1
        else:
            right = mid - 1
    return None, False

def binary_search_rekursif(barang_list, target, left, right):
    if left > right:
        return None, False
    mid = left + (right - left) // 2
    if barang_list[mid].nama.lower() == target:
        return barang_list[mid], True
    elif barang_list[mid].nama.lower() < target:
        return binary_search_rekursif(barang_list, target, mid + 1, right)
    else:
        return binary_search_rekursif(barang_list, target, left, mid - 1)

# Jumlah pengulangan untuk pengukuran waktu
iterations_list = [1000, 5000, 10000, 15000]

def measure_time(func, *args, iterations):
    start = time.time()
    for _ in range(iterations):
        func(*args)
    return time.time() - start

# Input nama barang yang ingin dicari
target = "gula"

# Mengukur waktu eksekusi untuk binary search iteratif dan rekursif
iterative_times = []
recursive_times = []

print("Barang yang dicari : Gula")
print("| Iterations | Iterative Time (s) | Recursive Time (s) |")
print("|------------|---------------------|--------------------|")

for ulangan in iterations_list:
    iterative_time = measure_time(binary_search_iteratif, barang_list, target, iterations=ulangan)
    recursive_time = measure_time(binary_search_rekursif, barang_list, target, 0, len(barang_list) - 1, iterations=ulangan)

    iterative_times.append(iterative_time)
    recursive_times.append(recursive_time)

    print(f"| {ulangan:<10} | {iterative_time:<19.4f} | {recursive_time:<18.4f} |")

# Membuat grafik perbandingan
plt.clf()  # Membersihkan grafik sebelumnya jika ada
plt.plot(iterations_list, iterative_times, marker='o', linestyle='-', color='orange', label='Iterative')
plt.plot(iterations_list, recursive_times, marker='o', linestyle='-', color='magenta', label='Recursive')

plt.xlabel("Number of Iterations")
plt.ylabel("Execution Time (seconds)")
plt.title("Comparison of Iterative and Recursive Binary Search")
plt.legend()

# Menampilkan nilai pada tiap titik
for i, ulangan in enumerate(iterations_list):
    plt.text(ulangan, iterative_times[i], f"{iterative_times[i]:.4f}s", ha='center', va='bottom', fontsize=10)
    plt.text(ulangan, recursive_times[i], f"{recursive_times[i]:.4f}s", ha='center', va='bottom', fontsize=10)

plt.grid(True)
plt.show()