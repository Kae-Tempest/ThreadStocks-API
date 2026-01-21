# threadStocks - Gestionnaire de Stock de Fils pour Point de Croix / Cross-stitch Thread Inventory Manager

[![Go Version](https://img.shields.io/github/go-mod/go-version/Kae-Tempest/threadStocks)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

---

## 🇫🇷 Français

### 📝 Description
`threadStocks` est une API backend conçue pour aider les passionnés de point de croix à gérer leur inventaire de fils (DMC, Anchor, etc.). Elle permet de suivre les quantités en stock, d'éviter les achats en double et de planifier les besoins pour les futurs projets.

### ✨ Fonctionnalités
- **Authentification sécurisée** : Inscription et connexion basées sur JWT (JSON Web Tokens).
- **Gestion des utilisateurs** : Consultation du profil utilisateur connecté.
- **Gestion de stock** : 
    - Création, lecture, mise à jour et suppression (CRUD) de fils.
    - Gestion de masse (mise à jour et suppression multiple).
    - Suivi des références (ID, Marque) et des quantités.
- **Base de données robuste** : Utilisation de PostgreSQL via l'ORM GORM.

### 🛣️ Routes API
| Méthode | Route | Description | Auth |
|---------|-------|-------------|------|
| POST | `/register` | Inscription d'un nouvel utilisateur | Non |
| POST | `/login` | Connexion et obtention du token JWT | Non |
| GET | `/users/me` | Récupérer les informations de l'utilisateur actuel | Oui |
| POST | `/threads/create` | Ajouter un nouveau fil au stock | Oui |
| GET | `/threads/{id}` | Récupérer les détails d'un fil | Oui |
| POST | `/threads/update/{id}` | Mettre à jour un fil spécifique | Oui |
| POST | `/threads/update` | Mise à jour multiple de fils | Oui |
| DELETE | `/threads/delete/{id}` | Supprimer un fil spécifique | Oui |
| DELETE | `/threads/delete` | Suppression multiple de fils | Oui |

### 🛠 Technologies
- **Langage** : [Go (Golang)](https://golang.org/)
- **Base de données** : [PostgreSQL](https://www.postgresql.org/)
- **ORM** : [GORM](https://gorm.io/)
- **Sécurité** : JWT, Bcrypt pour le hachage des mots de passe.

### 🚀 Installation
1. **Cloner le dépôt** :
   ```bash
   git clone https://github.com/Kae-Tempest/threadStocks.git
   cd threadStocks
   ```
2. **Configurer l'environnement** :
   Copiez le fichier `.env.example` vers `.env` et remplissez les informations de connexion à votre base de données.
   ```bash
   cp .env.example .env
   ```
3. **Lancer l'application** :
   ```bash
   go run main.go
   ```

---

## 🇺🇸 English

### 📝 Description
`threadStocks` is a backend API designed to help cross-stitch enthusiasts manage their thread inventory (DMC, Anchor, etc.). It allows tracking stock quantities, avoiding duplicate purchases, and planning requirements for future projects.

### ✨ Features
- **Secure Authentication**: Registration and login based on JWT (JSON Web Tokens).
- **User Management**: Access current user profile information.
- **Inventory Management**:
    - Full CRUD (Create, Read, Update, Delete) operations for threads.
    - Bulk operations (multiple update and delete).
    - Track thread references (ID, Brand) and quantities.
- **Robust Database**: Using PostgreSQL with GORM ORM.

### 🛣️ API Routes
| Method | Route | Description | Auth |
|--------|-------|-------------|------|
| POST | `/register` | Register a new user | No |
| POST | `/login` | Login and obtain JWT token | No |
| GET | `/users/me` | Get current user information | Yes |
| POST | `/threads/create` | Add a new thread to inventory | Yes |
| GET | `/threads/{id}` | Get details of a specific thread | Yes |
| POST | `/threads/update/{id}` | Update a specific thread | Yes |
| POST | `/threads/update` | Bulk update threads | Yes |
| DELETE | `/threads/delete/{id}` | Delete a specific thread | Yes |
| DELETE | `/threads/delete` | Bulk delete threads | Yes |

### 🛠 Tech Stack
- **Language**: [Go (Golang)](https://golang.org/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM**: [GORM](https://gorm.io/)
- **Security**: JWT, Bcrypt for password hashing.

### 🚀 Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/Kae-Tempest/threadStocks.git
   cd threadStocks
   ```
2. **Set up the environment**:
   Copy the `.env.example` file to `.env` and fill in your database connection details.
   ```bash
   cp .env.example .env
   ```
3. **Run the application**:
   ```bash
   go run main.go
   ```