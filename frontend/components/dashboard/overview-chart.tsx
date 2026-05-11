"use client";

import { Area, AreaChart, CartesianGrid, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";
import { Card } from "@/components/ui/card";
import { formatCurrency } from "@/lib/utils";
import { ChartPoint } from "@/types";

type OverviewChartProps = {
  data: ChartPoint[];
};

export function OverviewChart({ data }: OverviewChartProps) {
  return (
    <Card className="h-[380px]">
      <div className="mb-4 flex items-center justify-between">
        <div>
          <h3 className="text-base font-semibold">Cashflow Overview</h3>
          <p className="text-sm text-text-muted">Income vs expenses trend</p>
        </div>
      </div>

      <div className="chart-glow h-[300px]">
        <ResponsiveContainer width="100%" height="100%">
          <AreaChart data={data}>
            <defs>
              <linearGradient id="incomeGradient" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="var(--color-chart-green)" stopOpacity={0.6} />
                <stop offset="95%" stopColor="var(--color-chart-green)" stopOpacity={0} />
              </linearGradient>
              <linearGradient id="expenseGradient" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="var(--color-chart-blue)" stopOpacity={0.45} />
                <stop offset="95%" stopColor="var(--color-chart-blue)" stopOpacity={0} />
              </linearGradient>
            </defs>
            <CartesianGrid stroke="var(--color-border-primary)" vertical={false} />
            <XAxis dataKey="name" tick={{ fill: "var(--color-text-muted)", fontSize: 12 }} axisLine={false} tickLine={false} />
            <YAxis
              tickFormatter={(value) => `${Math.round(value / 1000)}k`}
              tick={{ fill: "var(--color-text-muted)", fontSize: 12 }}
              axisLine={false}
              tickLine={false}
            />
            <Tooltip
              contentStyle={{
                background: "var(--color-surface-secondary)",
                border: "1px solid var(--color-border-secondary)",
                borderRadius: "14px",
                color: "var(--color-text-primary)",
              }}
              formatter={(value) => formatCurrency(Number(value ?? 0))}
            />
            <Area type="monotone" dataKey="income" stroke="var(--color-chart-green)" fill="url(#incomeGradient)" strokeWidth={2.5} />
            <Area type="monotone" dataKey="expense" stroke="var(--color-chart-blue)" fill="url(#expenseGradient)" strokeWidth={2.5} />
          </AreaChart>
        </ResponsiveContainer>
      </div>
    </Card>
  );
}
