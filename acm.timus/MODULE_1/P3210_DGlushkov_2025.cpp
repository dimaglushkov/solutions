#include "stdio.h"
#include <cstdint>

int main()
{
    uint8_t t;
    uint16_t n, k, div, mod;
    scanf("%hhu", &t);
    for(uint8_t j = 0; j < t; j++)
    {
        scanf("%hu %hu", &n, &k);
        div = n / k;
        mod = n % k;
        printf("%d\n", mod * (mod - 1) / 2  * ( (div + 1) * (div + 1) )
        + (k - mod - 1) * (k - mod) / 2     * ( div * div )
        + (k - mod) * mod                   * ( div * div + 1 ) );
    }

    return 0;
}

