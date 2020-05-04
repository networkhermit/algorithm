<?php
    namespace fun\vac\data_structures\LinkedStack;

    require_once "fun/vac/data_structures/interfaces/IStack.php";

    use fun\vac\data_structures\interfaces\IStack;

    class Node {

        public $data;
        public $next = NULL;

        public function __construct(int $element) {
            $this->data = $element;
        }
    }

    class LinkedStack implements IStack\IStack {

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

            return $this->tail->data;
        }

        public function push(int $element): void {
            $node = new Node($element);

            if ($this->length == 0) {
                $this->head = $node;
            } else {
                $this->tail->next = $node;
            }

            $this->tail = $node;

            $this->length++;
        }

        public function pop(): void {
            if ($this->length == 0) {
                throw new RuntimeException("[PANIC - NoSuchElement]");
            }

            $target = $this->tail;

            if ($this->length == 1) {
                $this->head = NULL;
                $this->tail = NULL;
            } else {
                $cursor = $this->head;
                for ($i = 0, $bound = $this->length - 2; $i < $bound; $i++) {
                    $cursor = $cursor->next;
                }
                $cursor->next = NULL;
                $this->tail = $cursor;
            }

            $target->data = NULL;

            $this->length--;
        }
    }
?>
