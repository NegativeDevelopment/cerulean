import { JSX } from "preact";
import { Button } from "../components/Button.tsx";
import {Input} from "../components/Input.tsx";

export default function Login() {
  return (
    <div class="mx-auto">
      <h1 class="font-bold text-2xl">Login</h1>
      <form method="POST">
        <label for="username">Username:</label>
        <br />
        <Input type="text" id="username" name="username" />
        <br />
        <label for="password">Password:</label>
        <br />
        <Input type="text" id="password" name="password" />
        <br />
        <Input type="submit" value="Submit" />
      </form>
    </ div>
  );
}
