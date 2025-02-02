<?php

for ($a = 0; $a < 1000 / 3; $a++){
    for ($b = $a + 1; $b < 1000 / 2; $b++){
        $c = 1000 - $a - $b;
        if ($a * $a + $b * $b == $c * $c){
            echo "a = $a, b = $b, c = $c\n";
            echo "Произведение " . ($a * $b * $c);
            exit();
        }
    }
}
?>
