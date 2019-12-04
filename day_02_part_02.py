import sys

target = 19690720

input = sys.stdin.readlines()[0].strip()
input = [int(v) for v in input.split(',')]

def compute(commands):
  for i in range(0, len(commands), 4):
    opcode = commands[i]
    if opcode == 99:
      break
    val1_idx = commands[i+1]
    val2_idx = commands[i+2]
    store_idx = commands[i+3]
    if opcode == 1:
      commands[store_idx] = commands[val1_idx] + commands[val2_idx]
    elif opcode == 2:
      commands[store_idx] = commands[val1_idx] * commands[val2_idx]
    else:
      raise Exception(f"Unrecognized opcode at position {i}: {opcode}")
  return commands[0]

def findNounVerb(input):
  # just brute force it - 100 x 100 aint too bad bb
  for i in range(100):
    for j in range(100):
      commands = input[:]
      commands[1] = i
      commands[2] = j
      try:
        res = compute(commands)
        if res == target:
          return i, j
      except:
        continue


noun, verb = findNounVerb(input)
print(f"Solution: {noun} {verb}")
