<?php

declare(strict_types=1);

namespace muse\data_structures;

require_once "muse/data_structures/Stack.php";

use muse\data_structures\LinkedStack\Node;
use muse\data_structures\Stack;

class LinkedStack implements Stack
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

    public function peek(): int
    {
        if ($this->length == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        return $this->tail->data;
    }

    public function push(int $element): void
    {
        $node = new Node($element);

        if ($this->length == 0) {
            $this->head = $node;
        } else {
            $this->tail->next = $node;
        }

        $this->tail = $node;

        $this->length++;
    }

    public function pop(): void
    {
        if ($this->length == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        $target = $this->tail;

        if ($this->length == 1) {
            $this->head = null;
            $this->tail = null;
        } else {
            $cursor = $this->head;
            for ($i = 0, $bound = $this->length - 2; $i < $bound; $i++) {
                $cursor = $cursor->next;
            }
            $cursor->next = null;
            $this->tail = $cursor;
        }

        $target->data = null;

        $this->length--;
    }
}

namespace muse\data_structures\LinkedStack;

class Node
{
    public $data;
    public $next = null;

    public function __construct(int $element)
    {
        $this->data = $element;
    }
}
