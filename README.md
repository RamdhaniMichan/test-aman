## Fitur

- **Merchant** dapat membuat, mengupdate, menghapus, dan melihat produk.
- **Customer** dapat melihat dan membeli produk.
- Diskon 10% untuk transaksi di atas 50.000.
- Bebas ongkir untuk transaksi di atas 15.000.
- JWT Authentication.
- Semua response konsisten:  
  ```
  {
    "code": <http_status_code>,
    "message": "<message>",
    "data": { ... }
  }
  ```

---

## Struktur Project

- `src/domain` : Entity & repository interface
- `src/usecase` : Business logic
- `src/repository` : Implementasi repository (Postgres/GORM)
- `src/delivery/http` : Handler HTTP (Gin)
- `.env` : Konfigurasi environment
- `docker-compose.yaml` : Orkestrasi service Go & PostgreSQL
- `Dockerfile` : Build image Go yang kecil

---

## Menjalankan dengan Docker

1. **Clone repo https://github.com/RamdhaniMichan/test-aman.git**
2. **Buat file `.env`** (lihat contoh di bawah)
3. **Build & jalankan**
   ```sh
   docker-compose up --build
   ```
4. API berjalan di [http://localhost:8081](http://localhost:8081)

---

## Contoh `.env`

```
DB_HOST=db
DB_PORT=5432
DB_USER=testaman
DB_PASSWORD=testaman
DB_NAME=testaman
JWT_SECRET=your_jwt_secret
```

---

## Contoh Endpoint & CURL

### Register Merchant
```sh
curl -X POST http://localhost:8081/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Merchant 1","email":"merchant1@mail.com","password":"password123","role":"merchant"}'
```

### Login
```sh
curl -X POST http://localhost:8081/login \
  -H "Content-Type: application/json" \
  -d '{"email":"merchant1@mail.com","password":"password123"}'
```

### Get All Products
```sh
curl http://localhost:8081/products
```

### Protected Endpoint (contoh)
```sh
curl -X POST http://localhost:8081/products \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Produk A","price":20000,"stock":10,"description":"Deskripsi produk"}'
```