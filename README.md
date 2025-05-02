#  Mobile Store Backend

This is the **backend** of the Mobile Store App, built with **Go (Gin)** and **MySQL**.

It provides REST APIs for user authentication, product management, and role-based access for Admin and Customer roles.

---

##  About This Repo

This repository contains **only the backend** code of the full-stack Mobile Store application.

- Frontend is built separately using **React.js**
- Backend handles data storage, APIs, user roles, and business logic

---

##  Features

- User registration and login
- JWT-based authentication
- Role-based authorization (Admin / Customer)
- CRUD APIs for products (Admin only)
- Secure password handling

---

##  Tech Stack

- Go 1.20+
- Gin web framework
- MySQL 8+
- GORM (ORM)
- JWT for authentication
- RESTful API design

---

##  Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/your-username/mobile-store-backend.git
cd mobile-store-backend
