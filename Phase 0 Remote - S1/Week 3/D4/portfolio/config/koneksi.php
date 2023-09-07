<?php
error_reporting(E_ALL);
ini_set('display_errors', 1);

$server = 'localhost';
$user = 'root';
$password = 'root';
$database = 'portfolio';

$conn = new mysqli($server, $user, $password, $database) or die(mysqli_error($conn));

if (!$conn) {
    echo "Koneksi gagal";
}
