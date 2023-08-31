// Inisialisasi variabel
const btnTambahKaryawan = document.getElementById('btn-tambah-karyawan');
const formSection = document.getElementById('form-section');

// Event tambah karyawan
btnTambahKaryawan.addEventListener('click', () => {
    formSection.classList.add('show');
});

// Tampung data karyawan
let karyawanData = [];

// Karyawan object constructor
function karyawan(nama, posisi, gaji) {
    this.nama = nama;
    this.posisi = posisi;
    this.gaji = gaji;
}

// Form section submit event
formSection.addEventListener('submit', (e) => {
    e.preventDefault();

    // Ambil data dari form
    const nama = document.getElementById('nama').value;
    const posisi = document.getElementById('posisi').value;
    const gaji = document.getElementById('gaji').value;

    // Buat object karyawan
    const karyawanBaru = new karyawan(nama, posisi, gaji);

    // Tambahkan object karyawan ke array karyawanData
    karyawanData.push(karyawanBaru);

    // Render ulang table
    renderTable();

    // Reset form
    // formSection.reset();
});

// Render table
function renderTable() {
    // Ambil table body
    const tableBody = document.getElementById('table-body');
    tableBody.innerHTML = '';

    karyawanData.forEach((item, index) => {
        console.log(tableBody);
        tableBody.innerHTML += `
            <tr>
                <td>${index + 1}</td>
                <td>${item.nama}</td>
                <td>${item.posisi}</td>
                <td>${item.gaji}</td>
                <td>
                    <button>Edit</button>
                    <button>Hapus</button>
                </td>
            </tr>`
    });
}
