import { ReactNode } from "react";
import { MobileNav } from "@/components/layout/mobile-nav";
import { Sidebar } from "@/components/layout/sidebar";

type DashboardLayoutProps = {
  children: ReactNode;
};

export function DashboardLayout({ children }: DashboardLayoutProps) {
  return (
    <div className="min-h-screen bg-background-primary">
      <div className="mx-auto flex min-h-screen max-w-[1600px]">
        <Sidebar />
        <main className="w-full pb-24 lg:pb-0">
          <div className="mx-auto w-full max-w-7xl p-4 pt-5 sm:p-6 lg:p-10">{children}</div>
        </main>
      </div>
      <MobileNav />
    </div>
  );
}
