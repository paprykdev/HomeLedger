"use client";

import { Bar, BarChart, CartesianGrid, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";
import { Card } from "@/components/ui/card";
import { ChartPoint } from "@/types";

type SpendingTrendChartProps = {
  data: ChartPoint[];
};

export function SpendingTrendChart({ data }: SpendingTrendChartProps) {
  return (
    <Card className="h-[360px]">
      <h3 className="text-base font-semibold">Spending Trend</h3>
      <p className="mb-4 text-sm text-text-muted">Monthly comparison</p>
      <div className="chart-glow h-[270px]">
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={data}>
            <CartesianGrid stroke="var(--color-border-primary)" vertical={false} />
            <XAxis dataKey="name" tick={{ fill: "var(--color-text-muted)", fontSize: 12 }} axisLine={false} tickLine={false} />
            <YAxis tick={{ fill: "var(--color-text-muted)", fontSize: 12 }} axisLine={false} tickLine={false} />
            <Tooltip
              contentStyle={{
                background: "var(--color-surface-secondary)",
                border: "1px solid var(--color-border-secondary)",
                borderRadius: "14px",
              }}
            />
            <Bar dataKey="income" fill="var(--color-chart-green)" radius={[8, 8, 0, 0]} />
            <Bar dataKey="expense" fill="var(--color-chart-blue)" radius={[8, 8, 0, 0]} />
          </BarChart>
        </ResponsiveContainer>
      </div>
    </Card>
  );
}
