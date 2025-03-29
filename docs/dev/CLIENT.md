<div align="center">
  <img src="../logo.png" alt="Logo" width="290" height="251" />
</div>

<div align="center">

| [English](CLIENT.md) | [中文简体](CLIENT.zh-CN.md) |

</div>

# Client Development Guidelines

## Code Style Guidelines

### Naming Conventions
- Variables/Parameters: camelCase (e.g., userName, isLoading)
- Folders/Dart Files: lowercase_with_underscore (e.g., home_main)
- Constants: UPPER_CASE_WITH_UNDERSCORE (e.g., MAX_ITEMS)
- Resource Files: type_name (e.g., ic_name.png)
- Classes: PascalCase
  - Page files: suffix with Page (e.g., UserInfoDetailPage)
  - Dialog files: suffix with Dialog
  - Other components: generally suffix with Widget
- Utility Classes: suffix with Utils
- Method Names:

  | Method | Description |
  |--------|-------------|
  | initXXX() | Initialization |
  | isXX() | Returns boolean |
  | getXX() | Returns a value |
  | setXXX() | Sets a value |
  | updateXXX() | Updates data |
  Other method names should intuitively express their purpose

### Code Formatting
- Indentation: tab indentation
- Line Breaks:
  - One line between classes
  - One line between methods
  - No line breaks within logical blocks (use comments instead of line breaks)
- Import Statements: Can use aliases (e.g., import 'package:flutter/material.dart' as m;)

### Adaptive Scaling Guidelines
This project uses `flutter_screenutil` for adaptive scaling. Therefore, when setting font sizes, use `number.sp`, and when setting width, height, padding, or margin, use `number.w` and `number.h`.

In the project, all numbers need to have the corresponding suffix. One thing to note is that in the case of a square, the same unit should be used, such as `width: 100.w, height: 100.w` or `width: 100.h, height: 100.h`.

### Example
```dart
TextStyle agreementTextStyle = TextStyle(
  color: BaseColor.bg,
  fontSize: 16.sp,
);

child: Container(
  width: 285.w,
  height: 52.h,
);
```

### Dart Language Guidelines
- String Interpolation: Prefer String interpolation over + concatenation
  ```dart
  // Recommended
  Text('Hello, ${userName}!');
  ```
- Conditional Statements: Use ternary operators for simple conditions
  ```dart
  bool isValid = length >= 10 ? true : false;
  ```
- Collection Operations: Use isEmpty instead of length == 0

## Project Structure Guidelines

### Basic Directory Structure
```
assets/  
├── images/         # Images
├── voice/          # Voice files
lib/
├── app/            # Constants configuration
│   ├── apis/       # API paths
│   ├── events/     # Event handlers
│   ├── foundation/ # Mixin files
│   ├── notifier/   # State management
│   ├── constants.dart
│   ├── routes.dart
├── business/       # Feature modules
│   ├── auth/       # Authentication module
│   │   ├── dialog/    # Dialog components
│   │   ├── entity/    # Entity classes
│   │   ├── page/      # Page files
│   │   ├── service/   # Network requests
│   │   └── widgets/   # Encapsulated components
│   ├── common/     # Common module
├── common/         # Common configuration
│   ├── dio/        # Network request config
│   ├── material/    # Base components
│   ├── utils/       # Utility classes
├── main.dart       # Main entry
```

### Modular Design
- Organize directories by feature modules
- Each feature module should be developed independently
- Encapsulate and document components/methods used more than twice

## Performance Optimization Guidelines

### Reduce Widget Rebuilds
- Mark immutable components with const constructor
- Avoid setState or time-consuming operations in build method
- Use RepaintBoundary to wrap frequently updated local areas

### State Management Optimization
- Choose efficient state management solutions (e.g., provider, bloc)
- Avoid over-segmentation of states, update as needed

### Asynchronous Operations
- Handle async tasks using async/await or FutureBuilder
- Avoid time-consuming operations on the main thread

### Resource Management
- Use encapsulated image components that can cache common resources
- Prevent memory leaks by properly closing streams or timers

## Collaboration and Version Control

### Git Guidelines
- Follow commit message guidelines as specified in the git project

### Code Review
- Check for unused or unreferenced code before committing
- Add comments for complex logic
- Document field meanings in entity classes 
