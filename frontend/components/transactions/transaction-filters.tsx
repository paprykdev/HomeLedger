import { Dispatch, SetStateAction } from "react";
import { Card } from "@/components/ui/card";

type TransactionFiltersProps = {
  typeFilter: "all" | "income" | "expense";
  setTypeFilter: Dispatch<SetStateAction<"all" | "income" | "expense">>;
  categoryFilter: string;
  setCategoryFilter: Dispatch<SetStateAction<string>>;
};

export function TransactionFilters({
  typeFilter,
  setTypeFilter,
  categoryFilter,
  setCategoryFilter,
}: TransactionFiltersProps) {
  return (
    <Card className="mb-6">
      <div className="grid gap-3 sm:grid-cols-2">
        <label className="space-y-2 text-sm text-text-secondary">
          Type
          <select
            className="h-11 w-full rounded-2xl border border-border-secondary bg-surface-secondary px-3 text-text-primary outline-none"
            value={typeFilter}
            onChange={(event) => setTypeFilter(event.target.value as "all" | "income" | "expense")}
          >
            <option value="all">All</option>
            <option value="income">Income</option>
            <option value="expense">Expense</option>
          </select>
        </label>

        <label className="space-y-2 text-sm text-text-secondary">
          Category
          <input
            value={categoryFilter}
            onChange={(event) => setCategoryFilter(event.target.value)}
            placeholder="e.g. food, salary..."
            className="h-11 w-full rounded-2xl border border-border-secondary bg-surface-secondary px-3 text-text-primary placeholder:text-text-muted outline-none"
          />
        </label>
      </div>
    </Card>
  );
}
