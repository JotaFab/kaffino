import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

export const Login: React.FC = () => {
  const [email, setEmail] = useState("");
  const [emailError, setEmailError] = useState("");
  const [otp, setOtp] = useState("");
  const [otpError, setOtpError] = useState("");
  const [showOtpInput, setShowOtpInput] = useState(false);
  const [loginError, setLoginError] = useState("");
  const navigate = useNavigate();

  const validateEmail = (email: string): boolean => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };

  const handleEmailSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateEmail(email)) {
      setEmailError("Por favor, introduce un correo electrónico válido.");
      return;
    }

    setEmailError("");

    try {
      const response = await fetch("/api/v1/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log("OTP sent:", data);
      setShowOtpInput(true);
      setLoginError("");
      // Show success message
    } catch (error) {
      console.error("Login initiation failed:", error);
      setLoginError("No se pudo iniciar sesión. Por favor, inténtalo de nuevo.");
      setShowOtpInput(false);
      // Show error message
    }
  };

  const handleOtpSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!otp) {
      setOtpError("Por favor, introduce el OTP.");
      return;
    }

    setOtpError("");

    try {
      const response = await fetch("/api/v1/verify-otp", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, otp }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log("Login successful:", data);
      setLoginError("");
      // Redirect to home page
      navigate("/");
    } catch (error) {
      console.error("Login verification failed:", error);
      setLoginError("La verificación de inicio de sesión falló. Por favor, inténtalo de nuevo.");
      // Show error message
    }
  };

  return (
    <div className="flex justify-center items-center h-screen bg-isabeline">
      <div className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <h2 className="text-2xl font-bold mb-6 text-licorice">Acceder</h2>
        <p className="text-gray-700 text-sm italic mb-4">No creemos en las contraseñas. Ingresa tu correo electrónico para recibir una clave unica.</p>
        {loginError && <p className="text-red-500 text-xs italic">{loginError}</p>}
        {!showOtpInput ? (
          <form onSubmit={handleEmailSubmit}>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
                Correo Electrónico
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="email"
                type="email"
                placeholder="Correo Electrónico"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
              {emailError && <p className="text-red-500 text-xs italic">{emailError}</p>}
            </div>
            <div className="flex items-center justify-between">
              <button
                className="bg-licorice hover:bg-sepia text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="submit"
              >
                Enviar OTP
              </button>
            </div>
          </form>
        ) : (
          <form onSubmit={handleOtpSubmit}>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="otp">
                OTP
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="otp"
                type="text"
                placeholder="OTP"
                value={otp}
                onChange={(e) => setOtp(e.target.value)}
              />
              {otpError && <p className="text-red-500 text-xs italic">{otpError}</p>}
              <p className="text-gray-500 text-xs italic mt-1">Este OTP es válido solo por 5 minutos.</p>
            </div>
            <div className="flex items-center justify-between">
              <button
                className="bg-licorice hover:bg-sepia text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="submit"
              >
                Verificar OTP
              </button>
            </div>
          </form>
        )}
      </div>
    </div>
  );
};
