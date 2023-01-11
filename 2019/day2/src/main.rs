
fn runProgram(mut program: Vec<i32>) -> i32 {
    let mut pointer = 0;
    let len = program.len();

    while pointer < len && program[pointer] != 99 {
        let result: i32 = match program[pointer] {
            1 => program[program[pointer + 1] as usize] + program[program[pointer + 2] as usize],
            2 => program[program[pointer + 1] as usize] * program[program[pointer + 2] as usize],
            _ => -1
        };

        // println!("{},{} [{}]={}",program[pointer + 1], program[pointer + 2], program[pointer+3], result);

        if result == -1 {
            println!("{}", program[pointer]);
            panic!("result of computation is -1");
        }

        let set_loc  = program[pointer + 3];
        program[set_loc as usize] = result;
        pointer += 4;
    }

    program[0]
}

fn main() {
    let input = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,19,6,23,2,13,23,27,1,27,13,31,1,9,31,35,1,35,9,39,1,39,5,43,2,6,43,47,1,47,6,51,2,51,9,55,2,55,13,59,1,59,6,63,1,10,63,67,2,67,9,71,2,6,71,75,1,75,5,79,2,79,10,83,1,5,83,87,2,9,87,91,1,5,91,95,2,13,95,99,1,99,10,103,1,103,2,107,1,107,6,0,99,2,14,0,0";
    let program: Vec<i32> = input.split(",").map(|x| x.parse::<i32>().unwrap()).collect();

    let result = runProgram(program.clone());
    println!("part one: {}", result);

    for i in 0..99 {
        for j in 0..99 {
            let mut copy = program.clone();
            copy[1]=i;
            copy[2]=j;
            if runProgram(copy) == 19690720 {
                println!("part two: noun: {}, verb: {}, result: {}", i, j, 100 * i + j);
                return
            }
        }
    }
}
