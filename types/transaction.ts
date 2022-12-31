export interface Transaction {
  id: string;
  groupId: string;
  sender: string;
  amount: number;
  receivers: string[];
}
