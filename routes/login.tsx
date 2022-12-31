import { JSX } from "preact";
import Input from "../components/common/Input.tsx";
import BaseLayout from "../components/layouts/BaseLayout.tsx";

export default function Login() {
  return (
    <BaseLayout>
      <div class="mx-auto my-16">
        <h1 class="font-bold text-4xl">Login</h1>
        <form method="POST" class="flex flex-col gap-5 mt-8">
          <Input type="text" title="Username" id="username" name="username" />
          <Input type="text" title="Password" id="password" name="password" />
          <Input type="submit" value="Submit" class="btn btn-primary" />
        </form>
      </div>
    </BaseLayout>
  );
}
