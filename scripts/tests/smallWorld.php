<?php
$i = 0;
do {
    echo $i;
} while ($i > 0);
?>


<?php
$string = 'cup';
$name = 'coffee';
$str = 'This is a $string with my $name in it.';
echo $str. "\n";
eval("\$str = \"$str\";");
echo $str. "\n";
?>
