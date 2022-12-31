import { Head } from "$fresh/runtime.ts";
import RedirectButton from "../components/common/RedirectButton.tsx";
import BaseLayout from "../components/layouts/BaseLayout.tsx";

export default function Home() {
  return (
    <BaseLayout>
      <div class="mt-24 mx-16">
        <h1 class="text-3xl max-w-xl">
          A new way to manage expenses together with friends and family
        </h1>
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
