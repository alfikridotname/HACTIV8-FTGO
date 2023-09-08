## Persyaratan Basis Data Sistem Manajemen Perpustakaan Skenario

### Book Attributes

1. ISBN: Pengidentifikasi unik untuk setiap buku..
2. Title: Judul buku.
3. Author: Pengarang buku.
4. Publisher: Penerbit buku.
5. Year of Publication: Tahun terbit buku.

### Member Attributes

1. Member ID: Pengidentifikasi unik untuk setiap anggota perpustakaan.
2. First Name: Nama depan.
3. Last Name: Nama belakang.
4. Date of Birth: The birthdate of the member.
5. Address: The address of the member.

### Lending Transaction Attributes

1. Date of Lending: Tanggal pinjam.
2. Date of Return: Tanggal kembali.
3. Condition of the Book at Return: Kondisi buku saat dikembalikan ('Good,' 'Damaged,' 'Lost').

### Lending Rules

1. Setiap buku hanya dapat dipinjamkan kepada satu anggota dalam satu waktu.
2. Anggota dapat meminjam beberapa buku sekaligus, namun tidak lebih dari 5 buku sekaligus.

### Database Design Considerations

#### Tables

Berdasarkan persyaratan di atas, kami dapat menentukan tabel berikut di database kami:

**1. Books Table:**

 - List item
 - ISBN (Primary Key)
 - Title
 - Author
 - Publisher
 - Year of Publication
 - Availability Status (untuk melacak apakah buku tersebut saat ini tersedia untuk dipinjamkan)

**2. Members Table:**
- Member ID (Primary Key)
- First Name
- Last Name
- Date of Birth
- Address

**3. Lending Transactions Table:**
- Transaction ID (Primary Key)
- Member ID (Foreign Key referensi ke Members Table)
- ISBN (Foreign Key referensi ke Books Table)
- Date of Lending
- Date of Return
- Condition of the Book at Return

### Relationships

1. One-to-One: Setiap buku hanya dapat dipinjamkan kepada satu anggota dalam satu waktu.
2. One-to-Many: Setiap anggota dapat meminjam beberapa buku sekaligus, namun tidak lebih dari 5 buku sekaligus.

### Pertimbangan Ke Depannya

Seiring dengan berkembangnya sistem perpustakaan, penting untuk mempertimbangkan potensi kebutuhan di masa depan seperti reservasi buku, biaya keterlambatan, dan kategori buku. Persyaratan ini kemungkinan besar akan mengarah pada pembuatan tabel dan hubungan tambahan dalam database.

### Kesimpulan

Dengan mengikuti persyaratan ini dan merancang database yang sesuai, perpustakaan akan dapat mengelola buku, anggota, dan transaksi peminjamannya secara efisien, memastikan sistem manajemen perpustakaan yang lancar dan terorganisir.