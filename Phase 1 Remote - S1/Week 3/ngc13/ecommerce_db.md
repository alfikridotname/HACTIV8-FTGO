# Normalisasi Basis Data E-commerce

Diberikan tabel belanja online berikut:

| Order ID | Customer Name | Product       | Category      | Price  |
|----------|---------------|---------------|---------------|--------|
| 101      | John          | Laptop        | Electronics   | 800    |
| 102      | Jane          | Smartphone    | Electronics   | 500    |
| 103      | Mary          | T-Shirt       | Fashion       | 20     |
| 104      | Sarah         | Refrigerator  | Appliances    | 1000   |
| 105      | David         | Headphones    | Electronics   | 50     |

## Langkah 1: 1NF (First Normal Form)

Tabel awal sudah memenuhi 1NF karena setiap sel berisi satu nilai, tidak ada nilai berulang, dan semua atribut memiliki nama unik.

## Langkah 2: 2NF (Second Normal Form)

Dalam 2NF, kita memastikan bahwa semua atribut non-kunci sepenuhnya bergantung pada kunci utama. Saat ini, atribut "Category" tidak sepenuhnya tergantung pada kunci utama, karena kategori produk dapat ditentukan oleh nama produk. Kita akan memisahkan tabel menjadi dua tabel terpisah untuk mencapai 2NF.

**Tabel 1: Orders**

| Order ID | Customer Name |
|----------|---------------|
| 101      | John          |
| 102      | Jane          |
| 103      | Mary          |
| 104      | Sarah         |
| 105      | David         |

**Tabel 2: Products**

| Product       | Category      | Price  |
|---------------|---------------|--------|
| Laptop        | Electronics   | 800    |
| Smartphone    | Electronics   | 500    |
| T-Shirt       | Fashion       | 20     |
| Refrigerator  | Appliances    | 1000   |
| Headphones    | Electronics   | 50     |

Dengan membagi tabel asal menjadi dua tabel terpisah, kita telah mencapai 2NF karena semua atribut dalam masing-masing tabel bergantung pada kunci utama (Primary Key).

## Langkah 3: 3NF (Third Normal Form)

Dalam 3NF, kita memastikan bahwa setiap atribut non-kunci tidak transitif tergantung pada kunci utama. Dalam kasus ini, tidak ada ketergantungan transitif yang perlu diatasi.

Kami telah mencapai 3NF dengan memisahkan tabel asli menjadi dua tabel yang lebih kecil sesuai dengan dependensi fungsional. Sekarang data lebih efisien dan sesuai dengan prinsip normalisasi.
