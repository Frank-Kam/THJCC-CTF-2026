<?php
require_once 'config.php';

$err = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $user = $_POST['username'] ?? '';
    $pass = $_POST['password'] ?? '';
    $remember = !empty($_POST['remember']);

    if ($user === USERNAME && $pass === PASSWORD) {
        $_SESSION['username'] = $user;

        if ($remember) {
            setcookie(COOKIE_NAME, $user, time() + COOKIE_EXPIRE_DAYS*24*3600, '/');
        }

        header('Location: admin.php');
        exit;
    } else {
        $err = 'Wrong username or password';
    }
} else {
    if (!empty($_COOKIE[COOKIE_NAME])) {
        $_SESSION['username'] = $_COOKIE[COOKIE_NAME];
        header('Location: admin.php');
        exit;
    }
}
?>
<!doctype html>
<html>
<head><meta charset="utf-8"><title>👂️</title></head>
<body>
<h2>Log in</h2>
<?php if($err): ?><p style="color:red;"><?= htmlspecialchars($err) ?></p><?php endif; ?>
<form method="post">
  <label>Account: <input name="username" required></label><br>
  <label>Password: <input name="password" type="password" required></label><br>
  <button type="submit">login</button>
</form>
<h2>Source Code</h2>
<code>
<pre>
&lt;?php
require_once 'config.php';

if (empty($_SESSION['username'])) {
    header('Location: index.php');
}
?&gt;
&lt;!doctype html&gt;
&lt;html&gt;
&lt;head&gt;&lt;meta charset="utf-8"&gt;&lt;title&gt;Meow&lt;/title&gt;&lt;head&gt
Meow
</pre>
</code>
</body>
</html>