/**
source:https://acm.timus.ru/problem.aspx?num=1322
Пояснения к примененному алгоритму:
1) Докажем, что каждый символ побывает началом строки: 
По условию задачи для строки из n символов формируется n строк, для образования которой необходимо сдвигать символы влево с переносом. Таким образом, можно однозначно говорить, что в при рассматривании последовательности строк как матрицы символов, каждый символ побывает в каждом столбце ровно 1 раз. 
Из этого следует вывод, что каждый символ единожды побывает на первом месте.

2) Из условия лексиграфической отсортированности строк, а также из условия, доказанного в пункте 1 следует, что в первой колонке рассматриваемой матрицы символов расположены все символы из последней колонки, отсортированные в алфавитном порядке. 

3) Рассмотрим строку x, в которой n символов. Нам уже известен символ x_1 и символ x_n. Из алгоритма получения новой строки (сдвига символов) можно сделать вывод, что во всех других строках символы x_1 и x_n стоят рядом, более того, они обязательно расположены в порядке x_n,x_1. Однако для однозначного определения связи и сохранения исходного порядка пар x_i,x_j необходимо однозначно определить взаимосвязи между двумя буквами из разных столбцов. Для этого в программе используется вектор sym_pairs. В нем индекс - номер строки, буквы с конца, или первой буквы в паре, а значение по этому индексу - номер строки буквы с начала, или же второй буквы в паре.

*/
#include <iostream>
#include <algorithm>
#include <vector>
#include <stdint.h>

using namespace std;

const unsigned char SYM_ARR_SIZE = 127;
static string first_sym;

bool cmp_sym_pairs(uint32_t a, uint32_t b)
{
    return first_sym[a] != first_sym [b] ? first_sym[a] < first_sym[b] : a < b ;
}

int main()
{
    uint32_t cur_index, sym_counter[SYM_ARR_SIZE] = {0};
    vector<uint32_t> sym_pairs;
    string second_sym, result;
    char input_sym;



    cin >> cur_index;
    cur_index--;

    do
        scanf("%c", &input_sym);
    while ((input_sym < 'A' || input_sym > 'Z') && (input_sym < 'a' || input_sym > 'z') && input_sym !='_');

    do
    {
        first_sym += input_sym;
        sym_counter[input_sym]++;
        scanf("%c", &input_sym);
    }
    while ((input_sym >= 'A' && input_sym <= 'Z') || (input_sym >= 'a' && input_sym <= 'z') || input_sym =='_');

    for (char i = 0; i < SYM_ARR_SIZE; i++)
        if (sym_counter[i] > 0)
            for(uint32_t j = 0; j < sym_counter[i]; j++)
                second_sym += i;


    for(uint32_t i = 0; i < first_sym.length(); i++)
        sym_pairs.push_back(i);

    sort(sym_pairs.begin(), sym_pairs.end(), cmp_sym_pairs);


    for(uint32_t i = 0; i < first_sym.length(); i++)
    {
        result += second_sym[cur_index];
        cur_index = sym_pairs[cur_index];
    }

    cout << result;
    return 0;
}
