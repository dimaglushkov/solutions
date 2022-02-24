#include <iostream>
#include <algorithm>
#include <cstdint>
#include <list>

using namespace std;

int main()
{
    uint64_t n, * x, * y, sum = 0;

    cin >> n;
    x = new uint64_t[n];
    y = new uint64_t[n];

    for (uint64_t i = 0; i < n; i++)
        cin >> x[i] >> y[i];

    sort(x, x + n);
    sort(y, y + n);

    for (uint64_t i = 0; i < n - 1; i++)
        sum += (x[i + 1] - x[i]) * (i + 1) * (n - i - 1) + (y[i + 1] - y[i]) * (i + 1) * (n - i - 1);

    sum = 2 * sum / (n * (n - 1));
    cout << sum;

    delete[] x;
    delete[] y;
    return 0;
}