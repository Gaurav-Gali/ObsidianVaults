void main() {
  // List l = [10, 20, 30];

  // Dynamic lists : stores values of any type
  List dyn = [1, 2, 3, "Apple", true];

  // Type specific lists : stores values of the mentioned type
  List<int> l2 = [1, 2, 3];

  print(dyn);
  print(l2);

  // Map method
  dyn.map((e) {
    print(e);
  }).toList();
}
