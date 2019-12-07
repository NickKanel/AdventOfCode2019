from itertools import permutations
import time
import queue

def main():
    data = ''
    with open('input', 'r') as f:
        data = f.read()
    codes = list(map(lambda x: int(x), data.split(',')))

    perms = permutations([5,6,7,8,9])

    max_signal = 0
    for perm in list(perms):
        amplifiers = list()
        for i in range(0, 5):
            amp = Amp(codes)
            amp.add_input(perm[i])
            amplifiers.append(amp)

        amp_num = 0
        output_signal = 0
        while not amplifiers[4].stopped:
            amp = amplifiers[amp_num]
            amp.add_input(output_signal)
            amp.compute()
            output_signal = amp.stdout.get_nowait()
            amp_num += 1
            amp_num %= 5
        if output_signal > max_signal:
            max_signal = output_signal
    print(max_signal)

class Amp:
    def __init__(self, codes):
        self.memory = codes.copy()
        self.index = 0
        self.stopped = False
        self.stdin = queue.Queue()
        self.stdout = queue.Queue()

    def compute(self):
        while True:
            if get_op_code(self.memory[self.index]) == 99:
                self.stopped = True
                break
            if get_op_code(self.memory[self.index]) == 1:
                param0 = get_op_code_param(self.memory[self.index], 0)
                param1 = get_op_code_param(self.memory[self.index], 1)
                param2 = get_op_code_param(self.memory[self.index], 2)

                input0 = None
                input1 = None
                input2 = None

                if param0 == 0:
                    input0 = self.memory[self.memory[self.index+1]]
                else:
                    input0 = self.memory[self.index+1]

                if param1 == 0:
                    input1 = self.memory[self.memory[self.index+2]]
                else:
                    input1 = self.memory[self.index+2]

                if param2 == 0:
                    input2 = self.memory[self.index+3]
                else:
                    print('THIS SHOULD NEVER HAPPEN FOR ADD')
                    input2 = self.memory[self.index+3]

                self.add(input0, input1, input2)
                self.index += 4
            if get_op_code(self.memory[self.index]) == 2:
                param0 = get_op_code_param(self.memory[self.index], 0)
                param1 = get_op_code_param(self.memory[self.index], 1)
                param2 = get_op_code_param(self.memory[self.index], 2)

                input0 = None
                input1 = None
                input2 = None

                if param0 == 0:
                    input0 = self.memory[self.memory[self.index+1]]
                else:
                    input0 = self.memory[self.index+1]

                if param1 == 0:
                    input1 = self.memory[self.memory[self.index+2]]
                else:
                    input1 = self.memory[self.index+2]

                if param2 == 0:
                    input2 = self.memory[self.index+3]
                else:
                    print('THIS SHOULD NEVER HAPPEN FOR MULTIPLY')
                    input2 = self.memory[self.index+3]

                self.multiply(input0, input1, input2)
                self.index += 4
            if get_op_code(self.memory[self.index]) == 3:
                if self.stdin.qsize() == 0:
                    break
                self.read_value(self.memory[self.index+1])
                self.index += 2
            if get_op_code(self.memory[self.index]) == 4:
                self.print_value(self.memory[self.index+1])
                self.index += 2
            if get_op_code(self.memory[self.index]) == 5: # jump if true
                param0 = get_op_code_param(self.memory[self.index], 0)
                param1 = get_op_code_param(self.memory[self.index], 1)

                input0 = None
                input1 = None

                if param0 == 0:
                    input0 = self.memory[self.memory[self.index+1]]
                else:
                    input0 = self.memory[self.index+1]

                if param1 == 0:
                    input1 = self.memory[self.memory[self.index+2]]
                else:
                    input1 = self.memory[self.index+2]

                if input0 != 0:
                    self.index = input1
                else:
                    self.index += 3
            if get_op_code(self.memory[self.index]) == 6: # jump if false
                param0 = get_op_code_param(self.memory[self.index], 0)
                param1 = get_op_code_param(self.memory[self.index], 1)

                input0 = None
                input1 = None

                if param0 == 0:
                    input0 = self.memory[self.memory[self.index+1]]
                else:
                    input0 = self.memory[self.index+1]

                if param1 == 0:
                    input1 = self.memory[self.memory[self.index+2]]
                else:
                    input1 = self.memory[self.index+2]

                if input0 == 0:
                    self.index = input1
                else:
                    self.index += 3
            if get_op_code(self.memory[self.index]) == 7: # less than
                param0 = get_op_code_param(self.memory[self.index], 0)
                param1 = get_op_code_param(self.memory[self.index], 1)
                param2 = get_op_code_param(self.memory[self.index], 2)

                input0 = None
                input1 = None
                input2 = None

                if param0 == 0:
                    input0 = self.memory[self.memory[self.index+1]]
                else:
                    input0 = self.memory[self.index+1]

                if param1 == 0:
                    input1 = self.memory[self.memory[self.index+2]]
                else:
                    input1 = self.memory[self.index+2]

                if param2 == 0:
                    input2 = self.memory[self.index+3]
                else:
                    print('THIS SHOULD NEVER HAPPEN FOR LT')
                    input2 = self.memory[self.index+3]

                self.less_than(input0, input1, input2)
                self.index += 4
            if get_op_code(self.memory[self.index]) == 8: # equal to
                param0 = get_op_code_param(self.memory[self.index], 0)
                param1 = get_op_code_param(self.memory[self.index], 1)
                param2 = get_op_code_param(self.memory[self.index], 2)

                input0 = None
                input1 = None
                input2 = None

                if param0 == 0:
                    input0 = self.memory[self.memory[self.index+1]]
                else:
                    input0 = self.memory[self.index+1]

                if param1 == 0:
                    input1 = self.memory[self.memory[self.index+2]]
                else:
                    input1 = self.memory[self.index+2]

                if param2 == 0:
                    input2 = self.memory[self.index+3]
                else:
                    print('THIS SHOULD NEVER HAPPEN FOR EQ')
                    input2 = self.memory[self.index+3]

                self.equal_to(input0, input1, input2)
                self.index += 4
    
    def get_input(self):
        return self.stdin.get_nowait()

    def print_value(self, pos):
        print(self.memory[pos])
        self.stdout.put(self.memory[pos])

    def read_value(self, pos):
        self.memory[pos] = self.get_input()

    def add(self, input1, input2, pos3):
        self.memory[pos3] = input1 + input2

    def multiply(self, input1, input2, pos3):
        self.memory[pos3] = input1 * input2

    def less_than(self, input1, input2, pos3):
        if input1 < input2:
            self.memory[pos3] = 1
        else:
            self.memory[pos3] = 0

    def equal_to(self, input1, input2, pos3):
        if input1 == input2:
            self.memory[pos3] = 1
        else:
            self.memory[pos3] = 0

    def add_input(self, value):
        self.stdin.put(value)

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

if __name__ == '__main__':
    main()
