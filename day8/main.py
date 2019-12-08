def main():
    data = ''
    with open('input', 'r') as f:
        data = f.readline().strip()

    parsed_data = parse_data(data)

    layer_count = len(parsed_data) / (25 * 6)
    layer_count = int(layer_count)

    layers = []
    for i in range(0, layer_count):
        layer = read_layer(25, 6, parsed_data)
        layers.append(layer)

    fewest = 1000000
    fewest_layer = layers[0]
    for layer in layers:
        num = count_digits(layer, 0)
        if num < fewest:
            fewest = num
            fewest_layer = layer

    ones = count_digits(fewest_layer, 1)
    twos = count_digits(fewest_layer, 2)

    print(ones * twos)

    # image = []
    # for i in range(0, 6):
    #     row = []
    #     for j in range(0, 25):
    #         row.append(get_pixel_color(layers, i, j))
    #     image.append(row)

    # for row in image:
    #     print(row)

def get_pixel_color(layers, i, j):
    for k in range(0, 100):
        if layers[k][i][j] != 2:
            return layers[k][i][j]
    return 2

def count_digits(layer, num):
    count = 0
    for row in layer:
        for digit in row:
            if digit == num:
                count += 1
    return count

def parse_data(data):
    nums = []
    for c in data:
        nums.append(int(c))
    return nums

def read_layer(width, height, data):
    rows = []
    for i in range(0, height):
        row = []
        for j in range(0, width):
            value = data.pop(0)
            row.append(value)
        rows.append(row)
    return rows

if __name__ == '__main__':
    main()
