import { HTMLAttributes } from "react";
import { cn } from "@/lib/utils";

export function Card({ className, ...props }: HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn(
        "glass-card rounded-3xl border border-border-primary/80 bg-surface-primary p-5 shadow-xl md:p-6",
        className,
      )}
      {...props}
    />
  );
}
