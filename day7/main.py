from itertools import permutations
import time
import queue

memory = dict()

stdout = None
stdin = queue.Queue()

def main():
    data = ''
    with open('input', 'r') as f:
        data = f.read()
    codes = list(map(lambda x: int(x), data.split(',')))

    perms = permutations([0,1,2,3,4])

    max_signal = 0
    for perm in list(perms):
        output_signal = 0
        for i in perm:
            add_input(i)
            add_input(output_signal)
            init_memory(codes)
            compute()
            output_signal = stdout
        print('out:', stdout)
        if stdout > max_signal:
            max_signal = stdout
    print(max_signal)

def compute():
    index = 0
    while memory[index] != 99:
        if get_op_code(memory[index]) == 1:
            param0 = get_op_code_param(memory[index], 0)
            param1 = get_op_code_param(memory[index], 1)
            param2 = get_op_code_param(memory[index], 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = memory[memory[index+1]]
            else:
                input0 = memory[index+1]

            if param1 == 0:
                input1 = memory[memory[index+2]]
            else:
                input1 = memory[index+2]

            if param2 == 0:
                input2 = memory[index+3]
            else:
                print('THIS SHOULD NEVER HAPPEN FOR ADD')
                input2 = memory[index+3]

            add(input0, input1, input2)
            index += 4
        if get_op_code(memory[index]) == 2:
            param0 = get_op_code_param(memory[index], 0)
            param1 = get_op_code_param(memory[index], 1)
            param2 = get_op_code_param(memory[index], 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = memory[memory[index+1]]
            else:
                input0 = memory[index+1]

            if param1 == 0:
                input1 = memory[memory[index+2]]
            else:
                input1 = memory[index+2]

            if param2 == 0:
                input2 = memory[index+3]
            else:
                print('THIS SHOULD NEVER HAPPEN FOR MULTIPLY')
                input2 = memory[index+3]

            multiply(input0, input1, input2)
            index += 4
        if get_op_code(memory[index]) == 3:
            read_value(memory[index+1])
            index += 2
        if get_op_code(memory[index]) == 4:
            print_value(memory[index+1])
            index += 2
        if get_op_code(memory[index]) == 5: # jump if true
            param0 = get_op_code_param(memory[index], 0)
            param1 = get_op_code_param(memory[index], 1)

            input0 = None
            input1 = None

            if param0 == 0:
                input0 = memory[memory[index+1]]
            else:
                input0 = memory[index+1]

            if param1 == 0:
                input1 = memory[memory[index+2]]
            else:
                input1 = memory[index+2]

            if input0 != 0:
                index = input1
            else:
                index += 3
        if get_op_code(memory[index]) == 6: # jump if false
            param0 = get_op_code_param(memory[index], 0)
            param1 = get_op_code_param(memory[index], 1)

            input0 = None
            input1 = None

            if param0 == 0:
                input0 = memory[memory[index+1]]
            else:
                input0 = memory[index+1]

            if param1 == 0:
                input1 = memory[memory[index+2]]
            else:
                input1 = memory[index+2]

            if input0 == 0:
                index = input1
            else:
                index += 3
        if get_op_code(memory[index]) == 7: # less than
            param0 = get_op_code_param(memory[index], 0)
            param1 = get_op_code_param(memory[index], 1)
            param2 = get_op_code_param(memory[index], 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = memory[memory[index+1]]
            else:
                input0 = memory[index+1]

            if param1 == 0:
                input1 = memory[memory[index+2]]
            else:
                input1 = memory[index+2]

            if param2 == 0:
                input2 = memory[index+3]
            else:
                print('THIS SHOULD NEVER HAPPEN FOR LT')
                input2 = memory[index+3]

            less_than(input0, input1, input2)
            index += 4
        if get_op_code(memory[index]) == 8: # equal to
            param0 = get_op_code_param(memory[index], 0)
            param1 = get_op_code_param(memory[index], 1)
            param2 = get_op_code_param(memory[index], 2)

            input0 = None
            input1 = None
            input2 = None

            if param0 == 0:
                input0 = memory[memory[index+1]]
            else:
                input0 = memory[index+1]

            if param1 == 0:
                input1 = memory[memory[index+2]]
            else:
                input1 = memory[index+2]

            if param2 == 0:
                input2 = memory[index+3]
            else:
                print('THIS SHOULD NEVER HAPPEN FOR EQ')
                input2 = memory[index+3]

            equal_to(input0, input1, input2)
            index += 4

    # return memory[0]

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
    return stdin.get_nowait()

def print_value(pos):
    global stdout
    print(memory[pos])
    stdout = memory[pos]

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

def add_input(value):
    global stdin
    stdin.put(value)

if __name__ == '__main__':
    main()
