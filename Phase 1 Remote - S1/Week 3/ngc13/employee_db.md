# Normalisasi Basis Data Manajemen Karyawan

Diberikan tabel manajemen karyawan berikut:

| Employee ID | Name   | Department | Supervisor ID | Supervisor Name | Salary  |
|------------|--------|------------|---------------|-----------------|--------|
| 101        | John   | HR         | 201           | Mary            | 50000  |
| 102        | Jane   | Finance    | 202           | Sarah           | 55000  |
| 201        | Mary   | HR         | 203           | David           | 60000  |
| 202        | Sarah  | Finance    | 203           | David           | 65000  |
| 203        | David  | CEO        | NULL          | NULL            | 100000 |

## Langkah 1: 1NF (First Normal Form)

Tabel awal sudah memenuhi 1NF karena setiap sel berisi satu nilai, tidak ada nilai berulang, dan semua atribut memiliki nama unik.

## Langkah 2: 2NF (Second Normal Form)

Dalam 2NF, kita memastikan bahwa semua atribut non-kunci sepenuhnya bergantung pada kunci utama. Kunci utama dalam tabel ini adalah "Employee ID."

Atribut "Supervisor Name" tidak sepenuhnya tergantung pada "Employee ID," karena dapat ditentukan oleh "Supervisor ID." Oleh karena itu, kita akan memisahkan tabel menjadi dua tabel terpisah untuk mencapai 2NF.

**Tabel 1: Employees**

| Employee ID | Name   | Department | Supervisor ID | Salary  |
|------------|--------|------------|---------------|--------|
| 101        | John   | HR         | 201           | 50000  |
| 102        | Jane   | Finance    | 202           | 55000  |
| 201        | Mary   | HR         | 203           | 60000  |
| 202        | Sarah  | Finance    | 203           | 65000  |
| 203        | David  | CEO        | NULL          | 100000 |

**Tabel 2: Supervisors**

| Supervisor ID | Supervisor Name |
|---------------|-----------------|
| 201           | Mary            |
| 202           | Sarah           |
| 203           | David           |

Dengan membagi tabel asal menjadi dua tabel terpisah, kita telah mencapai 2NF karena semua atribut dalam masing-masing tabel bergantung pada kunci utama (Primary Key).

## Langkah 3: 3NF (Third Normal Form)

Dalam 3NF, kita memastikan bahwa setiap atribut non-kunci tidak transitif tergantung pada kunci utama.

Atribut "Department" tidak sepenuhnya tergantung pada "Employee ID," karena dapat ditentukan oleh "Department." Oleh karena itu, kita akan memisahkan atribut "Department" ke dalam tabel terpisah.

**Tabel 3: Departments**

| Department | Department Head |
|------------|-----------------|
| HR         | Mary            |
| Finance    | Sarah           |
| CEO        | David           |

Kami juga telah memindahkan atribut "Salary" ke tabel "Employees" karena informasi gaji bersifat pribadi dan harus diamankan.

Dengan langkah ini, kami telah mencapai 3NF dan mengamankan informasi gaji.

Sekarang, data lebih efisien dan sesuai dengan prinsip normalisasi.
