/*
 * @Description: Geo entities including continent and country entities
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-04-01 16:53:26
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-04-03 00:52:56
 */

class ContientEntity {
  final List<Continent> continents;
  final int total;

  ContientEntity({
    required this.continents,
    required this.total,
  });

  bool get hasContinents => continents.isNotEmpty;

  factory ContientEntity.fromJson(Map<String, dynamic> json) {
    final continentsJson = json['continents'] as List<dynamic>;
    final continents = continentsJson
        .map((item) => Continent.fromJson(item as Map<String, dynamic>))
        .toList();
    final total = int.tryParse(json['total'].toString()) ?? 0;
    return ContientEntity(
      continents: continents,
      total: total,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'continents': continents.map((continent) => continent.toJson()).toList(),
      'total': total,
    };
  }
}

class Continent {
  final int id;
  final String name;
  final String nameEn;
  final String nameCn;
  final String nameLocal;
  final String code;

  Continent({
    required this.id,
    required this.name,
    required this.nameEn,
    required this.nameCn,
    required this.nameLocal,
    required this.code,
  });

  factory Continent.fromJson(Map<String, dynamic> json) {
    return Continent(
      id: json['id'] as int,
      name: json['name'] as String,
      nameEn: json['nameEn'] as String,
      nameCn: json['nameCn'] as String,
      nameLocal: json['nameLocal'] as String,
      code: json['code'] as String,
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
    };
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

  factory CountriesEntity.fromJson(Map<String, dynamic> json) {
    final countriesJson = json['countries'] as List<dynamic>;
    final countries = countriesJson
        .map((item) => Country.fromJson(item as Map<String, dynamic>))
        .toList();
    final total = int.tryParse(json['total'].toString()) ?? 0;
    return CountriesEntity(
      countries: countries,
      total: total,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'countries': countries.map((country) => country.toJson()).toList(),
      'total': total,
    };
  }
}

class Country {
  final int id;
  final String code;
  final String name;
  final String nameCn;
  final String nameEn;
  final int continentId;
  final Continent? contient;
  final String nameLocal;
  final String? imageUrl;

  Country({
    required this.id,
    required this.code,
    required this.name,
    required this.nameCn,
    required this.nameEn,
    required this.continentId,
    this.contient,
    required this.nameLocal,
    this.imageUrl,
  });

  factory Country.fromJson(Map<String, dynamic> json) {
    // Currently, this picture is set as the default. If there is an update to the default picture later, it will be replaced.
    final imageUrl = json['imageUrl'] as String?;
    final validImageUrl =
        imageUrl?.isNotEmpty == true ? imageUrl : "img_china.png";
    final contientJson = json['contient'] as Map<String, dynamic>?;
    final contient =
        contientJson != null ? Continent.fromJson(contientJson) : null;
    return Country(
      id: json['id'] as int,
      code: json['code'] as String,
      name: json['name'] as String,
      nameCn: json['nameCn'] as String,
      nameEn: json['nameEn'] as String,
      continentId: json['continentId'] as int,
      contient: contient,
      nameLocal: json['nameLocal'] as String,
      imageUrl: validImageUrl,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'code': code,
      'name': name,
      'nameCn': nameCn,
      'nameEn': nameEn,
      'continentId': continentId,
      'contient': contient?.toJson(),
      'nameLocal': nameLocal,
      'imageUrl': imageUrl,
    };
  }
}
