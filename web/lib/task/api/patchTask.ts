import { Task } from "@/lib/task/types/task";

export async function patchTask(table: string, task: Task) {
  const response = await fetch(
    `http://localhost:8080/todo?table=${table}`,
    {
      method: "PATCH",
      body: null, // TODO: fix
    },    
  );

  if (!response.ok) throw new Error("failed to fetch (patch) from api");

  return response.json();
}
  