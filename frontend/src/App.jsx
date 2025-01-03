import React from "react";
import Header from "./components/Header";
import "./App.css";

function App() {
  return (
    <div>
      <Header />
      <div id="home" className="section">
        <h1>Домашняя страница</h1>
        <p>Добро пожаловать на наш сайт!</p>
      </div>
      <div id="about" className="section">
        <h1>О нас</h1>
        <p>Здесь вы узнаете о нашей компании или проекте.</p>
      </div>
      <div id="posts" className="section">
        <h1>Цели</h1>
        <p>Это список наших постов.</p>
      </div>
      <div id="work_reports" className="section">
        <h1>Отчеты о работе</h1>
        <p>Это список наших отчетов о работе.</p>
      </div>
      <div id="contact" className="section">
        <h1>Контакты</h1>
        <p>Контактная информация.</p>
      </div>
      
    </div>
  );
}

export default App;
