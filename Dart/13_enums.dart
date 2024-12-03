void main() {
  final type = Options.medium;

  switch (type) {
    case Options.small:
      print("SMALL");
      break;
    case Options.large:
      print("Large");
      break;
    case Options.medium:
      print("MEDIUM");
      print(type.name);
      break;
    default:
      print("NO SUCH TYPE");
  }
}

enum Options {
  small("sup"),
  medium("fun"),
  large("hola");

  final String name;
  const Options(this.name);
}
