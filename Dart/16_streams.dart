void main() {
  getData().listen((val) {
    print(val);
  }, onDone: () {
    print("Done with execution");
  }, onError: (e) {
    print("error : $e");
  });
}

// Defining a stream
Stream<int> getData() async* {
  for (int i = 1; i <= 5; i++) {
    yield i;
    if (i != 5) {
      await Future.delayed(Duration(seconds: 1));
    }
  }
} 
