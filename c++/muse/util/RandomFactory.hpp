#ifndef MUSE_UTIL_RANDOM_FACTORY_HPP
#define MUSE_UTIL_RANDOM_FACTORY_HPP 1

#include <cstdlib>
#include <ctime>

namespace RandomFactory {

    void seed() {
        std::srand(static_cast<unsigned>(std::time(nullptr)));
    }

    int genIntN(int min, int max) {
        return min + std::rand() % (max - min + 1);
    }

    int genInt() {
        return genIntN(0, 2'147'483'647);
    }

    int genEven() {
        return genInt() >> 1 << 1;
    }

    int genOdd() {
        return genInt() | 1;
    }
}

#endif
