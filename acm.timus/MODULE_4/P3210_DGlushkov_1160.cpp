#include <iostream>
#include <cstdint>
#include <map>
#include <list>
#include <vector>

template <class T, class S>
class Graph
{
public:

    S nof_nodes;
    std::multimap<T, std::pair<S, S>> mmap;
    std::list<std::pair<S,S>> mst;
    std::vector<S> mst_check;
    T result;

public:

    explicit Graph(S nof_nodes): nof_nodes(nof_nodes)
    {
        mst_check = std::vector<S>(nof_nodes);
        result = 0;
        for (S i = 0; i < nof_nodes; i++)
            mst_check[i] = i;
    }

    void insert(S from, S to, T weight)
    {
        mmap.insert({weight, {from, to}});
    }

    void kruskal()
    {
        S cur_from, cur_to;
        T cur_weight;
        for (S i = 0; i < nof_nodes; i++)
        {
            auto cur_node = mmap.begin();
            cur_from = cur_node.operator*().second.first;
            cur_to = cur_node.operator*().second.second;
            if (mst_check[cur_from] != mst_check[cur_to])
            {
                if (cur_node.operator*().first > result)
                    result = cur_node.operator*().first;
                mst.push_back({cur_from, cur_to});
                S old_id = mst_check[cur_from], new_id = mst_check[cur_to];
                for (S j = 0; j < nof_nodes; j++)
                    if (mst_check[j] == old_id)
                        mst_check[j] = new_id;
            }
            mmap.erase(cur_node);
        }
    }

    void print_result()
    {
        std::cout << result << std::endl;
        std::cout << mst.size() << std::endl;
        for (auto it : mst)
            std::cout << it.first << " " << it.second << std::endl;
    }

};

using namespace std;

int main()
{
    uint16_t nof_nodes, nof_edges;
    uint16_t node_from, node_to, weight;

    cin >> nof_nodes >> nof_edges;

    Graph<uint16_t, uint16_t> graph(nof_nodes);

    for (uint16_t i = 0; i < nof_edges; i++)
    {
        cin >> node_from >> node_to >> weight;
        graph.insert(node_from , node_to , weight);
    }

    graph.kruskal();
    graph.print_result();

    return 0;
}