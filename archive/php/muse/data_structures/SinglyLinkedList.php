<?php

declare(strict_types=1);

namespace muse\data_structures;

require_once "muse/data_structures/List.php";

use muse\data_structures\Lister;
use muse\data_structures\SinglyLinkedList\Node;

class SinglyLinkedList implements Lister
{
    private $head = null;
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

        $cursor = $this->head;

        for ($i = 0; $i < $index; $i++) {
            $cursor = $cursor->next;
        }

        return $cursor->data;
    }

    public function set(int $index, int $element): void
    {
        if ($index < 0 || $index >= $this->length) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        $cursor = $this->head;

        for ($i = 0; $i < $index; $i++) {
            $cursor = $cursor->next;
        }

        $cursor->data = $element;
    }

    public function insert(int $index, int $element): void
    {
        if ($index < 0 || $index > $this->length) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        $node = new Node($element);

        if ($index == 0) {
            if ($this->length != 0) {
                $node->next = $this->head;
            }
            $this->head = $node;
        } else {
            $cursor = $this->head;
            for ($i = 0, $bound = $index - 1; $i < $bound; $i++) {
                $cursor = $cursor->next;
            }
            $node->next = $cursor->next;
            $cursor->next = $node;
        }

        $this->length++;
    }

    public function remove(int $index): void
    {
        if ($index < 0 || $index >= $this->length) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        $target = null;

        if ($index == 0) {
            $target = $this->head;
            if ($this->length == 1) {
                $this->head = null;
            } else {
                $this->head = $target->next;
            }
        } else {
            $cursor = $this->head;
            for ($i = 0, $bound = $index - 1; $i < $bound; $i++) {
                $cursor = $cursor->next;
            }
            $target = $cursor->next;
            $cursor->next = $target->next;
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

namespace muse\data_structures\SinglyLinkedList;

class Node
{
    public $data;
    public $next = null;

    public function __construct(int $element)
    {
        $this->data = $element;
    }
}
