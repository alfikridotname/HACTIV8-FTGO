// Initialize object karyawan
let karyawan = {
    nama: "John Doe",
    posisi: "Staff",
    gaji: 5000000,
    tampilInfo: function () {
        document.getElementById('nama-karyawan').innerHTML = this.nama
        document.getElementById('posisi-karyawan').innerHTML = this.posisi
        document.getElementById('gaji-karyawan').innerHTML = this.gaji
    }
}

// Tampilkan data karyawan
karyawan.tampilInfo();


let ul = document.getElementById('detail-karyawan');
console.log(ul);
for (let prop in karyawan) {
    console.log(prop);
    console.log(karyawan[ prop ]);
    if (typeof karyawan[ prop ] !== 'function') {
        let li = document.createElement('li');
        li.innerHTML = prop + ': ' + karyawan[ prop ];
        ul.appendChild(li);
        console.log(li);
    }
}