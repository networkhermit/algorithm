<?php
    namespace vac\fun\algorithms\number_theory\Factorial;

    function iterativeProcedure(int $n): int {
        if ($n < 0) {
            return 0;
        }

        $result = 1;
        for ($i = 1; $i <= $n; $i++) {
            $result *= $i;
        }

        return $result;
    }

    function recursiveProcedure(int $n): int {
        if ($n < 0) {
            return 0;
        }

        if ($n == 0) {
            return 1;
        }
        return recursiveProcedure($n - 1) * $n;
    }
?>
