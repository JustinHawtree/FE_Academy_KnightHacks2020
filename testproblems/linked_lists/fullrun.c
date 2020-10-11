#include <stdlib.h>
#include <stdio.h>
#include "linked.h"


int main(void) {
  int array[10] = {2, 3, 4, 5, 6, 7, 8, 9, 10, 11};
  queue* q = createQueue(array, 10);
  node* head =  q->head;
  while (head != NULL) {
    printf("%d\n", head->data);
    if (head->next == NULL)
      break;
    head = head->next;
  }
  printf("\n");
  while (head != NULL) {
    printf("%d\n", head->data);
    head = head->prev;
  }

  printf("\n");
  int dequeuedVal = dequeue(q);
  printf("Dequeued: %d\n", dequeuedVal);
  printf("\n");
  
  head =  q->head;
  while (head != NULL) {
    printf("%d\n", head->data);
    if (head->next == NULL)
      break;
    head = head->next;
  }
}