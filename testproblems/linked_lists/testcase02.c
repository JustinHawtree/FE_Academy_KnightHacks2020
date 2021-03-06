#include <stdlib.h>
#include <stdio.h>
#include "linked.h"


int main(void) {
  // Testing for empty queue
  queue* q = createQueue(NULL, 0);
  int nullResult = dequeue(q);
  if (nullResult != 0) {
    printf("Failed Test Case: When queue is empty return 0\n");
    return 1;
  }

  printf("Success!");
  return 0;
}