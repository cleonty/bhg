#include <stdio.h>
#include <stdlib.h>

typedef struct {
  int size;
  char name[8];
} User;

typedef struct {
  int n;
  User users[];
} S;

int main(int argc, char *argv[]) {
  printf ("sizeof S = %d\n", sizeof (S));
  
  int n = 5;
  S *s = malloc(sizeof (S) + n * sizeof(User));
  s->n = n;
  for (int i = 0; i < n; i++) {
    s->users[i].size = i;
    sprintf (s->users[i].name, "user %d", i);
  }
  for (int i = 0; i < n; i++) {
    printf ("size %d name %s\n", s->users[i].size, s->users[i].name);
  }
  return 0;
}