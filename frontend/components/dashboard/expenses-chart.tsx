"use client";

import { Cell, Pie, PieChart, ResponsiveContainer, Tooltip } from "recharts";
import { Card } from "@/components/ui/card";
import { CategoryPoint } from "@/types";

type ExpensesChartProps = {
  data: CategoryPoint[];
};

export function ExpensesChart({ data }: ExpensesChartProps) {
  return (
    <Card className="h-[380px]">
      <h3 className="text-base font-semibold">Expense Categories</h3>
      <p className="text-sm text-text-muted">Current month split</p>

      <div className="chart-glow mt-2 h-[230px]">
        <ResponsiveContainer width="100%" height="100%">
          <PieChart>
            <Pie data={data} dataKey="value" nameKey="name" innerRadius={55} outerRadius={95} paddingAngle={2}>
              {data.map((entry) => (
                <Cell key={entry.name} fill={entry.color} />
              ))}
            </Pie>
            <Tooltip
              contentStyle={{
                background: "var(--color-surface-secondary)",
                border: "1px solid var(--color-border-secondary)",
                borderRadius: "14px",
                color: "var(--color-text-primary)",
              }}
              formatter={(value) => `${Number(value ?? 0)}%`}
            />
          </PieChart>
        </ResponsiveContainer>
      </div>

      <div className="grid grid-cols-2 gap-2 text-xs text-text-secondary">
        {data.map((item) => (
          <div key={item.name} className="flex items-center gap-2">
            <span className="size-2.5 rounded-full" style={{ background: item.color }} />
            <span>{item.name}</span>
            <span className="ml-auto text-text-muted">{item.value}%</span>
          </div>
        ))}
      </div>
    </Card>
  );
}
