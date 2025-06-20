import { useState } from "react";
import axios from "axios";

export default function AddProduct({ onBack }) {
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [price, setPrice] = useState("");
  const [message, setMessage] = useState(null);
  const [error, setError] = useState(null);

  async function handleAdd() {
    setMessage(null);
    setError(null);

    if (!code.trim() || !name.trim() || !price.trim()) {
      setError("Todos os campos são obrigatórios.");
      return;
    }

    if (isNaN(parseFloat(price))) {
      setError("Preço deve ser um número válido.");
      return;
    }

    try {
      const product = { code, name, price: parseFloat(price) };

      await axios.post("http://127.0.0.1:8080/product", product);

      setMessage("Produto adicionado com sucesso!");
      setCode("");
      setName("");
      setPrice("");
    } catch (err) {
      setError(err.response?.data?.error || "Erro ao adicionar produto.");
    }
  }

  return (
    <div>
      <h2>Adicionar novo produto</h2>
      <input
        type="text"
        placeholder="Código"
        value={code}
        onChange={(e) => setCode(e.target.value)}
        style={inputStyle}
      />
      <input
        type="text"
        placeholder="Nome"
        value={name}
        onChange={(e) => setName(e.target.value)}
        style={inputStyle}
      />
      <input
        type="text"
        placeholder="Preço"
        value={price}
        onChange={(e) => setPrice(e.target.value)}
        style={inputStyle}
      />
      <button onClick={handleAdd} style={btnStyle}>
        Adicionar Produto
      </button>
      <button onClick={onBack} style={{ ...btnStyle, backgroundColor: "#aaa", marginLeft: 10 }}>
        Voltar
      </button>

      {message && <p style={{ color: "green" }}>{message}</p>}
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}

const inputStyle = {
  display: "block",
  width: "100%",
  padding: "8px",
  marginBottom: "10px",
  fontSize: "16px",
};

const btnStyle = {
  padding: "10px 20px",
  fontSize: "16px",
  cursor: "pointer",
};
