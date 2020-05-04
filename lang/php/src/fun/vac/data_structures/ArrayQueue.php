<?php
    namespace fun\vac\data_structures\ArrayQueue;

    require_once "fun/vac/data_structures/interfaces/IQueue.php";

    use fun\vac\data_structures\interfaces\IQueue;

    class ArrayQueue implements IQueue\IQueue {

        private const DEFAULT_CAPACITY = 64;

        private $data = NULL;
        private $front = 0;
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

        public function peek(): int {
            if ($this->logicalSize == 0) {
                throw new RuntimeException("[PANIC - NoSuchElement]");
            }

            return $this->data[$this->front];
        }

        public function enqueue(int $element): void {
            if ($this->logicalSize == $this->physicalSize) {
                $newCapacity = self::DEFAULT_CAPACITY;

                if ($this->physicalSize >= self::DEFAULT_CAPACITY) {
                    $newCapacity = $this->physicalSize + ($this->physicalSize >> 1);
                }

                $temp = array_pad(array(), $newCapacity, 0);

                $cursor = $this->front;

                for ($i = 0, $length = $this->logicalSize; $i < $length; $i++) {
                    if ($cursor == $this->physicalSize) {
                        $cursor = 0;
                    }
                    $temp[$i] = $this->data[$cursor];
                    $cursor++;
                }

                $this->data = $temp;
                $this->front = 0;
                $this->physicalSize = $newCapacity;
            }

            $this->data[($this->front + $this->logicalSize) % $this->physicalSize] = $element;

            $this->logicalSize++;
        }

        public function dequeue(): void {
            if ($this->logicalSize == 0) {
                throw new RuntimeException("[PANIC - NoSuchElement]");
            }

            $this->data[$this->front] = NULL;

            $this->front = ($this->front + 1) % $this->physicalSize;

            $this->logicalSize--;
        }

        public function capacity(): int {
            return $this->physicalSize;
        }

        public function shrink(): void {
            $temp = array_pad(array(), $this->logicalSize, 0);

            $cursor = $this->front;

            for ($i = 0, $length = $this->logicalSize; $i < $length; $i++) {
                if ($cursor == $this->physicalSize) {
                    $cursor = 0;
                }
                $temp[$i] = $this->data[$cursor];
                $cursor++;
            }

            $this->data = $temp;
            $this->front = 0;
            $this->physicalSize = $this->logicalSize;
        }
    }
?>
