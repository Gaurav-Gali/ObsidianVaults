![Source](https://youtu.be/w2noCpc4ieM?list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm)

[Official Documentation](https://go.dev/learn/)

- Go which is formally known as Go Lang was launched to the public in 2009 by Ken Thompson(Creator of B , C) and 2 others at google.
- Mainly used for building scalable applications like servers , API's , Cloud and used in DevOps.
- It is faster and efficient compared to other programming languages.
- Supports cross platform , since the go code is directly bundled to individual system binaries which makes it easy to run anywhere (windows , mac , linux).
- ==It is a strong and statically typed programming language.==
### Features of GoLang
1. Concurrency : This can be implemented very easily in go , and each task doesn't have to wait till another task is completed.
2. Garbage Collection : After a piece of memory is allocated , it automatically de-allocated using it's garbage collections system.
3. Static typing : This ensures that a variable is declared with a type and it's type sticks on til it's lifetime. This is to help catch errors at compile time.
4. Lightweight : Go-routines take up only 8kb of memory and a lot of it can be created.
5. Fast Compilation : Go has a fast compiler that compiles the code in no time and creates specific binaries.
6. Zero dependencies : Since go is not built on any framework or libraries , hence such dependencies don't have to be installed in the target machine.
7. Built-in support for testing : Go includes built-in support for writing and running tests making it easy for running and testing code.
8. Strong community : Go has a string community of developers who contribute to the language and it's libraries.

<hr>
### Practise Routine (100 Days Challenge)
Starting a 100-day coding challenge for Go (Go-lang) can help you develop skills incrementally. Here's a rough plan divided into phases that progress in complexity:

#### **Phase 1: Basics (Days 1–20)**
1. **Day 1-2:** Install Go, setup IDE, and run your first "Hello, World!"
2. **Day 3-5:** Variables, constants, and data types (int, float, string).
3. **Day 6-7:** Conditional statements (if-else, switch).
4. **Day 8-9:** Loops (for, range).
5. **Day 10-12:** Functions: defining, passing arguments, return values.
6. **Day 13-14:** Arrays and slices (basic operations).
7. **Day 15-16:** Maps (CRUD operations on maps).
8. **Day 17-18:** Structs and methods.
9. **Day 19-20:** Pointers in Go.

#### **Phase 2: Intermediate Concepts (Days 21–40)**
1. **Day 21-23:** Error handling and custom errors.
2. **Day 24-26:** Packages and imports (create your own package).
3. **Day 27-29:** Interfaces and polymorphism.
4. **Day 30-32:** Goroutines and concurrency basics.
5. **Day 33-35:** Channels for communication between goroutines.
6. **Day 36-37:** Mutex and sync package for handling race conditions.
7. **Day 38-40:** File handling (read/write files).

#### **Phase 3: Advanced Concepts (Days 41–60)**
1. **Day 41-43:** HTTP package: basic web server setup.
2. **Day 44-46:** JSON handling (encoding/decoding).
3. **Day 47-49:** Building a REST API with Go.
4. **Day 50-53:** Working with databases (SQL package and GORM).
5. **Day 54-56:** Testing: unit tests with Go’s testing package.
6. **Day 57-60:** Advanced error handling (context package and timeouts).

#### **Phase 4: Projects and Libraries (Days 61–80)**
1. **Day 61-65:** Create a CLI tool using Go’s `flag` package.
2. **Day 66-70:** Build a CRUD app with Go and a web framework like Gin or Echo.
3. **Day 71-75:** Explore Docker and containerize a Go application.
4. **Day 76-80:** Work with gRPC or WebSockets in Go.

#### **Phase 5: Special Topics and Final Projects (Days 81–100)**
1. **Day 81-85:** Build a microservice with Go.
2. **Day 86-90:** Explore Go modules for dependency management.
3. **Day 91-95:** Implement JWT authentication in a Go web app.
4. **Day 96-100:** Final project: Build a full-stack app with Go (backend) and a frontend of your choice.

### Daily Routine:
- **Code for at least 1 hour daily**: Stick to a consistent coding schedule.
- **Document**: Keep track of what you learn each day in a blog or journal.
- **GitHub**: Push code daily to a GitHub repository to maintain accountability.
- **Review**: Spend time revisiting previous concepts if needed.

Would you like more in-depth resources or challenges for each phase?

<hr>

### Leetcode Plan
Here's a structured **beginner-to-advanced LeetCode plan** specifically for Go, designed to guide you through progressively more difficult problems while improving your knowledge of Go’s features and algorithms. This plan is divided into **3 levels**: Beginner, Intermediate, and Advanced, each with specific focus areas.
### **Beginner (Days 1-30)**
**Goal**: Learn basic Go syntax and solve fundamental LeetCode problems.

#### **Focus Areas**:
- Go Syntax: Loops, conditionals, slices, arrays, strings, maps.
- Basic Data Structures: Arrays, HashMaps (maps in Go).
- Problem-solving techniques: Brute force, simple optimizations.

#### **Problems**:
1. **Array and String Manipulation**:
   - [Two Sum](https://leetcode.com/problems/two-sum/)
   - [Palindrome Number](https://leetcode.com/problems/palindrome-number/)
   - [Reverse String](https://leetcode.com/problems/reverse-string/)
   - [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)

2. **Simple Math Problems**:
   - [Roman to Integer](https://leetcode.com/problems/roman-to-integer/)
   - [Add Binary](https://leetcode.com/problems/add-binary/)
   - [Plus One](https://leetcode.com/problems/plus-one/)
   - [Power of Two](https://leetcode.com/problems/power-of-two/)

3. **Basic HashMap Usage**:
   - [Valid Anagram](https://leetcode.com/problems/valid-anagram/)
   - [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)

4. **Intro to Sorting and Searching**:
   - [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)
   - [Binary Search](https://leetcode.com/problems/binary-search/)

5. **Basic Recursion**:
   - [Factorial](https://leetcode.com/problems/factorial-trailing-zeroes/)
   - [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

#### **Go Tips**:
- Focus on understanding Go’s syntax: `for` loops, `range`, and working with slices.
- Practice writing functions and handling maps and arrays effectively.
- Learn how Go handles error checking, even though LeetCode doesn't emphasize errors much.

---

### **Intermediate (Days 31-60)**
**Goal**: Deepen understanding of algorithms, recursion, dynamic programming, and improve problem-solving techniques.

#### **Focus Areas**:
- Recursion and Backtracking.
- Sorting, Searching, and Binary Search.
- Dynamic Programming (DP): Memoization, Tabulation.

#### **Problems**:
1. **Recursion and Backtracking**:
   - [Subsets](https://leetcode.com/problems/subsets/)
   - [Permutations](https://leetcode.com/problems/permutations/)
   - [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)
   - [Combination Sum](https://leetcode.com/problems/combination-sum/)

2. **Dynamic Programming (Beginner)**:
   - [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)
   - [House Robber](https://leetcode.com/problems/house-robber/)
   - [Unique Paths](https://leetcode.com/problems/unique-paths/)
   - [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

3. **Sorting and Searching**:
   - [Search Insert Position](https://leetcode.com/problems/search-insert-position/)
   - [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)
   - [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

4. **Linked Lists**:
   - [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)
   - [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)
   - [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)

5. **Sliding Window Technique**:
   - [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
   - [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)

#### **Go Tips**:
- Focus on solving problems with efficient recursion.
- Use Go slices for dynamic data and experiment with memoization for DP.
- Practice pointer manipulation with linked list problems.
- Get used to writing Go code for sorting and binary search.
  
---

### **Advanced (Days 61-100)**
**Goal**: Master advanced algorithms and data structures, improve efficiency, and solve complex problems.

#### **Focus Areas**:
- Dynamic Programming (Advanced).
- Graph Algorithms: BFS, DFS, Dijkstra’s.
- Greedy Algorithms and Divide and Conquer.

#### **Problems**:
1. **Advanced Dynamic Programming**:
   - [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)
   - [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)
   - [Edit Distance](https://leetcode.com/problems/edit-distance/)
   - [Coin Change](https://leetcode.com/problems/coin-change/)

2. **Graph Algorithms (BFS, DFS)**:
   - [Number of Islands](https://leetcode.com/problems/number-of-islands/)
   - [Clone Graph](https://leetcode.com/problems/clone-graph/)
   - [Course Schedule](https://leetcode.com/problems/course-schedule/)
   - [Word Ladder](https://leetcode.com/problems/word-ladder/)

3. **Greedy Algorithms**:
   - [Jump Game](https://leetcode.com/problems/jump-game/)
   - [Gas Station](https://leetcode.com/problems/gas-station/)
   - [Task Scheduler](https://leetcode.com/problems/task-scheduler/)

4. **Divide and Conquer**:
   - [Merge K Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)
   - [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)
   - [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)
   
5. **Heap and Priority Queue**:
   - [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)
   - [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)

#### **Go Tips**:
- Implement graph algorithms using Go’s `map` for adjacency lists.
- Learn Go’s concurrency features (e.g., goroutines) to optimize solutions that can benefit from parallel execution.
- Optimize memory and performance by understanding how Go manages memory with slices and structs.
  
---

### **Final Stretch (Days 91-100)**
**Goal**: Solve difficult LeetCode problems with Go, participate in contests, and push the limits of your knowledge.

#### **Focus Areas**:
- System Design and Larger Problem Sets.
- Competitive programming with Go.

#### **Problems**:
1. [Hard Problems Category on LeetCode](https://leetcode.com/problemset/all/?difficulty=HARD)
   - **Dynamic Programming**: [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)
   - **Graph Algorithms**: [Minimum Spanning Tree Problems]
   - **Advanced Linked Lists**: [LRU Cache](https://leetcode.com/problems/lru-cache/)
2. **LeetCode Contests**: Participate in weekly contests to apply what you've learned in Go.

---

### **Additional Tips**:
- **Consistency**: Solve at least 1-2 problems daily, but aim to solve 3-4 problems in different categories on weekends.
- **Code Reviews**: After solving, review other users’ solutions to see more efficient or Go-idiomatic approaches.
- **Track Progress**: Maintain a GitHub repo for all your solutions and reflect on optimizations.
  
By following this plan, you'll be well-equipped to handle a variety of Go problems on LeetCode, from basic syntax to advanced algorithms.