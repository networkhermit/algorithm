<?php

declare(strict_types=1);

namespace muse\data_structures;

require_once "muse/data_structures/List.php";

use muse\data_structures\DoublyLinkedList\Node;
use muse\data_structures\Lister;

class DoublyLinkedList implements Lister
{
    private $head = null;
    private $tail = null;
    private $length = 0;

    public function __construct()
    {
    }

    public function size(): int
    {
        return $this->length;
    }

    public function isEmpty(): bool
    {
        return $this->length == 0;
    }

    public function get(int $index): int
    {
        if ($index < 0 || $index >= $this->length) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        $cursor = null;

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

    public function set(int $index, int $element): void
    {
        if ($index < 0 || $index >= $this->length) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        $cursor = null;

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

    public function insert(int $index, int $element): void
    {
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
                $cursor = null;
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

    public function remove(int $index): void
    {
        if ($index < 0 || $index >= $this->length) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        $target = null;

        switch ($index) {
            case 0:
                $target = $this->head;
                if ($this->length == 1) {
                    $this->head = null;
                    $this->tail = null;
                } else {
                    $target->next->prev = null;
                    $this->head = $target->next;
                }
                break;
            case $this->length - 1:
                $target = $this->tail;
                $target->prev->next = null;
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

        $target->data = null;

        $this->length--;
    }

    public function front(): int
    {
        return $this->get(0);
    }

    public function back(): int
    {
        return $this->get($this->length - 1);
    }

    public function prepend(int $element): void
    {
        $this->insert(0, $element);
    }

    public function append(int $element): void
    {
        $this->insert($this->length, $element);
    }

    public function poll(): void
    {
        $this->remove(0);
    }

    public function eject(): void
    {
        $this->remove($this->length - 1);
    }
}

namespace muse\data_structures\DoublyLinkedList;

class Node
{
    public $data;
    public $next = null;
    public $prev = null;

    public function __construct(int $element)
    {
        $this->data = $element;
    }
}
