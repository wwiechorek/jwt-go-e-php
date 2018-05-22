<?php
$pathIndex = '/index.php';

$extensionsFiles = [
    'png',
    'jpg',
    'jpeg',
    'gif',
    'ico',
    'svg',
    'eot',
    'ttf',
    'woff',
    'css',
    'js',
];

if (preg_match('/\.(?:' . implode('|', $extensionsFiles) . ')$/', $_SERVER["SCRIPT_NAME"]))
    return false;

include __DIR__ . $pathIndex;
