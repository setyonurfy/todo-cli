# Todo CLI

Aplikasi To Do List berbasis terminal yang dibuat dengan Go. Tugas dikelola lewat menu interaktif dan tersimpan otomatis ke file JSON, jadi tidak hilang saat program ditutup.

Project ini dibuat sebagai latihan menggabungkan **struct**, **slice**, dan **interface** dalam satu aplikasi yang utuh.

## Fitur

- Tambah tugas baru
- Lihat daftar tugas beserta statusnya
- Tandai tugas selesai
- Edit judul tugas
- Hapus tugas
- Validasi input: judul kosong ditolak, ID harus berupa angka, ID yang tidak ada ditolak dengan pesan error
- Data tersimpan otomatis ke `tasks.json` dan dimuat kembali saat program dibuka

## Menjalankan

Prasyarat: [Go](https://go.dev/dl/) 1.16 atau lebih baru. Cek dengan `go version`.

```bash
git clone https://github.com/USERNAME/todo-cli.git
cd todo-cli
go run .
```

Untuk membuat program yang bisa dijalankan langsung:

```bash
go build -o todo .
./todo
```

## Unduh program jadi (tanpa install Go)

Tidak punya Go? Unduh program siap pakai dari halaman **[Releases](https://github.com/USERNAME/todo-cli/releases)**, lalu pilih file sesuai sistem operasi:

| Sistem operasi | File |
|---|---|
| Windows | `todo-windows.exe` |
| macOS (Intel) | `todo-mac-intel` |
| macOS (Apple Silicon: M1/M2/M3) | `todo-mac-m1` |
| Linux | `todo-linux` |

Untuk macOS dan Linux, beri izin eksekusi lalu jalankan dari terminal:

```bash
chmod +x todo-mac-m1
./todo-mac-m1
```

Untuk Windows, buka Command Prompt atau PowerShell di folder file tersebut, lalu jalankan `todo-windows.exe`.

> **Catatan:** program ini belum ditandatangani secara digital, jadi Windows SmartScreen atau macOS Gatekeeper dapat menampilkan peringatan keamanan. Itu wajar untuk program hasil build sendiri. Di Windows, klik *More info* lalu *Run anyway*. Di macOS, izinkan lewat pengaturan Privacy & Security. Jika ragu, jalankan dari kode sumber dengan `go run .` seperti di atas.

## Cara pakai

Setelah program berjalan, ketik angka sesuai menu lalu tekan Enter.

| Pilihan | Fungsi | Input yang diminta |
|---|---|---|
| `1` | Tambah | Judul tugas |
| `2` | Lihat | - |
| `3` | Selesai | ID tugas |
| `4` | Hapus | ID tugas |
| `5` | Edit | ID tugas, lalu judul baru |
| `6` | Keluar | - |

Contoh tampilan:

```
=== TO DO LIST ===
Ketikan Angka!
1.Tambah 2.Lihat 3.Selesai 4.Hapus 5.Edit 6.Keluar
Pilih: 1
Judul: Belajar interface Go
Task ditambahkan

Pilih: 2
[1] [ ] Belajar interface Go

Pilih: 3
ID: 1
Task ditandai selesai

Pilih: 2
[1] [x] Belajar interface Go
```

## Struktur kode

Seluruh kode ada di satu file, `main.go`, dengan pembagian peran berikut:

| Bagian | Peran |
|---|---|
| `Task` (struct) | Data satu tugas: `ID`, `Title`, `Done` |
| `TaskManager` (interface) | Kontrak operasi: `Add`, `List`, `Complete`, `Delete`, `Update` |
| `TodoList` (struct) | Implementasi `TaskManager`; menyimpan tugas dalam slice `[]Task` |
| `NewTodoList` | Membuat `TodoList` dan memuat data dari file JSON |
| `save()` | Menulis seluruh daftar tugas ke file JSON |
| `main()` | Menu interaktif dan pembacaan input |

Menu di `main()` hanya berbicara lewat interface `TaskManager`, bukan langsung ke `TodoList`. Dengan begitu, cara penyimpanan bisa diganti (misalnya ke database) tanpa mengubah menu.

## Konsep Go yang dipelajari

- **Struct** untuk memodelkan data
- **Slice**: `append`, menghapus elemen dengan `append(s[:i], s[i+1:]...)`, dan mengubah elemen lewat index
- **Interface** sebagai kontrak antara menu dan penyimpanan
- **Pointer receiver** agar method dapat mengubah data aslinya
- **Error handling** dengan pola `if err := ...; err != nil`
- **`encoding/json`** untuk menyimpan dan memuat data
- **`bufio.Scanner`** untuk membaca input dari terminal

## Format penyimpanan

Data disimpan di `tasks.json` pada folder tempat program dijalankan:

```json
[
  {
    "ID": 1,
    "Title": "Belajar interface Go",
    "Done": true
  },
  {
    "ID": 2,
    "Title": "Latihan slice",
    "Done": false
  }
]
```

ID tugas baru dihitung dari ID terbesar yang ada ditambah satu. Jika file belum ada atau masih kosong, program memulai dengan daftar kosong.

## Rencana pengembangan

- [ ] Unit test untuk `Add`, `Update`, `Complete`, dan `Delete`
- [ ] Memecah kode ke beberapa file
- [ ] Penyimpanan ke database lewat implementasi baru dari `TaskManager`

## Penulis

Dibuat oleh NAMA ANDA sebagai bagian dari proses belajar Go.