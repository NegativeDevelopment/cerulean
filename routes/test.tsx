import { Head } from "$fresh/runtime.ts";
import { BaseCard } from "../components/cards/BaseCard.tsx";
import { GroupCard } from "../components/cards/GroupCard.tsx";
import { TransactionCard } from "../components/cards/TransactionCard.tsx";

export default function Test() {
  return (
    <>
      <Head>
        <link
          href="https://cdn.jsdelivr.net/npm/daisyui@2.46.0/dist/full.css"
          rel="stylesheet"
          type="text/css"
        />
      </Head>

      {
        /* <div class="h-32 w-16 bg-red-200">
        <ProfilePicture />
      </div> */
      }
      <div class="flex flex-col gap-4 overflow-y-scroll">
        <h1>Groups</h1>
        <GroupCard
          imageUrl="https://picsum.photos/60"
          titel="Vacation 2022"
          description="3 Members"
          href="/groups/1"
        />
        <GroupCard
          titel="Group that has a really long name example 123"
          description="7 Members"
          href="/groups/1"
        />
        <h1>Transactions</h1>
        <TransactionCard
          titel="Carsharing Service"
          description="05.07.2022"
          info="95.00 $"
          infoTitel="345.00 $"
        />
        <TransactionCard
          titel="Gas Station"
          description="06.07.2022"
          info="15.00 $"
          infoTitel="75.00 $"
        />
        <h1>BaseCards</h1>
        <BaseCard
          imageUrl="https://picsum.photos/60"
          href="/groups/1"
        />
        <BaseCard
          imageUrl="https://picsum.photos/60"
          titel="Hello World"
          href="/groups/1"
        />
        <BaseCard
          imageUrl="https://picsum.photos/60"
          titel="Way to long titel looooooooooooooooooooooooooooooooooong"
          description="Tesstttt"
          href="/groups/1"
        />
        <BaseCard
          titel="Example"
          description="This description is wayy to looooooooooooooooooooooooooooooooooooooooooooooong"
          href="/groups/1"
        />
        <BaseCard
          description="This description is wayy to looooooooooooooooooooooooooooooooooooooooooooooong"
          href="/groups/1"
        />
        <BaseCard
          description="This description is wayy to looooooooooooooooooooooooooooooooooooooooooooooong"
          info="80$"
          infoTitel="Balance"
          href="/groups/1"
        />
      </div>
    </>
  );
}
