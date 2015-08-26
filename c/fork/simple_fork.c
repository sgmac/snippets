#include <unistd.h>
#include <stdio.h>
#include <sys/wait.h>

int main(int argc, char *argv[])
{
	pid_t p, w;
	int status;
	char *whoami = NULL;

	printf("\n");
	printf("%s is a process running since the moment you press enter\n", argv[0]);
	printf("if you run '%s &'  in the backgound you'll understand.\n\n\n", argv[0]);


	printf("First line before calling fork() only appears on the parent\n");
	printf("Let's create a child with p = fork()\n\n");
	p = fork(); /* From this line on, all the code is copied*/


	/* Am I parent  or child? */
	whoami = (p) ? "\e[0;31mPARENT\e[0m": "\e[0;32mCHILD\e[0m" ;
	/* Parent waits to collect its  child exit status. */
	w = wait(&status);

	printf("%s\n", whoami);
	if (p){
		/* This is only printed once in the parent */
	printf("This block is inside an if-clause with the value from p=fork()\n");
	printf("and thus it only appears on the parent.\n\n");
	printf("fork() call in the PARENT returns the child's pid %d\n", p);
	printf("getpid() call in the PARENT returns the PARENT's pid %d.\n", getpid());
	printf("\n");
	}
		/* These is printed in parent and child */
	printf("This line and below one, are printed in both, parent and child.\n\n");
	printf("fork() call returns %d on the %s.\n", p, whoami);

	/* IMPORTANT: If we do not wait the child process to gather
	 * the exit status of its child, chances are you will see
	 * the call to getpidd() return 1. How come? This is the
	 * PID of the 'init' process. It turns out the parent
	 * did not wait to the exit status and the process becomes 
	 * an orphan process.
	 */
	printf("getpid() call from %s returns %d PID\n", whoami, getpid());
	printf("\n");
	return 0;
}
