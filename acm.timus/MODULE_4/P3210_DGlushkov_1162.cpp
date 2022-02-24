#include <iostream>
#include <vector>
#define BORDER 0.0001

using namespace std;

template<class S, class T>
class Graph
{
    struct edge
    {
        S a, b;
        T rate, commission;
    };
    vector<edge> edges;
    vector<T> path;
    S nof_nodes, nof_edges;
    S start;
    T sum;

public:

    explicit Graph(S nof_nodes, S nof_edges, S currency, T sum): nof_nodes(nof_nodes), nof_edges(nof_edges), start(currency), sum(sum)
    {
        path = vector<T>(nof_nodes * 2);
        path[currency] = sum;
    }

    void insert(S a, S b, T rate, T commission)
    {
        edges.push_back({a, b, rate, commission});
    }

    bool solve()
    {
        for (S i = 0; i < nof_nodes - 1; i++)
            for (auto cur : edges)
                if (path[cur.b] - (path[cur.a] - cur.commission) * cur.rate < BORDER)
                    path[cur.b] = (path[cur.a] - cur.commission) * cur.rate;

        for (auto cur : edges)
            if ((path[cur.a] - cur.commission) * cur.rate - path[cur.b]  > BORDER)
                return true;

        return false;
    }


};


int main()
{
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);

    uint16_t nof_nodes, nof_edges, a, b, currency;
    double rab, rba, cab, cba, sum;

    cin >> nof_nodes >> nof_edges >> currency >> sum;
    Graph<uint16_t, double> graph = Graph<uint16_t, double>(nof_nodes, nof_edges, currency, sum);

    for (uint16_t i = 0; i < nof_edges; i++)
    {
        cin >> a >> b >> rab >> cab >> rba >> cba;
        graph.insert(a, b, rab, cab);
        graph.insert(b, a, rba, cba);
    }

    if ((graph.solve()))
        cout << "YES";
    else
        cout << "NO";

    return 0;
}