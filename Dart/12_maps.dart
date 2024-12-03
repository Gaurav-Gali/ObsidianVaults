void main() {
  // Defining a map
  Map<String, int> marks = {"Gaurav": 100, "Arav": 99};

  // Accessing elements of a map
  print(marks['Gaurav']);

  // Adding elements to the map
  marks['newName'] = 100;

  // Adding multiple elements to the map
  marks.addAll({"name2": 100, "name3": 80, "name4": 99});

  // Iterating through a map
  for (int i = 0; i < marks.length; i++) {
    print("${i + 1}) ${marks.keys.toList()[i]} : ${marks.values.toList()[i]}");
  }

  // Using forEach method
  marks.forEach((key, value) {
    print("$key : $value");
  });
}
