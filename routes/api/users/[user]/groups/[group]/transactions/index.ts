import { PageProps } from "$fresh/server.ts";
import { Transaction } from "../../../../../../../types/transaction.ts";

const transactions: Transaction[] = [
  { id: "1", groupId: "1", sender: "1", amount: 1, receivers: ["2"] },
  { id: "2", groupId: "1", sender: "2", amount: 1, receivers: ["1"] },
];

export const handler = (_req: Request, props: PageProps): Response => {
  const { group } = props.params;
  if (!group) {
    return new Response("No transaction ID provided", { status: 400 });
  }

  const groupTransactions = transactions.filter((transaction) =>
    transaction.groupId === group
  );
  if (!groupTransactions) {
    return new Response("Group not found", { status: 404 });
  }

  const body = JSON.stringify(groupTransactions, null, 2);
  return new Response(body, {
    headers: { "content-type": "application/json; charset=utf-8" },
  });
};
