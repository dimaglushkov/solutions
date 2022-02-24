#include <iostream>
#include <vector>
#include <set>
#include <algorithm>

using namespace std;

bool cmp_by_raws(pair<uint32_t, uint32_t> a, pair<uint32_t, uint32_t> b)
{
    if (a.first != b.first)
        return a.first < b.first;
    return a.second < b.second;
}

bool cmp_by_cols(pair<uint32_t, uint32_t> a, pair<uint32_t, uint32_t> b)
{
    if (a.second != b.second)
        return a.second < b.second;
    return a.first < b.first;
}

void fill_calendar(vector<pair<uint32_t, uint32_t>> & calendar, uint32_t x_size, uint32_t y_size, uint32_t & num_of_days)
{
    for (uint32_t i = 1; i <= x_size; i++)
    {
        calendar[num_of_days++] = {i, 0};
        calendar[num_of_days++] = {i, y_size + 1};
    }

    for (uint32_t i = 1; i <= y_size; i++)
    {
        calendar[num_of_days++] = {0, i};
        calendar[num_of_days++] = {x_size + 1, i};
    }
}


int main()
{
    uint32_t x_size, y_size, num_of_days, sum = 0;
    set<pair<uint32_t, uint32_t>> single_days;
    vector<pair<uint32_t, uint32_t>> calendar;

    cin >> x_size >> y_size >> num_of_days;
    calendar = vector<pair<uint32_t, uint32_t>>((y_size + 2) * (x_size + 2));

    for (uint32_t x, y, i = 0; i < num_of_days; ++i)
    {
        cin >> x >> y;
        calendar[i] = {x, y};
    }

    fill_calendar(calendar, x_size, y_size, num_of_days);

    sort(calendar.begin(), calendar.begin() + num_of_days, cmp_by_raws);
    for (uint32_t i = 0; i + 1 < num_of_days; i++)
    {
        uint32_t num = calendar[i + 1].second - calendar[i].second;
        if (calendar[i].first == calendar[i + 1].first && num >= 2)
        {
            if (num == 2)
                single_days.insert({calendar[i].first, calendar[i].second + 1});

            else ++sum;
        }
    }

    sort(calendar.begin(), calendar.begin() + num_of_days, cmp_by_cols);
    for (uint32_t i = 0; i + 1 < num_of_days; i++)
    {
        uint32_t num = calendar[i + 1].first - calendar[i].first;
        if (calendar[i].second == calendar[i + 1].second && num >= 2)
        {
            if (num == 2)
            {
                if (single_days.find({calendar[i].first + 1, calendar[i].second}) != single_days.end())
                    sum++;
            }
            else sum++;
        }
    }

    cout << sum << endl;
    return 0;
}