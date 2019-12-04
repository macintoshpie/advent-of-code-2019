import sys

input = sys.stdin.readlines()[0].strip()
input = [int(v) for v in input.split(',')]

# change idx 1 -> 12 and idx 2 -> 2 to restore gravity assist program
input[1], input[2] = 12, 2

for i in range(0, len(input), 4):
  opcode = input[i]
  if opcode == 99:
    break
  val1_idx = input[i+1]
  val2_idx = input[i+2]
  store_idx = input[i+3]
  if opcode == 1:
    input[store_idx] = input[val1_idx] + input[val2_idx]
  elif opcode == 2:
    input[store_idx] = input[val1_idx] * input[val2_idx]
  else:
    print(f"Unrecognized opcode at position {i}: {opcode}")

print(input[0])
