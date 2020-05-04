#ifndef FUN_VAC_UTIL_RANDOM_FACTORY_HPP
#define FUN_VAC_UTIL_RANDOM_FACTORY_HPP 1

#include <cstdlib>
#include <ctime>

namespace RandomFactory {

    void seed() {
        std::srand(static_cast<unsigned>(std::time(nullptr)));
    }

    int integerN(int min, int max) {
        return min + std::rand() % (max - min + 1);
    }

    int generateInteger() {
        return integerN(0, 2'147'483'647);
    }

    int generateEven() {
        return generateInteger() >> 1 << 1;
    }

    int generateOdd() {
        return generateInteger() | 1;
    }
}

#endif
