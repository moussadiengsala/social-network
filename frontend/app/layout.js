"use client";
import { AuthContextProvider } from "@/context/authContext";
import { Jost } from "next/font/google";
import { QueryClient, QueryClientProvider } from "react-query";
import "./globals.css";
import en from "javascript-time-ago/locale/en";
import TimeAgo from "javascript-time-ago";
import { SocketServiceProvider } from "@/context/realTimeContext";
TimeAgo.addDefaultLocale(en);

const inter = Jost({ subsets: ["latin"] });
export const queryClient = new QueryClient();

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <AuthContextProvider>
        <SocketServiceProvider>
          <QueryClientProvider client={queryClient}>
            <body className={inter.className}>{children}</body>
          </QueryClientProvider>
        </SocketServiceProvider>
      </AuthContextProvider>
    </html>
  );
}
