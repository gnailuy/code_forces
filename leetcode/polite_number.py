def polite_number(n):
    if n <= 1:
        return 0
    start = 1
    end = 2
    total = 1
    while total+end < n:
        total += end
        end += 1
    count = 0
    if total == n and end-start > 1:
        print([i for i in range(start, end)])
        count += 1
    while start < end:
        total += end
        end += 1
        while total > n:
            total -= start
            start += 1
        if total == n and end-start > 1:
            print([i for i in range(start, end)])
            count += 1
    return count

