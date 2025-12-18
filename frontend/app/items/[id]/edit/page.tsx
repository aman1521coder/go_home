'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import api from '@/lib/api';
import { isAuthenticated } from '@/lib/auth';
import { Item } from '@/types';

export default function EditItemPage() {
  const params = useParams();
  const router = useRouter();
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    price: '',
    selling_price: '',
    quantity: '1',
  });
  const [imageFiles, setImageFiles] = useState<File[]>([]);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);

  if (typeof window !== 'undefined' && !isAuthenticated()) {
    router.push('/login');
    return null;
  }

  useEffect(() => {
    fetchItem();
  }, [params.id]);

  const fetchItem = async () => {
    try {
      const response = await api.get(`/api/items/${params.id}`);
      const item = response.data;
      setFormData({
        name: item.name,
        description: item.description,
        price: item.price?.toString() || '0',
        selling_price: item.selling_price?.toString() || '0',
        quantity: item.quantity?.toString() || '1',
      });
    } catch (err: any) {
      setError('Failed to load item');
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setSaving(true);

    try {
      // Create FormData for multipart/form-data
      const formDataToSend = new FormData();
      formDataToSend.append('name', formData.name);
      formDataToSend.append('description', formData.description);
      formDataToSend.append('price', formData.price);
      formDataToSend.append('selling_price', formData.selling_price);
      formDataToSend.append('quantity', formData.quantity);

      // Append image files if new ones are selected
      if (imageFiles.length > 0) {
        imageFiles.forEach((file) => {
          formDataToSend.append('images', file);
        });
      }

      await api.put(`/api/items/${params.id}`, formDataToSend, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      router.push(`/items/${params.id}`);
    } catch (err: any) {
      const errorMessage = err.response?.data || err.message || 'Failed to update item';
      setError(typeof errorMessage === 'string' ? errorMessage : 'Failed to update item');
    } finally {
      setSaving(false);
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
    <div className="container mx-auto px-4 py-8 max-w-2xl">
      <h1 className="text-3xl font-bold mb-6">Edit Item</h1>

      <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow-lg p-6 space-y-4">
        {error && (
          <div className="bg-red-100 text-red-700 p-3 rounded">{error}</div>
        )}

        <div>
          <label htmlFor="name" className="block text-sm font-medium mb-1">
            Item Name *
          </label>
          <input
            id="name"
            type="text"
            required
            value={formData.name}
            onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label htmlFor="description" className="block text-sm font-medium mb-1">
            Description *
          </label>
          <textarea
            id="description"
            required
            value={formData.description}
            onChange={(e) => setFormData({ ...formData, description: e.target.value })}
            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            rows={4}
          />
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label htmlFor="price" className="block text-sm font-medium mb-1">
              Cost Price *
            </label>
            <input
              id="price"
              type="number"
              step="0.01"
              min="0"
              required
            value={formData.price}
            onChange={(e) => setFormData({ ...formData, price: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label htmlFor="selling_price" className="block text-sm font-medium mb-1">
              Selling Price *
            </label>
            <input
              id="selling_price"
              type="number"
              step="0.01"
              min="0"
              required
            value={formData.selling_price}
            onChange={(e) => setFormData({ ...formData, selling_price: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
        </div>

        <div>
          <label htmlFor="images" className="block text-sm font-medium mb-1">
            Update Images (Optional - leave empty to keep current images)
          </label>
          <input
            id="images"
            type="file"
            accept="image/*"
            multiple
            onChange={(e) => {
              const files = Array.from(e.target.files || []);
              if (files.length > 10) {
                setError('Maximum 10 images allowed');
                return;
              }
              // Validate each file
              for (const file of files) {
                if (file.size > 5 * 1024 * 1024) {
                  setError(`${file.name} exceeds 5MB limit`);
                  return;
                }
              }
              setImageFiles(files);
              setError(''); // Clear error if validation passes
            }}
            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          {imageFiles.length > 0 && (
            <div className="mt-2 grid grid-cols-4 gap-2">
              {imageFiles.map((file, index) => (
                <div key={index} className="relative">
                  <img
                    src={URL.createObjectURL(file)}
                    alt={`Preview ${index + 1}`}
                    className="w-full h-20 object-cover rounded"
                  />
                  <button
                    type="button"
                    onClick={() => setImageFiles(imageFiles.filter((_, i) => i !== index))}
                    className="absolute top-0 right-0 bg-red-500 text-white rounded-full w-5 h-5 text-xs flex items-center justify-center hover:bg-red-600"
                  >
                    ×
                  </button>
                </div>
              ))}
            </div>
          )}
          <p className="mt-1 text-xs text-gray-500">
            Select new images to replace existing ones. Leave empty to keep current images.
          </p>
        </div>

        <div>
          <label htmlFor="quantity" className="block text-sm font-medium mb-1">
            Quantity *
          </label>
          <input
            id="quantity"
            type="number"
            min="1"
            required
            value={formData.quantity}
            onChange={(e) => setFormData({ ...formData, quantity: e.target.value })}
            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div className="flex gap-4">
          <button
            type="submit"
            disabled={saving}
            className="flex-1 bg-blue-600 text-white py-2 rounded hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
          >
            {saving ? 'Saving...' : 'Save Changes'}
          </button>
          <button
            type="button"
            onClick={() => router.back()}
            className="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50 transition-colors"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}


