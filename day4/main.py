def main():
    start_val = 183564
    stop_val = 657474

    count = 0
    for i in range(start_val, stop_val + 1):
        if is_valid_password(i):
            count += 1
    print(count)

def is_valid_password(candidate):
    candidate_str = str(candidate)

    chars = dict()
    for index, c in enumerate(candidate_str):
        if c not in chars:
            chars[c] = list()
        chars[c].append(index)

    has_continuous_streak = False
    for _, value in chars.items():
        if len(value) == 2:
            if value[0] + 1 == value[1]:
                has_continuous_streak = True
        if longest_monotonic_streak(value) == 2:
            has_continuous_streak = True

    if not has_continuous_streak:
        return False

    if not ''.join(sorted(candidate_str)) == candidate_str:
        return False

    return True

def longest_monotonic_streak(values):
    longest = 0
    last = values[0] - 1
    values.append(-1)
    overall_longest = 0
    for value in values:
        if value == last + 1:
            longest += 1
        else:
            if longest > overall_longest:
                overall_longest = longest
            longest = 1
        last = value
    return overall_longest

if __name__ == '__main__':
    main()
