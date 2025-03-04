import React from "react";

export const Footer: React.FC = () => {
  return (
    <footer className="bg-licorice shadow-md mt-16">
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
          <div className="text-isabeline flex space-x-4">
            <a className=" hover:text-sepia" href="/products">
              Products
            </a>
            <a className="hover:text-sepia" href="/cart">
              Cart
            </a>
            <a className="hover:text-sepia" href="/contact-us">
              Contact Us
            </a>
            <a
              className="hover:text-sepia"
              href="/subscription"
            >
              Subscription
            </a>
            <a className="hover:text-sepia" href="/gifts">
              Gifts
            </a>
          </div>
        </div>
        <div className="mt-4 text-center text-white text-sm">
          © 2025 Jaqi Phaxi S.A.C. All rights reserved.
        </div>
      </div>
    </footer>
  );
};
