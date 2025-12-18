'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import api from '@/lib/api';
import { Item } from '@/types';
import { isAuthenticated, isAdmin } from '@/lib/auth';
import ItemCard from '@/components/ItemCard';

export default function Home() {
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const response = await api.get('/api/items');
      setItems(Array.isArray(response.data) ? response.data : []);
    } catch (err: any) {
      setError('Failed to load items');
      console.error('Error fetching items:', err);
      setItems([]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="hero-gradient">
      {/* Hero Section */}
      <section className="relative pt-20 pb-32 overflow-hidden">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
          <div className="text-center max-w-3xl mx-auto">
            <div className="inline-flex items-center px-4 py-2 rounded-full bg-blue-50 border border-blue-100 mb-8">
              <span className="w-2 h-2 rounded-full bg-blue-600 animate-pulse mr-2"></span>
              <span className="text-[10px] font-black text-blue-700 uppercase tracking-widest">Community Marketplace</span>
            </div>
            <h1 className="text-6xl md:text-7xl font-black text-gray-900 tracking-tighter leading-[0.9] mb-8">
              SHOP FROM THE <br />
              <span className="text-blue-600">MARKET</span>
            </h1>
            <p className="text-xl text-gray-500 font-medium leading-relaxed mb-10 px-4">
              A simple and reliable way to buy and sell items within your community.
              List your products and start trading today.
            </p>
            <div className="flex flex-col sm:flex-row items-center justify-center gap-4">
              {isAdmin() ? (
                <Link href="/items/new" className="w-full sm:w-auto bg-gray-900 text-white font-black px-10 py-5 rounded-2xl shadow-2xl shadow-gray-200 hover:bg-black transition-all active:scale-95 text-sm uppercase tracking-widest text-center">
                  Post an Item
                </Link>
              ) : isAuthenticated() ? (
                <Link href="/dashboard" className="w-full sm:w-auto bg-gray-900 text-white font-black px-10 py-5 rounded-2xl shadow-2xl shadow-gray-200 hover:bg-black transition-all active:scale-95 text-sm uppercase tracking-widest text-center">
                  View Dashboard
                </Link>
              ) : (
                <Link href="/login" className="w-full sm:w-auto bg-gray-900 text-white font-black px-10 py-5 rounded-2xl shadow-2xl shadow-gray-200 hover:bg-black transition-all active:scale-95 text-sm uppercase tracking-widest text-center">
                  Get Started
                </Link>
              )}
            </div>
          </div>
        </div>

        <div className="absolute top-0 left-1/2 -translate-x-1/2 w-full h-full pointer-events-none opacity-40">
           <div className="absolute top-40 left-10 w-64 h-64 bg-blue-200 rounded-full blur-3xl"></div>
           <div className="absolute bottom-10 right-10 w-96 h-96 bg-indigo-100 rounded-full blur-3xl"></div>
        </div>
      </section>

      {/* Main Content */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="flex flex-col md:flex-row justify-between items-end gap-6 mb-12">
          <div>
            <h2 className="text-xs font-black text-blue-600 uppercase tracking-[0.3em] mb-4">Market Catalog</h2>
            <h3 className="text-4xl font-black text-gray-900 tracking-tight uppercase">Available Items</h3>
          </div>
          <div className="flex gap-2">
             <div className="px-6 py-3 bg-white border border-gray-100 rounded-xl text-xs font-bold text-gray-500 shadow-sm">
                All Categories
             </div>
             <div className="px-6 py-3 bg-white border border-gray-100 rounded-xl text-xs font-bold text-gray-500 shadow-sm">
                Ending Soon
             </div>
          </div>
        </div>

        {error && (
          <div className="bg-red-50 border border-red-100 text-red-600 p-6 rounded-3xl mb-8 font-bold text-center">
            {error}
          </div>
        )}

        {loading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {[1, 2, 3].map((i) => (
              <div key={i} className="bg-gray-100 rounded-3xl aspect-[4/5] animate-pulse"></div>
            ))}
          </div>
        ) : (
          <>
            {items.length === 0 ? (
              <div className="text-center py-32 bg-white rounded-[3rem] border border-dashed border-gray-200">
                <p className="text-gray-400 font-bold text-xl uppercase tracking-widest">No items found</p>
                <p className="text-gray-400 mt-2">Check back later for new listings</p>
              </div>
            ) : (
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 lg:gap-12">
                {items.map((item) => (
                  <ItemCard key={item.id} item={item} />
                ))}
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}
