#include <iostream>
#include <vector>
#include <map>
#include <unordered_map>
#include <algorithm>
#include <deque>
#include <queue>

const uint16_t INF = 50010u;

using namespace std;

class Graph
{
    vector<uint16_t> price;
    unordered_map<string, uint16_t> num_to_id;
    vector<string> id_to_num;
    vector< vector< pair<uint16_t , uint16_t> > > edges;
    vector<uint16_t> length;
    vector<uint16_t> parents;
    uint16_t n, iter;


public:

    explicit Graph(uint16_t n): n(n)
    {
        length = vector<uint16_t>(n + 5, INF);
        parents = vector<uint16_t>(n + 5);
        id_to_num = vector<string>(n + 5);
        edges = vector< vector < pair<uint16_t, uint16_t > > >(n + 5);

        iter = 0;
    }

    void insert_price(uint16_t val)
    {
        price.push_back(val);
    }

    void insert_num(string & val_str)
    {
        num_to_id[val_str] = iter;
        id_to_num[iter++] = val_str;

        for (int16_t i = 0; i < 10; i++)
        {
            for(char e = 48; e < 58; e++)
            {
                if(val_str[i] != e)
                {
                    string temp = val_str;
                    temp[i] = e;
                    if(num_to_id.find(temp) != num_to_id.end())
                    {
                        uint16_t id1 = num_to_id[val_str], id2 = num_to_id[temp];
                        edges[id1].emplace_back(id2, price[i]);
                        edges[id2].emplace_back(id1, price[i]);
                    }
                }
            }

            for(int16_t j = i + 1; j < 10; j++)
            {
                string temp = val_str;
                swap(temp[i], temp[j]);
                if(num_to_id.find(temp) != num_to_id.end())
                {
                    uint16_t id1 = num_to_id[val_str], id2 = num_to_id[temp];
                    edges[id1].emplace_back(id2, price[i]);
                    edges[id2].emplace_back(id1, price[i]);
                }
            }

        }

    }

    void solve()
    {
        priority_queue <pair<uint16_t, uint16_t>> que;
        length[0] = 0;
        que.push({0, 0});
        while(!que.empty())
        {
            uint16_t v = que.top().second;
            que.pop();

            for (auto it : edges[v])
            {
                uint16_t to = it.first, len = it.second;
                if (length[v] + len < length[to])
                {
                    length[to] = length[v] + len;
                    parents[to] = v;
                    que.push({length[to], to});
                }
            }
        }
    }

    void print_result(uint16_t end)
    {
        deque<uint16_t> path;

        if (length[end] == INF)
        {
            cout << "-1";
            return;
        }

        for (int64_t i = end; i != 0; i = parents[i])
            path.push_front(i);
        path.push_front(0);

        cout << length[end] << '\n' << path.size() << '\n';
        for (auto i : path)
            cout << i + 1 << " ";
    }

};

int main()
{
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);

    uint16_t n = 0;
    uint16_t price;
    string num;

    cin >> n;
    Graph graph = Graph(n);

    for (int16_t i = 0; i < 10; i++)
    {
        cin >> price;
        graph.insert_price(price);
    }

    for (uint16_t i = 0; i < n; i++)
    {
        cin >> num;
        graph.insert_num(num);
    }

    graph.solve();
    graph.print_result(n - 1);

    return 0;
}