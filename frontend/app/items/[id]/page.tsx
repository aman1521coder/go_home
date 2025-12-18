'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import api, { getImageUrl } from '@/lib/api';
import { Item } from '@/types';
import { isAuthenticated, getUser } from '@/lib/auth';

export default function ItemDetailPage() {
  const params = useParams();
  const router = useRouter();
  const [item, setItem] = useState<Item | null>(null);
  const [loading, setLoading] = useState(true);
  const [activeImg, setActiveImg] = useState(0);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchItem();
  }, [params.id]);

  const fetchItem = async () => {
    try {
      const response = await api.get(`/api/items/${params.id}`);
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
      await api.delete(`/api/items/${params.id}`);
      router.push('/dashboard');
    } catch (err) {
      alert('Failed to delete item');
    }
  };

  if (loading) return <div className="flex items-center justify-center min-h-[60vh]"><div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-600"></div></div>;
  if (error || !item) return <div className="text-center py-20 text-gray-500 font-bold text-2xl tracking-widest uppercase">{error || 'Asset not found'}</div>;

  const images = item.images && item.images.length > 0 
    ? item.images.map(img => img.image_path)
    : [item.image];

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <button onClick={() => router.back()} className="mb-12 flex items-center text-[10px] font-black text-gray-400 hover:text-blue-600 transition-colors tracking-[0.3em] uppercase group">
        <span className="mr-3 text-lg group-hover:-translate-x-1 transition-transform">←</span> BACK TO GALLERY
      </button>

      <div className="lg:grid lg:grid-cols-2 lg:gap-x-20 lg:items-start">
        {/* Gallery */}
        <div className="flex flex-col">
          <div className="w-full aspect-square rounded-[3rem] overflow-hidden bg-white border border-gray-100 premium-shadow">
            <img
              src={getImageUrl(images[activeImg])}
              alt={item.name}
              className="w-full h-full object-contain p-12 transition-all duration-700"
            />
          </div>
          
          {images.length > 1 && (
            <div className="mt-8 grid grid-cols-4 gap-4">
              {images.map((img, idx) => (
                <button
                  key={idx}
                  onClick={() => setActiveImg(idx)}
                  className={`aspect-square rounded-2xl overflow-hidden border-2 transition-all duration-300 ${activeImg === idx ? 'border-blue-600 scale-95' : 'border-transparent opacity-40 hover:opacity-100'}`}
                >
                  <img src={getImageUrl(img)} className="w-full h-full object-cover" alt="" />
                </button>
              ))}
            </div>
          )}
        </div>

        {/* Info */}
        <div className="mt-12 lg:mt-0 flex flex-col">
          <div className="border-b border-gray-100 pb-10 mb-10">
            <div className="flex items-center gap-3 mb-6">
               <span className="px-3 py-1 bg-blue-50 text-blue-600 text-[10px] font-black uppercase tracking-widest rounded-lg">Available</span>
               <span className="text-[10px] font-bold text-gray-400 uppercase tracking-widest italic">Item ID: #{item.id?.slice(0,8) || 'N/A'}</span>
            </div>
            <h1 className="text-5xl md:text-6xl font-black text-gray-900 tracking-tighter mb-4 leading-tight uppercase">{item.name}</h1>
            <p className="text-gray-400 font-medium leading-relaxed max-w-lg">
               Browse item details and contact the seller to make a purchase.
            </p>
          </div>

          <div className="mb-12">
            <span className="text-gray-400 text-[10px] uppercase font-black tracking-widest block mb-2">Sale Price</span>
            <div className="flex items-baseline gap-3">
               <p className="text-6xl font-black text-gray-900 tracking-tighter">${item.selling_price?.toLocaleString()}</p>
               <span className="text-blue-600 font-bold text-sm">USD</span>
            </div>
          </div>

          <div className="space-y-10">
            <div className="bg-gray-50 rounded-[2rem] p-8 border border-gray-100">
              <h3 className="text-[10px] font-black text-gray-400 uppercase tracking-[0.2em] mb-6">Item Description</h3>
              <p className="text-gray-600 leading-relaxed font-medium whitespace-pre-line">
                {item.description || "The seller has not provided a specific description for this item."}
              </p>
            </div>

            <div className="grid grid-cols-2 gap-6">
              <div className="bg-white border border-gray-100 p-6 rounded-3xl shadow-sm">
                <span className="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">Quantity</span>
                <span className="text-2xl font-black text-gray-900 tracking-tight">{item.quantity} Units</span>
              </div>
              <div className="bg-white border border-gray-100 p-6 rounded-3xl shadow-sm">
                <span className="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">Availability</span>
                <span className={`text-2xl font-black tracking-tight ${item.is_sold ? 'text-red-500' : 'text-emerald-500'}`}>
                  {item.is_sold ? 'SOLD' : 'AVAILABLE'}
                </span>
              </div>
            </div>
          </div>

          <div className="mt-16 flex flex-col sm:flex-row gap-4">
            <button className="flex-grow bg-gray-900 hover:bg-black text-white font-black py-6 rounded-[2rem] shadow-2xl shadow-gray-200 active:scale-[0.98] transition-all uppercase tracking-widest text-sm">
              Buy Now
            </button>
            
            {isAuthenticated() && getUser()?.id === item.user_id && (
              <div className="flex gap-4">
                <button 
                  onClick={() => router.push(`/items/${item.id}/edit`)}
                  className="px-10 bg-white border border-gray-200 hover:bg-gray-50 text-gray-900 font-black rounded-[2rem] transition-all uppercase tracking-widest text-xs"
                >
                  Edit
                </button>
                <button 
                  onClick={handleDelete}
                  className="px-10 border-2 border-red-50 hover:bg-red-50 text-red-500 font-black rounded-[2rem] transition-all uppercase tracking-widest text-xs"
                >
                  Delete
                </button>
              </div>
            )}
          </div>
          
          <p className="mt-8 text-center sm:text-left text-[10px] font-bold text-gray-400 uppercase tracking-[0.2em]">
             GetAll Market &bull; Simple & Usable
          </p>
        </div>
      </div>
    </div>
  );
}


