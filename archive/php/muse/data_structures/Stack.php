<?php

declare(strict_types=1);

namespace muse\data_structures;

interface Stack
{
    public function size(): int;

    public function isEmpty(): bool;

    public function peek(): int;

    public function push(int $element): void;

    public function pop(): void;
}
