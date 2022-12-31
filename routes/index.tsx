import { Head } from "$fresh/runtime.ts";
import RedirectButton from "../components/common/RedirectButton.tsx";
import BaseLayout from "../components/layouts/BaseLayout.tsx";

export default function Home() {
  return (
    <BaseLayout>
      <div class="my-16 mx-12">
        <h1 class="text-6xl mb-8">Celurean</h1>
        <p>A new way to manage expanses together with Friends and Family</p>
        <div>
          <RedirectButton href="/login" title="Login" className="mt-12 mr-6" />
          <span class="text-gray-600">
            or <a class="underline text-gray-800">Sign up</a> now.
          </span>
        </div>
      </div>
    </BaseLayout>
  );
}
