# ERD Title: Beverage Store Database

---

## 1. Entities and Attributes

### A. Entity: Users
- **Attributes**:
- `user_id` (PK, AI) INT  
- `email` VARCHAR(255) NOT NULL UNIQUE  
- `password` VARCHAR(255) NOT NULL  
- `role` ENUM('admin', 'customer') NOT NULL  

### B. Entity: UserDetails
- **Attributes**:
- `user_detail_id` (PK, AI) INT  
- `user_id` (FK, UNIQUE) INT  
- `name` VARCHAR(255)  
- `phone_number` VARCHAR(20) 

### C. Entity: Categories
- **Attributes**:
- `category_id` (PK, AI) INT  
- `name` VARCHAR(255)  

### D. Entity: Products
- **Attributes**:
- `product_id` (PK, AI) INT  
- `product_name` VARCHAR(255) NOT NULL  
- `price` DECIMAL(10,2) CHECK (>= 0)  
- `stock` INT CHECK (>= 0)  
- `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP  

### E. Entity: ProductCategories
- **Attributes**:
- `product_id` (PK, AI) INT  
- `category_id` (PK, AI) INT  

### F. Entity: Orders
- **Attributes**:
- `order_id` (PK, AI) INT  
- `user_id` (FK) INT  
- `total_amount` DECIMAL(10,2) CHECK (>= 0)  
- `order_date` DATE NOT NULL  

### G. Entity: OrderItems
- **Attributes**:
- `order_item_id` (PK, AI) INT  
- `order_id` (FK) INT  
- `product_id` (FK) INT  
- `quantity` INT CHECK (>= 0)  
- `subtotal` DECIMAL(10,2) CHECK (>= 0)  

---

## 2. Relationships

### A. **Users to UserDetails**
- Type: One-to-One
- Description: Each user has exactly one associated detail record. The user_id in UserDetails references the primary key in Users and is unique to enforce the one-to-one relationship.

### B. **Users to Products **
- Type: One to Many
- Description: Each admin user can create multiple products. The created_by column in the Products table stores the user_id of the admin who created the product.

### C. **Users to Orders**
- Type: One to Many
- Description: A single user (customer) can place multiple orders. The user_id in the Orders table links each order to the corresponding user.

### D. **Orders to OrderItems**
- Type: One to Many
- Description: Each order can contain multiple order items. This is represented by order_id in the OrderItems table.

### E. **Products to OrderItems**
- Type: One to Many
- Description: EA product can appear in many order items (i.e., be ordered multiple times in different orders). This is represented by product_id in the OrderItems table.

### F. **Products to Categories (via ProductCategories)**
- Type: Many to Many
- Description: A product can belong to multiple categories, and a category can include multiple products. This many-to-many relationship is implemented through the productCategories join table.

### G. **Categories to Products (via ProductCategories)**
- Type: Many to Many
- Description: Same as above, the productCategories table acts as the bridge between products and categories.

---

## 3. Integrity Constraints

- Primary keys are auto-incremented and unique  
- Foreign keys reference valid records  
- `email` is unique and NOT NULL  
- `quantity`, `price`, `stock`, `subtotal` must be ≥ 0  
- `order_date`, `created_at` must be NOT NULL  
- `user_id` in `UserDetails` is UNIQUE for One-to-One relationship  

---

## 4. CLI Features Implemented

### ✅ Authentication
- Register
- Login

### ✅ Beverage Management
- List Beverages (User/Admin)
- Update Beverages (Admin)
- Delete Beverages (Admin)

### ✅ Category Management
- Add Categories (Admin)

### ✅ Transactions
- Buy Beverages (User)

### ✅ Reports
- User Report (User with most orders)
- Beverage Report (Most ordered beverage)
- Category Report (Category with most orders)

---

## 5. Reports Description

- **User Report**: Top users by number of orders
- **Beverage Report**: Most ordered beverage items
- **Category Report**: Categories with the highest product sales

---

## 6. Project Notes

- CLI application structured in folders: `cli`, `config`, `entity`, `handler`, `database`, `test`
- `.env` used for environment configuration (DB connection)
- SQL dump stored in `/database` folder
- Structured and scalable for future business analysis
