<?php
    declare(strict_types=1);

    if (count(debug_backtrace()) == 0) {
        while ($line = fgets(STDIN)) {
            $pair = preg_split("/ +/", trim($line));
            $a = $pair[0];
            $b = $pair[1];
            printf("%d\n", $a + $b);
        }
    }
?>
