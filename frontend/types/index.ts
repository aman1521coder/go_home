export interface User {
  id: string;
  username: string;
  email: string;
  password?: string;
  created_at: string;
  updated_at: string;
  is_admin: boolean;
}

export interface ItemImage {
  id: string;
  item_id: string;
  image_path: string;
  display_order: number;
  created_at: string;
}

export interface Item {
  id: string;
  user_id: string;
  name: string;
  description: string;
  price: number;
  selling_price: number;
  image: string;
  images?: ItemImage[];
  quantity: number;
  created_at: string;
  updated_at: string;
  is_sold: boolean;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface RegisterData {
  username: string;
  email: string;
  password: string;
}


