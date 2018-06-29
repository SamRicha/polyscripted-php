<?php
$i = 0;
do {
    echo "Hello, World: and Welcome to PolyScripting". "\n";
} while ($i > 0);

#echo this variable stays echo.
$echo = "The name of this variable is echo...";

echo $echo. "\nBut the name of the function calling it is a randomly generated string."


?>


<?php
$string = 'exploiting';
$name = 'a language';
$password = 'impossible';
$str = 'Polyscripting makes $string $name pretty much $password.';
echo $str. "\n";
eval("\$str = \"$str\";");
echo $str. "\n";
?>
