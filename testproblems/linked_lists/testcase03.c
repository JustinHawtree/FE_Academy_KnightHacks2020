#include <stdlib.h>
#include <stdio.h>
#include "linked.h"


char* getInput() {
  return "Input: NULL <-> [5] <-> [4] <-> [3] <-> [2] <-> [1] <-> NULL";
}

int main(void) {
  int array[5] = {5, 4, 3, 2, 1};
  queue* q = createQueue(array, 5);
  int result = dequeue(q);
  if (result != 5) {
    printf("Failed Test Case: Return value should be the value of the head dequeued\n");
    return 1;
  }

  if (q->head == NULL || q->head->data != 4) {
    printf("Failed Test Case: Head did not get moved over\n");
    return 1;
  }

  if (q->head->prev != NULL) {
    printf("Failed Test Case: Head Previous is not NULL\n");
    return 1;
  }

  if (q->size != 4) {
    printf("Failed Test Case: Queue size did not get decreased by one, expected size: %d, current size: %d\n", 4,  q->size);
    return 1;
  }

  printf("Success!");
  return 0;
}