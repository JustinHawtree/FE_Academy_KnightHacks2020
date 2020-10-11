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
    queue* q = malloc(sizeof(queue));
    q->size = 0;
    q->head = NULL;
    q->tail = NULL;
    return q;
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