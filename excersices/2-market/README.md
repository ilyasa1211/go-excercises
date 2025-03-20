# Tugas Simulasi Pembayaran Belanja di Toko Menggunakan Go

## Latar Belakang
Kamu adalah pemilik toko yang ingin mensimulasikan alur pembayaran belanja oleh pelanggan dengan skenario nyata. Sistem yang akan dibuat harus mencakup beberapa aspek, mulai dari pelayanan pelanggan oleh kasir secara paralel, penerapan diskon, dan penanganan transaksi yang gagal.

## Deskripsi Tugas
Buatlah sebuah program menggunakan bahasa **Go** yang mensimulasikan proses transaksi di toko dengan alur sebagai berikut:

1. **Pelayanan Kasir**
   - Terdapat beberapa kasir (maksimal 5 orang) yang melayani pelanggan.
   - Jumlah kasir yang aktif dapat diatur.

2. **Antrean Pelanggan**
   - Pelanggan datang dengan nomor antrian.
   - Minimal terdapat 100 pelanggan yang di-generate secara acak, dengan nama dihasilkan menggunakan [gofakeit](https://github.com/brianvoe/gofakeit).
   - Setiap pelanggan akan mendapatkan sejumlah barang secara acak (minimal 2 item)
   - Masing-masing pelanggan memiliki waktu proses yang dihasilkan secara random.

3. **Proses Transaksi**
   - **Setiap pelanggan akan mendapatkan sejumlah barang secara acak (minimal 2 item)**.
   - **Struktur Barang:** Buatlah struct dengan field `Name` dan `Price`, dan `Stock` untuk menggambarkan barang yang dijual di toko.
   - **Penerapan Diskon:**
     -  Tambahkan diskon pada barang jika pelanggan membeli item lebih dari yang kalian tentukan. 
   - **Timeout Transaksi:** 
     - Gunakan `context` untuk membatasi waktu transaksi oleh kasir.
     - Jika pelayanan pelanggan memakan waktu terlalu lama, maka transaksi akan dibatalkan.

4. **Pencatatan Transaksi**
   - Setelah transaksi selesai, kasir akan mencetak struk pesanan yang telah diproses.
   - Semua transaksi yang berhasil akan tercatat sebagai pemasukan toko.
  
5. **Stok Barang**
   - Jika stok barang habis maka pelangan tidak akan mendapatkan barang.
   - Jangan sampai pelanggan yang tidak membawa barang ke kasir.


## Hint Implementasi
- **Gunakan Prinsip OOP**
- **Jumlah Kasir:** Maksimal 5 orang (Configurable).
- **Jumlah Pelanggan:** Minimal 100 orang (gunakan [gofakeit](https://github.com/brianvoe/gofakeit) untuk generate nama pelanggan) (Configurable).
- **Sinkronisasi:** Gunakan `sync.WaitGroup` untuk menunggu semua kasir selesai.
- **Struktur Data:** Buatlah struct untuk barang dengan field `Name` dan `Price` dan `Stock`. 
- **Assign Barang:** Assign barang yang ada di toko ke pelanggan secara acak (seorang pelanggan bisa membawa 2 item atau lebih).
- **Buat mekanisme untuk running transsaksi yang gagal secara asynchronous.**


Selamat mengerjakan tugas dan semoga berhasil!