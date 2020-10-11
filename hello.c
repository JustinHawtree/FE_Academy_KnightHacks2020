#include <stdio.h>
#include <stdlib.h>


int testy(void) {
  return NULL;
}

int main(void){
  int num = 5;
  int num2 = 6;
  int i;
  for(i=0; i < 10; i++) {
    printf("Tom\n");
  }
  printf("Number is: %d", testy());
  return 0;
}