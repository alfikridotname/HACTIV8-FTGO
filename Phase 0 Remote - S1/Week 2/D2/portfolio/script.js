// Initialize 
const btnKontak = document.getElementById('btn-kontak');
const biodataName = document.getElementById('biodata-name');
const biodataPosition = document.getElementById('biodata-position');
const biodataAvailability = document.getElementById('biodata-availability');
const biodataUsia = document.getElementById('biodata-usia');
const biodataLokasi = document.getElementById('biodata-lokasi');
const biodataPengalaman = document.getElementById('biodata-pengalaman');
const biodataEmail = document.getElementById('biodata-email');
const formInput = document.getElementById('form-input');
const form = document.getElementsByTagName('form')[ 0 ];
const informasi = document.getElementById('informasi');

// Event Listener on load
window.addEventListener('load', () => {
    btnKontak.innerHTML = 'Edit';
    formInput.style.display = 'none';
});

// Value form
let bioname = document.getElementById('name');
let role = document.getElementById('role');
let availability = document.getElementById('availability');
let age = document.getElementById('age');
let lokasi = document.getElementById('lokasi');
let years = document.getElementById('years');
let email = document.getElementById('email');

// Event Listener on click on btnKontak
btnKontak.addEventListener('click', (e) => {
    e.preventDefault();
    if (btnKontak.innerHTML === 'Edit') {
        formInput.style.display = 'block';
        bioname.value = biodataName.innerHTML;
        role.value = biodataPosition.innerHTML;
        availability.value = biodataAvailability.innerHTML;
        age.value = biodataUsia.innerHTML;
        lokasi.value = biodataLokasi.innerHTML;
        years.value = biodataPengalaman.innerHTML;
        email.value = biodataEmail.innerHTML;
    } else {
        formInput.style.display = 'none';
    }
});

// Event Listener on submit form
formInput.addEventListener('submit', (e) => {
    e.preventDefault();

    // Mengubah isi biodata
    biodataName.innerHTML = document.getElementById('name').value;
    biodataPosition.innerHTML = document.getElementById('role').value;
    biodataAvailability.innerHTML = document.getElementById('availability').value;
    biodataUsia.innerHTML = document.getElementById('age').value;
    biodataLokasi.innerHTML = document.getElementById('lokasi').value;
    biodataPengalaman.innerHTML = document.getElementById('years').value;
    biodataEmail.innerHTML = document.getElementById('email').value;
    btnKontak.innerHTML = 'Edit';

    // Pesan berhasil
    informasi.style.display = 'block';
    informasi.innerHTML = 'Data berhasil diubah';

    // Informasi disembunyikan setelah 2 detik
    setTimeout(() => {
        informasi.style.display = 'none';
    }, 2000);

    // Form disembunyikan
    formInput.style.display = 'none';

    // Form dikosongkan
    form.reset();
});