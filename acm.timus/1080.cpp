/**
source:https://acm.timus.ru/problem.aspx?num=1080
Пояснения к примененному алгоритму:
Для данной задачи, очевидно, стоит применять графы, при чем за узлы брать страны, а за ребра брать наличие границы между двумя странами.
Из этого становится не менее очевидным, что получается неориентированный невзвешанный граф.
В своей реализации я представлял граф в виде матрицы смежности.
В условии требуется раскрасить все страны в 2 цвета, другими словами, 2 соседние страны должны быть разных цветов, а это задача на определение двудольности графа, это можно сделать с помощью BFS или DFS (ровно это было сказано на лекции 30 апреля, предпоследний слайд).
*/
#include <iostream>
#include <cstdint>

using namespace std;

template <class Size>
class Graph
{
private:

    Size size;
    bool ** matrix;
    Size * result;
    bool solvable = true;

public:

    explicit Graph(const Size size) : size(size)
    {
        matrix = new bool*[size];
        for (Size i = 0; i < size; i++)
            matrix[i] = new bool[size];

        result = new Size[size];

        for (Size i = 0; i < size; i++)
        {
            for(uint8_t j = 0; j < size; j++)
                matrix[i][j] = false;

            result[i] = -1;
        }
        result[0] = 0;
    }

    ~Graph()
    {
        for (int i = 0; i < size; ++i)
            delete[] matrix[i];
        delete[] matrix;
        delete[] result;
    }

    void add_edge(Size a, Size b)
    {
        matrix[a][b] = true;
        matrix[b][a] = true;
    }

    bool get_edge(Size a, Size b)
    {
        return matrix[a][b];
    }

    bool solve(Size cur_node = 0, Size color = 0)
    {
        result[cur_node] = color;
        for (Size i = 0; i < size; i++)
        {
            if (get_edge(cur_node, i))
            {
                if(result[i] == -1)
                    solve(i, !color);

                else if(result[i] == color)
                    return solvable = false;
            }
        }
        return solvable;
    }

    void print_result()
    {
        for (Size i = 0; i < size; i++)
            cout << result[i];
    }

};

int main()
{
    int16_t size, temp;
    cin >> size;
    Graph<int16_t> graph = Graph<int16_t  >(size);

    for (int16_t  i = 0; i < size; i++)
    {
        do
        {
            cin >> temp;
            if (temp != 0)
                graph.add_edge(i, temp - 1);
        }
        while(temp != 0);
    }

    if (graph.solve())
        graph.print_result();
    else
        cout << -1;

    graph.~Graph();
    return 0;
}