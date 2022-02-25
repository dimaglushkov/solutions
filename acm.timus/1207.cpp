/**
source:https://acm.timus.ru/problem.aspx?num=1207
Пояснения к примененному алгоритму:
Для нахождения требуемых точек рассмотрим 4 точки на плоскости:
(рисунок 1207_DESC_1.png)
Проведем прямую между точками 1 и 3, проведем линии между точками, и построим прямоугольные треугольники, где данная линия будет являться гипотенузой
(рисунок 1207_DESC_2.png)
Из рисунка видно, что если взять точку 1 как первую из двух требумых точек, то для определния второй точки достаточно найти
медианный угол между гипотенузой и катетом b треугольника. Таким образом, через отношение сторон a b находим тангенс этого угла, и
затем выбираем точку с медианным котангенсом.

*/
#include <stdint.h>
#include <stdio.h>
#include <limits.h>
#include <float.h>
#include <algorithm>

typedef struct{
    int x, y;
    int id;
    double ctg;
} point_t;

static int SORT_RUN = 32;

void set_ctg(point_t * points, int N, int chosen_id)
{
    for (int i = 0; i < N; i++)
    {
        if (chosen_id == i)
        {
            points[i].ctg = -FLT_MAX;
            continue;
        }
        if (points[chosen_id].x == points[i].x)
        {
            points[i].ctg = FLT_MAX;
            continue;
        }

        points[i].ctg = (points[i].y - points[chosen_id].y) / (double)(points[i].x - points[chosen_id].x);
    }

}


void insert_sort(point_t * points, int left, int right)
{
    for (int i = left + 1; i <= right; i++)
    {
        point_t temp = points[i];
        int j = i - 1;
        while (points[j].ctg > temp.ctg && j >= left)
        {
            points[j+1] = points[j];
            j--;
        }
        points[j+1] = temp;

    }
}

void merge(point_t * points, int  left, int  mid, int  right)
{
    int  left_len = mid - left + 1, right_len = right - mid;
    point_t * left_points = new point_t[left_len], * right_points = new point_t[right_len];

    for (int  i = 0; i < left_len; i++)
        left_points[i] = points[left + i];
    for (int i = 0; i < right_len; i++)
        right_points[i] = points[mid + 1 + i];

    int  i = 0;
    int  j = 0;
    int  k = left;

    while (i < left_len && j < right_len)
    {
        if (left_points[i].ctg <= right_points[j].ctg)
        {
            points[k] = left_points[i];
            i++;
        }
        else
        {
            points[k] = right_points[j];
            j++;
        }
        k++;
    }

    while (i < left_len)
    {
        points[k] = left_points[i];
        k++;
        i++;
    }

    while (j < right_len)
    {
        points[k] = right_points[j];
        k++;
        j++;
    }
    delete[] left_points;
    delete[] right_points;
}

void timsort(point_t * points, int n)
{
    for (int i = 0; i < n; i += SORT_RUN)
        insert_sort(points, i, (int)std::min((i + SORT_RUN - 1), (n - 1)));

    for (int size = SORT_RUN; size < n; size = 2 * size)
        for (int left = 0; left < n; left += 2*size)
            merge(points, left, left + size - 1, std::min((left + 2*size - 1), (n-1)));

}

int main() {
    point_t * points;
    int minY = INT_MAX, minX = INT_MAX;
    int N, chosen_id=0;
    scanf("%d", &N);
    points = new point_t[N];
    for (int i = 0; i < N; i++)
    {
        scanf("%d %d", &points[i].x, &points[i].y);
        points[i].id = i;

        if (( points[i].x == minX && points[i].y < minY ) || points[i].x < minX)
        {
            minX = points[i].x;
            minY = points[i].y;
            chosen_id = i;
        }
    }

    set_ctg(points, N, chosen_id);
    timsort(points, N);
    printf("%d %d", points[0].id + 1, points[N / 2].id + 1);
    delete[] points;
    return 0;
}
