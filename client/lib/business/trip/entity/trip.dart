class Trip {
  final String id;
  final DateTime createdAt;
  final DateTime updatedAt;
  final bool status;
  final String title;
  final String description;
  final DateTime startTs;
  final DateTime endTs;
  final List<DailyTrip> dailyTrips;

  Trip({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.status,
    required this.title,
    required this.description,
    required this.startTs,
    required this.endTs,
    required this.dailyTrips,
  });

  factory Trip.fromJson(Map<String, dynamic> json) {
    return Trip(
      id: json['id'] as String,
      createdAt: DateTime.parse(json['createdAt'] as String),
      updatedAt: DateTime.parse(json['updatedAt'] as String),
      status: json['status'] as bool,
      title: json['title'] as String,
      description: json['description'] as String,
      startTs: DateTime.parse(json['startTs'] as String),
      endTs: DateTime.parse(json['endTs'] as String),
      dailyTrips: (json['dailyTrips'] as List<dynamic>)
          .map((e) => DailyTrip.fromJson(e as Map<String, dynamic>))
          .toList(),
    );
  }
}

class DailyTrip {
  final String id;
  final DateTime createdAt;
  final DateTime updatedAt;
  final String tripId;
  final int day;
  final DateTime date;
  final String notes;
  final List<DailyItinerary> dailyItineraries;

  DailyTrip({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.tripId,
    required this.day,
    required this.date,
    required this.notes,
    required this.dailyItineraries,
  });

  factory DailyTrip.fromJson(Map<String, dynamic> json) {
    return DailyTrip(
      id: json['id'] as String,
      createdAt: DateTime.parse(json['createdAt'] as String),
      updatedAt: DateTime.parse(json['updatedAt'] as String),
      tripId: json['tripId'] as String,
      day: json['day'] as int,
      date: DateTime.parse(json['date'] as String),
      notes: json['notes'] as String,
      dailyItineraries: (json['dailyItineraries'] as List<dynamic>)
          .map((e) => DailyItinerary.fromJson(e as Map<String, dynamic>))
          .toList(),
    );
  }
}

class DailyItinerary {
  final String id;
  final String tripId;
  final String dailyTripId;
  final String poiId;
  final String notes;
  final DateTime createdAt;
  final DateTime updatedAt;
  final POI poi;

  DailyItinerary({
    required this.id,
    required this.tripId,
    required this.dailyTripId,
    required this.poiId,
    required this.notes,
    required this.createdAt,
    required this.updatedAt,
    required this.poi,
  });

  factory DailyItinerary.fromJson(Map<String, dynamic> json) {
    return DailyItinerary(
      id: json['id'] as String,
      tripId: json['tripId'] as String,
      dailyTripId: json['dailyTripId'] as String,
      poiId: json['poiId'] as String,
      notes: json['notes'] as String,
      createdAt: DateTime.parse(json['createdAt'] as String),
      updatedAt: DateTime.parse(json['updatedAt'] as String),
      poi: POI.fromJson(json['poi'] as Map<String, dynamic>),
    );
  }
}

class POI {
  final String id;
  final String name;
  final String nameEn;
  final String nameCn;
  final String description;
  final String address;
  final double latitude;
  final double longitude;
  final String type;
  final String category;
  final double rating;
  final String recommendedDurationMinutes;
  final String city;
  final String state;
  final String country;
  final String nameLocal;

  POI({
    required this.id,
    required this.name,
    required this.nameEn,
    required this.nameCn,
    required this.description,
    required this.address,
    required this.latitude,
    required this.longitude,
    required this.type,
    required this.category,
    required this.rating,
    required this.recommendedDurationMinutes,
    required this.city,
    required this.state,
    required this.country,
    required this.nameLocal,
  });

  factory POI.fromJson(Map<String, dynamic> json) {
    return POI(
      id: json['id'] as String,
      name: json['name'] as String,
      nameEn: json['nameEn'] as String,
      nameCn: json['nameCn'] as String,
      description: json['description'] as String,
      address: json['address'] as String,
      latitude: json['latitude'] as double,
      longitude: json['longitude'] as double,
      type: json['type'] as String,
      category: json['category'] as String,
      rating: json['rating'] as double,
      recommendedDurationMinutes: json['recommendedDurationMinutes'] as String,
      city: json['city'] as String,
      state: json['state'] as String,
      country: json['country'] as String,
      nameLocal: json['nameLocal'] as String,
    );
  }
} 