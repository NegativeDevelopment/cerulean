import { Handlers, PageProps } from "$fresh/server.ts";
import { GroupCard } from "../../components/cards/GroupCard.tsx";
import BaseLayout from "../../components/layouts/BaseLayout.tsx";
import { Group } from "../../types/group.ts";

interface Data {
  groups: Group[];
}

export const handler: Handlers<Data> = {
  async GET(req, ctx) {
    const url = new URL(req.url);
    const response = await fetch(`${url.origin}/api/users/1/groups/1`);
    const group: Group = await response.json().catch(console.error);
    const groups = [group];
    if (!groups) {
      return new Response("Groups not found", { status: 404 });
    }

    return ctx.render({ groups });
  },
};

export default function Groups({ data }: PageProps<Data>) {
  const { groups } = data;

  return (
    <BaseLayout>
      <h1 class="mb-4 text-2xl tracking-wider uppercase font-light">
        Your Groups
      </h1>
      <ul class="flex flex-col gap-4">
        {groups.map((group) => (
          <li>
            <GroupCard group={group} />
          </li>
        ))}
      </ul>
    </BaseLayout>
  );
}
