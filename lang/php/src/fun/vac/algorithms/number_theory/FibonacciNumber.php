<?php
    namespace fun\vac\algorithms\number_theory\FibonacciNumber;

    function iterativeProcedure(int $n): int {
        $sign = 1;

        if ($n < 0) {
            if (($n & 1) == 0) {
                $sign = -1;
            }
            $n = -$n;
        }

        if ($n < 2) {
            return $n;
        }

        $prev = 0;
        $curr = 1;

        $next = 0;
        for ($i = 2; $i <= $n; $i++) {
            $next = $prev + $curr;
            $prev = $curr;
            $curr = $next;
        }

        return $sign * $curr;
    }

    function recursiveProcedure(int $n): int {
        if ($n < 0) {
            if (($n & 1) == 0) {
                return -recursiveProcedure(-$n);
            } else {
                return recursiveProcedure(-$n);
            }
        } else if ($n < 2) {
            return $n;
        } else {
            return recursiveProcedure($n - 2) + recursiveProcedure($n - 1);
        }
    }
?>
