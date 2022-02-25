/**
source:https://acm.timus.ru/problem.aspx?num=1450
Пояснения к примененному алгоритму:
Данная задача представляет из себя классический вариант задачи о самом длинном пути в графе.
Исходные данные представляются как ациклический ориентированный взвешанный граф, где узлы - номера добывающих станций, а ребра - соединяющих их газопровод.
Для определения максимального пути можно исполльзовать алгоритм Беллмана-Мура (Беллмана-Форда), однако так как известно, что вес ребра неотрицателен, можно использовать алгоритм Дейкстры.
*/
#include <iostream>
#include <vector>

using namespace std;

template<class T, class S>
class Graph
{
    struct edge
    {
        S from, to;
        T weight;
    };
    vector<edge> edges;
    vector<T> max_path;
    S nof_nodes, nof_edges;

public:

    explicit Graph(S nof_nodes, S nof_edges): nof_nodes(nof_nodes), nof_edges(nof_edges)
    {
        max_path = vector<T>(nof_nodes, -1);
    }

    void insert(S from, S to, T weight)
    {
        edges.push_back({from, to, weight});
    }

    T solve(S start_node, S end_node)
    {
        max_path[start_node] = 0;

        for (S i = 0; i < nof_nodes - 1; i++)
            for (auto cur : edges)
                if (max_path[cur.to] < max_path[cur.from] + cur.weight && max_path[cur.from] != -1)
                    max_path[cur.to] = max_path[cur.from] + cur.weight;

        return max_path[end_node];
    }


};


int main()
{
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    uint32_t nof_nodes, nof_edges, node_from, node_to;
    int node_weight, result;

    cin >> nof_nodes >> nof_edges;
    Graph<int, uint32_t> graph = Graph<int, uint32_t>(nof_nodes, nof_edges);

    for (uint32_t i = 0; i < nof_edges; i++)
    {
        cin >> node_from >> node_to >> node_weight;
        graph.insert(node_from - 1, node_to - 1, node_weight);
    }

    cin >> node_from >> node_to;
    if ((result = graph.solve(node_from - 1, node_to - 1)) == -1)
        cout << "No solution";
    else
        cout << result;

    return 0;
}