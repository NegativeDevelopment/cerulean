import { JSX } from "preact";

type RedirectButtonProps = JSX.HTMLAttributes<HTMLButtonElement> & {
  href: string
  title: string
}

export default function RedirectButton(props: RedirectButtonProps) {
  return (
    <a href={props.href} >
      <button
        {...props}
        class="px-4 py-1 text-gray-50 bg-blue-700 rounded"
      >
        {props.title}
      </button>
    </a>
  );
}
