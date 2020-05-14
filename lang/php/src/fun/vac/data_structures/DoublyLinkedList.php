<?php
    namespace fun\vac\data_structures\DoublyLinkedList;

    require_once "fun/vac/data_structures/interfaces/IList.php";

    use fun\vac\data_structures\interfaces\IList;

    class Node {

        public $data;
        public $next = NULL;
        public $prev = NULL;

        public function __construct(int $element) {
            $this->data = $element;
        }
    }

    class DoublyLinkedList implements IList\IList {

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

        public function get(int $index): int {
            if ($index < 0 || $index >= $this->length) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

            $cursor = NULL;

            if ($index < $this->length >> 1) {
                $cursor = $this->head;
                for ($i = 0; $i < $index; $i++) {
                    $cursor = $cursor->next;
                }
            } else {
                $cursor = $this->tail;
                for ($i = $this->length - 1; $i > $index; $i--) {
                    $cursor = $cursor->prev;
                }
            }

            return $cursor->data;
        }

        public function set(int $index, int $element): void {
            if ($index < 0 || $index >= $this->length) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

            $cursor = NULL;

            if ($index < $this->length >> 1) {
                $cursor = $this->head;
                for ($i = 0; $i < $index; $i++) {
                    $cursor = $cursor->next;
                }
            } else {
                $cursor = $this->tail;
                for ($i = $this->length - 1; $i > $index; $i--) {
                    $cursor = $cursor->prev;
                }
            }

            $cursor->data = $element;
        }

        public function insert(int $index, int $element): void {
            if ($index < 0 || $index > $this->length) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

            $node = new Node($element);

            switch ($index) {
            case 0:
                if ($this->length != 0) {
                    $node->next = $this->head;
                    $this->head->prev = $node;
                } else {
                    $this->tail = $node;
                }
                $this->head = $node;
                break;
            case $this->length:
                $node->prev = $this->tail;
                $this->tail->next = $node;
                $this->tail = $node;
                break;
            default:
                $cursor = NULL;
                if ($index < $this->length >> 1) {
                    $cursor = $this->head;
                    for ($i = 0; $i < $index; $i++) {
                        $cursor = $cursor->next;
                    }
                } else {
                    $cursor = $this->tail;
                    for ($i = $this->length - 1; $i > $index; $i--) {
                        $cursor = $cursor->prev;
                    }
                }
                $node->next = $cursor;
                $node->prev = $cursor->prev;
                $cursor->prev->next = $node;
                $cursor->prev = $node;
            }

            $this->length++;
        }

        public function remove(int $index): void {
            if ($index < 0 || $index >= $this->length) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

            $target = NULL;

            switch ($index) {
            case 0:
                $target = $this->head;
                if ($this->length == 1) {
                    $this->head = NULL;
                    $this->tail = NULL;
                } else {
                    $target->next->prev = NULL;
                    $this->head = $target->next;
                }
                break;
            case $this->length - 1:
                $target = $this->tail;
                $target->prev->next = NULL;
                $this->tail = $target->prev;
                break;
            default:
                if ($index < $this->length >> 1) {
                    $target = $this->head;
                    for ($i = 0; $i < $index; $i++) {
                        $target = $target->next;
                    }
                } else {
                    $target = $this->tail;
                    for ($i = $this->length - 1; $i > $index; $i--) {
                        $target = $target->prev;
                    }
                }
                $target->prev->next = $target->next;
                $target->next->prev = $target->prev;
            }

            $target->data = NULL;

            $this->length--;
        }

        public function front(): int {
            return $this->get(0);
        }

        public function back(): int {
            return $this->get($this->length - 1);
        }

        public function prepend(int $element): void {
            $this->insert(0, $element);
        }

        public function append(int $element): void {
            $this->insert($this->length, $element);
        }

        public function poll(): void {
            $this->remove(0);
        }

        public function eject(): void {
            $this->remove($this->length - 1);
        }
    }
?>
