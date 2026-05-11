import { PageHeader } from "@/components/layout/page-header";
import { Card } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { BUDGETS } from "@/lib/constants";
import { formatCurrency } from "@/lib/utils";

export default function BudgetsPage() {
  return (
    <div>
      <PageHeader
        title="Budgets"
        description="Set monthly limits and monitor category progress."
        searchPlaceholder="Search budget category..."
      />

      <section className="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
        {BUDGETS.map((budget) => {
          const progress = (budget.spent / budget.total) * 100;
          return (
            <Card key={budget.name} className="space-y-4">
              <div className="flex items-center justify-between">
                <h3 className="font-semibold">{budget.name}</h3>
                <span className="text-xs text-text-muted">{Math.round(progress)}%</span>
              </div>
              <Progress value={progress} />
              <div className="flex items-center justify-between text-sm">
                <span className="text-text-secondary">{formatCurrency(budget.spent)} spent</span>
                <span className="text-text-muted">{formatCurrency(budget.total)} limit</span>
              </div>
            </Card>
          );
        })}
      </section>
    </div>
  );
}
