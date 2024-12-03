void main() {
  try {
    print(10 ~/ 0);
  } on UnsupportedError catch (e) {
    print("Cant divide by zero : $e");
  } finally {
    print("Exception handled");
  }
}
