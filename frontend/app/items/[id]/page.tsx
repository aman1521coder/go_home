'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import api from '@/lib/api';
import { Item } from '@/types';
import { isAuthenticated } from '@/lib/auth';

export default function ItemDetailPage() {
  const params = useParams();
  const router = useRouter();
  const [item, setItem] = useState<Item | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchItem();
  }, [params.id]);

  const fetchItem = async () => {
    try {
      const response = await api.get(`/api/items?id=${params.id}`);
      setItem(response.data);
    } catch (err: any) {
      setError('Item not found');
      console.error('Error fetching item:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!confirm('Are you sure you want to delete this item?')) return;

    try {
      await api.delete(`/api/items?id=${params.id}`);
      router.push('/dashboard');
    } catch (err) {
      alert('Failed to delete item');
    }
  };

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">Loading...</div>
      </div>
    );
  }

  if (error || !item) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center text-red-600">{error || 'Item not found'}</div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      <button
        onClick={() => router.back()}
        className="mb-4 text-blue-600 hover:underline"
      >
        ← Back
      </button>

      <div className="bg-white rounded-lg shadow-lg overflow-hidden">
        {item.image && (
          <img
            src={item.image}
            alt={item.name}
            className="w-full h-96 object-cover"
            onError={(e) => {
              (e.target as HTMLImageElement).src = '/placeholder-image.jpg';
            }}
          />
        )}
        <div className="p-6">
          <h1 className="text-3xl font-bold mb-4">{item.name}</h1>
          <p className="text-gray-700 mb-6">{item.description}</p>

          <div className="grid grid-cols-2 gap-4 mb-6">
            <div>
              <p className="text-gray-500 text-sm">Price</p>
              <p className="text-2xl font-bold text-blue-600">${item.selling_price}</p>
            </div>
            <div>
              <p className="text-gray-500 text-sm">Quantity</p>
              <p className="text-2xl font-semibold">{item.quantity}</p>
            </div>
          </div>

          {item.is_sold && (
            <div className="mb-4">
              <span className="px-3 py-1 bg-red-100 text-red-800 rounded">
                Sold
              </span>
            </div>
          )}

          {isAuthenticated() && (
            <div className="flex gap-4">
              <button
                onClick={() => router.push(`/items/${item.id}/edit`)}
                className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
              >
                Edit
              </button>
              <button
                onClick={handleDelete}
                className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
              >
                Delete
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}


