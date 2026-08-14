// Go package: main

// Go import: fmt

fn countTo(limit: i32) {
    let mut count = 0;
    while count < limit {
        println!(count);
        count = count + 1;
    }
}

fn countdown() {
    let mut count = 3;
    while count > 0 {
        println!(count);
        count = count - 1;
    }
}

fn main() {
    countTo(3);
    countdown();
}
