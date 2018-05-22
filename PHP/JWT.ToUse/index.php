<?php
require_once 'jwt.php';

$j = new jwt();

$j->setSecret('secret');
$j->setExpire(10);
$hash = $j->authorization(1);
echo $hash;
echo "<hr>";

if($j->verify($_GET['hash'])) {
    $payload = $j->getPayload();

    if($j->expired()) {
        echo $j->renew();
    } else {
        'its ok';
    }
}
