/**
source:https://acm.timus.ru/problem.aspx?num=1604
Пояснения к примененному алгоритму:
Будем хранить набор всех знаков как ограничение + количество.
Сортируем весь набор по количеству.
Найдем два максимальных по количеству знака. Если хотя осталось хотя бы по 1 знаку каждого типа, то последовательно выводим оба знака, иначе выводим только первый знак.
Повторяем, пока в наборе остался хотя бы один знак.

*Сортировать нужно для того, чтобы начинать вывод с тех знаков, которых больше.
*/
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

struct sign_t
{
    uint32_t id, num;

    bool operator< (const sign_t &b)
    {
        return num != b.num ? num < b.num : id < b.id;
    }

};


pair<size_t , size_t > find_two_maxes(vector<sign_t> & signs)
{
    uint32_t first = 0, second = 0;
    for (size_t i = 0; i < signs.size(); i++)
        if (signs[i].num >= signs[first].num)
        {
            second = first;
            first = i;
        }

    return {first, second};
}


int main()
{
    size_t k, sum = 0;
    vector<sign_t> signs;
    pair<size_t , size_t > maxes = {0, 1};

    cin >> k;
    signs = vector<sign_t>(k);

    for (size_t i = 0; i < k; i++)
    {
        signs[i].id = i + 1;
        cin >> signs[i].num;
        sum += signs[i].num;
    }

    sort(signs.begin(), signs.end());

    for (size_t i = 0; i < sum;)
    {

        maxes = find_two_maxes(signs);

        cout << signs[maxes.first].id << " ";
        signs[maxes.first].num--;
        i++;

        if(signs[maxes.second].num > 0 && maxes.second != maxes.first)
        {
            cout << signs[maxes.second].id << " ";
            signs[maxes.second].num--;
            i++;
        }

    }

    return 0;
}