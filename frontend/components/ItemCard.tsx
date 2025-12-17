'use client';

import Link from 'next/link';
import { Item } from '@/types';

interface ItemCardProps {
  item: Item;
}

export default function ItemCard({ item }: ItemCardProps) {
  return (
    <Link href={`/items/${item.id}`}>
      <div className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
        {item.image && (
          <img
            src={item.image || '/placeholder-image.jpg'}
            alt={item.name}
            className="w-full h-48 object-cover"
            onError={(e) => {
              (e.target as HTMLImageElement).src = '/placeholder-image.jpg';
            }}
          />
        )}
        <div className="p-4">
          <h3 className="text-xl font-semibold mb-2">{item.name}</h3>
          <p className="text-gray-600 text-sm mb-2 line-clamp-2">{item.description}</p>
          <div className="flex justify-between items-center">
            <div>
              <p className="text-gray-500 text-sm">Price</p>
              <p className="text-lg font-bold text-blue-600">${item.selling_price}</p>
            </div>
            <div className="text-right">
              <p className="text-gray-500 text-sm">Quantity</p>
              <p className="text-lg font-semibold">{item.quantity}</p>
            </div>
          </div>
          {item.is_sold && (
            <span className="inline-block mt-2 px-2 py-1 bg-red-100 text-red-800 text-xs rounded">
              Sold
            </span>
          )}
        </div>
      </div>
    </Link>
  );
}


