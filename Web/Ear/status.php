<?php
require_once 'config.php';

if (empty($_SESSION['username'])) {
    header('Location: index.php');
}
?>
<!doctype html>
<html>
<head><meta charset="utf-8"><title>Status page</title></head>
<body>
<p>Max Verstappen will win the 2025 World Championship</p>
<img src="max.png" alt="Max">
</body>
</html>