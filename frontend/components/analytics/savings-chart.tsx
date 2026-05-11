"use client";

import { Line, LineChart, CartesianGrid, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";
import { Card } from "@/components/ui/card";
import { ChartPoint } from "@/types";

type SavingsChartProps = {
  data: ChartPoint[];
};

export function SavingsChart({ data }: SavingsChartProps) {
  return (
    <Card className="h-[360px]">
      <h3 className="text-base font-semibold">Savings Momentum</h3>
      <p className="mb-4 text-sm text-text-muted">Monthly savings performance</p>
      <div className="chart-glow h-[270px]">
        <ResponsiveContainer width="100%" height="100%">
          <LineChart data={data}>
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
            <Line type="monotone" dataKey="savings" strokeWidth={3} stroke="var(--color-chart-purple)" dot={false} />
          </LineChart>
        </ResponsiveContainer>
      </div>
    </Card>
  );
}
