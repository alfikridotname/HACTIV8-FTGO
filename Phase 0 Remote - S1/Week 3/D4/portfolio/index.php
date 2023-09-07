<?php
require_once 'config/koneksi.php';
require_once 'data.php';
?>
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Portfolio</title>
    <link rel="stylesheet" href="style.css">
</head>

<body>
    <div class="wrapper">
        <nav class="navbar">
            <ul>
                <li>
                    <a href="home.html" target="contentFrame">HOME</a>
                </li>
                <li>
                    <a href="product.html" target="contentFrame">PRODUCT</a>
                </li>
                <li>
                    <a href="gallery.html" target="contentFrame">GALLERY</a>
                </li>
                <li>
                    <a href="blog.html" target="contentFrame">BLOG</a>
                </li>
                <li>
                    <a href="about.html" target="contentFrame">ABOUT</a>
                </li>
            </ul>
        </nav>
        <?php if (isset($_GET['sukses']) && $_GET['sukses'] == 'true') { ?>
            <div id="informasi" class="informasi">
                Biodata berhasil diupdate
            </div>
        <?php } ?>
        <div class="card">
            <div class="biodata-image">
                <img src="<?= $biodata->profile_photo; ?>" alt="Profile Photo" width="100px">
            </div>
            <div class="biodata-name">
                <h1 id="biodata-name"><?= $biodata->nama; ?></h1>
                <p id="biodata-position"><?= $biodata->role; ?></p>
                <a id="btn-kontak" class="btn btn-primary" href="">Kontak</a>
                <a class="btn btn-secondary" href="">Resume</a>
            </div>
            <div class="biodata-detail">
                <ul>
                    <li>
                        <span>Availability</span>
                        <span id="biodata-availability"><?= $biodata->availability; ?></span>
                    </li>
                    <li>
                        <span>Usia</span>
                        <span id="biodata-usia"><?= $biodata->age; ?></span>
                    </li>
                    <li>
                        <span>Lokasi</span>
                        <span id="biodata-lokasi"><?= $biodata->lokasi; ?></span>
                    </li>
                    <li>
                        <span>Pengalaman</span>
                        <span id="biodata-pengalaman"><?= $biodata->year_experience; ?></span>
                    </li>
                    <li>
                        <span>Email</span>
                        <span id="biodata-email"><?= $biodata->email; ?></span>
                    </li>
                </ul>
            </div>
        </div>
        <div id="form-input" class="form-group">
            <form action="proses.php" method="post">
                <div>
                    <label for="name" class="form-label">Nama</label>
                    <input type="text" name="name" id="name" value="<?= isset($biodata->nama) ? $biodata->nama : ''; ?>">
                </div>
                <div>
                    <label for="role">Role</label>
                    <input type="text" name="role" id="role" value="<?= isset($biodata->role) ? $biodata->role : ''; ?>">
                </div>
                <div>
                    <label for="availability">Availability</label>
                    <input type="text" name="availability" id="availability" value="<?= isset($biodata->availability) ? $biodata->availability : ''; ?>">
                </div>
                <div>
                    <label for="age">Age</label>
                    <input type="number" name="age" id="age" value="<?= isset($biodata->age) ? $biodata->age : ''; ?>">
                </div>
                <div>
                    <label for="lokasi">Lokasi</label>
                    <input type="text" name="lokasi" id="lokasi" value="<?= isset($biodata->lokasi) ? $biodata->lokasi : ''; ?>">
                </div>
                <div>
                    <label for="years">Years Experience</label>
                    <input type="number" name="years" id="years" value="<?= isset($biodata->year_experience) ? $biodata->year_experience : ''; ?>">
                </div>
                <div>
                    <label for="email">Email</label>
                    <input type="email" name="email" id="email" value="<?= isset($biodata->email) ? $biodata->email : ''; ?>">
                </div>
                <div>
                    <button type="submit" class="btn btn-primary" style="display: block;">Submit</button>
                </div>
            </form>
        </div>
    </div>
</body>

</html>