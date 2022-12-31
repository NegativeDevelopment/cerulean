import { Handlers, PageProps } from "$fresh/server.ts";
import { GroupList } from "../../components/group/GroupList.tsx";
import { Group } from "../../types/group.ts";

const NAMES = ["Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"];
const DEMO_GROUPS = [
  {
    id: "123-5132-43123",
    name: "Example Group",
    members: [],

  }
]

interface Data {
  groups: Group[];
}

export const handler: Handlers<Data> = {
  async GET(req, ctx) {
    const url = new URL(req.url);
    const response = await fetch(`${url.origin}/api/users/1/groups/1`);
    const group: Group =  await response.json().catch(console.error)
    const groups = [group]
    if(!groups)
      return new Response("Groups not found", {status: 404})

    return ctx.render({ groups });
  },
};



export default function Groups({ data }: PageProps<Data>) {
  const { groups } = data;

  return (
    <div>
      <GroupList groups={groups}/>
    </div>
  );
}