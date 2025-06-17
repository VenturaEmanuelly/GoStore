import { useState } from "react";
import axios from "axios";

export default function ConsultPrice({ onBack }) {
  const [code, setCode] = useState("");
  const [product, setProduct] = useState(null);
  const [error, setError] = useState(null);

  async function handleConsult() {
    setError(null);
    setProduct(null);

    if (!code.trim()) {
      setError("Informe o código do produto.");
      return;
    }

    try {
      const response = await axios.get("http://localhost:8080/product?code=" + encodeURIComponent(code));
      setProduct(response.data);
    } catch (err) {
      setError(err.response?.data?.error || "Erro ao buscar produto.");
    }
  }

  return (
    <div>
      <h2>Consultar preço pelo código</h2>
      <input
        type="text"
        placeholder="Código do produto"
        value={code}
        onChange={(e) => setCode(e.target.value)}
        style={inputStyle}
      />
      <button onClick={handleConsult} style={btnStyle}>
        Consultar
      </button>
      <button onClick={onBack} style={{ ...btnStyle, backgroundColor: "#aaa", marginLeft: 10 }}>
        Voltar
      </button>

      {product && (
        <div style={{ marginTop: 20 }}>
          <p>
            <strong>Nome:</strong> {product.name}
          </p>
          <p>
            <strong>Preço:</strong> R$ {parseFloat(product.price).toFixed(2)}
          </p>
        </div>
      )}
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}

const inputStyle = {
  width: "calc(100% - 22px)",
  padding: "10px",
  marginBottom: "10px",
  fontSize: "16px",
};

const btnStyle = {
  padding: "10px 20px",
  fontSize: "16px",
  cursor: "pointer",
};
