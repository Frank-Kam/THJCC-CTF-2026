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
<p>Attack the D Point :speaking_head:</p>
<img src="cat.png" alt="Cat">
</body>
</html>