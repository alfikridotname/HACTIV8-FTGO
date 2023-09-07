<?php
require_once 'config/koneksi.php';

// Submit Biodata
if ($_POST) {
    $name           = isset($_POST['name']) ? $_POST['name'] : '';
    $role           = isset($_POST['role']) ? $_POST['role'] : '';
    $availability   = isset($_POST['availability']) ? $_POST['availability'] : '';
    $age            = isset($_POST['age']) ? $_POST['age'] : '';
    $lokasi         = isset($_POST['lokasi']) ? $_POST['lokasi'] : '';
    $years          = isset($_POST['years']) ? $_POST['years'] : '';
    $email          = isset($_POST['email']) ? $_POST['email'] : '';

    // Update Biodata
    $query = "UPDATE biodata SET nama = ?, role = ?, availability = ?, age = ?, lokasi = ?, year_experience = ?, email = ? WHERE id = 1";
    $stmt = $conn->prepare($query);
    $stmt->bind_param("sssisis", $name, $role, $availability, $age, $lokasi, $years, $email);
    $result = $stmt->execute();

    // Check Query
    if (!$result) {
        die("Query Error : " . $conn->error);
    } else {
        header("Location: index.php?sukses=true");
    }
}
