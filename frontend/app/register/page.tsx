'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import api from '@/lib/api';
import { setToken, setUser } from '@/lib/auth';
import { RegisterData } from '@/types';

export default function RegisterPage() {
  const router = useRouter();
  const [formData, setFormData] = useState<RegisterData>({
    username: '',
    email: '',
    password: '',
  });
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      const response = await api.post('/api/auth/register', formData);
      setToken(response.data.token);
      setUser(response.data.user);
      router.push('/dashboard');
    } catch (err: any) {
      setError(err.response?.data || 'Registration failed. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-[80vh] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 hero-gradient">
      <div className="max-w-md w-full">
        <div className="bg-white rounded-[3rem] p-10 shadow-2xl shadow-blue-100 border border-gray-100 relative overflow-hidden">
          <div className="absolute top-0 left-0 w-full h-2 bg-blue-600"></div>
          
          <div className="text-center mb-10">
            <div className="w-16 h-16 bg-blue-50 rounded-2xl flex items-center justify-center text-blue-600 font-black text-2xl mx-auto mb-6 shadow-inner">
              G
            </div>
            <h2 className="text-3xl font-black text-gray-900 tracking-tighter uppercase">Join GetAll</h2>
            <p className="mt-3 text-sm font-bold text-gray-400 uppercase tracking-[0.2em]">
              Create your account
            </p>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            {error && (
              <div className="bg-red-50 border border-red-100 text-red-600 p-4 rounded-2xl text-xs font-bold uppercase tracking-widest text-center animate-shake">
                {error}
              </div>
            )}
            
            <div className="space-y-2">
              <label htmlFor="username" className="block text-[10px] font-black text-gray-400 uppercase tracking-widest ml-1">
                Username
              </label>
              <input
                id="username"
                type="text"
                required
                value={formData.username}
                onChange={(e) => setFormData({ ...formData, username: e.target.value })}
                className="w-full px-6 py-4 bg-gray-50 border border-gray-100 rounded-2xl focus:outline-none focus:ring-4 focus:ring-blue-50 focus:bg-white focus:border-blue-200 transition-all font-bold text-gray-900"
                placeholder="YourName"
              />
            </div>

            <div className="space-y-2">
              <label htmlFor="email" className="block text-[10px] font-black text-gray-400 uppercase tracking-widest ml-1">
                Email Address
              </label>
              <input
                id="email"
                type="email"
                required
                value={formData.email}
                onChange={(e) => setFormData({ ...formData, email: e.target.value })}
                className="w-full px-6 py-4 bg-gray-50 border border-gray-100 rounded-2xl focus:outline-none focus:ring-4 focus:ring-blue-50 focus:bg-white focus:border-blue-200 transition-all font-bold text-gray-900"
                placeholder="name@example.com"
              />
            </div>

            <div className="space-y-2">
              <label htmlFor="password" className="block text-[10px] font-black text-gray-400 uppercase tracking-widest ml-1">
                Password
              </label>
              <input
                id="password"
                type="password"
                required
                value={formData.password}
                onChange={(e) => setFormData({ ...formData, password: e.target.value })}
                className="w-full px-6 py-4 bg-gray-50 border border-gray-100 rounded-2xl focus:outline-none focus:ring-4 focus:ring-blue-50 focus:bg-white focus:border-blue-200 transition-all font-bold text-gray-900"
                placeholder="••••••••"
              />
            </div>

            <button
              type="submit"
              disabled={loading}
              className="w-full bg-gray-900 text-white font-black py-5 rounded-2xl shadow-xl shadow-gray-200 hover:bg-black transition-all active:scale-95 text-xs uppercase tracking-[0.2em] disabled:opacity-50"
            >
              {loading ? 'Processing...' : 'Create Account'}
            </button>
          </form>

          <div className="mt-10 text-center">
            <p className="text-[10px] font-bold text-gray-400 uppercase tracking-widest">
              Already have an account?{' '}
              <Link href="/login" className="text-blue-600 hover:text-blue-700 underline underline-offset-4 decoration-2">
                Sign In
              </Link>
            </p>
          </div>
        </div>
        
        <p className="mt-8 text-center text-[10px] font-black text-gray-300 uppercase tracking-[0.3em]">
          Secure Registration
        </p>
      </div>
    </div>
  );
}


