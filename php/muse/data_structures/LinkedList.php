<?php

declare(strict_types=1);

namespace muse\data_structures\LinkedList;

require_once "muse/data_structures/interfaces/IList.php";

use muse\data_structures\interfaces\IList;

class Node
{
    public $data;
    public $next = null;

    public function __construct(int $element)
    {
        $this->data = $element;
    }
}

class LinkedList implements IList\IList
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

        if ($index == $this->length - 1) {
            $cursor = $this->tail;
        } else {
            $cursor = $this->head;
            for ($i = 0; $i < $index; $i++) {
                $cursor = $cursor->next;
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

        if ($index == $this->length - 1) {
            $cursor = $this->tail;
        } else {
            $cursor = $this->head;
            for ($i = 0; $i < $index; $i++) {
                $cursor = $cursor->next;
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
                } else {
                    $this->tail = $node;
                }
                $this->head = $node;
                break;
            case $this->length:
                $this->tail->next = $node;
                $this->tail = $node;
                break;
            default:
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
                $this->tail = null;
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
            if ($index == $this->length - 1) {
                $this->tail = $cursor;
            }
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
