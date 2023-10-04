#include <ranges>
#include <iostream>
#include <functional>
#include <algorithm>
#include <string>
#include <stdio.h>
#include <execution>

struct depth
{
    int x, y, aim;

    depth &operator+=(depth const &other)
    {
        x += other.x;
        y += other.y;
        return *this;
    }

    depth operator+(depth const &other) const
    {
        return {x + other.x, y + other.y, aim + other.aim};
    }
};

void get_pos(std::vector<std::string> const &input)
{

    depth ans = std::accumulate(std::begin(input), std::end(input), depth{.x = 0, .y = 0}, [](depth const &prev, std::string const &cur)
                                {
        auto pos = cur.find_first_of(' ');
        int x = std::stoi(cur.substr(pos+1));
        if(cur.starts_with("forward")) {
            return prev + depth{x, x*prev.aim, 0};
        }
        if(cur.starts_with("down")) {
            return prev +depth {0,0 ,x};
        }
        if(cur.starts_with("up")){ 
            return prev + depth{0,0, -x};
        }
        if(cur.starts_with(""))
        return prev; });

    std::cout << ans.x * ans.y << '\n';
}

int main()
{

    std::string s;
    std::vector<std::string> input;
    while (getline(std::cin, s))
    {
        input.push_back(s);
    }
    std::cout <<"doign 2" << '\n';

    get_pos(input);
}