<?php

declare(strict_types=1);

namespace muse\data_structures\LinkedQueue;

require_once "muse/data_structures/interfaces/IQueue.php";

use muse\data_structures\interfaces\IQueue;

class Node
{
    public $data;
    public $next = null;

    public function __construct(int $element)
    {
        $this->data = $element;
    }
}

class LinkedQueue implements IQueue\IQueue
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

        return $this->head->data;
    }

    public function enqueue(int $element): void
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

    public function dequeue(): void
    {
        if ($this->length == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        $target = $this->head;

        if ($this->length == 1) {
            $this->head = null;
            $this->tail = null;
        } else {
            $this->head = $target->next;
        }

        $target->data = null;

        $this->length--;
    }
}
