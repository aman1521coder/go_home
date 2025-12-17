# Prime Auction Frontend

A modern Next.js frontend for the Prime Auction platform.

## Features

- 🔐 JWT Authentication (Login/Register)
- 📦 Item Management (Create, Read, Update, Delete)
- 👤 User Dashboard
- 🎨 Modern UI with Tailwind CSS
- 🔒 Security best practices

## Getting Started

### Prerequisites

- Node.js 18+ 
- npm or yarn

### Installation

1. Install dependencies:
```bash
npm install
```

2. Create `.env.local` file (already created):
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

3. Run the development server:
```bash
npm run dev
```

4. Open [http://localhost:3000](http://localhost:3000) in your browser.

## Project Structure

```
frontend/
├── app/                 # Next.js app directory
│   ├── login/          # Login page
│   ├── register/       # Registration page
│   ├── dashboard/      # User dashboard
│   ├── items/          # Item pages
│   └── page.tsx        # Home page
├── components/         # React components
│   ├── Navbar.tsx      # Navigation bar
│   └── ItemCard.tsx    # Item card component
├── lib/                # Utilities
│   ├── api.ts          # API client
│   └── auth.ts         # Auth utilities
└── types/              # TypeScript types
    └── index.ts        # Type definitions
```

## Security Features

- JWT token stored in localStorage
- Automatic token injection in API requests
- 401 error handling with auto-logout
- Security headers configured
- React Strict Mode enabled
- Input validation on forms

## API Integration

The frontend connects to the Go backend API running on `http://localhost:8080`.

### Available Routes

- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login user
- `GET /api/items` - Get all items
- `GET /api/items?id={id}` - Get item by ID
- `POST /api/items` - Create item (protected)
- `PUT /api/items?id={id}` - Update item (protected)
- `DELETE /api/items?id={id}` - Delete item (protected)

## Build for Production

```bash
npm run build
npm start
```

## Technologies Used

- **Next.js 16** - React framework
- **React 19** - UI library
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **Axios** - HTTP client
