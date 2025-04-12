import 'package:intl/intl.dart';

class DateUtil {
  static String formatDate(DateTime date) {
    final formatter = DateFormat('yyyy/MM/dd');
    return formatter.format(date);
  }

  static String formatDateWithWeekday(DateTime date) {
    final weekdays = ['周一', '周二', '周三', '周四', '周五', '周六', '周日'];
    final weekday = weekdays[date.weekday - 1];
    return '${formatDate(date)} $weekday';
  }

  static String formatDateRange(DateTime start, DateTime end) {
    final days = end.difference(start).inDays;
    return '${formatDate(start)} | ${days}天';
  }

  static String formatTime(DateTime date) {
    final formatter = DateFormat('HH:mm');
    return formatter.format(date);
  }
} 