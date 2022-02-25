/**
source:https://acm.timus.ru/problem.aspx?num=1650
Пояснения к примененному алгоритму:
Основная сложность данной задачи не в нахождении математического алгоритма, который бы давал правильный ответ, а именно в реализации этого алгоритма оптимальным способом.
Исходя из анализа услвоия можно сделать вывод, что для решения задачи требуются возможности быстро:
1) определять город по человеку;
2) определять город по сумме денег в этом городе;
3) определять сумму денег в городе и количество дней, которые этот город был самым богатым, по его названию;
Для выполнения этих условий будем использовать красно-черные деревья - 1 и 3 через std::map (для 1: ключ - фамилия, значение - текущий город; для 3: ключ - название города, значение - сумма денег в городе и количество дней, которые этот город был самым богатым), 2 условие через std::set т.к. возможны случаи, когда в двух разных городах находится одинаковая сумма денег.
Экспериментальным путем выяснилось, что все 3 условия должны выполнятся как минимум за logN, иначе 	Time limit exceeded :(
*/
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

