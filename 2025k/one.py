full_input = []
with open('input/one.txt', 'r') as f: # Open the file
    full_input = f.read().splitlines() # Turn the file into lines

sample_input = [
    "L68",
    "L30",
    "R48",
    "L5",
    "R60",
    "L55",
    "L1",
    "L99",
    "R14",
    "L82"
]

def get_instructions(input):
    instructions = [] # Create Empty List to add Instructions To
    for line in input: # Go through each line
        direction = line[0]  # First character (R or L) Example: > L < 68
        distance = int(line[1:])  # Rest of the string as intege L > 68 <
        
        instruction = (direction, distance) # Group direction and distance into instruction
        instructions.append(instruction) # Add instruction to list
    return instructions

instructions = get_instructions(sample_input)

current_point = 50

for direction, distance in instructions:
    print(f"{distance} to the {direction}")
     
