import { Handlers, PageProps } from "$fresh/server.ts";
import { GroupList } from "../../components/group/GroupList.tsx";

const NAMES = ["Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"];
const DEMO_GROUPS = [
  {
    id: "123-5132-43123",
    name: "Example Group",
    members: [],

  }
]

interface Data {
  results: string[];
  query: string;
}

export const handler: Handlers<Data> = {
  GET(req, ctx) {
    const url = new URL(req.url);
    const query = url.searchParams.get("q") || "";
    const results = NAMES.filter((name) => name.includes(query));
    return ctx.render({ results, query });
  },
};

export default function Groups({ data }: PageProps<Data>) {
  const { results, query } = data;
  return (
    <div>
      {/* <form>
        <input type="text" name="q" value={query} />
        <button type="submit">Search</button>
      </form>
      <ul>
        {results.map((name) => <li key={name}>{name}</li>)}
      </ul> */}
      <GroupList groups={[{name: "Example Group"},{name: "Example Group2"}]}/>
    </div>
  );
}