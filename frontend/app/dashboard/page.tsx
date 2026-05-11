import { Plus } from "lucide-react";
import { ExpensesChart } from "@/components/dashboard/expenses-chart";
import { OverviewChart } from "@/components/dashboard/overview-chart";
import { RecentTransactions } from "@/components/dashboard/recent-transactions";
import { StatsCard } from "@/components/dashboard/stats-card";
import { PageHeader } from "@/components/layout/page-header";
import { Button } from "@/components/ui/button";
import { DASHBOARD_STATS, EXPENSES_BY_CATEGORY, OVERVIEW_DATA, RECENT_TRANSACTIONS } from "@/lib/constants";

export default function DashboardPage() {
  return (
    <div>
      <PageHeader
        title="Dashboard"
        description="Track balances, spending trends and savings in one place."
        searchPlaceholder="Search transactions, categories..."
        action={
          <Button className="hidden sm:inline-flex">
            <Plus className="size-4" />
            Add Transaction
          </Button>
        }
      />

      <section className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        {DASHBOARD_STATS.map((item) => (
          <StatsCard key={item.label} label={item.label} value={item.value} trend={item.trend} trendUp={item.trendUp} />
        ))}
      </section>

      <section className="mt-6 grid gap-4 xl:grid-cols-[1.75fr_1fr]">
        <OverviewChart data={OVERVIEW_DATA} />
        <ExpensesChart data={EXPENSES_BY_CATEGORY} />
      </section>

      <section className="mt-6">
        <RecentTransactions transactions={RECENT_TRANSACTIONS} />
      </section>
    </div>
  );
}
