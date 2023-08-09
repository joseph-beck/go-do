import { Task } from "@/lib/types/task";

export async function patchTask(table: string, task: Task) {
  const response = await fetch(
    `http://localhost:8080/task/${table}`,
    {
      method: "PATCH",
      body: null, // TODO: fix
    },
  );

  if (!response.ok) throw new Error("failed to fetch (patch) from api");

  return response.json();
}
