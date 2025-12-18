'use client';

import Link from 'next/link';
import { Item } from '@/types';
import { getImageUrl } from '@/lib/api';

interface ItemCardProps {
  item: Item;
}

export default function ItemCard({ item }: ItemCardProps) {
  return (
    <Link href={`/items/${item.id}`} className="group">
      <div className="bg-white rounded-[2.5rem] border border-gray-100 premium-shadow hover:shadow-2xl hover:-translate-y-2 transition-all duration-500 overflow-hidden h-full flex flex-col p-3">
        <div className="relative aspect-[1/1] rounded-[2rem] bg-gray-50 overflow-hidden">
          <img
            src={getImageUrl(item.image)}
            alt={item.name}
            className="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700"
          />
          <div className="absolute top-4 right-4 bg-white/90 backdrop-blur-md px-4 py-2 rounded-2xl shadow-sm border border-white/50">
             <span className="text-[10px] font-black text-gray-900 uppercase tracking-widest leading-none block mb-1">Quantity</span>
             <span className="text-sm font-black text-blue-600 block leading-none">{item.quantity}</span>
          </div>
          {item.is_sold && (
            <div className="absolute inset-0 bg-black/40 backdrop-blur-[2px] flex items-center justify-center">
              <div className="bg-white text-black px-6 py-2 rounded-full text-xs font-black uppercase tracking-[0.2em] shadow-xl">
                Sold Out
              </div>
            </div>
          )}
        </div>
        
        <div className="p-6 flex flex-col flex-grow">
          <div className="flex justify-between items-start mb-3">
            <h3 className="text-xl font-black text-gray-900 group-hover:text-blue-600 transition-colors line-clamp-1 tracking-tight">
              {item.name}
            </h3>
          </div>
          <p className="text-gray-400 text-sm line-clamp-2 mb-8 font-medium leading-relaxed">
            {item.description || "No description provided for this exclusive asset."}
          </p>
          
          <div className="mt-auto flex items-center justify-between pt-6 border-t border-gray-50">
            <div>
              <span className="text-gray-400 text-[10px] uppercase font-black tracking-widest block mb-1">Price</span>
              <p className="text-2xl font-black text-gray-900 tracking-tighter">
                ${item.selling_price?.toLocaleString() || '0.00'}
              </p>
            </div>
            <div className="w-12 h-12 bg-gray-50 rounded-2xl flex items-center justify-center group-hover:bg-blue-600 group-hover:text-white transition-all duration-300 shadow-inner">
               <span className="text-xl font-bold">→</span>
            </div>
          </div>
        </div>
      </div>
    </Link>
  );
}


