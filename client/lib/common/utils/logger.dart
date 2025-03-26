import 'package:flutter/foundation.dart';
import 'package:logger/logger.dart';

enum LogLevel {
  debug,
  info,
  warn,
  error,
}

enum Environment {
  development,
  testing,
  production,
}

// Base logger class
class BaseLogger {
  BaseLogger._();

  static Environment _environment = _getDefaultEnvironment();
  static LogLevel _minLevel = LogLevel.debug;
  static bool _enable = true;

  static Environment _getDefaultEnvironment() {
    if (kReleaseMode) {
      return Environment.production;
    } else if (kDebugMode) {
      return Environment.development;
    } else {
      return Environment.testing;
    }
  }

  static final _logger = Logger(
    printer: PrettyPrinter(
      // Stack trace depth for normal logs, 0 means no stack trace
      methodCount: 0,
      // Stack trace depth for error logs
      errorMethodCount: 8,
      // Maximum width of log output
      lineLength: 5000,
      // Enable colored output
      colors: false,
      // Show emojis for different log levels
      printEmojis: true,
      // Time format setting, shows current time and app running duration
      dateTimeFormat: DateTimeFormat.onlyTimeAndSinceStart,
      // Disable box drawing by default
      noBoxingByDefault: true,
    ),
  );

  static bool get enable => _enable;
  static Environment get environment => _environment;
  static LogLevel get minLevel => _minLevel;

  static void setEnvironment(Environment env) {
    _environment = env;
    _updateMinLevel();
  }

  static void setEnable(bool enable) {
    _enable = enable;
  }

  static void _updateMinLevel() {
    switch (_environment) {
      case Environment.development:
        _minLevel = LogLevel.debug;
        break;
      case Environment.testing:
        _minLevel = LogLevel.info;
        break;
      case Environment.production:
        _minLevel = LogLevel.warn;
        break;
    }
  }

  static bool _shouldLog(LogLevel level) {
    if (!_enable) return false;
    return level.index >= _minLevel.index;
  }

  static void d(dynamic message, [dynamic error, StackTrace? stackTrace]) {
    if (_shouldLog(LogLevel.debug)) {
      _logger.d(message, error: error, stackTrace: stackTrace);
    }
  }

  static void i(dynamic message, [dynamic error, StackTrace? stackTrace]) {
    if (_shouldLog(LogLevel.info)) {
      _logger.i(message, error: error, stackTrace: stackTrace);
    }
  }

  static void w(dynamic message, [dynamic error, StackTrace? stackTrace]) {
    if (_shouldLog(LogLevel.warn)) {
      _logger.w(message, error: error, stackTrace: stackTrace);
    }
  }

  static void e(dynamic message, [dynamic error, StackTrace? stackTrace]) {
    if (_shouldLog(LogLevel.error)) {
      _logger.e(message, error: error, stackTrace: stackTrace);
    }
  }

  static void close() {
    _logger.close();
  }
}
