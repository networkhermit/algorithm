<?php
    require_once "fun/vac/util/RandomFactory.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\util\RandomFactory;
    use fun\vac\util\TestRunner;

    function testGenerateEven(): bool {
        RandomFactory\launch();

        $value = 0;
        for ($i = 0; $i < 8192; $i++) {
            $value = RandomFactory\generateEven();
            if (($value & 1) != 0) {
                return false;
            }
        }

        return true;
    }

    function testGenerateOdd(): bool {
        RandomFactory\launch();

        $value = 0;
        for ($i = 0; $i < 8192; $i++) {
            $value = RandomFactory\generateOdd();
            if (($value & 1) == 0) {
                return false;
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {

        TestRunner\parseTest(testGenerateEven());

        TestRunner\parseTest(testGenerateOdd());
    }
?>
