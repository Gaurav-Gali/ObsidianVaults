void main() {
  var s1 = Student<String>("Gaurav");
  s1.fn();
}

class Student<T> {
  final T data;

  Student(this.data);

  void fn() {
    print(data);
  }
}
