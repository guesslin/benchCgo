#include "m.h"

int my_bench_sum(int x, int y) {
	return x + y;
}

int callback_bench(int times) {
	int i = 0;
	int sum = 0;
	for (i = 0; i < times; i++) {
		sum += my_go_sum(i, i+1);
	}
	return sum;
}

int non_callback_bench(int times) {
	int i = 0;
	int sum = 0;
	for (i = 0; i < times; i++) {
		sum += my_bench_sum(i, i+1);
	}
	return sum;
}

