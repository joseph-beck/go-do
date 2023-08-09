export async function deleteTask(table: string, id: number) {
  const response = await fetch(
    `http://localhost:8080/task/${table}/${id}`,
    {
      method: "DELETE",
      body: null,
    },
  );

  if (!response.ok) throw new Error("failed to fetch (delete) from api");

  return response.json();
}
