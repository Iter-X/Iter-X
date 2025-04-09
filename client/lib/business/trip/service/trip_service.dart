import 'package:client/app/apis/trip.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/common/dio/http.dart';
import 'package:client/common/dio/http_result_bean.dart';
import 'package:flutter/material.dart';

class TripService extends ChangeNotifier {
  // 加载状态
  bool _isLoading = false;
  bool get isLoading => _isLoading;

  // 当前行程数据
  Trip? _trip;
  Trip? get trip => _trip;

  // 获取行程数据
  Future<void> fetchTripData({required String tripId}) async {
    _isLoading = true;
    notifyListeners();

    try {
      final HttpResultBean result = await Http.instance.get(
        TripApi.getTripDetailUrl(tripId),
      );
      
      if (result.isSuccess() && result.data['trip'] != null) {
        _trip = Trip.fromJson(result.data['trip']);
      }

      notifyListeners();
    } catch (e) {
      print('Error fetching trip data: $e');
      rethrow;
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  // 清空数据（用于刷新或重置）
  void clearData() {
    _trip = null;
    notifyListeners();
  }
}
