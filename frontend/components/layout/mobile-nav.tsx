"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { BarChart3, Landmark, LayoutGrid, Plus, ReceiptText, Wallet } from "lucide-react";
import { NAV_ITEMS } from "@/lib/constants";
import { cn } from "@/lib/utils";

const ICONS = {
  layout: LayoutGrid,
  receipt: ReceiptText,
  wallet: Wallet,
  chart: BarChart3,
  landmark: Landmark,
} as const;

export function MobileNav() {
  const pathname = usePathname();

  return (
    <>
      <button
        className="fixed bottom-24 right-5 z-40 inline-flex size-14 items-center justify-center rounded-full bg-brand-primary text-black shadow-[0_0_30px_rgba(34,197,94,.4)] transition hover:bg-brand-secondary lg:hidden"
        aria-label="Add transaction"
      >
        <Plus className="size-5" />
      </button>

      <nav className="fixed inset-x-0 bottom-0 z-40 border-t border-border-primary/80 bg-surface-primary/95 p-2 backdrop-blur lg:hidden">
        <ul className="grid grid-cols-5 gap-1">
          {NAV_ITEMS.map((item) => {
            const Icon = ICONS[item.icon];
            const active = pathname.startsWith(item.href);
            return (
              <li key={item.href}>
                <Link
                  href={item.href}
                  className={cn(
                    "flex flex-col items-center justify-center gap-1 rounded-xl py-2 text-[11px] transition",
                    active ? "bg-brand-primary/15 text-brand-primary" : "text-text-muted",
                  )}
                >
                  <Icon className="size-4" />
                  {item.label}
                </Link>
              </li>
            );
          })}
        </ul>
      </nav>
    </>
  );
}
