/*
 * @Description: Geo entities including continent and country entities
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-04-01 16:53:26
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-04-03 00:52:56
 */

class ContinentEntity {
  final List<Continent> continents;
  final int total;

  ContinentEntity({
    required this.continents,
    required this.total,
  });

  bool get hasContinents => continents.isNotEmpty;

  Map<String, dynamic> toJson() {
    return {
      'continents': continents.map((x) => x.toJson()).toList(),
      'total': total,
    };
  }

  factory ContinentEntity.fromJson(Map<String, dynamic> json) {
    return ContinentEntity(
      continents: (json['continents'] as List<dynamic>)
          .map((x) => Continent.fromJson(x as Map<String, dynamic>))
          .toList(),
      total: int.parse(json['total'].toString()),
    );
  }
}

class Continent {
  final int id;
  final String name;
  final String nameLocal;
  final String nameEn;
  final String nameCn;
  final String code;
  final bool isHot;

  Continent({
    required this.id,
    required this.name,
    required this.nameLocal,
    required this.nameEn,
    required this.nameCn,
    required this.code,
    this.isHot = false,
  });

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'nameLocal': nameLocal,
      'nameEn': nameEn,
      'nameCn': nameCn,
      'code': code,
      'isHot': isHot,
    };
  }

  factory Continent.fromJson(Map<String, dynamic> json) {
    return Continent(
      id: int.parse(json['id'].toString()),
      name: json['name'] as String,
      nameLocal: json['nameLocal'] as String,
      nameEn: json['nameEn'] as String,
      nameCn: json['nameCn'] as String,
      code: json['code'] as String,
      isHot: json['isHot'] as bool? ?? false,
    );
  }
}

class CountriesEntity {
  final List<Country> countries;
  final int total;

  CountriesEntity({
    required this.countries,
    required this.total,
  });

  bool get hasCountries => countries.isNotEmpty;

  Map<String, dynamic> toJson() {
    return {
      'countries': countries.map((x) => x.toJson()).toList(),
      'total': total,
    };
  }

  factory CountriesEntity.fromJson(Map<String, dynamic> json) {
    return CountriesEntity(
      countries: (json['countries'] as List<dynamic>)
          .map((x) => Country.fromJson(x as Map<String, dynamic>))
          .toList(),
      total: int.parse(json['total'].toString()),
    );
  }
}

class Country {
  final int id;
  final String name;
  final String nameLocal;
  final String nameEn;
  final String nameCn;
  final String code;
  final int continentId;
  final Continent? continent;
  final String imageUrl;
  final bool isHot;

  Country({
    required this.id,
    required this.name,
    required this.nameLocal,
    required this.nameEn,
    required this.nameCn,
    required this.code,
    required this.continentId,
    this.continent,
    required this.imageUrl,
    this.isHot = false,
  });

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'nameLocal': nameLocal,
      'nameEn': nameEn,
      'nameCn': nameCn,
      'code': code,
      'continentId': continentId,
      'continent': continent?.toJson(),
      'imageUrl': imageUrl,
      'isHot': isHot,
    };
  }

  factory Country.fromJson(Map<String, dynamic> json) {
    return Country(
      id: int.parse(json['id'].toString()),
      name: json['name'] as String,
      nameLocal: json['nameLocal'] as String,
      nameEn: json['nameEn'] as String,
      nameCn: json['nameCn'] as String,
      code: json['code'] as String,
      continentId: int.parse(json['continentId'].toString()),
      continent: json['continent'] != null
          ? Continent.fromJson(json['continent'] as Map<String, dynamic>)
          : null,
      imageUrl: json['imageUrl'] as String,
      isHot: json['isHot'] as bool? ?? false,
    );
  }
}

class StatesEntity {
  final List<GeoState> states;
  final int total;

  StatesEntity({
    required this.states,
    required this.total,
  });

  Map<String, dynamic> toJson() {
    return {
      'states': states.map((x) => x.toJson()).toList(),
      'total': total,
    };
  }

  factory StatesEntity.fromJson(Map<String, dynamic> json) {
    return StatesEntity(
      states: (json['states'] as List<dynamic>)
          .map((x) => GeoState.fromJson(x as Map<String, dynamic>))
          .toList(),
      total: int.parse(json['total'].toString()),
    );
  }
}

class GeoState {
  final int id;
  final String name;
  final String nameEn;
  final String nameCn;
  final String nameLocal;
  final String code;
  final int countryId;
  final Country? country;
  final String imageUrl;

  GeoState({
    required this.id,
    required this.name,
    required this.nameEn,
    required this.nameCn,
    required this.nameLocal,
    required this.code,
    required this.countryId,
    this.country,
    this.imageUrl = '',
  });

  factory GeoState.fromJson(Map<String, dynamic> json) {
    return GeoState(
      id: json['id'] as int,
      name: json['name'] as String,
      nameEn: json['nameEn'] as String,
      nameCn: json['nameCn'] as String,
      nameLocal: json['nameLocal'] as String,
      code: json['code'] as String,
      countryId: json['countryId'] as int,
      country:
          json['country'] != null ? Country.fromJson(json['country']) : null,
      imageUrl: json['imageUrl'] as String? ?? '',
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'nameEn': nameEn,
      'nameCn': nameCn,
      'nameLocal': nameLocal,
      'code': code,
      'countryId': countryId,
      'country': country?.toJson(),
      'imageUrl': imageUrl,
    };
  }
}

class CitiesEntity {
  final List<City> cities;
  final int total;

  CitiesEntity({
    required this.cities,
    required this.total,
  });

  Map<String, dynamic> toJson() {
    return {
      'cities': cities.map((x) => x.toJson()).toList(),
      'total': total,
    };
  }

  factory CitiesEntity.fromJson(Map<String, dynamic> json) {
    return CitiesEntity(
      cities: (json['cities'] as List<dynamic>)
          .map((x) => City.fromJson(x as Map<String, dynamic>))
          .toList(),
      total: int.parse(json['total'].toString()),
    );
  }
}

class City {
  final int id;
  final String name;
  final String nameLocal;
  final String nameEn;
  final String nameCn;
  final String code;
  final int stateId;
  final GeoState? state;
  final String imageUrl;

  City({
    required this.id,
    required this.name,
    required this.nameLocal,
    required this.nameEn,
    required this.nameCn,
    required this.code,
    required this.stateId,
    this.state,
    this.imageUrl = '',
  });

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'nameLocal': nameLocal,
      'nameEn': nameEn,
      'nameCn': nameCn,
      'code': code,
      'stateId': stateId,
      'state': state?.toJson(),
      'imageUrl': imageUrl,
    };
  }

  factory City.fromJson(Map<String, dynamic> json) {
    return City(
      id: int.parse(json['id'].toString()),
      name: json['name'] as String? ?? '',
      nameLocal: json['nameLocal'] as String? ?? '',
      nameEn: json['nameEn'] as String? ?? '',
      nameCn: json['nameCn'] as String? ?? '',
      code: json['code'] as String? ?? '',
      stateId: int.parse((json['stateId'] ?? 0).toString()),
      state: json['state'] != null
          ? GeoState.fromJson(json['state'] as Map<String, dynamic>)
          : null,
      imageUrl: json['imageUrl'] as String? ?? '',
    );
  }

  @override
  String toString() {
    return 'City(id: $id, name: $name, nameLocal: $nameLocal, nameEn: $nameEn, nameCn: $nameCn, code: $code, stateId: $stateId, state: ${state?.toString()}, imageUrl: $imageUrl)';
  }
}
