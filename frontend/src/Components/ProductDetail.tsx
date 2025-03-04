import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { Cart } from "./Cart";

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

interface ProductDetailProps {
  addToCart: (productId: string, size: string, quantity: number, schedule: string) => void;
}

export const ProductDetail: React.FC<ProductDetailProps> = ({ addToCart }) => {
  const { id } = useParams<{ id: string }>();
  const [product, setProduct] = useState<Product | null>(null);
  const [selectedSize, setSelectedSize] = useState<string>("");
  const [quantity, setQuantity] = useState<number>(1);
  const [dynamicPrice, setDynamicPrice] = useState<number>(0);
  const [selectedSchedule, setSelectedSchedule] = useState<string>("one-time");

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        const response = await fetch(`/api/v1/product/${id}`);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        setProduct(data);
      } catch (error) {
        console.error("Could not fetch product:", error);
      }
    };

    if (id) {
      fetchProduct();
    }
  }, [id]);

  useEffect(() => {
    if (product && selectedSize) {
      const sizePrice = product.map_size_price[selectedSize] || 0;
      let price = sizePrice * quantity;
      setDynamicPrice(price);
    }
  }, [selectedSize, quantity, product]);

  const handleSizeChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedSize(event.target.value);
  };

  const handleQuantityChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newQuantity = parseInt(event.target.value, 10);
    setQuantity(newQuantity > 0 ? newQuantity : 1);
  };

  const handleScheduleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedSchedule(event.target.value);
  };

  const handleAddToCartClick = () => {
    if (product && selectedSize) {
      addToCart(product.id, selectedSize, quantity, selectedSchedule);
    } else {
      alert("Please select a size before adding to cart.");
    }
  };

  if (!product) {
    return <div>Loading...</div>;
  }

  return (
    <div className="my-20">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div className="bg-isabeline shadow-md rounded-lg overflow-hidden">
          <div className="md:flex">
            <div className="md:w-1/2">
              <img
                alt={`Product image with placeholder text '${product.title}'`}
                className="w-full h-64 object-cover"
                src={product.images[0]}
              />
            </div>
            <div className="md:w-1/2 p-4">
              <h1 className="text-2xl font-bold text-licorice mb-2">
                {product.title}
              </h1>
              <p className="text-gray-700">{product.long_description}</p>
              <div className="mb-4">
                <label className="block text-sm font-bold mb-2">Purchase Options</label>
                <div className="grid grid-cols-1 gap-4">
                  <label className="flex items-center mb-2 cursor-pointer">
                    <input
                      type="radio"
                      name="purchase-option"
                      value="one-time"
                      checked={selectedSchedule === "one-time"}
                      onChange={() => setSelectedSchedule("one-time")}
                      className="mr-2 leading-tight"
                    />
                    <span className="text-sm">One-time Purchase</span>
                  </label>
                  <label className="flex items-center cursor-pointer">
                    <input
                      type="radio"
                      name="purchase-option"
                      value="1 week"
                      checked={selectedSchedule === "1 week"}
                      onChange={() => setSelectedSchedule("1 week")}
                      className="mr-2 leading-tight"
                    />
                    <span className="text-sm">Subscribe & Get it every 1 week</span>
                  </label>
                  <label className="flex items-center cursor-pointer">
                    <input
                      type="radio"
                      name="purchase-option"
                      value="2 weeks"
                      checked={selectedSchedule === "2 weeks"}
                      onChange={() => setSelectedSchedule("2 weeks")}
                      className="mr-2 leading-tight"
                    />
                    <span className="text-sm">Subscribe & Get it every 2 weeks</span>
                  </label>
                  <label className="flex items-center cursor-pointer">
                    <input
                      type="radio"
                      name="purchase-option"
                      value="3 weeks"
                      checked={selectedSchedule === "3 weeks"}
                      onChange={() => setSelectedSchedule("3 weeks")}
                      className="mr-2 leading-tight"
                    />
                    <span className="text-sm">Subscribe & Get it every 3 weeks</span>
                  </label>
                  <label className="flex items-center cursor-pointer">
                    <input
                      type="radio"
                      name="purchase-option"
                      value="4 weeks"
                      checked={selectedSchedule === "4 weeks"}
                      onChange={() => setSelectedSchedule("4 weeks")}
                      className="mr-2 leading-tight"
                    />
                    <span className="text-sm">Subscribe & Get it every 4 weeks</span>
                  </label>
                </div>
              </div>
              <div className="mb-4">
                {product.sizes && product.sizes.length > 0 && (
                  <>
                    <label className="block text-sm font-bold mb-2" htmlFor="size">
                      Size:
                    </label>
                    <select
                      className="border border-gray-300 p-2 w-full"
                      id="size"
                      name="size"
                      onChange={handleSizeChange}
                      value={selectedSize}
                    >
                      <option value="">Select Size</option>
                      {product.sizes.map((size) => (
                        <option key={size} value={size}>
                          {size} - ${product.map_size_price[size]}
                        </option>
                      ))}
                    </select>
                  </>
                )}
              </div>
              <div className="flex items-center justify-between space-x-4 mb-4">
                <label htmlFor="quantity" className="mr-2">
                  Quantity:
                </label>
                <div className="flex items-center border border-gray-300 p-2">
                  <input
                    className="w-16 text-center border-none focus:outline-none"
                    type="number"
                    id="quantity"
                    name="quantity"
                    value={quantity}
                    min="1"
                    onChange={handleQuantityChange}
                  />
                </div>
                <div className="text-2xl font-bold text-gray-900">
                  ${dynamicPrice.toFixed(2)}
                </div>
                <button className="bg-licorice text-white px-4 py-2 rounded-md hover:bg-sepia" onClick={handleAddToCartClick}>
                  Add to Cart
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProductDetail;