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
