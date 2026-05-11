"use client";

import { useEffect, useMemo, useState } from "react";
import { Plus } from "lucide-react";
import { PageHeader } from "@/components/layout/page-header";
import { TransactionFilters } from "@/components/transactions/transaction-filters";
import { TransactionList } from "@/components/transactions/transaction-list";
import { Button } from "@/components/ui/button";
import { RECENT_TRANSACTIONS } from "@/lib/constants";
import { getTransactions } from "@/lib/api";
import { Transaction } from "@/types";

export default function TransactionsPage() {
  const [transactions, setTransactions] = useState<Transaction[]>(RECENT_TRANSACTIONS);
  const [loadError, setLoadError] = useState<string | null>(null);
  const [search, setSearch] = useState("");
  const [typeFilter, setTypeFilter] = useState<"all" | "income" | "expense">("all");
  const [categoryFilter, setCategoryFilter] = useState("");

  useEffect(() => {
    async function load() {
      try {
        const payload = await getTransactions();
        setTransactions(payload);
      } catch {
        setLoadError("API unavailable, showing local sample data.");
      }
    }
    load();
  }, []);

  const filtered = useMemo(() => {
    return transactions.filter((transaction: Transaction) => {
      const matchesSearch =
        transaction.description.toLowerCase().includes(search.toLowerCase()) ||
        transaction.category.toLowerCase().includes(search.toLowerCase());
      const matchesType = typeFilter === "all" ? true : transaction.type === typeFilter;
      const matchesCategory = categoryFilter
        ? transaction.category.toLowerCase().includes(categoryFilter.toLowerCase())
        : true;
      return matchesSearch && matchesType && matchesCategory;
    });
  }, [transactions, search, typeFilter, categoryFilter]);

  return (
    <div>
      <PageHeader
        title="Transactions"
        description="Filter and inspect transaction history."
        searchPlaceholder="Search transactions..."
        searchValue={search}
        onSearchChange={setSearch}
        action={
          <Button className="hidden sm:inline-flex">
            <Plus className="size-4" />
            New Transaction
          </Button>
        }
      />

      <TransactionFilters
        typeFilter={typeFilter}
        setTypeFilter={setTypeFilter}
        categoryFilter={categoryFilter}
        setCategoryFilter={setCategoryFilter}
      />

      {loadError && <p className="mb-4 text-sm text-warning-primary">{loadError}</p>}

      <TransactionList transactions={filtered} />
    </div>
  );
}
