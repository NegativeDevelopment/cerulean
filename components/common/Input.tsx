import { JSX } from "preact";
import { IS_BROWSER } from "$fresh/runtime.ts";

export default function Input(props: JSX.HTMLAttributes<HTMLInputElement>) {
  if (props.title) {
    return (
      <label class="input-group">
        <span>{props.title}</span>
        <input type="text" class="input input-bordered" {...props} />
      </label>
    );
  }

  return <input type="text" class="input input-bordered" {...props} />;
}
