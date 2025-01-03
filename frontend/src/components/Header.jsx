import React from "react";
import { Link } from "react-router-dom";
import "./Header.css";

function Header() {
  return (
    <header className="header">
      <div className="container">
        <h1>Мой Сайт</h1>
        <nav>
          <ul className="nav-links">
            <li>
              <Link to="/">Домашняя</Link>
            </li>
            <li>
              <Link to="/about">О нас</Link>
            </li>
            <li>
              <Link to="/posts">Посты</Link>
            </li>
            <li>
              <Link to="/create">Создать пост</Link>
            </li>
          </ul>
        </nav>
      </div>
    </header>
  );
}

export default Header;
