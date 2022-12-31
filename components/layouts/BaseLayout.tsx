import { Head } from "$fresh/runtime.ts";
import { ComponentChildren } from "preact";

interface BaseLayoutProps {
  children: ComponentChildren
}

export default function BaseLayout({children}: BaseLayoutProps) {
  return (
    <>
      <Head>
        <title>Celurean - Money group manager</title>
      </Head>
      {children}
    </>
  )
}