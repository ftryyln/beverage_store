[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-2e0aaae1b6195c2367325f4f02e2d04e9abb55f0b24a779b69b11b9e10269abc.svg)](https://classroom.github.com/online_ide?assignment_repo_id=19700766&assignment_repo_type=AssignmentRepo)

# P1-Pair-Project

# 🧃 Beverage Store CLI App by BLACKMARKET Team

A command-line interface application for managing beverage store operations, including product inventory, customer orders, and reporting. Built with Go and MySQL.

---

## 📚 Features

### ✅ Authentication
- [x] Register
- [x] Login

### 🛍️ Beverage Management
- [x] List Beverages (User/Admin)
- [x] Update Beverages (Admin only)
- [x] Delete Beverages (Admin only)

### 🏷️ Category Management
- [x] Add Categories (Admin only)

### 🧾 Ordering System
- [x] Buy Beverages (User only)

### 📊 Reports
- [x] User Report: User with most orders
- [x] Beverage Report: Most ordered beverage
- [x] Category Report: Category with most orders

---

## 🧩 ERD Overview

Database includes the following entities:

- `Products`
- `Categories`
- `ProductCategories` (join table)
- `Users`
- `UserDetails`
- `Orders`
- `OrderItems`

![ERD](docs/ERD.png)

---

## 🗂️ Project Structure

```
/cli          → CLI app entry point & role-based menus
/config       → DB connection config
/database     → SQL (DDL and DML)
/docs         → Documentation, ERD diagrams
/entity       → Go structs mapped to DB tables
/handler      → Business logic handlers
/test         → Unit tests for core features
```

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone git@github.com:H8-FTGO-AOH-CLASSROOM-ALL-PHASE/p1-pair-project-beverage-store.git
cd beverage_program
```

### 2. Set up the database

- Import the SQL schema & sample data:

```bash
mysql -u root -p < database/beverage.sql
```

### 3. Configure `.env`

Create a `.env` file:

```env
<!-- LOCAL -->
DB_USER = root
DB_PASS = 
DB_HOST = 127.0.0.1
DB_PORT = 3306
DB_NAME = beverage_store

<!-- RAILWAY -->
mysql://root:jDWftbdHsppQTZXPEsTcLGGUMzbGMtXa@switchback.proxy.rlwy.net:12577/railway

DB_USER = root
DB_PASS = jDWftbdHsppQTZXPEsTcLGGUMzbGMtXa
DB_HOST = switchback.proxy.rlwy.net
DB_PORT = 12577
DB_NAME = railway

```

### 4. Run the application

```bash
go run main.go
```

---

## 🧪 Testing

Each feature includes at least:

- ✅ 1 success case
- ❌ 1 failure case

Run tests with:

```bash
go test ./handler       #run all test, with simple output
go test -v ./handler    #run all test, display every function test (vorbose log)
```

---

## 📦 Dependencies

- Go ≥ 1.20
- MySQL 8+
- [godotenv](https://github.com/joho/godotenv) for environment config

---

## 📄 Deliverables

- [x] ERD in `.png`
- [x] SQL dump (DDL + sample data)
- [x] CLI application with role-based access
- [x] Feature documentation & testing

---

## 📌 Notes

- Built for scalability and future analytical reporting
- Designed with clean database normalization and role-specific CLI UX

---

## 🧑‍💻 Author

Made with ❤️ by [Fitry Yuliani] and [Fahreza Alghifary] – Hacktiv8 Fulltime Golang Program

---