import React from "react";

export const Top: React.FC = () => {
    return (

        <div className="bg-isabeline">
            <div className="container mx-auto flex flex-col md:flex-row items-center my-12 md:my-24">
                <div className="flex flex-col w-full lg:w-1/3 justify-center items-start p-8">
                    <h1 className="text-3xl md:text-5xl p-2 text-licorice tracking-loose">
                        Kafino
                    </h1>
                    <h2 className="text-3xl md:text-5xl leading-relaxed md:leading-snug mb-2">
                        Peruvian Coffee Beans
                    </h2>
                    <p className="text-sm md:text-base text-gray-500 mb-4">
                        Experience the rich and bold flavors of our premium Peruvian coffee beans, sourced directly from the highlands of Peru.
                    </p>
                    <a className="bg-green-500 text-white px-4 py-2 rounded shadow-lg hover:bg-green-400 transition duration-300" href="#products">
                        Shop Now
                    </a>
                </div>
                <div className="p-8 mt-12 md:mt-0">
                    <img alt="A steaming cup of coffee with coffee beans scattered around" className="rounded-lg shadow-lg" height="400" src="https://storage.googleapis.com/a1aa/image/lXpqm6Y4HR54xc7idkIqWA4wTapohRL5y2c7GPQ7k_s.jpg" width="600" />
                </div>
            </div>
        </div>
    )
}
