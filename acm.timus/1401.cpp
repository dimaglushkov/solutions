/**
source:https://acm.timus.ru/problem.aspx?num=1401
Пояснение к примененному алгоритму:
Поле необходимо заполнить фигурами из 3 точек
Для любого n число 2^n * 2^n - 1 будет кратно 3, значит всегда можно будет расположить фигуры из 3 в квадрате со стороной 2^n таким образом, что останется единственная свободная точка. Следовательно, задача решаема для любого набора n, x, y
Так же следует заметить, что в поле 4*4 пустая точка может быть расположена в произвольном месте. Пример:
(файл 1401_DESC_1.png)
и так далее.
Рассмотрим площадь со стороной 8.
Разобьем площадь на 4 четверти - (1)..(4), каждая со стороной 4.
Зная, что в области 4x4 всегда будет одна свободная точка, пусть четверть, в которой находится точка, описанная координатами x y, это четверть (1).
Тогда остальные 3 четверти будут попарно смежны, а также каждая из них также будет иметь свободную точку.
Так как четверти попарно смежны, заполним их таким образом, чтобы их свободная точка находилась в центральном 2x2 квадрате исходной площади 8x8 (как на рисунке)
(файл 1401_DESC_2.png)
Тогда становится очевидным, что достаточно будет просто заполнить последние 3 свободные точки дополнительной фигурой.
Аналогичная ситуация получится и с областями большего, чем 8x8 размера.
Таким образом, решение задачи сводится к рекурсивному делению площади на четверти и заполнению четвертей необходимым образом.

*/
#include <stdio.h>
#include <math.h>

//number for next figure
static int num = 0;

typedef struct
{
    int raw;
    int col;
} coord_t;

void create_even_figure(int ** matrix, coord_t point)
{

    coord_t is_even = {point.raw % 2, point.col % 2};

    if(is_even.raw && is_even.col)
        matrix[point.raw - 1][point.col - 1]
                = matrix[point.raw - 1][point.col]
                = matrix[point.raw][point.col - 1]
                = ++num;

    else if(!is_even.raw && is_even.col)
        matrix[point.raw + 1][point.col]
                = matrix[point.raw][point.col - 1]
                = matrix[point.raw + 1][point.col - 1]
                = ++num;

    else if(is_even.raw && !is_even.col)
        matrix[point.raw - 1][point.col]
                = matrix[point.raw][point.col + 1]
                = matrix[point.raw - 1][point.col + 1]
                = ++num;

    else if(!is_even.raw && !is_even.col)
        matrix[point.raw + 1][point.col + 1]
                = matrix[point.raw + 1][point.col]
                = matrix[point.raw][point.col + 1]
                = ++num;

}

void create_figure(int ** matrix, coord_t point, int new_side_size, const int type)
{
    switch(type)
    {
        case (1):
            matrix[point.raw + new_side_size - 1][point.col + new_side_size - 1]
                    = matrix[point.raw + new_side_size][point.col + new_side_size - 1] 
                    = matrix[point.raw + new_side_size - 1][point.col + new_side_size]
                    = ++num;
            break;

        case (2):
            matrix[point.raw + new_side_size][point.col + new_side_size] 
                    = matrix[point.raw + new_side_size][point.col + new_side_size - 1] 
                    = matrix[point.raw + new_side_size - 1][point.col + new_side_size] 
                    = ++num;
            break;

        case (3):
            matrix[point.raw + new_side_size][point.col + new_side_size] 
                    = matrix[point.raw + new_side_size][point.col + new_side_size - 1] 
                    = matrix[point.raw + new_side_size - 1][point.col + new_side_size - 1] 
                    = ++num;
            break;

        case (4):
            matrix[point.raw + new_side_size][point.col + new_side_size]                    
                    = matrix[point.raw + new_side_size - 1][point.col + new_side_size]
                    = matrix[point.raw + new_side_size - 1][point.col + new_side_size - 1] 
                    = ++num;
            break;
    }
}

void solve(int ** matrix, int side_size, coord_t start, coord_t point)
{

    if(side_size == 2)
    {
        create_even_figure(matrix, {start.raw + point.raw % 2, start.col + point.col % 2});
        return;
    }

    int new_side_size = side_size / 2;

    if(point.raw >= new_side_size && point.col >= new_side_size)
    {
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col + new_side_size}, {point.raw - new_side_size, point.col - new_side_size});
        solve(matrix, new_side_size, {start.raw, start.col + new_side_size}, {new_side_size - 1, 0});
        solve(matrix, new_side_size, start, {new_side_size - 1, new_side_size - 1});
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col}, {0, new_side_size - 1});
        create_figure(matrix, start, new_side_size, 1);
    }
    else if(point.raw < new_side_size && point.col < new_side_size)
    {
        solve(matrix, new_side_size, start, point);
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col + new_side_size}, {0, 0});
        solve(matrix, new_side_size, {start.raw, start.col + new_side_size}, {new_side_size - 1, 0});
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col}, {0, new_side_size - 1});
        create_figure(matrix, start, new_side_size, 2);
    }
    else if(point.raw < new_side_size && point.col >= new_side_size)
    {
        solve(matrix, new_side_size, {start.raw, start.col + new_side_size}, {point.raw, point.col - new_side_size});
        solve(matrix, new_side_size, start, {new_side_size - 1, new_side_size - 1});
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col + new_side_size}, {0, 0});
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col}, {0, new_side_size - 1});
        create_figure(matrix, start, new_side_size, 3);
    }
    else
    {
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col}, {point.raw - new_side_size, point.col});
        solve(matrix, new_side_size, start, {new_side_size - 1, new_side_size - 1});
        solve(matrix, new_side_size, {start.raw + new_side_size, start.col + new_side_size}, {0, 0});
        solve(matrix, new_side_size, {start.raw, start.col + new_side_size}, {new_side_size - 1, 0});
        create_figure(matrix, start, new_side_size, 4);
    }

}

int main()
{

    int ** matrix, side_size;
    int n;
    coord_t point;
    scanf("%d %d %d", &n, &point.raw, &point.col);
    side_size = (int) pow(2, n);

    matrix = new int * [side_size];
    for(int i = 0; i < side_size; i++)
        matrix[i] = new int[side_size];

    point.raw--;
    point.col--;
    matrix[point.raw][point.col] = 0;

    solve(matrix, side_size, {0, 0}, point);

    for(int i = 0; i < side_size; i++)
    {
        for(int j = 0; j < side_size; j++)
            printf("%2d ", matrix[i][j]);
        printf("\n");
    }
    return 0;

}