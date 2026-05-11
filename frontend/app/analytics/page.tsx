import { ExpensesChart } from "@/components/dashboard/expenses-chart";
import { StatsCard } from "@/components/dashboard/stats-card";
import { SpendingTrendChart } from "@/components/analytics/spending-trend-chart";
import { SavingsChart } from "@/components/analytics/savings-chart";
import { PageHeader } from "@/components/layout/page-header";
import { DASHBOARD_STATS, EXPENSES_BY_CATEGORY, OVERVIEW_DATA } from "@/lib/constants";

export default function AnalyticsPage() {
  return (
    <div>
      <PageHeader title="Analytics" description="Deep insights into cashflow, savings and spending behavior." />

      <section className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        {DASHBOARD_STATS.map((item) => (
          <StatsCard key={item.label} label={item.label} value={item.value} trend={item.trend} trendUp={item.trendUp} />
        ))}
      </section>

      <section className="mt-6 grid gap-4 xl:grid-cols-2">
        <SpendingTrendChart data={OVERVIEW_DATA} />
        <SavingsChart data={OVERVIEW_DATA} />
      </section>

      <section className="mt-6 grid gap-4 xl:grid-cols-2">
        <ExpensesChart data={EXPENSES_BY_CATEGORY} />
      </section>
    </div>
  );
}
