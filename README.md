# threadStocks - Gestionnaire de Stock de Fils pour Point de Croix / Cross-stitch Thread Inventory Manager

[![Go Version](https://img.shields.io/github/go-mod/go-version/Kae-Tempest/threadStocks)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

---

## 🇫🇷 Français

### 📝 Description
`threadStocks` est une API backend conçue pour aider les passionnés de point de croix à gérer leur inventaire de fils (DMC, Anchor, etc.). Elle permet de suivre les quantités en stock, d'éviter les achats en double et de planifier les besoins pour les futurs projets.

### ✨ Fonctionnalités
- **Authentification sécurisée** : Inscription et connexion basées sur JWT (JSON Web Tokens).
- **Gestion des utilisateurs** : Création et gestion de profils.
- **Gestion de stock (En cours)** : Suivi des références de fils et des quantités.
- **Base de données robuste** : Utilisation de PostgreSQL pour une persistance fiable des données.

### 🛠 Technologies
- **Langage** : [Go (Golang)](https://golang.org/)
- **Base de données** : [PostgreSQL](https://www.postgresql.org/)
- **ORM** : [GORM](https://gorm.io/)
- **Sécurité** : JWT, Bcrypt pour le hachage des mots de passe.

### 🚀 Installation
1. **Cloner le dépôt** :
   ```bash
   git clone https://github.com/yourusername/threadStocks.git
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
- **User Management**: Profile creation and management.
- **Inventory Management (In progress)**: Tracking thread references and quantities.
- **Robust Database**: Using PostgreSQL for reliable data persistence.

### 🛠 Tech Stack
- **Language**: [Go (Golang)](https://golang.org/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM**: [GORM](https://gorm.io/)
- **Security**: JWT, Bcrypt for password hashing.

### 🚀 Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/threadStocks.git
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