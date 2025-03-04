import React from "react";

interface CartItem {
  productId: string;
  size: string;
  quantity: number;
  schedule: string;
}

interface CartProps {
  cartItems: CartItem[];
}

export const Cart: React.FC<CartProps> = ({ cartItems }) => {

  const handleCreateOrder = async () => {
    const orderData = cartItems.map(item => ({
      productId: item.productId,
      size: item.size,
      quantity: item.quantity,
      schedule: item.schedule
    }));

    try {
      const response = await fetch('/api/v1/order', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(orderData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log('Order created:', data);
      // Clear the cart after successful order creation
    } catch (error) {
      console.error("Could not create order:", error);
    }
  };

  return (
    <div className="my-20 bg-isabeline">
      <h2 className="text-licorice">Shopping Cart</h2>
      {cartItems.length === 0 ? (
        <p>Your cart is empty.</p>
      ) : (
        <ul>
          {cartItems.map((item, index) => (
            <li key={index}>
              Product ID: {item.productId}, Size: {item.size}, Quantity: {item.quantity}, Schedule: {item.schedule}
            </li>
          ))}
        </ul>
      )}
      <button className="bg-licorice text-white px-4 py-2 rounded-md hover:bg-sepia" onClick={handleCreateOrder} disabled={cartItems.length === 0}>
        Create Order
      </button>
    </div>
  );
};
