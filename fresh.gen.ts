// DO NOT EDIT. This file is generated by fresh.
// This file SHOULD be checked into source version control.
// This file is automatically updated during development when running `dev.ts`.

import config from "./deno.json" assert { type: "json" };
import * as $0 from "./routes/api/joke.ts";
import * as $1 from "./routes/groups/[group].tsx";
import * as $2 from "./routes/index.tsx";
import * as $3 from "./routes/login.tsx";
import * as $4 from "./routes/me/groups.tsx";
import * as $5 from "./routes/test.tsx";
import * as $$0 from "./islands/Counter.tsx";

const manifest = {
  routes: {
    "./routes/api/joke.ts": $0,
    "./routes/groups/[group].tsx": $1,
    "./routes/index.tsx": $2,
    "./routes/login.tsx": $3,
    "./routes/me/groups.tsx": $4,
    "./routes/test.tsx": $5,
  },
  islands: {
    "./islands/Counter.tsx": $$0,
  },
  baseUrl: import.meta.url,
  config,
};

export default manifest;
