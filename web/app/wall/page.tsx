//"use client";

import React from "react";

type Todo = {
  id: number;
  name: string;
  description: string;
  complete: boolean;
  deadline: string;
};

export default async function Page() {
  const response = await fetch("http://localhost:8080/todo?table=tasks", {
    method: "GET",
  });
  const todos: Todo[] = await response.json();
  const todoMap = todos.map((todo: Todo, index: number) => <li key={`list item ${index}`}>{todo.name} : {todo.description}</li>);

  return (
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <h1>Hello</h1>

        <ul key="list">
          {todoMap}
        </ul>
      </main>
  );
}