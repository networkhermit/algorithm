<?php
    declare(strict_types=1);

    if (count(debug_backtrace()) == 0) {
        print("When’s Weekend?\n");

        $today = idate("w");

        switch ($today) {
        case 0:
            print("Today.\n");
            break;
        case 6:
            print("Today.\n");
            break;
        case 5:
            print("Tomorrow.\n");
            break;
        case 4:
            print("In two days.\n");
            break;
        default:
            print("Too far away.\n");
        }
    }
?>
