#include <stdio.h>
#include <cstdint>
//size equals n
//rocks[i] equals w_i

using namespace std;

static uint32_t min_diff = 100000;
static uint8_t size;
static uint32_t * rocks;

void find_min_diff( uint32_t sum_heap1 = 0, uint32_t sum_heap2 = 0, uint8_t i = 0 )
{
    if (i != size)
    {
        find_min_diff(sum_heap1 + rocks[i], sum_heap2, i + 1);
        find_min_diff(sum_heap1, sum_heap2 + rocks[i], i + 1);
    }
    else
        min_diff = min_diff < sum_heap1 - sum_heap2 ? min_diff : sum_heap1 - sum_heap2;
}

int main()
{

    scanf("%hhu", &size);
    rocks = new uint32_t[size];
    for (uint8_t i = 0; i < size; i++)
        scanf("%u", &rocks[i]);

    find_min_diff();
    printf("%d\n", min_diff);

    return 0;
}
