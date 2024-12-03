void main() {
  Test t = Test();
  t.fn();
}

// Advantage of using mixin over inheritance is that , mixins don't introduce parent child relationships. 
mixin Stats {
  var version = 2;
}

class Test with Stats {
  void fn() {
    print(version);
  }
}
