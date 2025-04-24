class TripApi {
  static const String getTripDetail = '/api/v1/trips/{id}';
  static const String createTrip = '/api/v1/trips';
  static const String updateTrip = '/api/v1/trips/{id}';
  static const String deleteTrip = '/api/v1/trips/{id}';
  static const String listTrips = '/api/v1/trips';
  static const String listTripCollaborators = '/api/v1/trips/{id}/collaborators';
  static const String moveItineraryItem = '/api/v1/trips/{id}/itineraries/move';
  
  static String getTripDetailUrl(String id) => getTripDetail.replaceAll('{id}', id);
  static String updateTripUrl(String id) => updateTrip.replaceAll('{id}', id);
  static String deleteTripUrl(String id) => deleteTrip.replaceAll('{id}', id);
  static String listTripCollaboratorsUrl(String id) => listTripCollaborators.replaceAll('{id}', id);
  static String moveItineraryItemUrl(String id) => moveItineraryItem.replaceAll('{id}', id);
} 