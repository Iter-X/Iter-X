import 'package:logger/logger.dart';

// 日志输出
class BaseLogger {
  BaseLogger._();

  static final _logger = Logger(
    printer: PrettyPrinter(
      methodCount: 0,
      // 打印方法行数
      errorMethodCount: 8,
      // number of method calls if stacktrace is provided
      lineLength: 5000,
      // width of the output
      colors: true,
      // Colorful log messages
      printEmojis: false,
      // Print an emoji for each log message
      printTime: false,
      // Should each log print contain a timestamp
      noBoxingByDefault: true,
    ),
  );

  // 日志开关
  static bool _enable = true;

  static bool get enable => _enable;

  static setEnable(bool enable) {
    _enable = enable;
  }

  static v(dynamic message, [dynamic error, StackTrace? stackTrace]) {
    if (_enable) {
      _logger.d(message);
    }
  }

  static e(dynamic message, [dynamic error, StackTrace? stackTrace]) {
    if (_enable) {
      _logger.e(message);
    }
  }

  static close() {
    _logger.close();
  }
}
