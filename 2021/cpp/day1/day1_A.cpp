#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <vector>
#include <span>
#include <ranges>
#include <algorithm>
#include <limits>

void count_increasing(std::vector<int> const &v)
{
    int ans = std::ranges::count_if(
        v,
        [prev = std::numeric_limits<int>::max()](int cur) mutable
        {
            return std::exchange(prev, cur) < cur;
        });
    std::cout << ans << '\n';
}

void count_window_increasing(std::vector<int> const &v)
{
    auto sliding_view = v | std::ranges::views::transform([a = 0, b = 0, c = 0](int curr) mutable
                                                          {
                            a = std::exchange(b, std::exchange(c, curr));
                            return a+b+c; }) | std::ranges::views::drop(2);

    int ans = std::ranges::count_if(
        sliding_view,
        [prev = std::numeric_limits<int>::max()](int cur) mutable
        {
            return std::exchange(prev, cur) < cur;
        });
    std::cout << ans << '\n';
}

int main()
{
    int n;
    std::vector<int> inputs;
    while (std::cin >> n)
    {
        inputs.push_back(n);
    }
    count_increasing(inputs);
    count_window_increasing(inputs);
    return 0;
}