/**
source:https://acm.timus.ru/problem.aspx?num=1155
Пояснение к примененному алгоритму:
Разделим куб ABC...H на 2 области: AHFC — область 1 и EBDG — область 2 и обозначим их суммы как S(1) и S(2).
Таким образом, доступные операции xx+ или xx- будут влиять на суммы значений узлов в обеих областях — увеличивать .
Следовательно, можно утверждать, что если S(1) не равно S(2), то задача не может быть решена.
Так же следует, что для удаления значений во всех узлах достаточно обнулить значения лишь одной из областей.
Введем понятие смежности - два узла непосредственно соединены линией на рисунке, и относительной смежности - два узла лежат на разных концах диагонали куба и соединены через 2 других узла, находящихся в разных областях. Понятие относительной смежности указывает, что мы можем изменять значения 2 относительно смежных узлов через связующие их узлы. 
Так как все смежные и относительно смежный узлы относительно избранного находятся в другой области, достаточно обнулить отдельно каждый узел отдельный узел области, а так же его смежные и относительно смежные узлы.


*/
#include <stdio.h>
#include <cstdint>

static uint16_t num_of_iterations = 0;

void iterate(uint8_t * first, uint8_t * second, char first_char, char second_char,  char operation)
{

    if (operation == '-')
    {
        (* first)--;
        (* second)--;
    }
    else
    {
        (* first)++;
        (* second)++;
    }

    printf("%c%c%c\n", first_char, second_char, operation);
    num_of_iterations++;
}

int main()
{

    uint8_t A, B, C, D, E, F, G, H;
    scanf("%hhu %hhu %hhu %hhu %hhu %hhu %hhu %hhu", &A, &B, &C, &D, &E, &F, &G, &H);

    if ( (A + H + F + C ) != (E + B + D + G) )
    {
        printf("IMPOSSIBLE\n");
        return 0;
    }

    while (( A + H + F + C ) > 0)
    {

        if (A > 0)
        {
            if (E > 0) iterate(&A, &E, 'E', 'A', '-');
            else if (B > 0) iterate(&A, &B, 'A', 'B', '-');
            else if (D > 0) iterate(&A, &D, 'A', 'D', '-');
            else if (G > 0)
            {
                iterate(&F, &B, 'F', 'B', '+');
                iterate(&A, &B, 'A', 'B', '-');
                iterate(&F, &G, 'F', 'G', '-');
            }
        }


        if (F > 0)
        {
            if (E > 0) iterate(&F, &E, 'E', 'F', '-');
            else if (B > 0) iterate(&F, &B, 'F', 'B', '-');
            else if (G > 0) iterate(&F, &G, 'F', 'G', '-');
            else if (D > 0)
            {
                iterate(&G, &C, 'G', 'C', '+');
                iterate(&D, &C, 'D', 'C', '-');
                iterate(&F, &G, 'F', 'G', '-');
            }
        }

        if (H > 0)
        {
            if (E > 0) iterate(&H, &E, 'E', 'H', '-');
            else if (D > 0) iterate(&H, &D, 'H', 'D', '-');
            else if (G > 0) iterate(&H, &G, 'H', 'G', '-');
            else if (B > 0)
            {
                iterate(&G, &C, 'G', 'C', '+');
                iterate(&B, &C, 'B', 'C', '-');
                iterate(&H, &G, 'H', 'G', '-');
            }
        }

        if (C > 0)
        {
            if (D > 0) iterate(&C, &D, 'C', 'D', '-');
            else if (B > 0) iterate(&C, &B, 'C', 'B', '-');
            else if (G > 0) iterate(&C, &G, 'C', 'G', '-');
            else if (E > 0)
            {
                iterate(&F, &B, 'F', 'B', '+');
                iterate(&E, &F, 'E', 'F', '-');
                iterate(&B, &C, 'B', 'C', '-');
            }
        }


        if (num_of_iterations > 1000)
        {
            printf("IMPOSSIBLE\n");
            return 0;
        }

    }


    return 0;
}

/*
 * array version
 * NOT FINISHED

#include <stdio.h>
#include <cstdint>


static uint16_t num_of_iterations = 0;
static uint8_t val[8];


void iterate(uint8_t first_id, uint8_t second_id, char operation)
{

    if (operation == '-')
    {
        val[first_id]--;
        val[second_id]--;
    }
    else
    {
        val[first_id]++;
        val[second_id]++;
    }

    printf("%c%c%c\n", first_id + 65, second_id + 65, operation);
    num_of_iterations++;
}


int main()
{

    //reading in needed order
    scanf("%hhu %hhu %hhu %hhu %hhu %hhu %hhu %hhu", &val[0], &val[7], &val[1], &val[6], &val[5], &val[2], &val[4], &val[3]);

    if ( (val[0] + val[1] + val[2] + val[3]) != (val[4] + val[5] + val[6] + val[7]) )
    {
        printf("IMPOSSIBLE\n");
        return 0;
    }

    while (val[0] + val[1] + val[2] + val[3] > 0)
    {
        for (uint8_t i = 0; i < 4; i++)
        {
            if (val[i] > 0)
            {
                for (uint8_t j = 4; j < 7; j++)
                {
                    if (val[j] > 0)

                        if (j!= i + 4)
                            iterate(i, j, '-');
                        else if (i < 2)
                        {
                            iterate();
                        }
                }
            }
        }
    }


    return 0;
}
 */
