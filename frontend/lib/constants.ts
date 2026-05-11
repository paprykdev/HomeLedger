import { AccountItem, BudgetItem, CategoryPoint, ChartPoint, Stat, Transaction } from "@/types";

export const NAV_ITEMS = [
  { href: "/dashboard", label: "Dashboard", icon: "layout" },
  { href: "/transactions", label: "Transactions", icon: "receipt" },
  { href: "/budgets", label: "Budgets", icon: "wallet" },
  { href: "/analytics", label: "Analytics", icon: "chart" },
  { href: "/accounts", label: "Accounts", icon: "landmark" },
] as const;

export const DASHBOARD_STATS: Stat[] = [
  { label: "Total Balance", value: "82 540 zł", trend: "+8.1% vs last month", trendUp: true },
  { label: "Income", value: "23 800 zł", trend: "+4.3% this month", trendUp: true },
  { label: "Expenses", value: "11 220 zł", trend: "-1.8% this month", trendUp: true },
  { label: "Savings Rate", value: "41.2%", trend: "+2.4pp this month", trendUp: true },
];

export const OVERVIEW_DATA: ChartPoint[] = [
  { name: "Jan", income: 14300, expense: 9000, savings: 5300 },
  { name: "Feb", income: 15800, expense: 9800, savings: 6000 },
  { name: "Mar", income: 15100, expense: 9300, savings: 5800 },
  { name: "Apr", income: 16700, expense: 10200, savings: 6500 },
  { name: "May", income: 17400, expense: 11000, savings: 6400 },
  { name: "Jun", income: 18200, expense: 11200, savings: 7000 },
];

export const EXPENSES_BY_CATEGORY: CategoryPoint[] = [
  { name: "Housing", value: 34, color: "var(--color-chart-blue)" },
  { name: "Food", value: 18, color: "var(--color-chart-orange)" },
  { name: "Transport", value: 13, color: "var(--color-chart-purple)" },
  { name: "Leisure", value: 10, color: "var(--color-chart-pink)" },
  { name: "Utilities", value: 15, color: "var(--color-chart-green)" },
  { name: "Other", value: 10, color: "var(--color-chart-red)" },
];

export const RECENT_TRANSACTIONS: Transaction[] = [
  {
    id: "txn_1",
    type: "expense",
    amount: 219,
    description: "Grocery Market",
    category: "food",
    account_id: "acc_1",
    created_at: "2026-05-11T10:30:00Z",
    updated_at: "2026-05-11T10:30:00Z",
  },
  {
    id: "txn_2",
    type: "income",
    amount: 8600,
    description: "Salary",
    category: "salary",
    account_id: "acc_1",
    created_at: "2026-05-10T08:10:00Z",
    updated_at: "2026-05-10T08:10:00Z",
  },
  {
    id: "txn_3",
    type: "expense",
    amount: 430,
    description: "Electricity Bill",
    category: "utilities",
    account_id: "acc_2",
    created_at: "2026-05-09T07:20:00Z",
    updated_at: "2026-05-09T07:20:00Z",
  },
  {
    id: "txn_4",
    type: "expense",
    amount: 159,
    description: "Fuel",
    category: "transport",
    account_id: "acc_2",
    created_at: "2026-05-08T16:15:00Z",
    updated_at: "2026-05-08T16:15:00Z",
  },
];

export const BUDGETS: BudgetItem[] = [
  { name: "Housing", spent: 3400, total: 4200 },
  { name: "Food", spent: 1180, total: 1600 },
  { name: "Transport", spent: 720, total: 1100 },
  { name: "Utilities", spent: 680, total: 900 },
  { name: "Leisure", spent: 510, total: 950 },
];

export const ACCOUNTS: AccountItem[] = [
  { id: "acc_1", name: "Main Account", balance: 48500, currency: "PLN", change: "+6.2%", changeUp: true },
  { id: "acc_2", name: "Savings Vault", balance: 32100, currency: "PLN", change: "+9.4%", changeUp: true },
  { id: "acc_3", name: "Travel Card", balance: 1940, currency: "EUR", change: "-1.1%", changeUp: false },
];
