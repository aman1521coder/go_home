'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import api from '@/lib/api';
import { Item, User } from '@/types';
import { getUser, isAuthenticated, isAdmin } from '@/lib/auth';
import ItemCard from '@/components/ItemCard';

export default function DashboardPage() {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!isAuthenticated()) {
      router.push('/login');
      return;
    }

    const currentUser = getUser();
    setUser(currentUser);
    fetchUserItems();
  }, [router]);

  const fetchUserItems = async () => {
    try {
      // Note: You'll need to add a route to get items by user ID
      const response = await api.get('/api/items');
      const allItems = response.data;
      // Filter items by current user (client-side for now)
      // In production, use a backend endpoint like /api/items/my-items
      const userItems = allItems.filter((item: Item) => item.user_id === user?.id);
      setItems(userItems);
    } catch (err) {
      console.error('Error fetching items:', err);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">Loading...</div>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-end gap-6 mb-12">
        <div>
          <h1 className="text-xs font-black text-blue-600 uppercase tracking-[0.3em] mb-4">User Hub</h1>
          <h2 className="text-4xl font-black text-gray-900 tracking-tighter uppercase">Market Dashboard</h2>
          <p className="text-gray-400 font-bold text-sm mt-2 uppercase tracking-widest">Welcome back, {user?.username}</p>
        </div>
        
        {isAdmin() && (
          <Link
            href="/items/new"
            className="bg-gray-900 text-white font-black px-8 py-4 rounded-2xl shadow-xl shadow-gray-200 hover:bg-black transition-all active:scale-95 text-xs uppercase tracking-widest"
          >
            Post New Item
          </Link>
        )}
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-16">
         <div className="bg-white border border-gray-100 p-8 rounded-[2rem] premium-shadow">
            <span className="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">Total Value</span>
            <span className="text-3xl font-black text-gray-900 tracking-tighter">${items.reduce((acc, item) => acc + (item.selling_price || 0), 0).toLocaleString()}</span>
         </div>
         <div className="bg-white border border-gray-100 p-8 rounded-[2rem] premium-shadow">
            <span className="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">My Listings</span>
            <span className="text-3xl font-black text-gray-900 tracking-tighter">{items.length} Items</span>
         </div>
         <div className="bg-white border border-gray-100 p-8 rounded-[2rem] premium-shadow">
            <span className="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">Account Type</span>
            <span className="text-3xl font-black text-blue-600 tracking-tighter">MEMBER</span>
         </div>
      </div>

      <div>
        <div className="flex items-center gap-4 mb-8">
           <h3 className="text-xl font-black text-gray-900 tracking-tight uppercase">My Items</h3>
           <div className="h-px bg-gray-100 flex-grow"></div>
        </div>
        
        {items.length === 0 ? (
          <div className="text-center py-24 bg-white rounded-[3rem] border border-dashed border-gray-200">
            <p className="text-gray-400 font-black text-sm uppercase tracking-[0.2em] mb-4">
              {isAdmin() ? 'Your gallery is currently empty' : 'No items listed yet'}
            </p>
            {isAdmin() && (
              <Link
                href="/items/new"
                className="text-blue-600 font-black text-xs uppercase tracking-widest hover:text-blue-700 transition-colors underline underline-offset-8"
              >
                Post your first item
              </Link>
            )}
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {items.map((item) => (
              <ItemCard key={item.id} item={item} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}


