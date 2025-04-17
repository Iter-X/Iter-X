import 'package:client/app/apis/trip.dart';
import 'package:client/business/trip/entity/collaborator.dart';
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

  // 更新行程
  Future<void> updateTrip({
    required String tripId,
    String? title,
    String? description,
    DateTime? startTs,
    DateTime? endTs,
    int? duration,
    bool? status,
  }) async {
    _isLoading = true;
    notifyListeners();

    try {
      final HttpResultBean result = await Http.instance.put(
        TripApi.updateTripUrl(tripId),
        data: {
          if (title != null) 'title': title,
          if (description != null) 'description': description,
          if (startTs != null) 'start_ts': startTs.toIso8601String(),
          if (endTs != null) 'end_ts': endTs.toIso8601String(),
          if (duration != null) 'duration': duration,
          if (status != null) 'status': status,
        },
      );

      if (result.isSuccess() && result.data['trip'] != null) {
        _trip = Trip.fromJson(result.data['trip']);
      }

      notifyListeners();
    } catch (e) {
      print('Error updating trip: $e');
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

  // Move itinerary item to a new position
  Future<void> moveItineraryItem({
    required String tripId,
    required String dailyTripId,
    required String itineraryId,
    required int newDay,
    required int newIndex,
  }) async {
    try {
      final HttpResultBean result = await Http.instance.post(
        TripApi.moveItineraryItemUrl(tripId),
        data: {
          'daily_trip_id': dailyTripId,
          'daily_itinerary_id': itineraryId,
          'day': newDay,
          'after_order': newIndex,
        },
      );

      if (result.isSuccess() && _trip != null) {
        // Update local state with the returned data
        if (result.data['trip'] != null) {
          final updatedTrip = Trip.fromJson(result.data['trip']);

          // Update only the changed daily trips
          for (var updatedDailyTrip in updatedTrip.dailyTrips) {
            final existingDailyTrip = _trip!.dailyTrips.firstWhere(
              (dt) => dt.id == updatedDailyTrip.id,
              orElse: () => updatedDailyTrip,
            );

            // Update the daily itineraries
            existingDailyTrip.dailyItineraries.clear();
            existingDailyTrip.dailyItineraries
                .addAll(updatedDailyTrip.dailyItineraries);
          }

          notifyListeners();
        }
      } else if (!result.isSuccess()) {
        // Only refresh data if the move operation failed
        await fetchTripData(tripId: tripId);
      }
    } catch (e) {
      print('Error moving itinerary item: $e');
      // Only refresh data if there was an error
      await fetchTripData(tripId: tripId);
      rethrow;
    }
  }
}
