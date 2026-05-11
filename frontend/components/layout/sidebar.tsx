"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { BarChart3, Landmark, LayoutGrid, ReceiptText, Wallet } from "lucide-react";
import { NAV_ITEMS } from "@/lib/constants";
import { cn } from "@/lib/utils";

const ICONS = {
  layout: LayoutGrid,
  receipt: ReceiptText,
  wallet: Wallet,
  chart: BarChart3,
  landmark: Landmark,
} as const;

export function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="hidden h-screen w-72 shrink-0 border-r border-border-primary/80 bg-background-secondary/70 p-6 lg:flex lg:flex-col">
      <div className="mb-10 flex items-center gap-3">
        <div className="size-9 rounded-xl bg-brand-primary/90 shadow-[0_0_25px_rgba(34,197,94,.45)]" />
        <div>
          <p className="text-lg font-semibold tracking-tight">HomeLedger</p>
          <p className="text-xs text-text-muted">Finance OS</p>
        </div>
      </div>

      <nav className="space-y-2">
        {NAV_ITEMS.map((item) => {
          const Icon = ICONS[item.icon];
          const active = pathname.startsWith(item.href);
          return (
            <Link
              key={item.href}
              href={item.href}
              className={cn(
                "flex items-center gap-3 rounded-2xl border px-4 py-3 text-sm transition duration-200",
                active
                  ? "border-brand-primary/30 bg-brand-primary/15 text-text-primary"
                  : "border-transparent text-text-secondary hover:border-border-secondary hover:bg-surface-primary",
              )}
            >
              <Icon className="size-4" />
              {item.label}
            </Link>
          );
        })}
      </nav>
    </aside>
  );
}
