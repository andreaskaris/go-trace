#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <pthread.h>
#include <string.h>

#define NUM_THREADS 16

pthread_t tids[NUM_THREADS];
int n = NUM_THREADS;
int err;

void* method(void *arg) {
	int i = 0;
	while(1) {
		printf("%i\n", i);
		i++;
		sleep(1);
	}
}

int main()
{
  
    for(int i = 0; i < NUM_THREADS; i++) {
	err = pthread_create(&(tids[i]), NULL, &method, NULL);
	if(err != 0) {
		printf("\nCould not create thread :%s", strerror(err));
	}
    }

    while(1);
}
