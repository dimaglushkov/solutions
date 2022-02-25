/**
source:https://acm.timus.ru/problem.aspx?num=1296
Пояснение к примененному алгоритму:
Из условий задачи следует, что для обеспечения максимальной суммы первый элемент в результирующей сумме должен быть положительным (так как если сумма n_i ... n_k элементов отрицательна, то мы можем взять следующий положительный элемент n_j и он, очевидно, будет больше. Следовательно, сумма, содержащая n_j будет больше, чем та же сумма, содержащая сумму n_i ... n_k )
Значит, если после n-ой итерации сумма < 0, тогда достаточно просто обнулить сумму.


*/
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