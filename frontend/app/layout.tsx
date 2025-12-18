import type { Metadata } from "next";
import "./globals.css";
import Navbar from "@/components/Navbar";

export const metadata: Metadata = {
  title: "GetAll Market",
  description: "Simple Community Marketplace",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="antialiased font-sans selection:bg-blue-100 selection:text-blue-900">
        <Navbar />
        <main className="min-h-screen pt-20 pb-20">{children}</main>
        
        <footer className="bg-white border-t border-gray-100 py-12">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex flex-col md:flex-row justify-between items-center gap-6">
              <div className="flex items-center gap-2">
                <div className="w-8 h-8 bg-gray-900 rounded-lg flex items-center justify-center text-white font-black text-sm">G</div>
                <span className="text-sm font-black tracking-tighter text-gray-900">GET<span className="text-blue-600">ALL</span></span>
              </div>
              <p className="text-xs font-bold text-gray-400 uppercase tracking-widest">
                &copy; {new Date().getFullYear()} GetAll Market. Simple & Usable.
              </p>
              <div className="flex gap-6 text-[10px] font-black text-gray-400 uppercase tracking-widest">
                <a href="#" className="hover:text-blue-600 transition-colors">Privacy</a>
                <a href="#" className="hover:text-blue-600 transition-colors">Terms</a>
              </div>
            </div>
          </div>
        </footer>
      </body>
    </html>
  );
}
