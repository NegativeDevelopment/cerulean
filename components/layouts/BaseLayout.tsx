import { Head } from "$fresh/runtime.ts";
import { ComponentChildren } from "preact";
import Footer from "../Footer.tsx";
import Header from "../Header.tsx";

interface BaseLayoutProps {
  children: ComponentChildren;
}

export default function BaseLayout({ children }: BaseLayoutProps) {
  return (
    <div class="min-h-screen flex flex-col">
      <Head>
        <title>Celurean - Money group manager</title>
        <link
          href="https://cdn.jsdelivr.net/npm/daisyui@2.46.0/dist/full.css"
          rel="stylesheet"
          type="text/css"
        />
      </Head>
      <Header />
      <div class="mx-6">
        {children}
      </div>
      <Footer />
    </div>
  );
}
