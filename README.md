# GoStore

Sistema simples de cadastro, consulta e registro de vendas de produtos.  
Feito com **Go + Fiber** no backend e **React** no frontend.

---

## 🚀 Como rodar o projeto

### 1. Requisitos

- Go 1.20 ou superior
- Node.js 18+ (para o frontend)
- PostgreSQL rodando localmente (ou em container)

---

### 2. Clonando o projeto

```bash
git clone https://seurepo.com/store-manager.git
cd store-manager
```

---

### 3. Backend (Go + Fiber)

📦 **Instalar dependências**
```bash
go mod tidy
```

🧾 **Configurar banco de dados**

Você precisa de um banco **PostgreSQL** com a seguinte tabela:

```sql
CREATE TABLE products (
  code TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  price NUMERIC(10,2) NOT NULL
);
```

▶️ **Rodar o backend**
```bash
go run main.go
```

Por padrão, ele roda em: [http://localhost:8080](http://localhost:8080)

---

### 4. Frontend (React)

```bash
cd frontend
npm install
npm run dev
```

A aplicação React estará disponível em: [http://localhost:5173](http://localhost:5173)

---

## 📌 Funcionalidades

✅ Cadastrar produtos  
🔍 Consultar produto por código  
🔄 Atualizar ou deletar produto  
🧾 Registrar vendas e gerar total

---

## 🔌 Endpoints da API

### ➕ Criar Produto

**POST /product**

```json
{
  "code": "123",
  "name": "Caneta Azul",
  "price": 2.50
}
```

---

### 🔍 Consultar Produto

**GET /product?code=123**

**Resposta:**

```json
{
  "code": "123",
  "name": "Caneta Azul",
  "price": 2.50
}
```

---

### ✏️ Atualizar Produto

**PUT /product**

```json
{
  "code": "123",
  "name": "Caneta Vermelha",
  "price": 3.00
}
```

---

### ❌ Deletar Produto

**DELETE /product/{code}**

---

## 🖥️ Telas do sistema (frontend)

- Menu Principal com 4 opções:
  - Registrar Vendas
  - Consultar Preço
  - Adicionar Produto
  - Atualizar / Deletar Produto

---

## 🛠 Tecnologias usadas

### Backend
- Go 1.20+
- Fiber (Web Framework)
- PostgreSQL

### Frontend
- React
- Vite
- Axios

---

## 👩‍💻 Autora

Desenvolvido por **Emanuelly Ventura**