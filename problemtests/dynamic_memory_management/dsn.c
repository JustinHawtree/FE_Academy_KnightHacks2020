#include <stdlib.h>
#include <string.h>
#include <stdio.h>

char ** make_grocery_list (FILE *ifp, int numItems) {
  char buffer[128];
  char **list = NULL;
  int i;

  list = malloc(sizeof(char *) * numItems);

  for (i=0; i < numItems; i++) {
    fscanf(ifp, "%s", buffer);
    list[i] = malloc(sizeof(char) * (strlen(buffer) + 1));
    strcpy(list[i], buffer);
  }
  return list;
}

int main(void) {
  FILE *fp;
  fp = fopen("test.txt", "r");
  char **output = NULL;
  output = make_grocery_list(fp, 3);
  int i = 0;
  while (i < 3) {
    printf("%s\n", output[i]);
    i++;
  }
  fclose(fp);
}