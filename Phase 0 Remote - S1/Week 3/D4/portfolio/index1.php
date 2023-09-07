<?php

echo "<h1>Test debug </h1><br>";

$nama = "Alfikri";
$umur = 33;
$nilai = 80.5;

echo "Nama : " . $nama . "<br>";
echo "Umur : " . $umur . "<br>";
echo "Nilai : " . $nilai . "<br>";

echo "<hr>";

echo '<h1>Looping</h1>';
// Perulangan
for ($i = 0; $i < 5; $i++) {
    echo "Hello World! <br>";
}

echo "<hr>";

echo '<h1>Array</h1>';
// Array
$hari = array("Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu");
var_dump($hari);

echo "<br>";

echo "<hr>";

echo '<h1>Associative Array</h1>';
// Associative Array
$hari = array(
    "senin" => "Monday",
    "selasa" => "Tuesday",
    "rabu" => "Wednesday",
    "kamis" => "Thursday",
    "jumat" => "Friday",
    "sabtu" => "Saturday",
    "minggu" => "Sunday"
);

var_dump($hari);

// Percabangan
echo "<hr>";

echo '<h1>Conditional</h1>';

$x = 10;
if ($x < 20) {
    echo "Benar";
} else {
    echo "Salah";
}

echo "<hr>";
require 'config/koneksi.php';

echo $mantap;
