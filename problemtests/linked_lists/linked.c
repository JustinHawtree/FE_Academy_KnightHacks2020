#include <stdlib.h>
#include <stdio.h>


typedef struct node {
  int data;
  struct node* next, *prev;
} node;

typedef struct queue {
  int size;
  struct node *head, *tail;
} queue;

int dequeue (queue *thisQ);

queue* createQueue(int* array, int size) {
  if (size == 0) {
    printf("Gave 0 amount to createQueue");
    return NULL;
  }

  queue* q = malloc(sizeof(queue));
  q->size = 0;
  node* head = malloc(sizeof(node));
  head->prev = NULL;
  head->next = NULL;
  head->data = array[0];
  q->head = head;
  q->size++;

  int i;
  for (i=1; i < size; i++) {
    head->next = malloc(sizeof(node));
    q->size++;
    head->next->data = array[i];
    head->next->prev = head;
    head->next->next = NULL;
    head = head->next;
  }
  q->tail = head;
  return q;
} 


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












