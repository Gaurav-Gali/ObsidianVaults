void main() {
  Child c = Child();
  c.eyeColor = "Brown";
  c.height = 173;
  c.info();
}

class Parent {
  String? eyeColor;
  int? height;
}

class Child extends Parent {
  void info() {
    print("Eye color : ${eyeColor}");
    print("Eye color : ${height}");
  }
}
