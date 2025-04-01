/*
 * @Description: 
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-04-01 16:53:26
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-04-01 18:21:36
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
  final String code;

  Continent({
    required this.id,
    required this.name,
    required this.nameEn,
    required this.nameCn,
    required this.code,
  });

  factory Continent.fromJson(Map<String, dynamic> json) {
    return Continent(
      id: json['id'] as int,
      name: json['name'] as String,
      nameEn: json['nameEn'] as String,
      nameCn: json['nameCn'] as String,
      code: json['code'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'nameEn': nameEn,
      'nameCn': nameCn,
      'code': code,
    };
  }
}
