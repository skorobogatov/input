
#include <stdio.h>

int scanverbatim(char *fmt)
{
	return scanf(fmt);
}

int scanint(char *fmt, int *res)
{
	return scanf(fmt, res);
}
