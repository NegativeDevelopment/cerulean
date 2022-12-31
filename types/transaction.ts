export interface Transaction {
  id: string;
  sender: string;
  amount: number;
  recievers: string[];
}
