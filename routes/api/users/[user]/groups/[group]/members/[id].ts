import { PageProps } from "$fresh/server.ts";
import { Member } from "../../../../../../../types/member.ts";

const members: Member[] = [
  { userId: "1", groupId: "1" },
  { userId: "2", groupId: "1" },
];

export const handler = (_req: Request, props: PageProps): Response => {
  const { group, id } = props.params;
  if (!id || !group) {
    return new Response("No member ID provided", { status: 400 });
  }

  const member = members.find((member) =>
    member.userId === id && member.groupId === group
  );
  if (!member) {
    return new Response("Member not found", { status: 404 });
  }

  const body = JSON.stringify(member, null, 2);
  return new Response(body, {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
};
