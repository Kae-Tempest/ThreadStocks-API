# threadStocks - Gestionnaire de Stock de Fils pour Point de Croix / Cross-stitch Thread Inventory Manager

[![Go Version](https://img.shields.io/github/go-mod/go-version/Kae-Tempest/threadStocks-API)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

---

## 🇫🇷 Français

### 📝 Description
`threadStocks` est une API backend conçue pour aider les passionnés de point de croix à gérer leur inventaire de fils (DMC, Anchor, etc.). Elle permet de suivre les quantités en stock, d'éviter les achats en double et de planifier les besoins pour les futurs projets.

### ✨ Fonctionnalités
- **Authentification sécurisée** : Inscription, connexion, déconnexion et gestion du mot de passe (oublié/réinitialisation) basées sur JWT (JSON Web Tokens).
- **Gestion des utilisateurs** : Consultation du profil utilisateur connecté.
- **Gestion de stock** : 
    - Création, lecture, mise à jour et suppression (CRUD) de fils.
    - Gestion de masse (suppression multiple).
    - Suivi des références (Marque, ID) et des quantités.
- **Base de données robuste** : Utilisation de PostgreSQL via l'ORM GORM.
- **Observabilité** : Intégration d'OpenTelemetry pour le traçage.

### 🛣️ Routes API
| Méthode | Route | Description | Auth |
|---------|-------|-------------|------|
| POST | `/register` | Inscription d'un nouvel utilisateur | Non |
| POST | `/login` | Connexion et obtention du token JWT | Non |
| POST | `/logout` | Déconnexion | Non |
| POST | `/forgot-password` | Demande de réinitialisation de mot de passe | Non |
| POST | `/reset-password` | Réinitialisation du mot de passe | Non |
| POST | `/contact` | Formulaire de contact | Non |
| GET | `/users/me` | Récupérer les informations de l'utilisateur actuel | Oui |
| PUT | `/users/update-password` | Mettre à jour le mot de passe | Oui |
| GET | `/threads` | Récupérer tous les fils de l'utilisateur | Oui |
| POST | `/threads/create` | Ajouter un nouveau fil au stock | Oui |
| PUT | `/threads/update/{id}` | Mettre à jour un fil spécifique | Oui |
| DELETE | `/threads/delete/{id}` | Supprimer un fil spécifique | Oui |
| DELETE | `/threads/delete` | Suppression multiple de fils | Oui |

### 🛠 Technologies
- **Langage** : [Go (Golang)](https://golang.org/)
- **Base de données** : [PostgreSQL](https://www.postgresql.org/)
- **ORM** : [GORM](https://gorm.io/)
- **Sécurité** : JWT, Bcrypt pour le hachage des mots de passe.
- **Observabilité** : OpenTelemetry.

---

## 🇺🇸 English

### 📝 Description
`threadStocks` is a backend API designed to help cross-stitch enthusiasts manage their thread inventory (DMC, Anchor, etc.). It allows tracking stock quantities, avoiding duplicate purchases, and planning requirements for future projects.

### ✨ Features
- **Secure Authentication**: Registration, login, logout, and password management (forgot/reset) based on JWT (JSON Web Tokens).
- **User Management**: Access current user profile information.
- **Inventory Management**:
    - Full CRUD (Create, Read, Update, Delete) operations for threads.
    - Bulk operations (multiple delete).
    - Track thread references (Brand, ID) and quantities.
- **Robust Database**: Using PostgreSQL with GORM ORM.
- **Observability**: OpenTelemetry integration for tracing.

### 🛣️ API Routes
| Method | Route | Description | Auth |
|--------|-------|-------------|------|
| POST | `/register` | Register a new user | No |
| POST | `/login` | Login and obtain JWT token | No |
| POST | `/logout` | Logout | No |
| POST | `/forgot-password` | Forgot password request | No |
| POST | `/reset-password` | Reset password | No |
| POST | `/contact` | Contact form | No |
| GET | `/users/me` | Get current user information | Yes |
| PUT | `/users/update-password` | Update user password | Yes |
| GET | `/threads` | Get all threads for the user | Yes |
| POST | `/threads/create` | Add a new thread to inventory | Yes |
| PUT | `/threads/update/{id}` | Update a specific thread | Yes |
| DELETE | `/threads/delete/{id}` | Delete a specific thread | Yes |
| DELETE | `/threads/delete` | Bulk delete threads | Yes |

### 🛠 Tech Stack
- **Language**: [Go (Golang)](https://golang.org/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM**: [GORM](https://gorm.io/)
- **Security**: JWT, Bcrypt for password hashing.
- **Observability**: OpenTelemetry.

### 🚀 Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/Kae-Tempest/threadStocks.git
   cd threadStocks/api
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