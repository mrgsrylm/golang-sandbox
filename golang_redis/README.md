# golang-redis

## Pengertian

Redis adalah sebuah sistem manajemen basis data (database management system - DBMS) open-source dan berkinerja tinggi yang dapat digunakan sebagai penyimpanan data berkinerja tinggi, cache, dan broker pesan. Redis berfokus pada kinerja tinggi dan mendukung berbagai struktur data, termasuk string, hash, list, set, sorted set dengan rentang query, bitmaps, hyperloglogs, dan geospatial indexes dengan query radius.

## Karakteristik utama

1. In-Memory Storage: Redis menyimpan data di dalam memori, yang membuatnya sangat cepat karena menghindari kebutuhan untuk membaca atau menulis ke disk secara berulang.

2. Data Structures: Redis mendukung berbagai jenis struktur data, memungkinkan pengembang untuk memilih struktur data yang paling sesuai dengan kebutuhan aplikasi.

3. Persistence Options: Meskipun Redis menyimpan data di dalam memori, Anda dapat mengkonfigurasi opsi persistensi untuk menyimpan snapshot periodik dari dataset ke disk atau menggunakan jurnal transaksi untuk merekam perubahan data.

4. Pub/Sub (Publish/Subscribe): Redis mendukung sistem publikasi dan berlangganan di mana klien dapat berlangganan ke saluran tertentu dan menerima pembaruan saat ada perubahan di saluran tersebut.

5. Replication: Redis mendukung replikasi data, memungkinkan untuk membuat salinan dari satu server Redis ke server lain untuk tujuan toleransi kesalahan atau meningkatkan kapasitas baca sistem.

6. Tingkat Konsistensi: Redis biasanya mengutamakan kinerja daripada konsistensi mutlak. Ini berarti, dalam beberapa kasus, Redis mungkin mengorbankan konsistensi untuk mempercepat operasi baca/tulis.

## Cara Kerja

1. In-Memory Database:
   - Redis menyimpan semua data di dalam memori utama (RAM) sistem, yang membuatnya sangat cepat.
   - Data diakses langsung dari memori, menghindari keterlambatan yang mungkin terjadi jika perlu membaca atau menulis ke disk.
2. Key-Value Store:
   - Redis adalah database key-value, di mana setiap data disimpan sebagai pasangan key-value.
   - Key dapat berupa string atau beberapa tipe data lainnya, dan nilainya dapat berupa berbagai jenis struktur data (string, hash, list, set, dll.).
3. Struktur Data:
   - Redis mendukung berbagai struktur data seperti string, hash, list, set, sorted set, bitmaps, hyperloglogs, dan geospatial indexes.
   - Pengguna dapat memilih struktur data yang paling sesuai dengan kebutuhan aplikasi.
4. Operasi Atomik:
    - Redis mendukung operasi atomik pada tingkat perintah tunggal, yang berarti operasi tertentu dieksekusi sepenuhnya atau tidak sama sekali. 
5. Pub/Sub (Publish/Subscribe):
    - Redis mendukung sistem publikasi dan berlangganan di mana klien dapat berlangganan ke saluran tertentu dan menerima pembaruan saat ada perubahan di saluran tersebut. 
6. Persistence Options:
    - Redis menyediakan opsi persistensi, yang memungkinkan pengguna untuk mengonfigurasi cara data disimpan pada disk, baik melalui snapshot periodik atau jurnal transaksi. 
7. Replication:
    - Redis mendukung replikasi, memungkinkan pembuatan salinan dari satu server Redis ke server lain untuk tujuan toleransi kesalahan atau meningkatkan kapasitas baca sistem. 
8. Konsistensi:
    - Redis biasanya mengutamakan kinerja daripada konsistensi mutlak. Ini berarti dalam beberapa kasus, Redis mungkin mengorbankan konsistensi untuk mempercepat operasi baca/tulis.

## Referensi

- [Official Documentation](https://redis.io/docs/get-started/data-store/)