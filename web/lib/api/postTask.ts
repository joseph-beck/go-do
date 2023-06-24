export async function postTask(table: string, task: Task) {
  const response = await fetch(
    `http://localhost:8080/todo?table=${table}`,
    {
      method: "POST",
      body: null,  
    },    
  );

  if (!response.ok) throw new Error("failed to fetch (post) from api");

  return response.json();
}
