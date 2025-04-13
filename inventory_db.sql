CREATE USER 'inventory_user'@'localhost' IDENTIFIED BY 'invpass123';
GRANT ALL PRIVILEGES ON inventory_db.* TO 'inventory_user'@'localhost';
FLUSH PRIVILEGES;

-- Membuat tabel produk
CREATE TABLE produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    harga DECIMAL(10,2) NOT NULL,
    kategori VARCHAR(50)
);

-- Membuat tabel inventaris
CREATE TABLE inventaris (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_produk INT,
    jumlah INT NOT NULL,
    lokasi VARCHAR(100),
    FOREIGN KEY (id_produk) REFERENCES produk(id)
);

-- Membuat tabel pesanan
CREATE TABLE pesanan (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_produk INT,
    jumlah INT NOT NULL,
    tanggal_pesanan DATE,
    FOREIGN KEY (id_produk) REFERENCES produk(id)
);

-- Menambahkan data sample ke tabel produk
INSERT INTO produk (nama, deskripsi, harga, kategori) VALUES
('Laptop XYZ', 'Laptop performa tinggi', 15000000, 'Elektronik'),
('Kursi Kantor', 'Kursi ergonomis', 1200000, 'Furniture'),
('Mouse Wireless', 'Mouse tanpa kabel', 250000, 'Aksesoris');

-- Menambahkan data sample ke tabel inventaris
INSERT INTO inventaris (id_produk, jumlah, lokasi) VALUES
(1, 10, 'Gudang A'),
(2, 5, 'Gudang B'),
(3, 25, 'Gudang A');

-- Menambahkan data sample ke tabel pesanan
INSERT INTO pesanan (id_produk, jumlah, tanggal_pesanan) VALUES
(1, 2, '2025-04-10'),
(3, 5, '2025-04-11'),
(1, 1, '2025-04-12');

-- Query untuk menampilkan semua produk
SELECT * FROM produk;

-- Query untuk menampilkan inventaris dengan nama produk
SELECT 
    p.nama AS nama_produk,
    i.jumlah,
    i.lokasi
FROM inventaris i
JOIN produk p ON i.id_produk = p.id;

-- Query untuk menampilkan detail pesanan
SELECT 
    p.nama AS nama_produk,
    ps.jumlah,
    ps.tanggal_pesanan
FROM pesanan ps
JOIN produk p ON ps.id_produk = p.id;

-- Query agregasi: total pesanan per produk
SELECT 
    p.nama AS nama_produk,
    SUM(ps.jumlah) AS total_terjual
FROM pesanan ps
JOIN produk p ON ps.id_produk = p.id
GROUP BY ps.id_produk;

-- Query agregasi: total stok per lokasi
SELECT 
    lokasi,
    SUM(jumlah) AS total_stok
FROM inventaris
GROUP BY lokasi;
