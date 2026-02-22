<?php
declare(strict_types=1);
session_start();

define('USERNAME', bin2hex(random_bytes(32)));
define('PASSWORD', bin2hex(random_bytes(32)));

define('COOKIE_NAME', 'USER_ID');
define('COOKIE_EXPIRE_DAYS', 30);
