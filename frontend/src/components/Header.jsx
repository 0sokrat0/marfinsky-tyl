import React, { useState } from "react";
import "./Header.css";
import logo from "../assets/logo2.png";

function Header() {
  const [menuOpen, setMenuOpen] = useState(false);

  const toggleMenu = () => {
    setMenuOpen(!menuOpen);
  };

  const scrollToSection = (id) => {
    document.getElementById(id).scrollIntoView({ behavior: "smooth" });
    setMenuOpen(false); // Закрываем меню после перехода
  };

  return (
    <header className="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
      <div className="container">
        <a className="navbar-brand d-flex align-items-center" href="#home">
          <img src={logo} alt="Логотип" className="logo me-2" />
          <span>Марфинский Тыл</span>
        </a>
        <button
          className="navbar-toggler"
          type="button"
          onClick={toggleMenu}
          aria-expanded={menuOpen}
          aria-label="Переключить навигацию"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
        <div
          className={`collapse navbar-collapse ${menuOpen ? "show" : ""}`}
          id="navbarNav"
        >
          <ul className="navbar-nav ms-auto">
            <li className="nav-item">
              <button
                onClick={() => scrollToSection("about")}
                className="nav-link btn btn-link"
              >
                О нас
              </button>
            </li>
            <li className="nav-item">
              <button
                onClick={() => scrollToSection("posts")}
                className="nav-link btn btn-link"
              >
                Цели
              </button>
            </li>
            <li className="nav-item">
              <button
                onClick={() => scrollToSection("work_reports")}
                className="nav-link btn btn-link"
              >
                Отчеты
              </button>
            </li>
            <li className="nav-item">
              <button
                onClick={() => scrollToSection("contact")}
                className="nav-link btn btn-link"
              >
                Контакты
              </button>
            </li>
          </ul>
        </div>
      </div>
    </header>
  );
}

export default Header;
