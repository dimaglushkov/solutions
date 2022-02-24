#include <iostream>
#include <map>
#include <cstdint>
#include <set>

using namespace std;

typedef uint64_t money_t;

struct human
{
    money_t money;
    string cur_city;

    void init(money_t & money_, string & cur_city_)
    {
        human::money = money_;
        human::cur_city = cur_city_;
    }
};

struct city
{
    uint16_t num_of_days = 0;
    money_t sum = 0;
};


void increase_num_of_days(map<string, city> & world, set<pair<money_t, city*>> & cream, uint16_t diff)
{
    auto richest_city = cream.end();
    richest_city--;
    if (world.size() > 1)
    {
        auto rich_city = cream.end();
        rich_city--;
        rich_city--;
        if(richest_city.operator*().first != rich_city.operator*().first)
            richest_city.operator*().second->num_of_days += diff;
    }
    else
        richest_city.operator*().second->num_of_days += diff;
}

void print_cities(map<string, city> & world)
{
    for (const auto& iterator : world)
    {
        if(iterator.second.num_of_days > 0)
            cout << iterator.first << " " << iterator.second.num_of_days << endl;
    }
}

int main()
{
    map<string, city> world;
    map<string, human> humans;
    set<pair<money_t, city*>> cream;
    uint16_t num_of_humans;
    string cur_name, cur_city;
    money_t cur_money;
    uint32_t num_of_moves, num_of_days, prev_day, cur_day = 0;

    cin >> num_of_humans;
    for (uint16_t i = 0; i < num_of_humans; i++)
    {
        cin >> cur_name >> cur_city >> cur_money;
        world[cur_city].sum += cur_money;
        humans[cur_name].init(cur_money, cur_city);
    }

    for (auto &it : world)
        cream.insert({it.second.sum, &it.second});


    cin >> num_of_days >> num_of_moves;
    for (uint16_t i = 0; i < num_of_moves; i++)
    {

        prev_day = cur_day;
        cin >> cur_day >> cur_name >> cur_city;
        increase_num_of_days(world, cream, cur_day - prev_day);

        string old_city = humans[cur_name].cur_city;

        //changing city
        cream.erase({world[old_city].sum, &world[old_city]});
        world[old_city].sum -= humans[cur_name].money;
        cream.insert({world[old_city].sum, &world[old_city]});

        /*if  (world[temp_city].sum == 0 && world[temp_city].num_of_days == 0)
            world.erase(temp_city);*/

        // updating new city
        humans[cur_name].cur_city = cur_city;
        cream.erase({world[cur_city].sum, &world[cur_city]});
        world[cur_city].sum += humans[cur_name].money;
        cream.insert({world[cur_city].sum, &world[cur_city]});
    }

    increase_num_of_days(world, cream, num_of_days - cur_day);
    print_cities(world);



    return 0;
}

