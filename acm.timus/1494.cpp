/**
source:https://acm.timus.ru/problem.aspx?num=1494
Пояснения к примененному алгоритму:
Рассмотрим положение шаров, при котором результат будет однозначно "Not a proof". Заметим, что существуют наборы, которые можно трактовать как "Cheater", однако на самом деле они "Not a proof".
Из этого следует, что нужно проверять именно на соответствие набору для "Not a proof" - если будут расхождения, то вывести "Cheater", если ни одно расхождения не будет - "Not a proof".
Набор для "Not a proof" - такой набор, что, если бы мы не доставали ни один шар, то номера шаров шли бы в порядке возрастания.
Для проверки на соответствие такому набору при каждом изъятии шара будем предполагать, что все предыдущие шары уже находятся в лунке. Если при изъятии шара будет нарушено условие для "Not a proof", значит игрок - "Cheater".

*/
#include <iostream>
#include <stack>
#include <cstdint>

template <class T, class Size>
class Stack{

    std::stack<T> stack;

public:
    Size max_size = 0;

    void push(T value)
    {
        if (stack.size() < max_size)
            stack.push(value);
    }

    Size size()
    {
        return stack.size();
    }

    T pop()
    {
        T temp = stack.top();
        stack.pop();
        return temp;
    }

    T top()
    {
        return stack.top();
    }

    bool empty()
    {
        return stack.empty();
    }

    bool full()
    {
        return stack.size() == max_size;
    }

    void fill(T begin, T end)
    {
        for (T i = begin; i < end; i++)
            stack.push(i);
    }

};


int main()
{
    Stack<uint32_t, uint32_t> stack;
    uint32_t prev, cur;

    std::cin >> stack.max_size;

    std::cin >> prev;
    stack.fill(1, prev);

    for (uint32_t i = 1; i < stack.max_size; i++)
    {
        std::cin >> cur;

        if (cur > prev)
        {
            stack.fill(prev + 1, cur);
            prev = cur;
            continue;
        }
        else
        if (stack.pop() != cur)
        {
            std::cout << "Cheater";
            return 0;
        }
    }

    std::cout << "Not a proof";
    return 0;
}
