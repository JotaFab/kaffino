import React, { useState } from "react";
import { Products } from "./Components/Products";
import { Navbar } from "./Components/Navbar";
import { Footer } from "./Components/Footer";
import { Top } from "./Components/Top";
import { Cart } from "./Components/Cart";
import { Login } from "./Components/Login";
import "./index.css";
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { ProductDetail } from "./Components/ProductDetail";

interface CartItem {
  productId: string;
  size: string;
  quantity: number;
  schedule: string; // "one-time", "1 week", "2 weeks", "3 weeks", "4 weeks"
}

export const App: React.FC = () => {
  const [cartItems, setCartItems] = useState<CartItem[]>([]);

  const addToCart = (productId: string, size: string, quantity: number, schedule: string) => {
    setCartItems([...cartItems, { productId, size, quantity, schedule }]);
  };

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<><Navbar /><Top /><Products /><Footer /></>} />
        <Route path="/product/:id" element={<><Navbar /><ProductDetail addToCart={addToCart} /><Footer /> </>} />
        <Route path="/cart" element={<><Navbar /><Cart cartItems={cartItems} /><Footer /></>} />
        <Route path="/products" element={<><Navbar /><Products /><Footer /></>} />
        <Route path="/login" element={<><Navbar /><Login /><Footer /></>} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
