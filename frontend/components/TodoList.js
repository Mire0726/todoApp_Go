import React from "react";

const TodoList = ({ todos }) => {
  console.log(todos);
  // todos が未定義または空の場合、メッセージを表示
  if (!todos || todos.length === 0) {
    return <p>No todos found.</p>;
  }

  return (
    <ul>
      {todos.map((todo) => (
        <li key={todo.id}>
          {" "}
          {/* ここで一意のkeyを割り当てる */}
          <strong>ID:</strong><p>{todo.ID}</p><strong>Title:</strong> <p>{todo.Title}</p>
        </li>
      ))}
    </ul>
  );
};

export default TodoList;
