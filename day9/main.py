import time
memory = dict()

def main():
    data = ''
    with open('input', 'r') as f:
        data = f.read()
    codes = list(map(lambda x: int(x), data.split(',')))

    init_memory(codes)

    compute()

def compute():
    index = 0
    relative_index = 0
    while get_memory(index) != 99:
        print(get_memory(index))
        time.sleep(0.05)
        if get_op_code(get_memory(index)) == 1:
            param0 = get_op_code_param(get_memory(index), 0)
            param1 = get_op_code_param(get_memory(index), 1)
            param2 = get_op_code_param(get_memory(index), 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = get_memory(get_memory(index+1))
            elif param0 == 2:
                input0 = get_memory(get_memory(index+1)+relative_index)
            else:
                input0 = get_memory(index+1)

            if param1 == 0:
                input1 = get_memory(get_memory(index+2))
            elif param1 == 2:
                input1 = get_memory(get_memory(index+2)+relative_index)
            else:
                input1 = get_memory(index+2)

            if param2 == 0:
                input2 = get_memory(index+3)
            elif param2 == 2:
                input2 = get_memory(get_memory(index+3)+relative_index)
            else:
                print('THIS SHOULD NEVER HAPPEN FOR ADD')
                input2 = get_memory(index+3)

            add(input0, input1, input2)
            index += 4
        if get_op_code(get_memory(index)) == 2:
            param0 = get_op_code_param(get_memory(index), 0)
            param1 = get_op_code_param(get_memory(index), 1)
            param2 = get_op_code_param(get_memory(index), 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = get_memory(get_memory(index+1))
            elif param0 == 2:
                input0 = get_memory(get_memory(index+1)+relative_index)
            else:
                input0 = get_memory(index+1)

            if param1 == 0:
                input1 = get_memory(get_memory(index+2))
            elif param1 == 2:
                input1 = get_memory(get_memory(index+2)+relative_index)
            else:
                input1 = get_memory(index+2)

            if param2 == 0:
                input2 = get_memory(index+3)
            elif param2 == 2:
                input2 = get_memory(get_memory(index+3)+relative_index)
            else:
                print('THIS SHOULD NEVER HAPPEN FOR MULTIPLY')
                input2 = get_memory(index+3)

            multiply(input0, input1, input2)
            index += 4
        if get_op_code(get_memory(index)) == 3:
            read_value(get_memory(index+1))
            index += 2
        if get_op_code(get_memory(index)) == 4:
            print_value(get_memory(index+1))
            index += 2
        if get_op_code(get_memory(index)) == 5: # jump if true
            param0 = get_op_code_param(get_memory(index), 0)
            param1 = get_op_code_param(get_memory(index), 1)

            input0 = None
            input1 = None

            if param0 == 0:
                input0 = get_memory(get_memory(index+1))
            elif param0 == 2:
                input0 = get_memory(get_memory(index+1)+relative_index)
            else:
                input0 = get_memory(index+1)

            if param1 == 0:
                input1 = get_memory(get_memory(index+2))
            elif param1 == 2:
                input1 = get_memory(get_memory(index+2)+relative_index)
            else:
                input1 = get_memory(index+2)

            if input0 != 0:
                index = input1
            else:
                index += 3
        if get_op_code(get_memory(index)) == 6: # jump if false
            param0 = get_op_code_param(get_memory(index), 0)
            param1 = get_op_code_param(get_memory(index), 1)

            input0 = None
            input1 = None

            if param0 == 0:
                input0 = get_memory(get_memory(index+1))
            elif param0 == 2:
                input0 = get_memory(get_memory(index+1)+relative_index)
            else:
                input0 = get_memory(index+1)

            if param1 == 0:
                input1 = get_memory(get_memory(index+2))
            elif param1 == 2:
                input1 = get_memory(get_memory(index+2)+relative_index)
            else:
                input1 = get_memory(index+2)

            if input0 == 0:
                index = input1
            else:
                index += 3
        if get_op_code(get_memory(index)) == 7: # less than
            param0 = get_op_code_param(get_memory(index), 0)
            param1 = get_op_code_param(get_memory(index), 1)
            param2 = get_op_code_param(get_memory(index), 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = get_memory(get_memory(index+1))
            elif param0 == 2:
                input0 = get_memory(get_memory(index+1)+relative_index)
            else:
                input0 = get_memory(index+1)

            if param1 == 0:
                input1 = get_memory(get_memory(index+2))
            elif param1 == 2:
                input1 = get_memory(get_memory(index+2)+relative_index)
            else:
                input1 = get_memory(index+2)

            if param2 == 0:
                input2 = get_memory(index+3)
            elif param2 == 2:
                input2 = get_memory(get_memory(index+3)+relative_index)
            else:
                print('THIS SHOULD NEVER HAPPEN FOR LT')
                input2 = get_memory(index+3)

            less_than(input0, input1, input2)
            index += 4
        if get_op_code(get_memory(index)) == 8: # equal to
            param0 = get_op_code_param(get_memory(index), 0)
            param1 = get_op_code_param(get_memory(index), 1)
            param2 = get_op_code_param(get_memory(index), 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = get_memory(get_memory(index+1))
            elif param0 == 2:
                input0 = get_memory(get_memory(index+1)+relative_index)
            else:
                input0 = get_memory(index+1)

            if param1 == 0:
                input1 = get_memory(get_memory(index+2))
            elif param1 == 2:
                input1 = get_memory(get_memory(index+2)+relative_index)
            else:
                input1 = get_memory(index+2)

            if param2 == 0:
                input2 = get_memory(index+3)
            elif param2 == 2:
                input2 = get_memory(get_memory(index+3)+relative_index)
            else:
                print('THIS SHOULD NEVER HAPPEN FOR EQ')
                input2 = get_memory(index+3)

            equal_to(input0, input1, input2)
            index += 4
        if get_op_code(get_memory(index)) == 9:
            input0 = get_memory(index + 1)
            relative_index += input0
            print('increasing')

    # return get_memory(0]

def get_memory(index):
    if index in memory:
        return memory[index]
    return 0

def left_pad(string, length, pad = '0'):
    if len(string) >= length:
        return string
    return ''.join([pad]*(length-len(string))) + string

def get_op_code(code):
    return int(left_pad(str(code), 2)[-2:])

def get_op_code_param(code, index):
    padded_code = left_pad(str(code), 2 + index + 1)
    params = padded_code[:len(padded_code)-2]
    return int(params[len(params) - 1 - index])

def init_memory(codes):
    global memory
    memory = dict()
    for index, code in enumerate(codes):
        memory[index] = code

def get_input():
    return 5

def print_value(pos):
    print(memory[pos])


def init_memory(codes):
    global memory
    memory = dict()
    for index, code in enumerate(codes):
        memory[index] = code

def get_input():
    return 5

def print_value(pos):
    print(memory[pos])

def read_value(pos):
    memory[pos] = get_input()

def add(input1, input2, pos3):
    memory[pos3] = input1 + input2

def multiply(input1, input2, pos3):
    memory[pos3] = input1 * input2

def less_than(input1, input2, pos3):
    if input1 < input2:
        memory[pos3] = 1
    else:
        memory[pos3] = 0

def equal_to(input1, input2, pos3):
    if input1 == input2:
        memory[pos3] = 1
    else:
        memory[pos3] = 0

def get_input_lines():
    with open('input', 'r') as f:
        return f.readlines()

def get_input_line_characters():
    chars = []
    with open('input', 'r') as f:
        for line in f:
            line_chars = []
            for c in line.strip():
                line_chars.append(c)
            chars.append(line_chars)
    return chars

def get_input_chars():
    l = []
    with open('input', 'r') as f:
        for c in f.read():
            l.append(c)
    return l

def get_input_chars_no_newlines():
    l = []
    with open('input', 'r') as f:
        for line in f:
            for c in line.strip():
                l.append(c)
    return l

def get_input_stripped_lines():
    return get_input_separated('\n')

def get_input_separated(delim):
    with open('input', 'r') as f:
        return f.read().split(delim)

def get_input_raw():
    with open('input', 'r') as f:
        return f.read()

def get_input_lines_ints():
    ints = []
    with open('input', 'r') as f:
        for line in f:
            ints.append(int(line.strip()))
    return ints

if __name__ == '__main__':
    main()
