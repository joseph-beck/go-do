export default async function getAllTasks(): Promise<Task[]> {
  const response = await fetch("http://localhost:8080/todo", {
    method: "GET",
  });

  if (!response.ok) throw new Error("failed to fetch from api");

  return response.json();
}
