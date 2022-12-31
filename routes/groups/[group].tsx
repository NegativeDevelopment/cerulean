import { Handlers, PageProps } from "$fresh/server.ts";
import ProfilePicture from "../../components/profile/ProfilePicture.tsx";
import { Group } from "../../types/group.ts";

export default function GroupPage(props: PageProps<Group[]>) {
  console.log(props.data);

  const group = {
    name: "Example group",
  };

  const transaction = {
    titel: "Test",
    user: {
      name: "Example User",
    },
    amount: 10,
    myAmount: 5,
  };

  return (
    <>
      <div class="sm:w-[640px] mx-8 my-3 sm:mx-auto">
        <div class="mb-4 tracking-wider uppercase font-light flex justify-between">
          <h1 class="text-xl text-gray-700">{group.name}</h1>
          <div class="flex text-xs items-end text-gray-400">
            <a>Sort</a>
            <div class="w-1 bg-black"></div>
            <a>Sort</a>
          </div>
        </div>
        <h2 class="my-8 tracking-wider uppercase font-light flex justify-between">
          Pending
        </h2>
        <ul>
          <li class="flex w-full gap-6">
            <ProfilePicture />
            <div class="">
              <div class="text-gray-600 ">{transaction.user.name}</div>
            </div>
            <div class="ml-auto">
              <div class="text-gray-900 font-bold">{transaction.myAmount}</div>
            </div>
          </li>
        </ul>

        <h2 class="my-8 tracking-wider uppercase font-light flex justify-between">
          Transactions
        </h2>
        <ul>
          <li class="flex w-full gap-6">
            <ProfilePicture />
            <div class="">
              <div class="text-gray-600 ">{transaction.user.name}</div>
              <div class="font-light text-gray-400 text-sm">
                {transaction.titel}
              </div>
              <div>Example User 1, 2 3</div>
            </div>
            <div class="ml-auto">
              <div class="text-gray-900 font-bold">{transaction.myAmount}</div>
              <br />
              <div class="text-yellow-500">{transaction.amount}</div>
            </div>
          </li>
        </ul>
      </div>
    </>
  );
}
