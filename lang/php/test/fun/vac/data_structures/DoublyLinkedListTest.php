<?php
    require_once "fun/vac/data_structures/DoublyLinkedList.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\data_structures\DoublyLinkedList\DoublyLinkedList;
    use fun\vac\util\TestRunner;

    function testLinkedList(): bool {
        $size = 8192;

        $list = new DoublyLinkedList();

        for ($i = 1; $i <= $size; $i++) {
            $list->append($i);
        }

        if ($list->size() != $size) {
            return false;
        }

        for ($i = 0; $i < $size; $i++) {
            if ($list->get($i) != $i + 1) {
                return false;
            }
        }

        for ($i = 0; $i < $size; $i++) {
            $list->set($i, $size - $i);
        }

        for ($i = 0; $i < $size; $i++) {
            if ($list->get($i) != $size - $i) {
                return false;
            }
        }

        $x = 0;
        $y = 0;

        for ($i = 0, $j = $size - 1; $i < $j; $i++, $j--) {
            $x = $list->get($i);
            $y = $list->get($j);

            $list->remove($i);
            $list->insert($i, $y);
            $list->remove($j);
            $list->insert($j, $x);
        }

        for ($i = 0; $i < $size; $i++) {
            if ($list->get($i) != $i + 1) {
                return false;
            }
        }

        for ($i = $size; $i >= 1; $i--) {
            if ($list->back() != $i) {
                return false;
            } else {
                $list->eject();
            }
        }

        if (!$list->isEmpty()) {
            return false;
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testLinkedList());
    }
?>
