void main() {
  Map<String, String> jsonDummy = {
    'userName': "Gaurav",
    "age": "19",
    "month": "August"
  };

  var {'userName': userName, 'age': age} = jsonDummy;
  print(userName);
  print(age);

  // Pattern matching for classes
  final h1 = Human("Gaurav", "male");

  final Human(:name, :gender) = h1;
  print("Name : $name, Gender : $gender");
}

class Human {
  String name = "Gaurav";
  String gender = "male";

  Human(this.name, this.gender);
}
