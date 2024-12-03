void main() {
  String text = "this is a text";
  print(text.capitalise());
}

// Creating an extension
// Extensions work as additional methods to easily operate on objects
extension StringExtensions on String {
  String capitalise() {
    return this[0].toUpperCase() + this.substring(1);
  }
}
