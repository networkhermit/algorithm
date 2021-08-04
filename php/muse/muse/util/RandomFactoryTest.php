<?php
    require_once "muse/util/RandomFactory.php";
    require_once "muse/util/TestRunner.php";

    use muse\util\RandomFactory;
    use muse\util\TestRunner;

    function testIntegerN(): bool {
        RandomFactory\seed();

        $value = 0;
        for ($i = 0; $i < 8192; $i++) {
            if (RandomFactory\integerN(0, 0) != 0) {
                return false;
            }

            if (RandomFactory\integerN(1, 1) != 1) {
                return false;
            }

            $value = RandomFactory\integerN(0, 1);
            if ($value < 0 || $value > 1) {
                return false;
            }

            $value = RandomFactory\integerN(100, 10000);
            if (RandomFactory\integerN($value, $value) != $value) {
                return false;
            }
            if ($value < 100 || $value > 10000) {
                return false;
            }
        }

        return true;
    }

    function testGenerateEven(): bool {
        RandomFactory\seed();

        for ($i = 0; $i < 8192; $i++) {
            if ((RandomFactory\generateEven() & 1) != 0) {
                return false;
            }
        }

        return true;
    }

    function testGenerateOdd(): bool {
        RandomFactory\seed();

        for ($i = 0; $i < 8192; $i++) {
            if ((RandomFactory\generateOdd() & 1) == 0) {
                return false;
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testIntegerN());

        TestRunner\parseTest(testGenerateEven());

        TestRunner\parseTest(testGenerateOdd());
    }
?>
