import { Task } from "@/lib/task/types/task";

export async function getAllTasks(table: string): Promise<Task[]> {
  const response = await fetch(
    `http://localhost:8080/todo?table=${table}`, 
    {
      method: "GET",
    },
  );

  if (!response.ok) throw new Error("failed to fetch (get) from api");

  return response.json();
}

export async function getTask(table: string, id: number) : Promise<Task> {
  const response = await fetch(
    `http://localhost:8080/todo?table=${table}&id=${id}`, 
    {
      method: "GET",
    },
  );

  if (!response.ok) throw new Error("failed to fetch (get) from api");

  return response.json();
}
