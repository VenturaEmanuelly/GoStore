import { useState } from "react";
import axios from "axios";

export default function EditProduct({ onBack }) {
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [price, setPrice] = useState("");

  const handleUpdate = async () => {
    if (!code.trim()) {
      alert("Informe o código do produto para atualizar.");
      return;
    }

    try {
      await axios.put("http://127.0.0.1:8080/product", {
        code,
        name,
        price: parseFloat(price),
      });

      alert("Produto atualizado com sucesso!");
      setCode("");
      setName("");
      setPrice("");
    } catch (err) {
      alert("Erro ao atualizar produto.");
      console.error(err);
    }
  };

  const handleDelete = async () => {
    if (!code.trim()) {
      alert("Informe o código do produto para deletar.");
      return;
    }

    try {
      
      const response = await axios.delete(`http://127.0.0.1:8080/product/${encodeURIComponent(code)}`);
      
      alert(response.data.message || "Produto deletado com sucesso!");
      setCode("");
      setName("");
      setPrice("");
    } catch (err) {
      alert("Erro ao deletar produto.");
      console.error(err);
    }
  };

  return (
    <div>
      <h2>Editar ou Deletar Produto</h2>
      <input
        type="text"
        placeholder="Código do produto"
        value={code}
        onChange={(e) => setCode(e.target.value)}
        style={inputStyle}
      />
      <input
        type="text"
        placeholder="Novo nome"
        value={name}
        onChange={(e) => setName(e.target.value)}
        style={inputStyle}
      />
      <input
        type="number"
        placeholder="Novo preço"
        value={price}
        onChange={(e) => setPrice(e.target.value)}
        style={inputStyle}
      />
      <button onClick={handleUpdate} style={btnStyle}>
        Atualizar Produto
      </button>
      <button onClick={handleDelete} style={{ ...btnStyle, backgroundColor: "#d9534f", marginLeft: 10 }}>
        Deletar Produto
      </button>
      <br />
      <button onClick={onBack} style={{ ...btnStyle, backgroundColor: "#aaa", marginTop: 15 }}>
        Voltar ao Menu
      </button>
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
