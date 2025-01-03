import React from "react";
import Header from "./components/Header";
import Carousel from "./components/Carousel";
import "./App.css";

function App() {
  return (
    <div className="bg-dark text-white min-vh-100">
      <Header />
      <div className="photo">
        <Carousel />
      </div>
      <div id="about" className="section container text-center py-5">
        <h1>О нас</h1>
        <p>Здесь вы узнаете о нашей компании или проекте.</p>
      </div>
      <div id="posts" className="section container text-center py-5">
        <h1>Цели</h1>
        <p>Это список наших постов.</p>
      </div>
      <div id="work_reports" className="section container text-center py-5">
        <h1>Отчеты о работе</h1>
        <p>Это список наших отчетов о работе.</p>
      </div>
      <div id="contact" className="section container text-center py-5">
        <h1>Контакты</h1>
        <p>Контактная информация.</p>
      </div>
    </div>
  );
}

export default App;
