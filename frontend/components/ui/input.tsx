import { InputHTMLAttributes } from "react";
import { cn } from "@/lib/utils";

export function Input({ className, ...props }: InputHTMLAttributes<HTMLInputElement>) {
  return (
    <input
      className={cn(
        "h-11 w-full rounded-2xl border border-border-secondary bg-surface-secondary px-4 text-sm text-text-primary placeholder:text-text-muted outline-none transition focus:border-brand-primary/70",
        className,
      )}
      {...props}
    />
  );
}
