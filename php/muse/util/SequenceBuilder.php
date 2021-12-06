<?php
    declare(strict_types=1);

    namespace muse\util\SequenceBuilder;

    require_once "muse/util/RandomFactory.php";
    require_once "muse/util/Sequences.php";

    use muse\util\RandomFactory;
    use muse\util\Sequences;

    function packIncreasing(array &$arr): void {
        $arr[0] = RandomFactory\genIntN(1, 3);
        for ($i = 1, $length = count($arr); $i < $length; $i++) {
            $arr[$i] = $arr[$i - 1] + RandomFactory\genIntN(1, 3);
        }
    }

    function packRandom(array &$arr): void {
        foreach ($arr as &$v) {
            $v = RandomFactory\genInt();
        }
    }

    function packDecreasing(array &$arr): void {
        packIncreasing($arr);
        Sequences.reverse($arr);
    }
?>
