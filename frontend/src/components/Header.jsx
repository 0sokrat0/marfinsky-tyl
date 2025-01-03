    import React from "react";

    function Header() {
    const scrollToSection = (id) => {
        document.getElementById(id).scrollIntoView({ behavior: "smooth" });
    };

    return (
        <header style={{ position: "fixed", top: 0, width: "100%", background: "#333", color: "#fff", padding: "10px" }}>
        <nav style={{ display: "flex", gap: "15px" }}>
            <button onClick={() => scrollToSection("about")} style={{ color: "#fff", background: "none", border: "none" }}>
            О нас
            </button>
            <button onClick={() => scrollToSection("posts")} style={{ color: "#fff", background: "none", border: "none" }}>
            Цели
            </button>
            <button onClick={() => scrollToSection("work_reports")} style={{ color: "#fff", background: "none", border: "none" }}>
            Отчеты о работе
            </button>
            <button onClick={() => scrollToSection("contact")} style={{ color: "#fff", background: "none", border: "none" }}>
            Контакты
            </button>
        </nav>
        </header>
    );
    }

    export default Header;
