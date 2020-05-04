<?php
    namespace fun\vac\data_structures\LinkedQueue;

    require_once "fun/vac/data_structures/interfaces/IQueue.php";

    use fun\vac\data_structures\interfaces\IQueue;

    class Node {

        public $data;
        public $next = NULL;

        public function __construct(int $element) {
            $this->data = $element;
        }
    }

    class LinkedQueue implements IQueue\IQueue {

        private $head = NULL;
        private $tail = NULL;
        private $length = 0;

        public function __construct() {}

        public function size(): int {
            return $this->length;
        }

        public function isEmpty(): bool {
            return $this->length == 0;
        }

        public function peek(): int {
            if ($this->length == 0) {
                throw new RuntimeException("[PANIC - NoSuchElement]");
            }

            return $this->head->data;
        }

        public function enqueue(int $element): void {
            $node = new Node($element);

            if ($this->length == 0) {
                $this->head = $node;
            } else {
                $this->tail->next = $node;
            }

            $this->tail = $node;

            $this->length++;
        }

        public function dequeue(): void {
            if ($this->length == 0) {
                throw new RuntimeException("[PANIC - NoSuchElement]");
            }

            $target = $this->head;

            if ($this->length == 1) {
                $this->head = NULL;
                $this->tail = NULL;
            } else {
                $this->head = $target->next;
            }

            $target->data = NULL;

            $this->length--;
        }
    }
?>
