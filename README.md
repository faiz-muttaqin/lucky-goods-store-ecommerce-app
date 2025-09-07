
# LGS - Lucky Good Store

LGS (Lucky Good Store) adalah aplikasi web toko online yang dibangun menggunakan React, Vite, TypeScript, dan SWC. Proyek ini menggunakan pnpm sebagai package manager untuk instalasi dan manajemen dependensi yang lebih cepat dan efisien.

## Fitur Awal

- Inisialisasi aplikasi React dengan Vite dan TypeScript
- Hot Module Replacement (HMR) untuk pengembangan yang cepat
- Konfigurasi linting dengan ESLint
- Pengujian dengan Vitest

## Teknologi yang Digunakan

- [React](https://react.dev/) ^19
- [Vite](https://vitejs.dev/) ^7
- [TypeScript](https://www.typescriptlang.org/) ^5
- [SWC](https://swc.rs/) (melalui @vitejs/plugin-react-swc)
- [pnpm](https://pnpm.io/) sebagai package manager

## Instalasi & Menjalankan Aplikasi

1. **Clone repository**
   ```bash
   git clone https://github.com/faiz-muttaqin/lgs.git
   cd lgs
   ```

2. **Install dependencies**
   ```bash
   pnpm install
   ```

3. **Menjalankan aplikasi (development)**
   ```bash
   pnpm run dev
   ```

4. **Build aplikasi untuk produksi**
   Output build akan berada di folder `docs` (bukan `dist`).
   ```bash
   pnpm run build
   ```

5. **Preview hasil build**
   ```bash
   pnpm run preview
   ```

## Struktur Folder

- `src/` : Source code aplikasi
- `docs/` : Output build produksi
- `public/` : File statis

## Pengembangan

- **Linting**
  ```bash
  pnpm run lint
  ```
- **Type Checking**
  ```bash
  pnpm run typecheck
  ```
- **Testing**
  ```bash
  pnpm run test
  ```

## Konfigurasi Build ke Folder docs

Untuk mengubah output build ke folder `docs`, pastikan konfigurasi berikut ada di `vite.config.ts`:

```ts
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';

export default defineConfig({
  plugins: [react()],
  build: {
    outDir: 'docs',
  },
});
```

---
LGS - Lucky Good Store © 2025
