import { ArrowDownLeft, ArrowUpRight } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
import { formatCurrency } from "@/lib/utils";
import { Transaction } from "@/types";

type TransactionListProps = {
  transactions: Transaction[];
};

export function TransactionList({ transactions }: TransactionListProps) {
  return (
    <Card>
      <div className="mb-4 flex items-center justify-between">
        <h2 className="text-base font-semibold">Transactions</h2>
        <p className="text-xs text-text-muted">{transactions.length} total</p>
      </div>

      <div className="space-y-2">
        {transactions.map((transaction) => (
          <article
            key={transaction.id}
            className="flex items-center justify-between rounded-2xl border border-border-primary/70 bg-surface-secondary/80 px-4 py-3 transition hover:border-border-secondary"
          >
            <div className="flex items-center gap-3">
              <div
                className={`flex size-9 items-center justify-center rounded-xl ${
                  transaction.type === "income" ? "bg-income-bg text-income-primary" : "bg-expense-bg text-expense-primary"
                }`}
              >
                {transaction.type === "income" ? <ArrowDownLeft className="size-4" /> : <ArrowUpRight className="size-4" />}
              </div>
              <div>
                <p className="text-sm font-medium">{transaction.description}</p>
                <p className="text-xs text-text-muted">{new Date(transaction.created_at).toLocaleDateString()}</p>
              </div>
            </div>

            <div className="flex items-center gap-3">
              <Badge tone={transaction.type}>{transaction.category}</Badge>
              <p className={transaction.type === "income" ? "text-income-primary" : "text-expense-primary"}>
                {transaction.type === "income" ? "+" : "-"}
                {formatCurrency(transaction.amount)}
              </p>
            </div>
          </article>
        ))}
      </div>
    </Card>
  );
}
