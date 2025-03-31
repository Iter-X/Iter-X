import 'package:flutter/material.dart';

class PoiSearchService extends ChangeNotifier {
  bool _isLoading = false;
  String _currentCity = '洛杉矶Los Angeles';
  List<PoiItem> _poiList = [];
  String _searchQuery = '';
  final Set<PoiItem> _selectedPois = {};

  bool _isPoiSelected(PoiItem poi) {
    return _selectedPois.any((item) => item.name == poi.name);
  }

  bool get isLoading => _isLoading;

  String get currentCity => _currentCity;

  List<PoiItem> get poiList => _poiList;

  String get searchQuery => _searchQuery;

  Set<PoiItem> get selectedPois => _selectedPois;

  void togglePoiSelection(PoiItem poi) {
    if (_isPoiSelected(poi)) {
      _selectedPois.removeWhere((item) => item.name == poi.name);
    } else {
      _selectedPois.add(poi);
    }
    notifyListeners();
  }

  void removeSelectedPoi(PoiItem poi) {
    _selectedPois.removeWhere((item) => item.name == poi.name);
    notifyListeners();
  }

  Future<void> searchPoi(String query) async {
    _searchQuery = query;
    _isLoading = true;
    notifyListeners();

    try {
      // TODO: Replace with actual API call
      await Future.delayed(Duration(seconds: 1));
      _poiList = [
        PoiItem(
          name: '好莱坞环球影城',
          englishName: 'Universal Studios Hollywood',
          duration: '平均游玩时间: 1天',
          popularity: '60%的人选择去',
          reviews: '28条评论',
          rating: 4,
          imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/us.png',
        ),
        PoiItem(
          name: '格里菲斯天文台',
          englishName: 'Griffith Observatory',
          duration: '平均游玩时间: 3h',
          popularity: '80%的人选择去',
          reviews: '999+评论',
          rating: 5,
          imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/cd.png',
        ),
        PoiItem(
          name: '好莱坞标志',
          englishName: 'Hollywood Sign',
          duration: '平均游玩时间: 1h',
          popularity: '50%的人选择去',
          reviews: '3条评论',
          rating: 3,
          imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/us.png',
        ),
        PoiItem(
          name: '圣莫妮卡沙滩',
          englishName: 'Santa Monica Beach',
          duration: '平均游玩时间: 1.5h',
          popularity: '82%的人选择去',
          reviews: '998条评论',
          rating: 3.5,
          imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/cd.png',
        ),
        PoiItem(
          name: '比弗利山庄超长名称超长名称超长名称超长名称超长名称超长名称',
          englishName:
              'Griffith Observatory Long Name Long Name Long Name Long Name Long Name Long Name',
          duration: '平均游玩时间: 1.5h',
          popularity: '82%的人选择去',
          reviews: '998条评论',
          rating: 3.5,
          imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/cd.png',
        ),
        PoiItem(
          name: '加利福尼亚大学洛杉矶分校',
          englishName: 'University Of California, Los Angeles',
          duration: '平均游玩时间: 1.5h',
          popularity: '82%的人选择去',
          reviews: '998条评论',
          rating: 2,
          imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/us.png',
        ),
      ];
    } catch (e) {
      debugPrint('Error searching POIs: $e');
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  void setCurrentCity(String city) {
    _currentCity = city;
    notifyListeners();
  }
}

class PoiItem {
  final String name;
  final String englishName;
  final String duration;
  final String popularity;
  final String reviews;
  final double rating;
  final String imageUrl;

  PoiItem({
    required this.name,
    required this.englishName,
    required this.duration,
    required this.popularity,
    required this.reviews,
    required this.rating,
    required this.imageUrl,
  });
}
