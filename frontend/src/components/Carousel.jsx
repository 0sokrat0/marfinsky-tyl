// src/components/Carousel.js
import React, { useState, useEffect } from "react";
import "./Carousel.css";

import image1 from "../assets/photo_1.jpg";
import image2 from "../assets/photo_2.jpg";
import image3 from "../assets/photo_3.jpg";
import image4 from "../assets/photo_4.jpg";
import image5 from "../assets/photo_5.jpg";

const images = [image1, image2, image3, image4, image5];

function Carousel() {
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentIndex((prevIndex) => (prevIndex + 1) % images.length);
    }, 8000); // Меняем изображение каждые 8 секунд

    return () => clearInterval(interval);
  }, [images.length]);

  return (
    <div className="carousel-container">
      <div className="carousel">
        {images.map((image, index) => (
          <div
            key={index}
            className={`carousel-slide ${index === currentIndex ? "active" : ""}`}
          >
            <img
              src={image}
              alt={`Slide ${index + 1}`}
              className="carousel-image"
              onError={(e) => {
                console.error(`Ошибка загрузки изображения ${index + 1}:`, e);
                e.target.src = "https://via.placeholder.com/800x600?text=Fallback+Image";
              }}
            />
          </div>
        ))}
      </div>
    </div>
  );
}

export default Carousel;
