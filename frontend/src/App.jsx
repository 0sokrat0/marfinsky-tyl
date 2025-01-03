import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Header from "./components/Header";
import Home from "./pages/Home";
import About from "./pages/About";
import Posts from "./pages/Posts";
import CreatePost from "./pages/CreatePost";

function App() {
  console.log("App rendered"); // Лог для проверки работы
  return  (
    <Router>
      <Header />
      <div className="container">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/posts" element={<Posts />} />
          <Route path="/create" element={<CreatePost />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
