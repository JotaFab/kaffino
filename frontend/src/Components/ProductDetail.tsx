import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";

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

export const ProductDetail: React.FC = () => {
  console.log(document.URL)
  const { id } = useParams<{ id: string }>();
  const [product, setProduct] = useState<Product | null>(null);

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        const response = await fetch(`/api/v1/product/${id}`);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        console.log(data)
        setProduct(data);
      } catch (error) {
        console.error("Could not fetch product:", error);
      }
    };

    if (id) {
      fetchProduct();
    }
  }, [id]);

  if (!product) {
    return <div>Loading...</div>;
  }

  return (
    <main className="mt-20">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div className="bg-white shadow-md rounded-lg overflow-hidden">
          <img
            alt={`Product image with placeholder text '${product.title}'`}
            className="w-full h-64 object-cover"
            src={product.images[0]}
          />
          <div className="p-4">
            <h1 className="text-2xl font-bold text-gray-900 mb-2">
              {product.title}
            </h1>
            <p className="text-gray-700">{product.long_description}</p>
            <p className="text-gray-600 mt-4">
              Price: ${product.map_size_price["Full Bag (12oz)"]}
            </p>
            <button className="mt-4 bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600">
              Add to Cart
            </button>
          </div>
        </div>
      </div>
    </main>
  );
};
