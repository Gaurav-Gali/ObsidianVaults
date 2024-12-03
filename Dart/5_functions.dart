void main() {
  greet("Gaurav");
  print(add(1, 2));

  var (ad, sub, mul, div) = calc(1, 2);
  print(ad);
  print(sub);
  print(mul);
  print(div);

  // Calling a function with named arguments
  info("Address", name: "Gaurav", age: 19);

  // Calling function with named return
  final data = printData(1, 2);
  print(data.sum);
  print(data.concat);

  // Anonymous functions
  (String name) {
    print("Hello " + name);
  }("Gaurav");

  // Flat Arrow / Arrow functions
  void arrorFunc() => print("Arrow Func");
}

// Function with no return type
void greet(String name) {
  print("Hello " + name);
}

// Function with return type
int add(int a, int b) {
  return a + b;
}

// Function with multiple return types
(int, int, int, double) calc(int a, int b) {
  return (a + b, a - b, a * b, a / b);
}

// Multiple named elements can be returned
({int sum, String concat}) printData(int a, int b) {
  return (sum: a + b, concat: a.toString() + b.toString());
}

// Passing named arguments
// To add positional arguments, just add them before the named arguments
void info(String address, {required String name, required int age}) {
  print("Name : " + name);
  print("Age : " + age.toString());
}


// Syntax
/*
<dataType> functionName(**arguments) {
  Function content
  return <value of type dataType>
}
*/