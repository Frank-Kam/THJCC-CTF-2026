<?php
require_once 'config.php';

if (empty($_SESSION['username'])) {
    header('Location: index.php');
}


$token = bin2hex(random_bytes(8));
$flag = "THJCC{" . $token . "_U_kNoW-HOw-t0_uSe-EaR}";
?>
<!doctype html>
<html>
<head><meta charset="utf-8"><title>Admin Panel</title></head>
<body>
<p>System settings</p>
<p><?= htmlspecialchars($flag) ?></p>
</body>
</html>
