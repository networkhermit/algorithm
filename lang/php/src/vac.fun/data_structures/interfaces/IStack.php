<?php
    namespace vac\fun\data_structures\interfaces\IStack;

    interface IStack {
        public function size(): int;

        public function isEmpty(): bool;

        public function peek(): int;

        public function push(int $element): void;

        public function pop(): void;
    }
?>
