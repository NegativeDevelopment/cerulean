import { Handlers, PageProps } from "$fresh/server.ts";
import { TransactionCard } from "../../components/cards/TransactionCard.tsx";
import BaseLayout from "../../components/layouts/BaseLayout.tsx";
import ProfilePicture from "../../components/profile/ProfilePicture.tsx";
import { Group } from "../../types/group.ts";
import { Transaction } from "../../types/transaction.ts";

export default function GroupPage(props: PageProps<Group[]>) {
  const pendingPayments = [];

  const transactions: Transaction[] = [{
    id: "123",
    groupId: "1",
    sender: "Example User 1",
    amount: 50,
    receivers: ["Example User 1", "Example Receiver 2"],
  }, {
    id: "123",
    groupId: "1",
    sender: "Example User 1",
    amount: 345,
    receivers: ["Example Receiver 2"],
  }];

  return (
    <BaseLayout>
      <h1 class="mb-4 text-2xl tracking-wider uppercase font-light">
        Pending Payments
      </h1>
      <ul class="flex flex-col gap-4">
        <li>
          <span>Nothing to show yet</span>
        </li>
      </ul>

      <h1 class="my-4 text-2xl tracking-wider uppercase font-light">
        Transactions
      </h1>
      <ul class="flex flex-col gap-4">
        {transactions.map((transaction) => (
          <li>
            <TransactionCard transaction={transaction} />
          </li>
        ))}
      </ul>
    </BaseLayout>
  );
}
