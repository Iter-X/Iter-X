class TripApi {
  // 创建行程
  static const String createTripFromCard = '/api/v1/trips/card';
  static const String createTripManually = '/api/v1/trips/manual';
  static const String createTripFromExternalLink =
      '/api/v1/trips/external-link';
  // 获取行程
  static const String getTrip = '/api/v1/trips/{id}';
  // 更新行程
  static const String updateTrip = '/api/v1/trips/{id}';
  // 删除行程
  static const String deleteTrip = '/api/v1/trips/{id}';
  // 获取行程列表
  static const String listTrips = '/api/v1/trips';
  // 创建每日行程
  static const String createDailyTrip = '/api/v1/trips/{trip_id}/daily';
  // 获取每日行程
  static const String getDailyTrip = '/api/v1/trips/{trip_id}/daily/{daily_id}';
  // 更新每日行程
  static const String updateDailyTrip =
      '/api/v1/trips/{trip_id}/daily/{daily_id}';
  // 删除每日行程
  static const String deleteDailyTrip =
      '/api/v1/trips/{trip_id}/daily/{daily_id}';
  // 获取每日行程列表
  static const String listDailyTrips = '/api/v1/trips/{trip_id}/daily';
}
