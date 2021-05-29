<?php
    namespace vac\fun\util\SequenceBuilder;

    require_once "vac.fun/util/RandomFactory.php";
    require_once "vac.fun/util/Sequences.php";

    use vac\fun\util\RandomFactory;
    use vac\fun\util\Sequences;

    function packIncreasing(array &$arr): void {
        RandomFactory\seed();
        $arr[0] = RandomFactory\integerN(1, 3);
        for ($i = 1, $length = count($arr); $i < $length; $i++) {
            $arr[$i] = $arr[$i - 1] + RandomFactory\integerN(1, 3);
        }
    }

    function packRandom(array &$arr): void {
        RandomFactory\seed();
        foreach ($arr as &$v) {
            $v = RandomFactory\generateInteger();
        }
    }

    function packDecreasing(array &$arr): void {
        packIncreasing($arr);
        Sequences.reverse($arr);
    }
?>
