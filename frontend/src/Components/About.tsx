export const AboutUs: React.FC = () => {
    return (

        <section className="bg-gray-100 py-8" id="about">
            <div className="container mx-auto px-6">
                <h2 className="text-3xl font-semibold text-center text-gray-800 mb-8">
                    About Us
                </h2>
                <div className="flex flex-col md:flex-row items-center">
                    <div className="md:w-1/2">
                        <img alt="A picturesque view of Peruvian coffee plantations" className="rounded-lg shadow-lg" height="300" src="https://storage.googleapis.com/a1aa/image/bVL5dFprM_bR0plkbF0qnf6WhhgXqLDzuIoKzMgqT5U.jpg" width="500" />
                    </div>
                    <div className="md:w-1/2 md:pl-8 mt-8 md:mt-0">
                        <p className="text-gray-600 mb-4">
                            Kafino is dedicated to bringing you the finest coffee beans from the heart of Peru. Our beans are carefully selected and roasted to perfection, ensuring a rich and flavorful cup every time.
                        </p>
                        <p className="text-gray-600 mb-4">
                            We work directly with local farmers to source our beans, ensuring fair trade practices and supporting the local community. Our commitment to quality and sustainability is at the core of everything we do.
                        </p>
                        <p className="text-gray-600">
                            Join us on a journey to discover the unique flavors of Peruvian coffee. Whether you're a coffee connoisseur or just starting your coffee journey, Kafino has something for everyone.
                        </p>
                    </div>
                </div>
            </div>
        </section>
    )
}