import React, { useState } from "react";
import { createPost } from "../services/api";

function CreatePost() {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    createPost({ title, body })
      .then(() => {
        alert("Пост успешно создан!");
        setTitle("");
        setBody("");
      })
      .catch(console.error);
  };

  return (
    <div>
      <h2>Создать пост</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Заголовок:</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
        </div>
        <div>
          <label>Текст:</label>
          <textarea
            value={body}
            onChange={(e) => setBody(e.target.value)}
          />
        </div>
        <button type="submit">Создать</button>
      </form>
    </div>
  );
}

export default CreatePost;
