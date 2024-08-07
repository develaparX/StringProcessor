# String Processor

## Deskripsi

Fungsi `processStrings` adalah sebuah utilitas JavaScript yang memproses array of strings untuk menghasilkan string unik berdasarkan frekuensi karakter dan aturan pengurutan tertentu.

## Cara Kerja

### 1. Penggabungan String

- Semua string dalam array input digabungkan menjadi satu string tunggal.

### 2. Penghitungan Frekuensi

- Setiap karakter dalam string gabungan dihitung frekuensinya.
- Hasil penghitungan disimpan dalam objek, dengan karakter sebagai key dan frekuensi sebagai value.

### 3. Pembuatan Array Karakter

- Objek frekuensi dikonversi menjadi array of objects.
- Setiap objek dalam array ini memiliki properti `char` (karakter) dan `freq` (frekuensi).

### 4. Pengurutan

Array karakter diurutkan berdasarkan kriteria berikut:

1. Frekuensi tertinggi ke terendah.
2. Jika frekuensi sama:
   - Huruf besar didahulukan dari huruf kecil untuk karakter yang sama.
   - Jika karakter berbeda, huruf besar didahulukan.
   - Jika keduanya huruf besar atau keduanya huruf kecil, diurutkan berdasarkan abjad.

### 5. Pembuatan String Hasil

- String hasil dibuat dengan mengambil satu kemunculan dari setiap karakter unik.
- Menggunakan `Set` untuk melacak karakter yang sudah dilihat, menghindari duplikasi.

### 6. Pengembalian Hasil

- Fungsi mengembalikan string hasil akhir.

## Penggunaan

```javascript
let inputArray = ['Pendanaan', 'Terproteksi', 'Untuk', 'Dampak', 'Berarti']
let output = processStrings(inputArray)
console.log(output) // Output: "aenrktipBDPTUdmosu"
```

```

## Cara Menjalankan

1. Pastikan Node.js sudah terinstal di sistem Anda.
2. Simpan kode dalam file bernama `processStrings.js`.
3. Buka terminal atau command prompt.
4. Navigasikan ke direktori tempat file disimpan.
5. Jalankan perintah: `node processStrings.js`

## Catatan

- Fungsi ini case-sensitive, membedakan huruf besar dan kecil.
- Karakter non-alfabet juga diproses dan diurutkan sesuai aturan yang sama.
```
