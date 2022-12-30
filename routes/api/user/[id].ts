import { PageProps } from "$fresh/server.ts";
import { User } from "../../../types/user.ts";

const users: User[] = [
  { id: "1", name: "John Doe" },
  { id: "2", name: "Jane Doe" },
];

export const handler = (_req: Request, props: PageProps): Response => {
  const { id } = props.params;
  if (!id) {
    return new Response("No ID provided", { status: 400 });
  }

  const user = users.find((user) => user.id === id);
  if (!user) {
    return new Response("User not found", { status: 404 });
  }

  const body = JSON.stringify(user, null, 2);
  return new Response(body, {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
};
