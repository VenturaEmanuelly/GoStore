import { useState } from 'react';
import Routine from './Routine';
import ConsultPrice from './ConsultPrice';
import AddProduct from './AddProduct';

export default function App() {
  const [screen, setScreen] = useState('menu');

  return (
    <div style={{ padding: 20, fontFamily: 'Arial, sans-serif' }}>
      {screen === 'menu' && (
        <>
          <h1>Gestão de Produtos</h1>
          <button onClick={() => setScreen('routine')}>
            Registrar vendas (Bipar códigos)
          </button><br /><br />
          <button onClick={() => setScreen('consult')}>
            Consultar preço por código
          </button><br /><br />
          <button onClick={() => setScreen('add')}>
            Adicionar novo produto
          </button>
        </>
      )}

      {screen === 'routine' && <Routine onBack={() => setScreen('menu')} />}
      {screen === 'consult' && <ConsultPrice onBack={() => setScreen('menu')} />}
      {screen === 'add' && <AddProduct onBack={() => setScreen('menu')} />}
    </div>
  );
}

