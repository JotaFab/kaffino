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
    <div className="my-20">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <h2 className="text-2xl font-extrabold text-gray-900 mb-6">
          Our Products
        </h2>
        <div className="overflow-x-auto hide-scrollbar">
          <div className="flex space-x-6 snap-x snap-mandatory w-full">
            {products.map((product) => (
              <a href={`/product/${product.id}`} key={product.id}>
                <div className="bg-white snap-center shadow-md rounded-lg overflow-hidden w-full min-w-[300px]">
                  <img
                    alt={`Product image with placeholder text '${product.title}'`}
                    className="w-full h-48 object-cover"
                    src={product.images[0]}
                  />
                  <div className="p-4">
                    <h3 className="text-lg font-semibold text-gray-900">
                      {product.title}
                    </h3>
                    <p className="text-gray-600">
                      ${product.map_size_price["Full Bag (12oz)"]}
                    </p>
                    <button className="mt-2 bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600">
                      View Details
                    </button>
                  </div>
                </div>
              </a>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};
