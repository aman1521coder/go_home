'use client';

import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import { isAuthenticated, removeToken, getUser } from '@/lib/auth';

export default function Navbar() {
  const router = useRouter();
  const [authenticated, setAuthenticated] = useState(false);
  const [user, setUserState] = useState<any>(null);

  useEffect(() => {
    setAuthenticated(isAuthenticated());
    setUserState(getUser());
  }, []);

  const handleLogout = () => {
    removeToken();
    setAuthenticated(false);
    setUserState(null);
    router.push('/');
  };

  return (
    <nav className="bg-white shadow-lg">
      <div className="container mx-auto px-4">
        <div className="flex justify-between items-center py-4">
          <Link href="/" className="text-xl font-bold text-blue-600">
            Prime Auction
          </Link>
          <div className="flex gap-4 items-center">
            {authenticated ? (
              <>
                <span className="text-gray-700">Hello, {user?.username}</span>
                <Link href="/dashboard" className="text-gray-700 hover:text-blue-600">
                  Dashboard
                </Link>
                <button
                  onClick={handleLogout}
                  className="text-gray-700 hover:text-blue-600"
                >
                  Logout
                </button>
              </>
            ) : (
              <>
                <Link href="/login" className="text-gray-700 hover:text-blue-600">
                  Login
                </Link>
                <Link
                  href="/register"
                  className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
                >
                  Register
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}


