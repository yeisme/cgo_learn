#include "libshare.h"

#include <iostream>

int main() {
    GoInt32 a = 5;
    GoInt32 b = 10;
    GoInt32 result = Add(a, b);
    std::cout << "Result: " << result << std::endl;

}