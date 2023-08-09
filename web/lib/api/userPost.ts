import { User } from "@/lib/types/user";

export async function postUser(user: User) {
  const response = await fetch(
    `http://localhost:8080/user`,
    {
      method: "POST",
      body: null, // TODO: fix
    },
  );

  if (!response.ok) throw new Error("failed to fetch (post) from api");

  return response.json();
}
