#include <stdio.h>
#include <cstdint>

int main()
{
    uint16_t N;
    int16_t * array;
    int32_t max_sum = 0, sum = 0;

    scanf("%hu", &N);
    array = new int16_t[N];
    for (uint16_t i = 0; i < N; i++)
        scanf("%hd", &array[i]);

    for (uint16_t i = 0; i < N; i++)
    {
        sum = array[i] + sum > 0 ? array[i] + sum : 0;
        if (sum > max_sum)
            max_sum = sum;
    }

    delete[] array;
    printf("%d", max_sum);
    return  0;
}