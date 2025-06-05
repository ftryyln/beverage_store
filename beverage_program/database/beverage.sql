-- Membuat database utama
CREATE DATABASE beverage_store;

-- Menggunakan database yang baru dibuat
USE beverage_store;

-- ======================================================================
-- DDL
-- ======================================================================

-- Tabel untuk menyimpan informasi user (admin & customer)
CREATE TABLE
  Users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,                               -- ID unik user
    email VARCHAR(255) NOT NULL UNIQUE,                                   -- Email harus unik dan tidak boleh kosong
    password VARCHAR(255) NOT NULL,                                       -- Password user (disimpan sebagai string terenkripsi)
    role ENUM ('admin', 'customer') NOT NULL                              -- Role user hanya bisa 'admin' atau 'customer'
  );

-- Tabel untuk menyimpan detail tambahan user
CREATE TABLE
  UserDetails (
    user_detail_id INT PRIMARY KEY AUTO_INCREMENT,                        -- ID unik untuk detail user
    user_id INT UNIQUE,                                                   -- Hubungan one-to-one ke Users
    name VARCHAR(255),                                                    -- Nama lengkap user
    phone_number VARCHAR(20),                                             -- Nomor HP user
    FOREIGN KEY (user_id) REFERENCES Users (user_id) ON DELETE CASCADE    -- Jika user dihapus, detailnya juga terhapus
  );

-- Tabel kategori produk (misalnya: kopi, teh, soda)
CREATE TABLE
  Categories (
    category_id INT PRIMARY KEY AUTO_INCREMENT,                           -- ID unik kategori
    name VARCHAR(255)                                                     -- Nama kategori
  );

-- Tabel produk minuman
CREATE TABLE
  Products (
    product_id INT PRIMARY KEY AUTO_INCREMENT,                            -- ID unik produk
    created_by INT,                                                       -- ID user (admin) yang menambahkan produk
    product_name VARCHAR(255) NOT NULL,                                   -- Nama produk
    price DECIMAL(10, 2) CHECK (price >= 0),                              -- Harga produk (tidak boleh negatif)
    stock INT CHECK (stock >= 0),                                         -- Jumlah stok tersedia (tidak boleh negatif)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,              -- Waktu produk ditambahkan
    FOREIGN KEY (created_by) REFERENCES Users (user_id)                   -- Admin yang menambahkan produk
  );

-- Tabel relasi many-to-many antara produk dan kategori
CREATE TABLE
  productCategories (
    product_id INT NOT NULL,                                              -- ID produk
    category_id INT NOT NULL,                                             -- ID kategori
    PRIMARY KEY (product_id, category_id),                                -- Kombinasi unik produk dan kategori
    FOREIGN KEY (product_id) REFERENCES Products (product_id),
    FOREIGN KEY (category_id) REFERENCES Categories (category_id)
  );

-- Tabel order / pesanan dari user
CREATE TABLE
  Orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,                              -- ID unik order
    user_id INT,                                                          -- ID user yang melakukan pemesanan
    total_amount DECIMAL(10, 2) CHECK (total_amount >= 0),                -- Total harga pesanan
    order_date DATE NOT NULL,                                             -- Tanggal order dilakukan
    FOREIGN KEY (user_id) REFERENCES Users (user_id) ON DELETE CASCADE    -- Jika user dihapus, order juga ikut dihapus
  );

-- Tabel detail item yang dibeli dalam setiap order
CREATE TABLE
  OrderItems (
    order_item_id INT PRIMARY KEY AUTO_INCREMENT,                         -- ID unik untuk tiap item pesanan
    order_id INT,                                                         -- Order terkait
    product_id INT,                                                       -- Produk yang dipesan
    quantity INT NOT NULL CHECK (quantity >= 0),                          -- Jumlah produk yang dipesan
    subtotal DECIMAL(10, 2) CHECK (subtotal >= 0),                        -- Harga total produk * jumlah
    FOREIGN KEY (order_id) REFERENCES Orders (order_id),                  -- Foreign key ke Orders
    FOREIGN KEY (product_id) REFERENCES Products (product_id)             -- Foreign key ke Products
  );

-- ======================================================================
-- DML
-- ======================================================================

-- Insert sample users (admin & customer)
INSERT INTO Users (email, password, role) VALUES
('admin@mail.com', '123123', 'admin'),
('customer1@mail.com', '111222', 'customer'),
('customer2@mail.com', '222333', 'customer');

-- Insert user details
INSERT INTO UserDetails (user_id, name, phone_number) VALUES
(1, 'Admin Beverage', '081234567890'),
(2, 'Customer One', '082345678901'),
(3, 'Customer Two', '083456789012');

-- Insert categories
INSERT INTO Categories (name) VALUES
('Kopi'),
('Teh'),
('Soda');

-- Insert products (created_by = admin user_id = 1)
INSERT INTO Products (created_by, product_name, price, stock) VALUES
(1, 'Kopi Hitam', 15000.00, 50),
(1, 'Teh Manis', 10000.00, 40),
(1, 'Soda Gembira', 12000.00, 30);

-- Link products to categories (many-to-many)
INSERT INTO productCategories (product_id, category_id) VALUES
(1, 1),  -- Kopi Hitam ke Kategori Kopi
(2, 2),  -- Teh Manis ke Kategori Teh
(3, 3);  -- Soda Gembira ke Kategori Soda

-- Insert orders by customers
INSERT INTO Orders (user_id, total_amount, order_date) VALUES
(2, 45000.00, '2025-06-01'),
(3, 15000.00, '2025-06-02');

-- Insert order items for each order
INSERT INTO OrderItems (order_id, product_id, quantity, subtotal) VALUES
(1, 1, 2, 30000.00),  -- Customer1 beli 2 Kopi Hitam
(1, 2, 1, 10000.00),  -- Customer1 beli 1 Teh Manis
(2, 1, 1, 15000.00);  -- Customer2 beli 1 Kopi Hitam
