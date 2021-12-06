<?php
    declare(strict_types=1);

    namespace muse\data_structures\interfaces\IQueue;

    interface IQueue {
        public function size(): int;

        public function isEmpty(): bool;

        public function peek(): int;

        public function enqueue(int $element): void;

        public function dequeue(): void;
    }
?>
