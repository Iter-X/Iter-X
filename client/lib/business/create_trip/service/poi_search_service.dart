import 'package:client/app/apis/geo_api.dart';
import 'package:client/common/dio/http.dart';
import 'package:client/common/dio/http_result_bean.dart';
import 'package:flutter/material.dart';

class PoiSearchService extends ChangeNotifier {
  bool _isLoading = false;
  String _currentCity = '';
  int? _currentCityId;
  List<int>? _allCityIds; // 存储所有选中城市的ID
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

  Future<void> searchPoi(String query, {int? cityId}) async {
    _searchQuery = query;
    _isLoading = true;
    notifyListeners();

    try {
      final HttpResultBean result = await Http.instance.get(
        GeoApi.getPois,
        params: {
          if (cityId != null)
            'city_id': cityId
          else if (_allCityIds != null && _allCityIds!.isNotEmpty)
            'city_ids': _allCityIds,
          if (query.isNotEmpty) 'keyword': query,
          'size': 20,
          'page': 0,
        },
      );

      if (result.isSuccess()) {
        final List<dynamic> poisData = result.data['pois'] ?? [];
        _poiList = poisData.map((poi) => PoiItem.fromJson(poi)).toList();
      } else {
        _poiList = [];
      }
    } catch (e) {
      debugPrint('Error searching POIs: $e');
      _poiList = [];
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  void setCurrentCity(String cityName, {int? cityId, List<int>? allCityIds}) {
    _currentCity = cityName;
    _currentCityId = cityId;
    _allCityIds = allCityIds;
    searchPoi('', cityId: cityId);
    notifyListeners();
  }
}

class PoiItem {
  final String id;
  final String name;
  final String englishName;
  final String duration;
  final String popularity;
  final String reviews;
  final double rating;
  final String imageUrl;

  PoiItem({
    required this.id,
    required this.name,
    required this.englishName,
    required this.duration,
    required this.popularity,
    required this.reviews,
    required this.rating,
    required this.imageUrl,
  });

  factory PoiItem.fromJson(Map<String, dynamic> json) {
    final int durationMinutes = json['duration'] is String
        ? int.tryParse(json['duration'] as String) ?? 0
        : json['duration'] as int? ?? 0;
    final hours = (durationMinutes / 60).floor();
    final minutes = durationMinutes % 60;
    String durationText = '';
    if (hours > 0) {
      durationText = '平均游玩时间: $hours天';
    } else if (minutes > 0) {
      durationText = '平均游玩时间: $minutes分钟';
    } else {
      durationText = '平均游玩时间: 未知';
    }

    final double popularityValue = json['popularity'] is String
        ? double.tryParse(json['popularity'] as String) ?? 0
        : (json['popularity'] as num?)?.toDouble() ?? 0;
    final popularityText = '${(popularityValue).round()}%的人选择去';

    final reviewsCount = json['reviews_count'] is String
        ? int.tryParse(json['reviews_count'] as String) ?? 0
        : json['reviews_count'] as int? ?? 0;
    final reviewsText = reviewsCount > 999 ? '999+评论' : '$reviewsCount条评论';

    return PoiItem(
      id: json['id'] as String? ?? '',
      name: json['name_cn'] as String? ?? json['name'] as String? ?? '',
      englishName: json['name_en'] as String? ?? '',
      duration: durationText,
      popularity: popularityText,
      reviews: reviewsText,
      rating: json['rating'] is String
          ? double.tryParse(json['rating'] as String) ?? 0
          : (json['rating'] as num?)?.toDouble() ?? 0,
      imageUrl: json['image_url'] as String? ?? '',
    );
  }
}
