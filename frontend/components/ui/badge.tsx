import { HTMLAttributes } from "react";
import { cn } from "@/lib/utils";

type BadgeProps = HTMLAttributes<HTMLSpanElement> & {
  tone?: "default" | "income" | "expense";
};

export function Badge({ className, tone = "default", ...props }: BadgeProps) {
  return (
    <span
      className={cn(
        "inline-flex items-center rounded-full border px-2.5 py-1 text-xs font-medium",
        tone === "income" && "border-income-primary/30 bg-income-bg text-income-secondary",
        tone === "expense" && "border-expense-primary/30 bg-expense-bg text-expense-secondary",
        tone === "default" && "border-border-secondary bg-surface-secondary text-text-secondary",
        className,
      )}
      {...props}
    />
  );
}
