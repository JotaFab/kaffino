import React, { useState } from "react";
import { Link } from "react-router-dom";
import CartIcon from "../assets/cart.svg";
import UserIcon from "../assets/user.svg";
import MenuIcon from "../assets/menu.svg";
import LogoIcon from "../assets/logobean.svg";

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
                                src={LogoIcon}
                                width="100"
                            />
                        </Link>
                    </div>
                    <div className="text-sepia hidden md:flex md:items-center md:space-x-4">
                        <a
                            className="hover:text-licorice px-3 py-2 rounded-md text-sm font-medium"
                            href="/products"
                        >
                            Products
                        </a>
                        <a
                            className="hover:text-licorice px-3 py-2 rounded-md text-sm font-medium"
                            href="/cart"
                        >
                            Cart
                        </a>
                        <a
                            className="hover:text-licorice px-3 py-2 rounded-md text-sm font-medium"
                            href="/contact-us"
                        >
                            Contact Us
                        </a>
                        <a
                            className="hover:text-licorice px-3 py-2 rounded-md text-sm font-medium"
                            href="/subscription"
                        >
                            Subscription
                        </a>
                        <a
                            className="hover:text-licorice px-3 py-2 rounded-md text-sm font-medium"
                            href="/gifts"
                        >
                            Gifts
                        </a>
                    </div>
                    <div className="flex items-center">
                        <Link to="/cart" className="mr-4">
                            <img src={CartIcon} alt="Cart" className="h-5 w-5" />
                        </Link>
                        <Link to="/login" className="mr-4" >
                            <img src={UserIcon} alt="User" className="h-5 w-5" />
                        </Link>

                        <button
                            className="md:hidden"
                            onClick={toggleMenu}
                        >
                            <img src={MenuIcon} alt="Menu" className="h-5 w-5" />
                        </button>
                    </div>
                </div>
            </div>
            <div
                className={`bg-licorice px-2 shadow-md md:hidden ${isMenuOpen ? "" : "hidden"}`}
                id="menu"
            >
                <div className="text-isabeline px-2 pt-2 pb-3 space-y-1 sm:px-3">
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium hover:text-sepia "
                        href="/products"
                    >
                        Products
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium hover:text-sepia "
                        href="/cart"
                    >
                        Cart
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium hover:text-sepia "
                        href="/contact-us"
                    >
                        Contact Us
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium hover:text-sepia "
                        href="/subscription"
                    >
                        Subscription
                    </a>
                    <a
                        className="block px-3 py-2 rounded-md text-base font-medium hover:text-sepia "
                        href="/gifts"
                    >
                        Gifts
                    </a>
                </div>
            </div>
        </nav>
    );
};
