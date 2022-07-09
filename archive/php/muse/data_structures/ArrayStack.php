<?php

declare(strict_types=1);

namespace muse\data_structures;

require_once "muse/data_structures/Stack.php";

use muse\data_structures\Stack;

class ArrayStack implements Stack
{
    private const DEFAULT_CAPACITY = 64;

    private $data = null;
    private $logicalSize = 0;
    private $physicalSize = self::DEFAULT_CAPACITY;

    public function __construct(int $physicalSize = 0)
    {
        if ($physicalSize > 1) {
            $this->physicalSize = $physicalSize;
        }
        $this->data = array_pad(array(), $this->physicalSize, 0);
    }

    public function size(): int
    {
        return $this->logicalSize;
    }

    public function isEmpty(): bool
    {
        return $this->logicalSize == 0;
    }

    public function peek(): int
    {
        if ($this->logicalSize == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        return $this->data[$this->logicalSize - 1];
    }

    public function push(int $element): void
    {
        if ($this->logicalSize == $this->physicalSize) {
            $newCapacity = self::DEFAULT_CAPACITY;

            if ($this->physicalSize >= self::DEFAULT_CAPACITY) {
                $newCapacity = $this->physicalSize + ($this->physicalSize >> 1);
            }

            $temp = array_pad(array(), $newCapacity, 0);

            for ($i = 0, $length = $this->logicalSize; $i < $length; $i++) {
                $temp[$i] = $this->data[$i];
            }

            $this->data = $temp;
            $this->physicalSize = $newCapacity;
        }

        $this->data[$this->logicalSize] = $element;

        $this->logicalSize++;
    }

    public function pop(): void
    {
        if ($this->logicalSize == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        $this->logicalSize--;

        $this->data[$this->logicalSize] = null;
    }

    public function capacity(): int
    {
        return $this->physicalSize;
    }

    public function shrink(): void
    {
        $temp = array_pad(array(), $this->logicalSize, 0);

        for ($i = 0, $length = $this->logicalSize; $i < $length; $i++) {
            $temp[$i] = $this->data[$i];
        }

        $this->data = $temp;
        $this->physicalSize = $this->logicalSize;
    }
}
