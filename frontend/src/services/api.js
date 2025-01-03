const API_URL = "http://localhost:8080/api";

export const fetchPosts = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/posts");
      console.log("RAW response:", response); // Лог для проверки сырых данных
      const result = await response.json();
      console.log("Parsed response:", result); // Лог для проверки распарсенных данных
      return result.data; // Возвращаем только массив постов
    } catch (error) {
      console.error("Ошибка в fetchPosts:", error);
      throw error;
    }
  };
  

// Создание поста
export const createPost = async (post) => {
  const response = await fetch(`${API_URL}/posts`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(post),
  });
  if (!response.ok) {
    throw new Error("Ошибка при создании поста");
  }
  const result = await response.json();
  return result;
};
