void main() {
  // For loop
  String str = "Hello World";
  for (int i = 0; i <= str.length; i++) {
    print(str.substring(
        0, i)); // This method returns the substring in the given range
  }

  // While loop
  int i = 0;
  while (i < str.length) {
    print(str[i]);
    i++;
  }

  // Do while loop
  i = 0;
  do {
    print(i);
    i++;
  } while (i < 5);

  // Break and Continue
  for (int i = 0; i < 10; i++) {
    if (i.isEven) {
      continue;
    }
    print(i);
  }

  for (int i = 0; i < 10; i+=2) {
    if (i > 5) {
      break;
    }
    print(i);
  }
}
