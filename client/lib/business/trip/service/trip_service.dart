import 'package:client/app/apis/trip.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/business/trip/entity/collaborator.dart';
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

  // 合作者数据
  List<Collaborator> _collaborators = [];
  List<Collaborator> get collaborators => _collaborators;
  bool _loadingCollaborators = false;
  bool get loadingCollaborators => _loadingCollaborators;

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
        // 获取行程数据后，获取合作者
        await fetchCollaborators(tripId);
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

  // 获取合作者
  Future<void> fetchCollaborators(String tripId) async {
    _loadingCollaborators = true;
    notifyListeners();

    try {
      final HttpResultBean result = await Http.instance.get(
        TripApi.listTripCollaboratorsUrl(tripId),
      );
      
      if (result.isSuccess() && result.data['collaborators'] != null) {
        _collaborators = (result.data['collaborators'] as List)
            .map((e) => Collaborator.fromJson(e))
            .toList();
      }

      notifyListeners();
    } catch (e) {
      print('Error fetching collaborators: $e');
      rethrow;
    } finally {
      _loadingCollaborators = false;
      notifyListeners();
    }
  }

  // 清空数据（用于刷新或重置）
  void clearData() {
    _trip = null;
    _collaborators = [];
    notifyListeners();
  }
}
