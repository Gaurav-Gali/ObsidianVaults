void main() {
  String? name = null;
  int age = 18;

  if (age >= 18) {
    print("Hey ${name ?? "Human"}, you can vote");
  } else if (age == 17) {
    print("Wait for another year");
  } else {
    print("Wait for ${18 - age} more years");
  }

  // Use '&&' for logical and operation
  // Use '||' for logical or operation
  if (age.isEven && age.isOdd) {
    print("Impossible");
  }

  if (age.isEven || age.isOdd) {
    print("Possible");
  }

  // Ternary Operators
  String? name1 = age > 0 ? "Harry" : null;
  print("Hello there, ${name1 ?? "User"}" );

  // Switch Case Statements
  int day = 7;
  switch (day) {
    case 1:
    case 2:
    case 3:
    case 4:
    case 5:
      print("Weekday");
      break;
    default:
      print("Weekend");
  }

  // Using logical statements with switch case
  switch (age) {
    case 18 when name != null:
      print("Valid person");
      break;
    default:
      print("In human");
  }
}
