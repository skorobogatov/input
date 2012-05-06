
#include <stdio.h>

int scanverbatim(char *fmt)
{
	return scanf(fmt);
}

int scanint(char *fmt, int *res)
{
	return scanf(fmt, res);
}

int scanstring(char *fmt, char **res)
{
	int n = scanf(fmt, res);
	printf("scanned %s\n", *res);
	return n;
}
