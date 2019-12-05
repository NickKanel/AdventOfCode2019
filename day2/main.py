memory = dict()

def main():
    data = ''
    with open('input', 'r') as f:
        data = f.read()
    codes = list(map(lambda x: int(x), data.split(',')))

    init_memory(codes)

    found = False
    for i in range(0, 100):
        memory[1] = i
        for j in range(0, 100):
            memory[2] = j
            result = compute(codes)
            if result == 19690720:
                print(memory[1])
                print(memory[2])
                break
        init_memory(codes)

def compute(codes):
    index = 0
    while codes[index] != 99:
        if codes[index] == 1:
            add(codes[index+1], codes[index+2], codes[index+3])
            index += 4
        if codes[index] == 2:
            multiply(codes[index+1], codes[index+2], codes[index+3])
            index += 4
    return memory[0]

def init_memory(codes):
    global memory
    memory = dict()
    for index, code in enumerate(codes):
        memory[index] = code

def add(pos1, pos2, pos3):
    memory[pos3] = memory[pos1] + memory[pos2]

def multiply(pos1, pos2, pos3):
    memory[pos3] = memory[pos1] * memory[pos2]
    

if __name__ == '__main__':
    main()
