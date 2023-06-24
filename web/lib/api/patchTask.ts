export async function patchTask(table: string, task: Task) {
  const response = await fetch(
    `http://localhost:8080/todo?table=${table}`,
    {
      method: "PATCH",
      body: null,  
    },    
  );

  if (!response.ok) throw new Error("failed to fetch (patch) from api");

  return response.json();
}
  