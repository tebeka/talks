#include <stdio.h>

int fact(int n) {
  int total = 1;
  while (--n > 0) {
    total *= n;
  }
  return total;
}

int main() { printf("%d\n", fact(10)); }
