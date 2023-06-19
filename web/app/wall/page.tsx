"use client";

import { useState } from "react";

type Todo = {
  id: number;
  name: string;
  description: string;
  complete: boolean;
  deadline: string;
};

export default function Wall() {
  const [todos, setTodos] = useState([]);
 
  const fetchTodos = async () => {
    const response = await fetch("http://localhost:8080/todo?table=tasks&id=1", {
      method: "GET",
    });
    const data = await response.json();
    setTodos(data);
  };

  return (
    <div>
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <h1>
          Hello
        </h1>

        <button onClick={fetchTodos}>Get todos</button>
        <ul>
            {todos.map((todo: Todo) => {
              return (
                <li
                  style={{ color: `${todo.complete ? "green" : "red"}` }}
                  key={todo.id}
                >
                  {todo.name}:{todo.complete}.
                </li>
              );
            })};
          </ul>
      </main>
    </div>
  );
}