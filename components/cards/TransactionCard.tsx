import { Transaction } from "../../types/transaction.ts";
import { BaseCard } from "./BaseCard.tsx";

export interface TransactionCardProps {
  transaction: Transaction;
}

export function TransactionCard(
  { transaction }: TransactionCardProps,
) {
  const perReceiverAmount = transaction.amount / transaction.receivers.length;

  return (
    <BaseCard
      titel={transaction.sender}
      description={"01.01.2023"}
      info={transaction.amount + " €"}
      infoTitel={perReceiverAmount + " €"}
    />
  );
}
