import type { Metadata } from "next";
import {inter,lusitana,josefin} from "./ui/font"
import "./ui/globals.css";
import Navbar from "../components/Fragments/navbar"

export const metadata: Metadata = {
  title: "Stress Race",
  description: "Tugas Besar 2 Strategi Algoritma",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={josefin.className}>
      <Navbar />
        {children}
        </body>
    </html>
  );
}
