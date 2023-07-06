import { User } from "@/lib/user/types/user";

export async function getUser(id: number) : Promise<User> {
  const response = await fetch(
    `http://localhost:8080/user?id=${id}`, 
    {
      method: "GET",
    },
  );
  
  if (!response.ok) throw new Error("failed to fetch (get) from api");
  
  return response.json();
}