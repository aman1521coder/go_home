'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import api from '@/lib/api';
import { Item, User } from '@/types';
import { getUser, isAuthenticated } from '@/lib/auth';
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
    <div className="container mx-auto px-4 py-8">
      <div className="mb-6">
        <h1 className="text-3xl font-bold mb-2">Dashboard</h1>
        <p className="text-gray-600">Welcome back, {user?.username}!</p>
      </div>

      <div className="mb-6">
        <Link
          href="/items/new"
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 inline-block"
        >
          Create New Item
        </Link>
      </div>

      <div>
        <h2 className="text-2xl font-semibold mb-4">My Items</h2>
        {items.length === 0 ? (
          <div className="text-center py-12 bg-white rounded-lg">
            <p className="text-gray-500 mb-4">You haven't created any items yet.</p>
            <Link
              href="/items/new"
              className="text-blue-600 hover:underline"
            >
              Create your first item
            </Link>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {items.map((item) => (
              <ItemCard key={item.id} item={item} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}


