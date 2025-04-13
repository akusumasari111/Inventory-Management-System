# Inventory-Management-System
Deskripsi Assignment:
Dalam tugas ini, student akan membangun sistem backend untuk mengelola aplikasi manajemen inventaris. Proyek ini akan melibatkan perancangan dan query pada database relasional untuk mengelola produk, inventaris, dan pesanan. Student akan membuat dan mengintegrasikan RESTful API untuk berinteraksi dengan database, memungkinkan operasi CRUD untuk data inventaris menggunakan Golang.

Selain itu, student akan mengimplementasikan penanganan file untuk mengunggah dan mengunduh gambar produk, serta menyusun route menggunakan framework Gin untuk menyajikan API secara efektif. Dengan menyelesaikan tugas ini, student akan mendapatkan pengalaman langsung dalam pengembangan backend, integrasi API, dan penanganan penyimpanan file, serta menerapkan konsep pemrograman inti dalam skenario dunia nyata.

Detail Assignment
1. Database Design and Queries
   - Desain sebuah database relasional dengan setidaknya tabel-tabel berikut:
     (1) Produk: Menyimpan detail produk (ID, nama, deskripsi, harga, dan kategori);
     (2) Inventaris: Melacak tingkat stok dan lokasi produk (ID produk, jumlah, dan lokasi);
     (3) Pesanan: Mencatat pesanan pelanggan (ID pesanan, ID produk, jumlah, dan tanggal pesanan).
   - Write SQL scripts to:
     (1) Memasukkan data sampel ke dalam tabel-tabel;
     (2) Melakukan query untuk produk, inventaris, dan detail pesanan;
     (3) Melakukan agregasi seperti total pesanan untuk suatu produk atau tingkat stok di lokasi tertentu.

3. RESTful API Development
   Buat API untuk hal-hal berikut:
   - Produk:
     (1) Menambahkan, melihat, memperbarui, dan menghapus produk;
     (2) Melihat detail produk berdasarkan ID atau kategori.
   - Inventaris:
     (1) Melihat tingkat stok untuk suatu produk;
     (2) Memperbarui tingkat stok (menambah atau mengurangi stok).
   - Pesanan
     (1) Membuat pesanan baru;
     (2) Mengambil detail pesanan berdasarkan ID.

5. API Integration
   - Uji API untuk memastikan integrasi dan fungsionalitas berjalan lancar.
   - Tulis skrip atau gunakan alat (misalnya, Postman, Curl) untuk menguji semua endpoint.
 
7. Web Server and Routing
   Gunakan framework Gin di Golang untuk:
   - Menyusun route untuk semua endpoint API (misalnya, /products, /inventory, /orders).
   - Menangani metode HTTP seperti GET, POST, PUT, dan DELETE
