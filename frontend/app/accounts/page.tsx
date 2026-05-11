import { ArrowDownRight, ArrowUpRight, CreditCard } from "lucide-react";
import { PageHeader } from "@/components/layout/page-header";
import { Card } from "@/components/ui/card";
import { ACCOUNTS } from "@/lib/constants";
import { formatCurrency } from "@/lib/utils";

export default function AccountsPage() {
  return (
    <div>
      <PageHeader title="Accounts" description="Overview of balances and account performance." />

      <section className="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
        {ACCOUNTS.map((account) => (
          <Card key={account.id} className="space-y-4">
            <div className="flex items-center justify-between">
              <div className="inline-flex size-10 items-center justify-center rounded-2xl bg-info-bg text-info-secondary">
                <CreditCard className="size-5" />
              </div>
              <span
                className={`inline-flex items-center gap-1 rounded-full px-2 py-1 text-xs ${
                  account.changeUp ? "bg-income-bg text-income-secondary" : "bg-expense-bg text-expense-secondary"
                }`}
              >
                {account.changeUp ? <ArrowUpRight className="size-3.5" /> : <ArrowDownRight className="size-3.5" />}
                {account.change}
              </span>
            </div>
            <div>
              <p className="text-sm text-text-secondary">{account.name}</p>
              <p className="mt-1 text-2xl font-semibold tracking-tight">{formatCurrency(account.balance, account.currency)}</p>
            </div>
          </Card>
        ))}
      </section>
    </div>
  );
}
