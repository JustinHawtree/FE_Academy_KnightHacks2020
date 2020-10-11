#include <stdlib.h>
#include <stdio.h>
#include "linked.h"


int main(void) {
  // Testing for NULL queue
  int nullResult = dequeue(NULL);
  if (nullResult != 0) {
    printf("Failed Test Case: When queue is NULL return 0");
    return 1;
  }
}












