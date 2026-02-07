use std::fs::DirEntry;
use chrono::{Local, Utc};

// main.rs
fn main() {

    // -----------------------------
    // 1. Basic variable declaration
    // -----------------------------
    let a = 10;                 // immutable variable
    let mut b = 20;             // mutable variable

    b = 25;

    println!("a = {}", a);
    println!("b = {}", b);

    // -----------------------------
    // 2. Explicit type annotation
    // -----------------------------
    let x: i32 = 100;
    let y: f64 = 3.1415;
    let z: bool = true;
    let ch: char = 'R';

    println!("x = {}, y = {}, z = {}, ch = {}", x, y, z, ch);

    // -----------------------------
    // 3. Shadowing
    // -----------------------------
    let s = "hello";
    let s = s.len();   // new variable, different type

    println!("shadowed s = {}", s);

    // -----------------------------
    // 4. Constants
    // -----------------------------
    const MAX_USERS: u32 = 1000;
    println!("MAX_USERS = {}", MAX_USERS);

    // -----------------------------
    // 5. Integer types
    // -----------------------------
    let i1: i8 = -128;
    let i2: i16 = 32000;
    let i3: i32 = 2_000_000;
    let i4: i64 = 9_000_000_000;
    let i5: isize = 10; // arch dependent

    let u1: u8 = 255;
    let u2: u16 = 65535;
    let u3: u32 = 4_000_000;
    let u4: u64 = 18_000_000_000;
    let u5: usize = 20;

    println!("{}, {}, {}, {}, {}", i1, i2, i3, i4, i5);
    println!("{}, {}, {}, {}, {}", u1, u2, u3, u4, u5);

    // -----------------------------
    // 6. Floating point types
    // -----------------------------
    let f1: f32 = 2.5;
    let f2: f64 = 99.99;

    println!("f1 = {}, f2 = {}", f1, f2);

    // -----------------------------
    // 7. Boolean
    // -----------------------------
    let is_active: bool = false;
    println!("is_active = {}", is_active);

    // -----------------------------
    // 8. Character
    // -----------------------------
    let c1: char = 'A';

    println!("chars: {}", c1);

    // -----------------------------
    // 9. Tuples
    // -----------------------------
    let tup: (i32, f64, char) = (10, 3.14, 'x');
    let (t1, t2, t3) = tup;

    println!("tuple values: {}, {}, {}", t1, t2, t3);
    println!("tuple index access: {}", tup.1);

    // -----------------------------
    // 10. Arrays (fixed size)
    // -----------------------------
    let arr: [i32; 5] = [1, 2, 3, 4, 5];
    let same_value_arr = [0; 4];

    println!("arr[0] = {}", arr[0]);
    println!("same_value_arr = {:?}", same_value_arr);

    // -----------------------------
    // 11. String types
    // -----------------------------
    let s1: &str = "hello";          // string slice (immutable)
    let mut s2: String = String::from("hello");

    s2.push_str(" world");

    println!("s1 = {}", s1);
    println!("s2 = {}", s2);

    // -----------------------------
    // 12. Struct
    // -----------------------------

    #[derive(Debug)]
    struct User {
        username: String,
        email: String,
        age: i16
    }

    impl User {
        fn get_name(&self) -> &str {
            &self.username
        }
    }

    let user1 = User {
        username: String::from("Gaurav"),
        email: String::from("email@email.com"),
        age: 20
    };

    println!(">> {:#?} \n>> {:#?}", user1, user1.get_name()); // Pretty printed

    // -----------------------------
    // 13. Enum
    // -----------------------------
    enum Shape {
        Rectangle { width: f64, height: f64 },
        Circle { radius: f64 }
    }

    fn area(shape:Shape) -> f64 {
        // Pattern Matching
        let ar = match shape {
            Shape::Rectangle { width, height } => width * height,
            Shape::Circle { radius } => 3.14 * radius * radius,
        };

        ar
    }

    let rect = Shape::Rectangle {width: 3.0, height: 4.0};
    println!("area = {}", area(rect));

    // -----------------------------
    // 14. Option type
    // -----------------------------
    let some_number: Option<i32> = Some(5);
    let no_number: Option<i32> = None;

    println!("some_number = {:?}", some_number);
    println!("no_number = {:?}", no_number);

    fn find_value(x: i32) -> Option<i32> {
        if x >= 3 {
            return Some(3i32);
        } else {
            return None;
        }
    }

    let value = find_value(-5);
    println!("value = {:?}", value);

    // -----------------------------
    // 15. Result type
    // -----------------------------
    let success: Result<i32, &str> = Ok(200);
    let failure: Result<i32, &str> = Err("error occurred");

    println!("success = {:?}", success);
    println!("failure = {:?}", failure);

    // -----------------------------
    // 16. Crates (cargo add <crate-name>)
    // -----------------------------
    use chrono::Local;

    let local = Local::now();
    println!("utc = {:?}", local);

    // -----------------------------
    // 17. Loops
    // -----------------------------

    // -----------------------------
    // 1. loop (infinite loop)
    // -----------------------------
    let mut i = 0;

    loop {
        i += 1;

        if i == 3 {
            continue; // skip this iteration
        }

        if i == 5 {
            break; // exit loop
        }

        println!("loop i = {}", i);
    }

    // loop can return a value
    let result = loop {
        let x = 10;
        break x * 2;
    };

    println!("loop result = {}", result);

    // -----------------------------
    // 2. while loop
    // -----------------------------
    let mut n = 5;

    while n > 0 {
        println!("while n = {}", n);
        n -= 1;
    }

    // -----------------------------
    // 3. for loop (most common)
    // -----------------------------

    // range
    for i in 0..5 {
        println!("for i = {}", i);
    }

    // inclusive range
    for i in 1..=5 {
        println!("inclusive i = {}", i);
    }

    // iterating over array
    let arr = [10, 20, 30, 40];

    for val in arr {
        println!("array value = {}", val);
    }

    // -----------------------------
    // 4. Iterators with references
    // -----------------------------
    let mut numbers = [1, 2, 3, 4];

    // immutable borrow
    for n in numbers.iter() {
        println!("iter n = {}", n);
    }

    // mutable borrow
    for n in numbers.iter_mut() {
        *n *= 2;
    }

    println!("numbers after mut = {:?}", numbers);

    // -----------------------------
    // 5. Enumerate
    // -----------------------------
    let letters = ['a', 'b', 'c'];

    for (index, value) in letters.iter().enumerate() {
        println!("{} -> {}", index, value);
    }

    // -----------------------------
    // 6. Loop labels (nested loops)
    // -----------------------------
    'outer: for i in 1..=3 {
        for j in 1..=3 {
            if i == 2 && j == 2 {
                break 'outer;
            }
            println!("i = {}, j = {}", i, j);
        }
    }

    // -----------------------------
    // 7. for_each (functional style)
    // -----------------------------
    let nums = vec![1, 2, 3];

    nums.iter().for_each(|x| {
        println!("for_each {}", x);
    });

    // -----------------------------------

    // 1. Ownership & Moving
    let a1 = String::from("Hello");
    let a2 = a1.clone();

    // NOTE : Without clone(), the ownership of "Hello" will be moved to a2, and a1 becomes a dangling pointer, where it points to nothing since after a2 took ownership.
    println!("a1 = {:?}", a1);
    println!("a2 = {:?}", a2);

    // 2. Borrowing
    let s1 = String::from("Hello");
    let s2 = &s1;

    // NOTE : Here s2 is borrowing the value of s1, but the ownership of that value is still bound to s1 only.
    println!("s1 = {:?}", s1);
    println!("s2 = {:?}", s2);

    // 3. Mutable reference
    let mut k1 = String::from("Hello");
    println!("k1 = {:?}", k1);

    let k2 = &mut k1;
    k2.push_str(" World");

    println!("k2 = {:?}", k2);

}
