import { User } from "@/lib/user/types/user";

export async function patchUser(user: User) {
  const response = await fetch(
    `http://localhost:8080/user`,
    {
      method: "PATCH",
      body: null,  // TODO: fix
    },    
  );

  if (!response.ok) throw new Error("failed to fetch (post) from api");

  return response.json();
}