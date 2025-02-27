import React from "react";
import { Products } from "./Components/Products";
import { Navbar } from "./Components/Navbar";
import { Footer } from "./Components/Footer";
import { Top } from "./Components/Top";
import "./index.css";
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { ProductDetail } from "./Components/ProductDetail";

export const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<><Navbar /><Top /><Products /><Footer /></>} />
        <Route path="/product/:id" element={<><Navbar /><ProductDetail /><Footer /> </>} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
