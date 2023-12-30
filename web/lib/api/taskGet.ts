import { Task, Tasks } from "@/lib/types/task";

export async function listTasks(table: string): Promise<Tasks> {
  const response = await fetch(
    `http://localhost:8080/api/v1/tasks/${table}`,
    {
      method: "GET",
    },
  );

  if (!response.ok) throw new Error("failed to fetch (get) from api");

  return response.json();
}

export async function getTask(table: string, id: number) : Promise<Task> {
  const response = await fetch(
    `http://localhost:8080/task/${table}/${id}`,
    {
      method: "GET",
    },
  );

  if (!response.ok) throw new Error("failed to fetch (get) from api");

  return response.json();
}
