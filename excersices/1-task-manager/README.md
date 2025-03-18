# Tugas OOP pada Golang

## Deskripsi Tugas

Buatlah sebuah aplikasi manajemen tugas berbasis OOP di Golang. Setiap tugas memiliki logika eksekusi sendiri, dan dikelola oleh sebuah manager yang bertanggung jawab untuk menjalankan, memonitor, serta mencatat (logging) eksekusi tugas-tugas tersebut.

## Spesifikasi Tugas

### 1. Definisikan Interface `Task`

- Buat interface `Task` dengan minimal method:
  - `Execute() error` – Method untuk mengeksekusi tugas.
  - `Info() string` – Method untuk mengembalikan informasi tentang tugas (misalnya, nama tipe tugas, deskripsi, dan waktu pembuatan).

### 2. Buat Struct Dasar untuk Tugas

- Buatlah struct dasar (misalnya, `BaseTask`) yang menyimpan field umum seperti:
  - `ID` (string)
  - `CreatedAt` (time.Time)
  - `Description` (string)
- Gunakan **struct embedding** agar tipe tugas lainnya dapat mewarisi field dasar ini.

### 3. Implementasikan Beberapa Jenis Tugas 

Buat setidaknya tiga tipe tugas berbeda yang mengimplementasikan interface `Task` (buatlah waktu eksekusi random), misalnya:
- **EmailTask** – Mensimulasikan pengiriman email.
- **SMSTask** – Mensimulasikan pengiriman SMS.
- **ReportTask** – Mensimulasikan pembuatan laporan.

Setiap tugas harus mengimplementasikan method `Execute()` dengan logika yang berbeda (misalnya, menampilkan output yang berbeda atau mensimulasikan durasi eksekusi yang berbeda menggunakan `time.Sleep`).

### 4. Manager untuk Menangani Tugas

- Buat sebuah struct `TaskManager` yang memiliki field untuk menyimpan daftar tugas (misalnya, slice `[]Task`).
- `TaskManager` harus menyediakan method untuk:
  - Menambahkan tugas baru.
  - Mengeksekusi semua tugas (gunakan for loop).
  - Memonitor status eksekusi (misalnya, mencatat waktu mulai dan selesai eksekusi setiap tugas).

### 5. Penggunaan Reflection

- Di dalam `TaskManager` , gunakan package `reflect` untuk mengambil informasi dinamis mengenai tipe tugas yang sedang dieksekusi (misalnya, mendapatkan nama tipe dari objek tugas) dan tampilkan bersama log eksekusi.


## Alur Program (Contoh)

1. Program utama membuat instance `TaskManager`.
2. Program utama menambahkan berbagai tipe tugas (misalnya, beberapa `EmailTask`, `SMSTask`, dan `ReportTask`) ke dalam manager.
3. `TaskManager` menjalankan semua tugas sekaligus.
4. Selama eksekusi, program mencetak log yang memuat informasi tipe tugas (diperoleh menggunakan refleksi), ID, deskripsi, dan durasi eksekusi.
5. Program selesai ketika semua tugas telah dieksekusi.

#### Poin Bonus jika bisa mensimulasikan task yang failed, dan buat fungsi untuk kembali menjalankan task yang failed