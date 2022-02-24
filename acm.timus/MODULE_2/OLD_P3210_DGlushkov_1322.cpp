//this solution works fine but it's kinda slow. Use the one without OLD in its name. Thanks for idea of optimization to Viktor Khrulev - github.com/Demevag

#include <iostream>
#include <algorithm>
#include <vector>

const int SYM_ARR_SIZE = 127;

typedef struct {
    char sym;
    int id;
} pair_t;


using namespace std;

bool pair_check_less(pair_t const &a, pair_t const &b)
{
    if (a.sym < b.sym)
        return true;
    else if (a.sym > b.sym)
        return false;
    else
        return a.id < b.id;
}

int main()
{
    vector<pair_t> first_sym, second_sym;
    int src_str_id, str_size;
    char input_letter, * result;
    int first_sym_counter[SYM_ARR_SIZE];

    for (int i = 0; i < SYM_ARR_SIZE; i++)
        first_sym_counter[i] = 0;

    scanf("%d", &src_str_id);
    src_str_id--;

    //skip spaces
    do
        scanf("%c", &input_letter);
    while ((input_letter < 'A' || input_letter > 'Z') && (input_letter < 'a' || input_letter > 'z') && input_letter !='_');

    do
    {
        first_sym.insert( first_sym.end(), {input_letter, first_sym_counter[input_letter]++} );
        scanf("%c", &input_letter);
    }
    while ((input_letter >= 'A' && input_letter <= 'Z') || (input_letter >= 'a' && input_letter <= 'z') || input_letter =='_');

    str_size = static_cast<int>(first_sym.size());
    for (char i = 0; i < SYM_ARR_SIZE; i++)
        if (first_sym_counter[i] > 0)
            for(int j = 0; j < first_sym_counter[i]; j++)
                second_sym.insert(second_sym.end(), { i, j} );

    result = new char[str_size+1];
    result[str_size] = '\0';


    vector<pair_t>::iterator cur_sym_iter;
    for(int i = str_size - 1; i > -1; i--)
    {
        cur_sym_iter = lower_bound(second_sym.begin(), second_sym.end(), first_sym[src_str_id], pair_check_less);
        result[i] = cur_sym_iter.operator*().sym;
        src_str_id = static_cast<int>(cur_sym_iter - second_sym.begin());
    }

    printf("%s", result);
    delete[] result;
    return 0;
}
