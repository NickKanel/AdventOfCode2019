import math

def main():
	total = 0
	with open('input', 'r') as f:
		for line in f:
			mass = int(line)
			module_fuel = get_fuel(mass)
			last_added = module_fuel
			while True:
				last_added = get_fuel(last_added)
				if last_added <= 0:
					break
				module_fuel += last_added
			total += module_fuel
		print(total)

def get_fuel(mass):
	return math.floor(mass/3) - 2

if __name__ == '__main__':
	main()
