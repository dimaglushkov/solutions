/// this problem is pure evil, have fun
#include <iostream>
#include <cstdint>

enum direction{
    NONE,
    ROOT,
    TOP,
    LEFT,
    RIGHT
};

template <class T, class Index>
class Tree
{
    T * values;
    Index last_level;
    Index size;
    Index pop_step;
    Index current_node;

    void remove(Index cur_node)
    {
        do
            values[cur_node]--;
        while(to_father(cur_node));
    }

    bool to_left_son(Index& cur_node)
    {
        if (2 * cur_node + 2 < size)
        {
            cur_node = 2 * cur_node + 1;
            return true;
        }
        else return false;
    }

    bool to_right_son(Index& cur_node)
    {
        if (2 * cur_node + 2 < size)
        {
            cur_node = 2 * cur_node + 2;
            return true;
        }
        else return false;
    }

    bool to_father(Index& cur_node)
    {
        if (cur_node != 0)
        {
            cur_node = (cur_node - 1) / 2;
            return true;
        }
        else return false;
    }

public:

    Tree (Index N, Index K)
    {
        last_level = 1;
        while (last_level < N)
            last_level = last_level << 1u;

        size = (last_level << 1u) - 1;
        pop_step = K;
        values = new T[size];
        current_node = 0;

        for (Index i = 0; i < N; i++)
            values[i + last_level - 1] = 1;
        for (Index i = N; i < last_level; i++)
            values[i + last_level - 1] = 0;

        for (Index temp, i = last_level - 2; i > 0; i--)
        {
            values[i] = 0;
            if (to_left_son(temp = i))
                values[i] += values[temp];
            if (to_right_son(temp = i))
                values[i] += values[temp];
        }

    }

    
    T pop()
    {
        if (!current_node)
        {
            current_node = last_level - 2 + pop_step;
            return current_node - last_level + 2;
        }

        remove(current_node);

        Index counter = pop_step;
        direction dir = NONE;

        Index father;

        while(counter > 0)
        {
            Index temp_left = current_node, temp_right = current_node;
            to_left_son(temp_left);
            to_right_son(temp_right);

            switch(dir)
            {

                case NONE:
                {
                    father = current_node;
                    to_father(father);
                    to_right_son(father);
                    dir = father == current_node ? RIGHT : LEFT;
                    current_node = father;
                    break;
                }

                case ROOT:
                {
                    counter -= values[current_node];
                    if (counter > 0)
                    {
                        dir = LEFT;
                        to_father(current_node);
                    }
                    break;
                }

                case TOP:
                {
                    if(current_node > last_level - 1)
                        counter--;
                    else if(values[temp_left] < counter )
                    {
                        counter -= values[temp_left];
                        dir = TOP;
                        current_node = temp_right;
                    }
                    else
                    {
                        dir = TOP;
                        current_node = temp_left;
                    }
                    break;
                }

                case LEFT:
                {
                    if (values[temp_right] < counter)
                    {
                        counter -= values[temp_right];
                        if (current_node == 0)
                        {
                            current_node = last_level - 1;
                            dir = ROOT;
                        }
                        else
                        {

                            to_father(father = current_node);
                            to_right_son(father);
                            dir = father == current_node ? RIGHT : LEFT;
                            to_father(current_node);
                        }
                    }
                    else
                    {
                        dir = TOP;
                        current_node = temp_right;
                    }
                    break;
                }

                case RIGHT:
                {
                    if (current_node == 0)
                    {
                        current_node = last_level - 1 ;
                        dir = ROOT;
                    }
                    else
                    {
                        father = current_node;
                        to_father(father);
                        to_right_son(father);
                        dir = father == current_node ? RIGHT : LEFT;
                        to_father(current_node);
                    }
                    break;
                }

            }

        }
        return current_node - last_level + 2;
    }

};

using namespace std;

int main()
{
    int n, k;

    cin >> n >> k;

    if (k == 1)
    {
        for (uint32_t i = 1; i <= n; i++)
            cout << i << " ";
        return 0;
    }

    Tree<uint32_t, uint32_t> tree = Tree<uint32_t, uint32_t>(n, k);

    for (int i = 0; i < n; i++)
        cout << tree.pop() << " ";

    return 0;
}

