import axios, { AxiosError } from 'axios';
import { getToken, removeToken } from './auth';

export const BACKEND_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: BACKEND_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

/**
 * Transforms a relative storage path into a full public URL pointing to the backend.
 * Uses a remote placeholder if no image exists to avoid 404s.
 */
export const getImageUrl = (path: string | null | undefined): string => {
  if (!path) return 'https://placehold.co/600x400/e2e8f0/1e293b?text=No+Image';
  if (path.startsWith('http')) return path;
  
  // Normalize path: remove leading slash if exists to avoid double slashes
  const cleanPath = path.startsWith('/') ? path.slice(1) : path;
  return `${BACKEND_URL}/${cleanPath}`;
};

// Add token to requests
api.interceptors.request.use((config) => {
  const token = getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  // Don't set Content-Type for FormData, let browser set it with boundary
  if (config.data instanceof FormData) {
    delete config.headers['Content-Type'];
  }
  return config;
});

// Handle 401 errors (unauthorized)
api.interceptors.response.use(
  (response) => response,
  (error: AxiosError) => {
    if (error.response?.status === 401) {
      removeToken();
      if (typeof window !== 'undefined') {
        window.location.href = '/login';
      }
    }
    return Promise.reject(error);
  }
);

export default api;


