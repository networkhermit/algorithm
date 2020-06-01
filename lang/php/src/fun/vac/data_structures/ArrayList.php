<?php
    namespace fun\vac\data_structures\ArrayList;

    require_once "fun/vac/data_structures/interfaces/IList.php";

    use fun\vac\data_structures\interfaces\IList;

    class ArrayList implements IList\IList {
        private const DEFAULT_CAPACITY = 64;

        private $data = NULL;
        private $logicalSize = 0;
        private $physicalSize = self::DEFAULT_CAPACITY;

        public function __construct(int $physicalSize = 0) {
            if ($physicalSize > 1) {
                $this->physicalSize = $physicalSize;
            }
            $this->data = array_pad(array(), $this->physicalSize, 0);
        }

        public function size(): int {
            return $this->logicalSize;
        }

        public function isEmpty(): bool {
            return $this->logicalSize == 0;
        }

        public function get(int $index): int {
            if ($index < 0 || $index >= $this->logicalSize) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

            return $this->data[$index];
        }

        public function set(int $index, int $element): void {
            if ($index < 0 || $index >= $this->logicalSize) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

            $this->data[$index] = $element;
        }

        public function insert(int $index, int $element): void {
            if ($index < 0 || $index > $this->logicalSize) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

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

            for ($i = $this->logicalSize; $i > $index; $i--) {
                $this->data[$i] = $this->data[$i - 1];
            }

            $this->data[$index] = $element;

            $this->logicalSize++;
        }

        public function remove(int $index): void {
            if ($index < 0 || $index >= $this->logicalSize) {
                throw new RuntimeException("[PANIC - IndexOutOfBounds]");
            }

            for ($i = $index + 1; $i < $this->logicalSize; $i++) {
                $this->data[$i - 1] = $this->data[$i];
            }

            $this->logicalSize--;

            $this->data[$this->logicalSize] = NULL;
        }

        public function front(): int {
            return $this->get(0);
        }

        public function back(): int {
            return $this->get($this->logicalSize - 1);
        }

        public function prepend(int $element): void {
            $this->insert(0, $element);
        }

        public function append(int $element): void {
            $this->insert($this->logicalSize, $element);
        }

        public function poll(): void {
            $this->remove(0);
        }

        public function eject(): void {
            $this->remove($this->logicalSize - 1);
        }

        public function capacity(): int {
            return $this->physicalSize;
        }

        public function shrink(): void {
            $temp = array_pad(array(), $this->logicalSize, 0);

            for ($i = 0, $length = $this->logicalSize; $i < $length; $i++) {
                $temp[$i] = $this->data[$i];
            }

            $this->data = $temp;
            $this->physicalSize = $this->logicalSize;
        }
    }
?>
