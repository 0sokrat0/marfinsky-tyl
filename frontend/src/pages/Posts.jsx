import React, { useEffect, useState } from "react";
import { fetchPosts } from "../services/api";

function Posts() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    fetchPosts().then(setPosts).catch(console.error);
  }, []);

  return (
    <div>
      <h2>Список постов</h2>
      {posts.length > 0 ? (
        <ul>
          {posts.map((post) => (
            <li key={post.id}>
              <h3>{post.title}</h3>
              <p>{post.body}</p>
            </li>
          ))}
        </ul>
      ) : (
        <p>Постов пока нет.</p>
      )}
    </div>
  );
}

export default Posts;
