import { useState } from 'react';
import axios from 'axios';

export default function Routine({ onBack }) {
  const [code, setCode] = useState('');
  const [items, setItems] = useState([]);
  const [total, setTotal] = useState(0);
  const [isFinalized, setIsFinalized] = useState(false);

  const handleAdd = async () => {
    if (!code) return;
    try {
      const response = await axios.get(`http://127.0.0.1:8080/product?code=${code}`);
      const product = response.data;
      setItems([...items, product]);
      setTotal((prev) => prev + product.price);
      setCode('');
    } catch (err) {
      alert("Produto não encontrado");
    }
  };

  const handleFinalize = () => {
    setIsFinalized(true);
  };

  const handlePrint = () => {
    window.print();
  };

  return (
    <div>
      {!isFinalized && (
        <>
          <h2>Registrar Vendas</h2>
          <input
            type="text"
            placeholder="Digite o código do produto"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter') handleAdd();
            }}
          />
          <button onClick={handleAdd}>Adicionar</button>
          <button onClick={onBack}>Voltar ao Menu</button>
        </>
      )}

      {items.length > 0 && (
        <>
          {!isFinalized && <h3>Itens do Pedido</h3>}
          <table>
            <thead>
              <tr>
                <th>Código</th>
                <th>Produto</th>
                <th>Preço (R$)</th>
              </tr>
            </thead>
            <tbody>
              {items.map((item, i) => (
                <tr key={i}>
                  <td>{item.code}</td>
                  <td>{item.name}</td>
                  <td>R$ {item.price.toFixed(2)}</td>
                </tr>
              ))}
            </tbody>
          </table>
          <p><strong>Total:</strong> R$ {total.toFixed(2)}</p>
        </>
      )}

      {!isFinalized && items.length > 0 && (
        <button onClick={handleFinalize}>Finalizar Venda</button>
      )}

      {isFinalized && (
        <div style={{ marginTop: "1rem" }}>
          <button onClick={onBack}>Voltar ao Menu</button>
          <button onClick={handlePrint} style={{ marginLeft: '10px' }}>Imprimir</button>
        </div>
      )}
    </div>
  );
}
