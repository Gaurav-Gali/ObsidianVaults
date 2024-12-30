![Source](https://youtu.be/CzRQ9mnmh44?t=35053)

### Scaffold
The `Scaffold` widget is used as the root widget of the application.
- The `AppBar` contains a title, which can be customised with different configurations like `title`, `leading`, `actions`, and `bottomLeading`.
- The `body` property holds the content that will be displayed in the main area of the app.
- The `floatingActionButton` provides a button that can be used to trigger an action.

The `Scaffold` widget is highly customisable and can be extended with various other widgets such as `Drawer`, `BottomNavigationBar`, `TabBar`, etc., depending 
on your specific requirements.

### Text
 In Flutter, the `Text` widget is a fundamental component used to display text on the screen. It can be used to render plain text or formatted text, and it 
supports various styling options such as font size, color, alignment, and more.

Here's an example of how to use the `Text` widget:

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text('Flutter Text Example')),
        body: Center(
          child: Text(
            'Hello, Flutter!',
            style: TextStyle(
              fontSize: 24.0,
              color: Colors.blue,
              fontWeight: FontWeight.bold,
            ),
          ),
        ),
      ),
    );
  }
}
```

In this example:
- The `Text` widget is used to display the text "Hello, Flutter!".
- The `style` property of the `Text` widget allows for customization of the font size, color, and weight. You can also add other styling properties like 
`fontSize`, `color`, `fontWeight`, `fontStyle`, etc.

### Column
In Flutter, the `Column` widget is used to stack its child widgets vertically in a column direction. It is often used when you want to arrange multiple 
widgets vertically on the screen.
`mainAxisAlignment` property is set to `MainAxisAlignment.center`, which centres the child widgets horizontally in the vertical direction.
- `crossAxisAlignment` property is set to `CrossAxisAlignment.center`, which centres the child widgets vertically in the horizontal direction.
- `SizedBox` is used to add some space between the two text widgets.

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text('Column Example')),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text('Hello, Flutter!'),
              SizedBox(height: 20),
              Text('This is a column with two text widgets.'),
            ],
          ),
        ),
      ),
    );
  }
}
```

### ColoredBox
In Flutter, `ColoredBox` is a widget that allows you to create boxes with a specific color or gradient. It's similar to the `Container` widget but provides 
more flexibility and control over the color and gradient.
The `boxShadow` property is used to add a subtle shadow effect around the box.
- The `Text` widget inside `ColoredBox` displays the text "Hello, Flutter!" with white color.

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text('Color Box Example')),
        body: Center(
          child: ColoredBox(
            color: Colors.blue, // Set the background color to blue
            boxShadow: [
              BoxShadow(
                offset: Offset(2.0, 2.0), // Add a shadow effect
                blurRadius: 5.0,
              ),
            ],
            child: Text('Hello, Flutter!', style: TextStyle(color: Colors.white)),
          ),
        ),
      ),
    );
  }
}
```

### TextField
