void main() async {
  // While using Future, always handle errors

  // Using await
  try {
    final result = await greetafterdelay();
    print(result);
  } catch (err) {
    print("error");
  }

  // Using the then method
  greetafterdelay().then((val) {
    print(val);
  }).catchError((err) {
    print("promise error");
  });
}

Future<String> greetafterdelay() async {
  return "Hey";
}

// Alternatively
Future<String> greetafterdelay2() {
  return Future(() {
    return "Hello";
  });
}
