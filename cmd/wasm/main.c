#include <stdio.h>
#include <emscripten.h>
typedef int bool;
#define true 1
#define false 0


bool isPrime(int n) {
  for (int i = 2; i*i <= n; i++) {
    if (n % i == 0) {
      return false;
    }
  }
  return true;
}

#ifdef __cplusplus
extern "C" {
#endif

void EMSCRIPTEN_KEEPALIVE firstNPrimeC(int n) {
  if (n < 1) {
    printf("N should be >= 1\n");
  }

  for (int i = 2; n > 0; i++) {
    if (isPrime(i)) {
      printf("%d,",i);
      n--;
    }
  }
  printf("\n");
}

#ifdef __cplusplus
}
#endif