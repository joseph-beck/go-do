import { Task } from "@/lib/types/task";

export async function postTask(table: string, task: Task) {
  const response = await fetch(
    `http://localhost:8080/task/${table}`,
    {
      method: "PUT",
      body: null, // TODO:
    },
  );

  if (!response.ok) throw new Error("failed to fetch (post) from api");

  return response.json();
}
