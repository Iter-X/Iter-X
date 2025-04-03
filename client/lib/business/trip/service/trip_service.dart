import 'package:flutter/material.dart';

class TripService extends ChangeNotifier {
  // 加载状态
  bool _isLoading = false;

  bool get isLoading => _isLoading;

  // 行程标题
  final String _title = '美西自驾10日游';

  String get title => _title;

  // 行程参与者列表
  final List<TripParticipant> _participants = [];

  List<TripParticipant> get participants => _participants;

  // 待规划的景点列表
  final List<UnplannedSpot> _unplannedSpots = [];

  List<UnplannedSpot> get unplannedSpots => _unplannedSpots;

  // 行程日程列表
  final List<TripDay> _days = [];

  List<TripDay> get days => _days;

  // 模拟从服务器获取数据
  Future<void> fetchTripData() async {
    _isLoading = true;
    notifyListeners();

    try {
      // 模拟网络延迟
      await Future.delayed(const Duration(milliseconds: 500));

      // 模拟获取参与者数据
      _participants.addAll([
        TripParticipant(
          id: '1',
          name: '用户1',
          avatar: 'https://robohash.org/1.png?size=200x200',
        ),
        TripParticipant(
          id: '2',
          name: '用户2',
          avatar: 'https://robohash.org/2.png?size=200x200',
        ),
        TripParticipant(
          id: '3',
          name: '用户3',
          avatar: 'https://robohash.org/3.png?size=200x200',
        ),
        TripParticipant(
          id: '4',
          name: '用户4',
          avatar: 'https://robohash.org/4.png?size=200x200',
        ),
        TripParticipant(
          id: '5',
          name: '用户5',
          avatar: 'https://robohash.org/5.png?size=200x200',
        ),
        TripParticipant(
          id: '6',
          name: '用户6',
          avatar: 'https://robohash.org/6.png?size=200x200',
        ),
      ]);

      // 模拟获取待规划景点数据
      _unplannedSpots.addAll([
        UnplannedSpot(
          id: '1',
          name: '优胜美地国家公园',
          description: '国家公园',
        ),
        UnplannedSpot(
          id: '2',
          name: '洛杉矶',
          description: '城市',
        ),
      ]);

      // 模拟获取行程日数据
      _days.addAll([
        TripDay(
          id: '1',
          day: 'Day 1',
          date: '2025/01/25 周六',
          cities: ['厦门', '旧金山'],
        ),
        TripDay(
          id: '2',
          day: 'Day 2',
          date: '2025/01/26 周日',
          cities: ['旧金山'],
          spots: ['金门公园', '双峰', '渔人码头', '金门大桥'],
        ),
        TripDay(
          id: '3',
          day: 'Day 3',
          date: '2025/01/27 周一',
          cities: ['旧金山', '红杉树国家公园'],
          spots: ['斯坦福大学', '红杉树国家公园', 'Apple', 'Google'],
        ),
      ]);

      notifyListeners();
    } catch (e) {
      // Handle any exceptions that might occur during data fetching
      print('Error fetching trip data: $e');
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  // 清空数据（用于刷新或重置）
  void clearData() {
    _participants.clear();
    _unplannedSpots.clear();
    _days.clear();
    notifyListeners();
  }

  // 添加参与者
  void addParticipant(TripParticipant participant) {
    _participants.add(participant);
    notifyListeners();
  }

  // 移除参与者
  void removeParticipant(TripParticipant participant) {
    _participants.remove(participant);
    notifyListeners();
  }

  // 添加待规划景点
  void addUnplannedSpot(UnplannedSpot spot) {
    _unplannedSpots.add(spot);
    notifyListeners();
  }

  // 移除待规划景点
  void removeUnplannedSpot(UnplannedSpot spot) {
    _unplannedSpots.remove(spot);
    notifyListeners();
  }

  // 添加行程日
  void addDay(TripDay day) {
    _days.add(day);
    notifyListeners();
  }

  // 移除行程日
  void removeDay(TripDay day) {
    _days.remove(day);
    notifyListeners();
  }
}

// 行程参与者模型
class TripParticipant {
  final String id;
  final String name;
  final String avatar;

  TripParticipant({
    required this.id,
    required this.name,
    required this.avatar,
  });
}

// 待规划景点模型
class UnplannedSpot {
  final String id;
  final String name;
  final String description;

  UnplannedSpot({
    required this.id,
    required this.name,
    required this.description,
  });
}

// 行程日模型
class TripDay {
  final String id;
  final String day;
  final String date;
  final List<String> cities; // 主要城市列表
  final List<String>? spots; // 景点列表

  TripDay({
    required this.id,
    required this.day,
    required this.date,
    required this.cities,
    this.spots,
  });
}
