#include <stdlib.h>
#include <stdio.h>
#include "linked.h"


int main(void) {
  int array[1] = {1};
  queue* q = createQueue(array, 1);
  int nullResult = dequeue(q);
  if (nullResult != 1) {
    printf("Failed Test Case: Return dequeued node data\n");
    return 1;
  }

  if (q->size != 0) {
    printf("Failed Test Case: Queue size did not get decreased by one, expected size: %d, current size: %d\n", 0,  q->size);
    return 1;
  }

  if (q->tail != NULL) {
    printf("Failed Test Case: When queue becomes empty tail node needs to point to NULL");
    return 1;
  }

  printf("Success!");
  return 0;
}











