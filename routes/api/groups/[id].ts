import { PageProps } from "$fresh/server.ts";
import { Group } from "../../../types/group.ts";

const groups: Group[] = [
  { id: "1", name: "Group 1" },
  { id: "2", name: "Group 2" },
];

export const handler = (_req: Request, props: PageProps): Response => {
  const { id } = props.params;
  if (!id) {
    return new Response("No ID provided", { status: 400 });
  }

  const group = groups.find((group) => group.id === id);
  if (!group) {
    return new Response("Group not found", { status: 404 });
  }

  const body = JSON.stringify(group, null, 2);
  return new Response(body, {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
};
