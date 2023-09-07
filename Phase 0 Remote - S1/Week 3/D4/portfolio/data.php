<?php
// Get Biodata
$query      = "SELECT * FROM biodata";
$result     = $conn->query($query);
$biodata    = $result->fetch_object();
