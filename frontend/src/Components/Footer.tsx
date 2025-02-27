import React from "react";

export const Footer: React.FC = () => {
  return (
    <footer className="bg-white shadow-md mt-16">
      <div className="max-w-auto mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div className="flex justify-between items-center">
          <div className="flex-shrink-0">
            <img
              alt="Company logo with placeholder text 'Logo'"
              className="h-8 w-auto"
              height="40"
              src="/imgs/logo.jpg"
              width="100"
            />
          </div>
          <div className="flex space-x-4">
            <a className="text-gray-700 hover:text-gray-900" href="/products">
              Products
            </a>
            <a className="text-gray-700 hover:text-gray-900" href="/cart">
              Cart
            </a>
            <a className="text-gray-700 hover:text-gray-900" href="/contact-us">
              Contact Us
            </a>
            <a
              className="text-gray-700 hover:text-gray-900"
              href="/subscription"
            >
              Subscription
            </a>
            <a className="text-gray-700 hover:text-gray-900" href="/gifts">
              Gifts
            </a>
          </div>
        </div>
        <div className="mt-4 text-center text-gray-500 text-sm">
          © 2023 Company Name. All rights reserved.
        </div>
      </div>
    </footer>
  );
};
