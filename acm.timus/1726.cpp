/**
source:https://acm.timus.ru/problem.aspx?num=1726
Пояснения к примененному алгоритму:
Сперва найдем количество путей для n домов. Очевидно, что это число будет равно сумме арифметической прогрессии
n-1, n-2, n-3, ..., n - (n - 1), 0
и будет определятся формулой k = n*(n + 1) / 2
Из условий, что между 2 домами выбирается кратчайший маршрут, при этом двигаться можно только вертикально или горизонтально, следует, что можно рассматривать отдельно участки, пройденные горизонтально, и участки, пройденные вертикально.
Рассмотрим горизонтальные участки. Сперва необходимо отсортировать отдельно все координаты.
Из набора координат домов выберем произвольный i-ый дом.
Так как набор отсортирован, горизонтальным путем этого дома воспользуется i (перед текущей включительно) и n - i (после текущей) раз. Значит вклад этого дома в сумму путей по горизонтали равен dx * i * (n-i)
Проведем подобные действия для координат всех домов, после чего поделим суммарный путь на количество уникальных путей.

*/
#include <iostream>
#include <algorithm>
#include <cstdint>
#include <list>

using namespace std;

int main()
{
    uint64_t n, * x, * y, sum = 0;

    cin >> n;
    x = new uint64_t[n];
    y = new uint64_t[n];

    for (uint64_t i = 0; i < n; i++)
        cin >> x[i] >> y[i];

    sort(x, x + n);
    sort(y, y + n);

    for (uint64_t i = 0; i < n - 1; i++)
        sum += (x[i + 1] - x[i]) * (i + 1) * (n - i - 1) + (y[i + 1] - y[i]) * (i + 1) * (n - i - 1);

    sum = 2 * sum / (n * (n - 1));
    cout << sum;

    delete[] x;
    delete[] y;
    return 0;
}