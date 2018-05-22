<?php
function b64encode($string) {
    $string = base64_encode($string);
    $string = strtr($string, '+/', '-_');
    return rtrim($string, '='); 
} 
 
function b64decode($string) {
    $string = strtr($string, '-_', '+/');
    $string = str_pad($string, strlen($string) % 4, '=', STR_PAD_RIGHT); 
    return base64_decode($string);
}

function signature( $message, $secret ) {
    $signatured =  hash_hmac('sha256', $message, $secret, true);
    return b64encode($signatured);
} 


$header = ["alg" => "HS256", "typ" => "JWT"];
$payload = [
    "iss" => "spr",
    "sub" => "1",
    "exp" => 1];
$header = json_encode($header);
$payload = json_encode($payload);
$header = b64encode($header);
$payload = b64encode($payload);

$message = $header.'.'.$payload;

$signatured = signature($message, 'secret');
echo $message.'.'.$signatured;