void main() {
  // <data type> <variable name> = <value>;
  int a = -10;
  double b = 10.5;
  String c = "Hello";
  bool d = true;
  dynamic e = "Hi"; // Takes any data type
  String f = "${c} - ${e}";

  print(a);
  print(b);
  print(c);
  print(d);
  print(e);
  print(f + "\n");

  // Variable, Constants and Final
  var name = "John"; // Type is automatically inferred
  final time = DateTime.now(); // Value is Immutable and can be set in runtime.
  const pi = 3.14; // Value is Immutable and should be set in compile time.

  print(name);
  print(time);
  print(pi);

  // Optional  Variables
  int? age = null; // Can have an int value as well as null

  // Using the '!' operator to ensure that the value is not null
  // print(age!.isEven); // Throws an exception if age is null

  // Using the '?' operator to ensure that the value could be null
  print(age?.isEven);

  // If the value of an optional variable is null, then in order to return a valid output instead of null the '??' operator can be used.
  print(age ?? 10);
  print(age?.isEven ?? 1);
}
