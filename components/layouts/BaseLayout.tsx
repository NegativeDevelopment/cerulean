import { Head } from "$fresh/runtime.ts";
import { ComponentChildren } from "preact";
import Footer from "../Footer.tsx";
import Header from "../Header.tsx";

interface BaseLayoutProps {
  children: ComponentChildren;
}

export default function BaseLayout({ children }: BaseLayoutProps) {
  return (
    <>
      <Head>
        <title>Celurean - Money group manager</title>
      </Head>
      <Header />
      {children}
      <Footer />
    </>
  );
}
