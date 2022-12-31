import { ComponentChildren } from "preact";

export default function Footer() {
  const menus = [
    {
      title: "Links",
      children: [
        { name: "GitHub", href: "#" },
        { name: "Discord", href: "#" },
      ],
    },
    {
      title: "Techstack",
      children: [
        { name: "Deno", href: "https://deno.land/" },
        { name: "Fresh", href: "https://fresh.deno.dev/" },
        { name: "Redis", href: "#" },
      ],
    },
  ];

  return (
    <div class="bg-white flex flex-col md:flex-row w-full max-w-screen-lg gap-8 md:gap-16 px-8 py-8 text-sm bottom-0 fixed">
      <div class="flex-1">
        <div class="flex items-center gap-1">
          <div class="font-bold text-2xl">
            Celurean
          </div>
        </div>
        <div class="text-gray-500">
          Expense management for friends and family
        </div>
      </div>

      {menus.map((item) => (
        <div class="mb-4" key={item.title}>
          <div class="font-bold">{item.title}</div>
          <ul class="mt-2">
            {item.children.map((child) => (
              <li class="mt-2" key={child.name}>
                <a
                  class="text-gray-500 hover:text-gray-700"
                  href={child.href}
                >
                  {child.name}
                </a>
              </li>
            ))}
          </ul>
        </div>
      ))}

      <div class="text-gray-500 space-y-2">
        <div class="text-xs">
        </div>
      </div>
    </div>
  );
}
