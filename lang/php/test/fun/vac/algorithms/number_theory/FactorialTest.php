<?php
    require_once "fun/vac/algorithms/number_theory/Factorial.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\algorithms\number_theory\Factorial;
    use fun\vac\util\TestRunner;

    function testFactorial(): bool {
        $mapping = [
            [ 0,         1],
            [ 1,         1],
            [ 2,         2],
            [ 3,         6],
            [ 4,        24],
            [ 5,       120],
            [ 6,       720],
            [ 7,      5040],
            [ 8,     40320],
            [ 9,    362880],
            [10,   3628800],
            [11,  39916800],
            [12, 479001600],
        ];

        $instances = count($mapping);

        for ($i = 0; $i < $instances; $i++) {
            if (Factorial\iterativeProcedure($mapping[$i][0]) != $mapping[$i][1]) {
                return false;
            }
        }

        for ($i = 0; $i < $instances; $i++) {
            if (Factorial\recursiveProcedure($mapping[$i][0]) != $mapping[$i][1]) {
                return false;
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testFactorial());
    }
?>
