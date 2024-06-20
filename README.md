# Ecojourney Backend

Ecojourney Backend adalah aplikasi backend yang dirancang untuk mendukung platform Ecojourney, yang mencakup fitur pengenalan sampah, penyimpanan histori pengambilan gambar, serta pencarian artikel terkait dengan pengelolaan sampah.

## Fitur

1. **Waste Recognition**
   - Mengunggah gambar sampah dan mengklasifikasikan jenis sampah.
   - Menyimpan URL gambar dan jenis sampah yang diunggah ke Firebase Firestore.

2. **Waste History**
   - Melihat histori pengenalan sampah yang telah dilakukan sebelumnya.

3. **User Management**
   - Mengelola informasi pengguna dan mengunggah gambar profil pengguna ke Firebase Storage.

## Teknologi yang Digunakan

- Go (Golang)
- Firebase (Firestore, Storage, Authentication)
- Gin (HTTP web framework)
- News API (untuk pencarian artikel)

## Struktur Proyek

```bash
.
├── config
│   └── connection.go
├── controller
│   ├── user_controller.go
│   └── waste_recognition_controller.go
├── helper
│   ├── history_helper.go
│   └── article_helper.go
├── middleware
│   └── auth_middleware.go
├── models
│   └── user.go
├── routes
│   └── routes.go
├── main.go
└── README.md
```

# Konfigurasi
## Environment Variables
Buat file .env dan tambahkan variabel berikut:
```bash
GOOGLE_APPLICATION_CREDENTIALS=path/to/your/service-account.json
PORT=8080
GOOGLE_BUCKET_NAME=your-bucket-name
NEWS_API_KEY=your-news-api-key
```

## Firebase Setup
1. Buat project di Firebase Console.
2. Tambahkan file kredensial (service-account.json) ke root proyek Anda.
3. Aktifkan Firestore, Firebase Storage, dan Authentication.

## Instalasi
1. Clone repository ini:
```bash
git clone https://github.com/your-username/ecojourney-backend.git
cd ecojourney-backend
```
2. Instal dependensi:
```bash
go mod tidy
```
3. Jalankan server:
```bash
go run main.go
```

# API endpoint
Technical document spesification dapat anda lihat melalui tautan berikut:
```bash
ristek.link/technical-spesification-documentristek.link/technical-spesification-document
```
