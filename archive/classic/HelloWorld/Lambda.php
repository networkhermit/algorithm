<?php

declare(strict_types=1);

/*
 *  Just for Fun
 *
 *  Author:       vac
 *  Organization: STEM
 *  LICENCE:      MIT
 *  Copyright:    2014-2018
 */

// function declaration
function sayHello(): void
{
    // output some text
    print("hello, world\n");
}

// entry point of the program
if (count(debug_backtrace()) == 0) {
    // function invocation
    sayHello();
}
