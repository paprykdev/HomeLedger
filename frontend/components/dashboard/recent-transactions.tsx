import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
import { formatCurrency } from "@/lib/utils";
import { Transaction } from "@/types";

type RecentTransactionsProps = {
  transactions: Transaction[];
};

export function RecentTransactions({ transactions }: RecentTransactionsProps) {
  return (
    <Card>
      <div className="mb-4 flex items-center justify-between">
        <h3 className="text-base font-semibold">Recent Transactions</h3>
        <p className="text-xs text-text-muted">{transactions.length} entries</p>
      </div>

      <div className="space-y-2">
        {transactions.map((transaction) => (
          <div
            key={transaction.id}
            className="flex items-center justify-between rounded-2xl border border-border-primary/70 bg-surface-secondary/80 px-3 py-3 transition hover:border-border-secondary"
          >
            <div>
              <p className="text-sm font-medium text-text-primary">{transaction.description}</p>
              <p className="mt-1 text-xs text-text-muted">{transaction.category}</p>
            </div>

            <div className="text-right">
              <p className={transaction.type === "income" ? "text-income-primary" : "text-expense-primary"}>
                {transaction.type === "income" ? "+" : "-"}
                {formatCurrency(transaction.amount)}
              </p>
              <div className="mt-1">
                <Badge tone={transaction.type}>{transaction.type}</Badge>
              </div>
            </div>
          </div>
        ))}
      </div>
    </Card>
  );
}
