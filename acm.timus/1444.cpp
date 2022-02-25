/**
source:https://acm.timus.ru/problem.aspx?num=1444
Пояснения к примененному алгоритму:
1. При защите задачи 1207 было доказано, что для любой точки из произвольного множества точек на плоскости можно определить численную характеристику, описывающую относительное положение выбранной точки ко всем остальным (в нашем случае использовался котангенс).
2. Будем называть первую точку на входе избранной.
Исходя из пункта 1 следует, что условие начала пути в избранной точке выполнимо, так как мы можем оценить относительное положение любой точки. В реализации решения избранная точка рассматривается как условный центр тригонометрического круга и с помощью арктангенса и ряда проверок высчитывается значения углов на этом круге для всех остальных точек.
После этого, через проверки значений координат избранной точки определяется значение ее угла - если она имеет наибольшее значение y, то она занимает угол pi/2, если наименьшее - угол 3pi/2, и так далее.
Затем все точки сортируются по значению угла, среди них ищется избранная точка, после чего на вывод подаются  все точки, начиная с избранной в получившемся после сортировки порядке.
*/
#include <iostream>
#include <cstdint>
#include <vector>
#include <algorithm>
#define _USE_MATH_DEFINES
#include <cmath>

using namespace std;

struct coord_t{
    long int x, y;
    double angle;
    size_t id;

    static bool cmp(const coord_t &a, const coord_t &b)
    {
        if (a.angle != b.angle)
            return a.angle < b.angle;
        else
            return a.x < b.x;
    }

    void set_angle_by_point(coord_t base_point)
    {
        long int  dy =  y - base_point.y , dx = x - base_point.x;
        angle = atan((double) dx / dy);
        if (dx > 0 && dy > 0)
            angle = M_PI_2 - angle;
        else if (dx < 0 && dy > 0)
        {
            angle *= -1;
            angle += M_PI_2;
        }
        else if (dx <= 0 && dy < 0)
            angle = 3 * M_PI_2 - angle;
        else if (dx > 0 && dy < 0)
        {
            angle *= -1;
            angle += 3 * M_PI_2;
        }
    }
};

void print_path( vector<coord_t>& points)
{
    size_t id_of_first = 0, N = points.size();
    for (size_t i = 0; i < N; i++)
        if (points[i].id == 1)
        {
            id_of_first = i;
            break;
        }

    for (size_t i = id_of_first; i < N; i++)
        cout << points[i].id << endl;

    for (size_t i = 0; i < id_of_first; i++)
        cout << points[i].id << endl;

}

int main()
{
    size_t N;
    vector<coord_t> points;
    coord_t max = {-1001, -1001}, min = {1001, 1001};
    coord_t * left, * right, * bot, * top;

    cin >> N;
    points = vector<coord_t>(N);


    for (size_t i = 0; i < N; i++)
    {
        long int x, y;
        cin >> x >> y;

        if (x > max.x)
            max.x = x;
        if (x < min.x)
            min.x = x;

        if (y > max.y)
            max.y = y;
        if (y < min.y)
            min.y = y;

        points[i] = {x, y};

        points[i].id = i + 1;

        if (!i)
        {
            points[i].angle = -361;
            continue;
        }
        if (points[i].x == points[0].x)
        {
            points[i].angle = (points[i].y > points[0].y) ? M_PI_2 : 3 * M_PI_2;
            continue;
        }

        if (points[i].y == points[0].y)
        {
            points[i].angle = (points[i].x > points[0].x) ? 0 : M_PI;
            continue;
        }

        points[i].set_angle_by_point(points[0]);

    }

    if (points[0].x == max.x)
        points[0].angle = 0;
    if (points[0].y == max.y)
        points[0].angle = M_PI_2;
    if (points[0].x == min.x)
        points[0].angle = M_PI;
    if (points[0].x == min.y)
        points[0].angle = 3 * M_PI_2;


    sort(points.begin(), points.end(), coord_t::cmp);

    cout << N << endl;
    print_path(points);

    return 0;
}
