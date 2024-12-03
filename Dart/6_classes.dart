void main() {
  Car seltos = Car("KIA", "Red"); // Creating the class objects
  seltos.modifyVersion = 10; // Using setter
  seltos.getInfo(); // Calling the class methods
  print(seltos.version); // Using getter
}

class Car {
  // Class variables
  String? brand;
  String? color;

  // Private Variables
  // Note that private variables in dart are mentioned with an underscore prefixed to the variable name
  int _version = 4;

  // Class Constructor
  // Car(String brand, String color) {
  //   this.brand = brand;
  //   this.color = color;
  // }

  // Alternatively
  Car(this.brand, this.color);

  // Getters and Setters
  int get version => _version;

  set modifyVersion(int v) {
    _version = v;
  }

  // Class Methods
  void getInfo() {
    brand != null ? print("Brand : ${this.brand}") : print("No Brand");
    color != null ? print("Color : ${this.color}") : print("No Color");
    print("Version : ${this._version}");
  }

  void changeColor(String newColor) {
    this.color = newColor;
  }
}
