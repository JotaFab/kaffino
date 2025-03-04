import React, { useState, useEffect } from "react";

interface Product {
  id: string;
  code: string;
  images: string[];
  discount: number;
  title: string;
  description: string;
  long_description: string;
  reviews: string[];
  map_size_price: { [key: string]: number };
  shedules: string[];
  tags: string[];
  created_at: string;
  updated_at: string;
  stock_quantity: number;
  sizes: string[];
}

export const Products: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await fetch("/api/v1/products");
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        setProducts(data);
      } catch (error) {
        console.error("Could not fetch products:", error);
      }
    };

    fetchProducts();
  }, []);

  return (
    <div className="bg-isabeline py-12">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <h2 className="text-3xl font-bold text-licorice text-center mb-8">
          Our Products
        </h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {products.map((product) => (
            <div
              key={product.id}
              className="bg-white rounded-lg overflow-hidden shadow-md"
            >
              <img
                alt={`Product image with placeholder text '${product.title}'`}
                className="w-full h-64 object-cover"
                src={product.images[0]}
              />
              <div className="p-4">
                <h3 className="text-xl font-semibold text-gray-900 mb-2">
                  {product.title}
                </h3>
                <p className="text-gray-700">{product.description}</p>
                <div className="mt-4 flex justify-between items-center">
                  <span className="text-gray-600">${product.map_size_price["Full Bag (12oz)"]}</span>
                  <button className="bg-licorice hover:bg-sepia text-white font-bold py-2 px-4 rounded">
                    Add to Cart
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
