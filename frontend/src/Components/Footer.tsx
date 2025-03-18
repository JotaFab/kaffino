import React from "react";
import LogoIcon from "../assets/logobean.svg";

export const Footer: React.FC = () => {
  return (
    <footer className="bg-bistre shadow-md mt-16">
      <div className="max-w-auto mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div className="flex justify-between items-center">
          <div className="flex-shrink-0">
            <img
              alt="Company logo with placeholder text 'Logo'"
              className="h-8 w-auto"
              height="40"
              src={LogoIcon}
              width="100"
            />
          </div>
          <div className="text-isabeline flex space-x-4">
            <a href="/products">
              Products
            </a>
            <a href="/cart">
              Cart
            </a>
            <a href="/contact-us">
              Contact Us
            </a>
            <a href="/subscription">
              Subscription
            </a>
            <a href="/gifts">
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
