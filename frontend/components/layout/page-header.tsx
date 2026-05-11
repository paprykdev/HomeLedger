"use client";

import { ReactNode } from "react";
import { Search } from "lucide-react";
import { Input } from "@/components/ui/input";

type PageHeaderProps = {
  title: string;
  description: string;
  action?: ReactNode;
  searchPlaceholder?: string;
  searchValue?: string;
  onSearchChange?: (value: string) => void;
};

export function PageHeader({
  title,
  description,
  action,
  searchPlaceholder,
  searchValue,
  onSearchChange,
}: PageHeaderProps) {
  return (
    <header className="mb-8 space-y-4">
      <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <h1 className="text-2xl font-semibold tracking-tight text-text-primary md:text-3xl">{title}</h1>
          <p className="mt-1 text-sm text-text-secondary">{description}</p>
        </div>
        {action}
      </div>

      {(searchPlaceholder || typeof onSearchChange === "function") && (
        <div className="relative max-w-md">
          <Search className="pointer-events-none absolute left-3 top-1/2 size-4 -translate-y-1/2 text-text-muted" />
          <Input
            value={searchValue}
            onChange={(event) => onSearchChange?.(event.target.value)}
            placeholder={searchPlaceholder ?? "Search..."}
            className="pl-9"
          />
        </div>
      )}
    </header>
  );
}
