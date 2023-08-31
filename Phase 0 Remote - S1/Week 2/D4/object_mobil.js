// Literal Object
let mobil = {
    merk: 'Toyota',
    model: 'Avanza',
    tahun: 2019,
    warna: 'Putih',
    start: function () {
        console.log('Mobil dinyalakan');
    }, stop: function () {
        console.log('Mobil dimatikan');
    }, getInfoMobile: function () {
        console.log(`Mobil ${this.merk} ${this.model} tahun ${this.tahun} berwarna ${this.warna}`);
    }
}

// Pengambilan menggunakan notation
console.log(mobil.merk);

// Pengambilan menggunakan array notation
console.log(mobil[ 'merk' ]);

// Get Mobil Info
mobil.getInfoMobile();

// Object Constructor
function Mobil(merk, model, tahun) {
    this.merk = merk;
    this.model = model;
    this.tahun = tahun;
    this.warna = 'Putih';
    this.start = function () {
        console.log('Mobil dinyalakan');
    }
}

// Object Instance
let avanza = new Mobil('Toyota', 'Avanza', 2019);
let xenia = new Mobil('Toyota', 'Xenia', 2018);

// Start mobil avanza
avanza.start();
// Start mobil xenia
xenia.start();

// Pengambilan menggunakan notation
console.log(avanza);
console.log(xenia);

// Pengambilan menggunakan array notation
console.log(avanza[ 'merk' ]);
console.log(xenia[ 'merk' ]);