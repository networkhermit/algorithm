<?php

declare(strict_types=1);

namespace muse\data_structures;

interface Lister
{
    public function size(): int;

    public function isEmpty(): bool;

    public function get(int $index): int;

    public function set(int $index, int $element): void;

    public function insert(int $index, int $element): void;

    public function remove(int $index): void;

    public function front(): int;

    public function back(): int;

    public function prepend(int $element): void;

    public function append(int $element): void;

    public function poll(): void;

    public function eject(): void;
}
