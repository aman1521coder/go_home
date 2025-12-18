'use client';

import Link from 'next/link';
import { useRouter, usePathname } from 'next/navigation';
import { useEffect, useState } from 'react';
import { isAuthenticated, removeToken, getUser } from '@/lib/auth';

export default function Navbar() {
  const router = useRouter();
  const pathname = usePathname();
  const [authenticated, setAuthenticated] = useState(false);
  const [user, setUserState] = useState<any>(null);
  const [scrolled, setScrolled] = useState(false);

  useEffect(() => {
    setAuthenticated(isAuthenticated());
    setUserState(getUser());

    const handleScroll = () => {
      setScrolled(window.scrollY > 10);
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, [pathname]);

  const handleLogout = () => {
    removeToken();
    setAuthenticated(authenticated); // Trigger re-render
    setAuthenticated(false);
    setUserState(null);
    router.push('/');
    router.refresh();
  };

  const isActive = (path: string) => pathname === path;

  return (
    <nav className={`fixed top-0 left-0 right-0 z-50 transition-all duration-300 ${scrolled ? 'glass border-b border-gray-100 py-3 shadow-sm' : 'bg-transparent py-5'}`}>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center">
          <Link href="/" className="flex items-center gap-2 group">
            <div className="w-10 h-10 bg-blue-600 rounded-xl flex items-center justify-center text-white font-black text-xl shadow-lg shadow-blue-200 group-hover:scale-105 transition-transform">
              G
            </div>
            <span className="text-xl font-black tracking-tighter text-gray-900 group-hover:text-blue-600 transition-colors">
              GET<span className="text-blue-600">ALL</span>
            </span>
          </Link>

          <div className="hidden md:flex items-center gap-8">
            <Link href="/" className={`text-sm font-bold tracking-wide transition-colors ${isActive('/') ? 'text-blue-600' : 'text-gray-500 hover:text-gray-900'}`}>
              MARKET
            </Link>
            <Link href="/items" className={`text-sm font-bold tracking-wide transition-colors ${isActive('/items') ? 'text-blue-600' : 'text-gray-500 hover:text-gray-900'}`}>
              ALL ITEMS
            </Link>
          </div>

          <div className="flex gap-4 items-center">
            {authenticated ? (
              <div className="flex items-center gap-6">
                <Link href="/dashboard" className={`text-sm font-bold tracking-wide transition-colors ${isActive('/dashboard') ? 'text-blue-600' : 'text-gray-500 hover:text-gray-900'}`}>
                  DASHBOARD
                </Link>
                <div className="flex items-center gap-3 pl-6 border-l border-gray-100">
                  <div className="text-right hidden sm:block">
                    <p className="text-xs font-black text-gray-900 leading-none mb-1 capitalize">{user?.username}</p>
                    <button onClick={handleLogout} className="text-[10px] font-bold text-gray-400 hover:text-red-500 transition-colors uppercase tracking-widest">
                      Sign Out
                    </button>
                  </div>
                  <div className="w-10 h-10 rounded-full bg-gray-100 border-2 border-white shadow-sm flex items-center justify-center text-gray-400 font-bold overflow-hidden">
                    {user?.username?.[0]?.toUpperCase() || 'U'}
                  </div>
                </div>
              </div>
            ) : (
              <div className="flex items-center gap-3">
                <Link href="/login" className="text-sm font-bold text-gray-500 hover:text-gray-900 px-4 py-2 transition-colors uppercase tracking-widest">
                  Login
                </Link>
                <Link
                  href="/register"
                  className="bg-gray-900 text-white text-xs font-black px-6 py-3 rounded-xl hover:bg-black transition-all shadow-lg shadow-gray-200 active:scale-95 uppercase tracking-widest"
                >
                  Join Now
                </Link>
              </div>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}


