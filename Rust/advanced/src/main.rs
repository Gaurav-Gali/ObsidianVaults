use std::collections::HashMap;
use std::thread;
use std::time::Duration;
use std::sync::mpsc;

fn main() {
    // 1. Collections API

    // 1.1 Vectors
    let mut vec:Vec<i32> = Vec::new();

    for i in 1..=10 {
        vec.push(i);
    }

    println!("vec = {:?}", vec);

    let mut my_vec = vec![1,2,3,4,5];

    let mut i:usize = 0;

    while i < my_vec.len() {
        if my_vec[i] %2 != 0 {
            my_vec.remove(i);
        } else {
            i += 1;
        }
    }

    println!("my_vec = {:#?}", my_vec);

    // 1.2 HashMaps
    let mut map = HashMap::new();
    map.insert("a", 1);
    map.insert("b", 2);

    match map.get("a") {
        Some(value) => {
            println!("a = {}", value);
        },
        None => {
            println!("value not in map");
        }
    }

    // 2. Iterator

    // 2.1 .iter()
    let mut vec1 = vec![1,2,3,4,5];

    for x in vec1.iter() {
        println!("x = {}", x*10);
    }

    // OR
    let mut vec1_iter = vec1.iter();
    while let Some(x) = vec1_iter.next() {
        println!("x = {}", x);
    }

    // 2.2 .mut_iter()
    let mut vec2 = vec![1,2,3,4,5];

    for x in vec2.iter_mut() {
        *x *= 2;
    }

    println!("vec2 = {:?}", vec2);

    // 2.3 .into_iter()
    let mut vec3 = vec![1,2,3,4,5];

    let vec3_iter = vec3.into_iter();

    for x in vec3_iter {
        println!("(INTO ITER) x = {}", x);
    }

    // NOTE : INTO ITER will take complete ownership of the iterable and the original iterable cannot be used later.

    // 3. Consuming Adaptors

    // 3.1 .sum()
    let mut vec4 = vec![1,2,3,4,5];

    let sum:i32 = vec4.iter().sum();
    println!("sum = {}", sum);

    // 4. Iterator Adaptors
    let mut vec5 = vec![1,2,3,4,5];

    let iter = vec5.iter().map(|x| x*x);

    let iter2  = iter.filter(|x| *x % 2 == 0);

    for x in iter2 {
        println!("x = {}", x);
    }

    // 5. Collector - converts iterators back to vectors
    let mut vec6 = vec![1,2,3,4,5];

    vec6 = vec6.iter().filter(|x| *x % 2 == 0).map(|x| *x).collect();
    println!("vec6 = {:?}", vec6);

    // BETTER WAY
    let mut vec7 = vec![1,2,3,4,5];

    vec7 = vec7.into_iter()
        .filter(|x| *x % 2 == 0)
        .collect();

    println!("vec7 = {:#?}", vec7);

    // 6. Strings & Slices

    // 6.1 Strings
    let mut s = String::from("hello");
    s.push_str(", world!");
    s.replace_range(5..s.len(), "");
    println!("s = {}", s);

    // 6.2 Slices
    s = String::from("hello world");
    println!("s = {}", s);
    let sliced = &s[0..5];
    println!("sliced = {}", sliced);

    // 7. Traits
    pub trait Summary {
        fn summarize(&self) -> String {
            String::from("(Read more...)")
        }
    }

    #[derive(Debug)]
    pub struct NewsArticle {
        pub headline: String,
    }

    impl Summary for NewsArticle {

    }

    let article = NewsArticle {
        headline: String::from("Awesome"),
    };

    println!("{:#?}",article);
    println!("{:#?}",article.summarize());

    // 8. Lifetimes in Structs
    #[derive(Debug)]
    struct User<'a> {
        name: &'a str
    };
    let name = String::from("user-name");
    let user = User { name: &name };
    println!("Lifetime : {:#?}", user);

    // 9. Multithreading
    let handler = thread::spawn(|| {
        for i in 1..=10 {
            println!("(thread) {}", i);
            Duration::from_millis(10);
        }
    });

    handler.join().unwrap();

    // 10. Move closure - allows the thread to take ownership of external variables.
    let vec = vec![1,2,3,4,5];
    let handle = thread::spawn(move || {
        for i in vec.iter() {
            println!("(thread2) {}", i);
        }
    });
    handle.join().unwrap();

    // 11. Message Passer - allows thread to pass messages between threads
    let (tx, rx) = mpsc::channel();
    thread::spawn(move || {
        let val = String::from("hi");
        tx.send(val).unwrap();
    });

    let received = rx.recv();
    match received {
        Ok(val) => {
            println!("val = {}", val);
        },
        Err(_) => {
            println!("recv err");
        }
    }
}
