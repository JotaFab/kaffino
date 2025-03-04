import React, { useState } from "react";
import { Link } from "react-router-dom";
import CartIcon from "../assets/cart.svg";
import UserIcon from "../assets/user.svg";
import MenuIcon from "../assets/menu.svg";

export const Navbar: React.FC = () => {
    const [isMenuOpen, setIsMenuOpen] = useState(false);

    const toggleMenu = () => {
        setIsMenuOpen(!isMenuOpen);
    };

    return (
        <nav className="bg-isabeline fixed w-full z-30 top-0">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between h-16">
                    <div className="flex-shrink-0 flex items-center">
                        <Link to="/">
                            <img
                                alt="Company logo with placeholder text 'Logo'"
                                className="h-8 w-auto"
                                height="40"
                                src="https://storage.googleapis.com/a1aa/image/kslKXCnxdVEwxwUEiHe8nG2gvJGXnl7lF8y7OYBzec0.jpg"
                                width="100"
                            />
                        </Link>
                    </div>
                    <div className="text-licorice hidden md:flex md:items-center md:space-x-4">
                        <a
                            className="hover:text-sepia px-3 py-2 rounded-md text-sm font-medium"
                            href="/products"
                        >
                            Products
                        </a>
                        <a
                            className="hover:text-sepia px-3 py-2 rounded-md text-sm font-medium"
                            href="/cart"
                        >
                            Cart
                        </a>
                        <a
                            className="hover:text-sepia px-3 py-2 rounded-md text-sm font-medium"
                            href="/contact-us"
                        >
                            Contact Us
                        </a>
                        <a
                            className="hover:text-sepia px-3 py-2 rounded-md text-sm font-medium"
                            href="/subscription"
                        >
                            Subscription
                        </a>
                        <a
                            className="hover:text-sepia px-3 py-2 rounded-md text-sm font-medium"
                            href="/gifts"
                        >
                            Gifts
                        </a>
                    </div>
                    <div className="flex items-center">
                        <Link to="/cart" className="mr-4">
                            <img src={CartIcon} alt="Cart" className="fill-licorice hover:fill-sepia h-5 w-5" />
                        </Link>
                        <Link to="/login" className="mr-4" >
                            <img src={UserIcon} alt="User" className="h-5 w-5" />
                        </Link>

                        <button
                            className="text-gray-700 hover:text-gray-900 focus:outline-none focus:text-gray-900 md:hidden"
                            onClick={toggleMenu}
                        >
                            <img src={MenuIcon} alt="Menu" className="h-5 w-5" />
                        </button>
                    </div>
                </div>
            </div>
            <div
                className={`bg-isabeline shadow-md md:hidden ${isMenuOpen ? "" : "hidden"}`}
                id="menu"
            >
                <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                        href="/products"
                    >
                        Products
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                        href="/cart"
                    >
                        Cart
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                        href="/contact-us"
                    >
                        Contact Us
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                        href="/subscription"
                    >
                        Subscription
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                        href="/gifts"
                    >
                        Gifts
                    </a>
                </div>
            </div>
        </nav>
    );
};
