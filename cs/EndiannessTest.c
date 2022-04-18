#include <stdbool.h>
#include <stdio.h>

bool isLittleEndian(void) {
  int val = 0x00FF;

  unsigned char *p = (unsigned char *)&val;

  return p[0] == (unsigned char)0xFF;
}

int main(void) {
  printf("Byte Order: ");
  if (isLittleEndian()) {
    puts("Little Endian");
  } else {
    puts("Big Endian");
  }
}
