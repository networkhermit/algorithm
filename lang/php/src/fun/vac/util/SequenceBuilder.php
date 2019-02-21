<?php
    namespace fun\vac\util\SequenceBuilder;

    require_once "fun/vac/util/RandomFactory.php";
    require_once "fun/vac/util/Sequences.php";

    use fun\vac\util\RandomFactory;
    use fun\vac\util\Sequences;

    function packIncreasing(array &$arr): void {
        RandomFactory\launch();
        $arr[0] = RandomFactory\integerN(1, 4);
        for ($i = 1, $length = count($arr); $i < $length; $i++) {
            $arr[$i] = $arr[$i - 1] + RandomFactory\integerN(1, 4);
        }
    }

    function packRandom(array &$arr): void {
        RandomFactory\launch();
        foreach ($arr as &$v) {
            $v = RandomFactory\generateInteger();
        }
    }

    function packDecreasing(array &$arr): void {
        packIncreasing($arr);
        Sequences.reverse($arr);
    }
?>
