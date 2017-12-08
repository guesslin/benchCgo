#ifndef __CGO_BENCH_MY_TEST__
#define __CGO_BENCH_MY_TEST__

extern int my_go_sum(int x, int y);
int my_bench_sum(int x, int y);
int callback_bench(int times);
int non_callback_bench(int times);

#endif
