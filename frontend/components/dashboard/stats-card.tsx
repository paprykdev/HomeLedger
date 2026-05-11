import { ArrowDownRight, ArrowUpRight } from "lucide-react";
import { Card } from "@/components/ui/card";
import { cn } from "@/lib/utils";

type StatsCardProps = {
  label: string;
  value: string;
  trend: string;
  trendUp?: boolean;
};

export function StatsCard({ label, value, trend, trendUp = true }: StatsCardProps) {
  return (
    <Card className="space-y-3">
      <p className="text-sm text-text-secondary">{label}</p>
      <p className="text-2xl font-semibold tracking-tight">{value}</p>
      <p
        className={cn(
          "inline-flex items-center gap-1 rounded-full px-2 py-1 text-xs",
          trendUp ? "bg-income-bg text-income-secondary" : "bg-expense-bg text-expense-secondary",
        )}
      >
        {trendUp ? <ArrowUpRight className="size-3.5" /> : <ArrowDownRight className="size-3.5" />}
        {trend}
      </p>
    </Card>
  );
}
