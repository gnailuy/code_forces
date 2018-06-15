class RomanToInt:
    c_map = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    def romanToInt(self, s):
        result = 0
        large = -1

        for c in reversed(s):
            cn = self.c_map[c]
            if cn >= large:
                large = cn
                result += cn
            else:
                result -= cn

        return result


ri = RomanToInt()

tests = [
    'III',
    'IV',
    'IX',
    'LVIII',
    'MCMXCIV',
]

for test in tests:
    print(ri.romanToInt(test))

