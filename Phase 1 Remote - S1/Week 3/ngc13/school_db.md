# Normalisasi Basis Data Sekolah

Diberikan tabel informasi siswa berikut:

| Student ID | Student Name | Course         | Course Instructor |
|------------|--------------|----------------|-------------------|
| 101        | John         | HTML           | Mrs. Ayu          |
| 102        | Jane         | CSS            | Mrs. Amel         |
| 103        | Mary         | Basic Go       | Mr. Aldi          |
| 104        | Sarah        | Handling Error | Mr. Safran        |
| 105        | David        | REST API       | Mr. Tugas         |

## Langkah 1: 1NF (First Normal Form)

Tabel awal sudah memenuhi 1NF karena setiap sel berisi satu nilai, tidak ada nilai berulang, dan semua atribut memiliki nama unik.

## Langkah 2: 2NF (Second Normal Form)

Dalam 2NF, kita memastikan bahwa semua atribut non-kunci sepenuhnya bergantung pada kunci utama. Saat ini, atribut "Course Instructor" tidak sepenuhnya tergantung pada kunci utama. Karena itu, kita memisahkan tabel menjadi dua tabel terpisah untuk mencapai 2NF.

**Tabel 1: StudentCourses**

| Student ID | Course         |
|------------|----------------|
| 101        | HTML           |
| 102        | CSS            |
| 103        | Basic Go       |
| 104        | Handling Error |
| 105        | REST API       |

**Tabel 2: CourseInstructors**

| Course         | Course Instructor |
|----------------|-------------------|
| HTML           | Mrs. Ayu          |
| CSS            | Mrs. Amel         |
| Basic Go       | Mr. Aldi          |
| Handling Error | Mr. Safran        |
| REST API       | Mr. Tugas         |

Dengan membagi tabel asal menjadi dua tabel, kita telah mencapai 2NF karena semua atribut dalam masing-masing tabel bergantung pada kunci utama (Primary Key).

## Langkah 3: 3NF (Third Normal Form)

Dalam 3NF, kita memastikan bahwa setiap atribut non-kunci tidak transitif tergantung pada kunci utama. Dalam kasus ini, tidak ada ketergantungan transitif yang perlu diatasi.

Kami telah mencapai 3NF dengan memisahkan tabel asli menjadi dua tabel yang lebih kecil sesuai dengan dependensi fungsional. Sekarang data lebih efisien dan sesuai dengan prinsip normalisasi.
