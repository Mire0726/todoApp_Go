import React, { useEffect, useState } from "react";
import TodoList from "../components/TodoList";
import TodoForm from "../components/TodoForm";

const IndexPage = () => {
  const [todos, setTodos] = useState([]);

  const fetchTodos = async () => {
    try {
      const response = await fetch("http://localhost:8080/todos");
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      setTodos(data);
    } catch (error) {
      console.error("Failed to fetch todos:", error);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const handleNewTodo = () => {
    fetchTodos(); // ToDoリストを再取得
  };

  return (
    <div>
      <h1>ToDo List</h1>
      <TodoForm onNewTodo={handleNewTodo} />
      <TodoList todos={todos} />
    </div>
  );
};

export default IndexPage;
