#include <iostream>
#include <cstdint>
#include <algorithm>

using namespace std;

struct sym_t
{
    uint32_t id, num;

    static bool cmp(const sym_t &a, const sym_t &b)
    {
        if (a.num != b.num)
            return a.num > b.num;

        else
            return a.id > b.id;
    }

};

int main()
{
    size_t k;
    uint32_t last_id = 0, sum = 0;
    sym_t * n;
    cin >> k;
    n = new sym_t[k];

    for (size_t i = 0; i < k; i++)
    {
        n[i].id = i + 1;
        cin >> n[i].num;
        sum += n[i].num;
    }

    if (k == 1)
    {
        sort(n, n + k, sym_t::cmp);

        for(int i = 0; i < sum; i++)
            cout << n[0].id << " ";
        delete[] n;

        return 0;
    }

    if (k == 2)
    {
        sort(n, n + k, sym_t::cmp);

        size_t last = 0;
        for (int i = 0; i < sum; i++)
            for (size_t j = 0; j < k; j++)
            {
                if(n[j].num > 0)
                {
                    cout << j + 1 << " ";
                    n[j].num--;
                }
            }
        delete[] n;

        return 0;
    }


    size_t cur = 0;
    for (int i = 0; i < sum; i++)
    {
        sort(n, n + k, sym_t::cmp);
        if (n[0].id == last_id || n[0].num == n[1].num)
            cur = 1;

        if (n[0].id != last_id)
            cur = 0;

        cout << n[cur].id << " ";
        last_id = n[cur].id;
        n[cur].num--;


    }
    delete[] n;
    return 0;
}