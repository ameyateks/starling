import type { Metadata } from "next";
import { Abhaya_Libre } from "next/font/google";
import "./globals.css";

const inter = Abhaya_Libre({ weight: "400", subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Pocket Watcher",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  );
}
