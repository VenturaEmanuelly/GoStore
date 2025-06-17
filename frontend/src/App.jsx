import { useState } from "react";
import Routine from "./Routine";
import ConsultPrice from "./ConsultPrice";
import AddProduct from "./AddProduct";
import EditProduct from "./EditProduct";

export default function App() {
  const [screen, setScreen] = useState("menu");

  return (
    <div style={{ padding: 20, fontFamily: "Arial, sans-serif", maxWidth: 600, margin: "auto" }}>
      {screen === "menu" && (
        <>
          <h1>Gestão de Produtos</h1>

          <button onClick={() => setScreen("routine")} style={btnStyle}>
            Registrar vendas (Bipar códigos)
          </button>

          <button onClick={() => setScreen("consult")} style={btnStyle}>
            Consultar preço por código
          </button>

          <button onClick={() => setScreen("add")} style={btnStyle}>
            Adicionar novo produto
          </button>

          <button onClick={() => setScreen("updateDelete")} style={btnStyle}>
            Atualizar / Deletar produto
          </button>
        </>
      )}

      {screen === "routine" && <Routine onBack={() => setScreen("menu")} />}
      {screen === "consult" && <ConsultPrice onBack={() => setScreen("menu")} />}
      {screen === "add" && <AddProduct onBack={() => setScreen("menu")} />}
      {screen === "updateDelete" && <EditProduct onBack={() => setScreen("menu")} />}
    </div>
  );
}

const btnStyle = {
  display: "block",
  width: "100%",
  padding: "10px",
  marginBottom: "15px",
  fontSize: "16px",
  cursor: "pointer",
};


