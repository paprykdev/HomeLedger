export type TransactionType = "income" | "expense";

export type Transaction = {
  id: string;
  user_id?: string;
  type: TransactionType;
  amount: number;
  description: string;
  category: string;
  account_id: string;
  created_at: string;
  updated_at: string;
};

export type Stat = {
  label: string;
  value: string;
  trend: string;
  trendUp?: boolean;
};

export type ChartPoint = {
  name: string;
  income: number;
  expense: number;
  savings?: number;
};

export type CategoryPoint = {
  name: string;
  value: number;
  color: string;
};

export type BudgetItem = {
  name: string;
  spent: number;
  total: number;
};

export type AccountItem = {
  id: string;
  name: string;
  balance: number;
  currency: "USD" | "EUR" | "PLN";
  change: string;
  changeUp: boolean;
};
