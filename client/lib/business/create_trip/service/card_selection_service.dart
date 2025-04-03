import 'package:client/app/apis/geo_api.dart';
import 'package:client/common/dio/http.dart';
import 'package:client/common/dio/http_result_bean.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../entity/geo_entity.dart';
import 'dart:convert';

// The data cache time for continents and countries is one month, in milliseconds. Here it is set to 30 days.
const oneMonthInMilliseconds = 30 * 24 * 60 * 60 * 1000;

class CardSelectionService {
  static Future<ContientEntity?> getContinentsData() async {
    final prefs = await SharedPreferences.getInstance();
    final cachedData = prefs.getString('continents_data');
    final cacheTime = prefs.getInt('continents_cache_time');

    if (cachedData != null && cacheTime != null) {
      final currentTime = DateTime.now().millisecondsSinceEpoch;
      if (currentTime - cacheTime < oneMonthInMilliseconds) {
        try {
          final jsonData = json.decode(cachedData) as Map<String, dynamic>;
          return ContientEntity.fromJson(jsonData);
        } catch (e) {
          print('Failed to parse cached continents data: $e');
        }
      }
    }

    HttpResultBean result = await Http.instance.get(
      GeoApi.getContinents,
      isShowLoading: false,
    );
    if (result.isSuccess()) {
      final data = result.data as Map<String, dynamic>;
      final entity = ContientEntity.fromJson(data);
      final newContinent = {
        'id': 0,
        'name': '热门',
        'nameEn': 'hot',
        'nameCn': '热门',
        'nameLocal': '热门',
        'code': 'hot'
      };
      final newContinents = [
        Continent.fromJson(newContinent),
        ...entity.continents
      ];
      final total = int.tryParse(data['total'].toString()) ?? 0;
      final newEntity =
          ContientEntity(continents: newContinents, total: total + 1);
      prefs.setString('continents_data', json.encode(newEntity.toJson()));
      prefs.setInt(
          'continents_cache_time', DateTime.now().millisecondsSinceEpoch);

      return newEntity;
    }
    return null;
  }

  static Future<List<Map<String, dynamic>>> getCountriesData({
    int? continentId,
    int? size = 100,
    int? page = 1,
  }) async {
    final prefs = await SharedPreferences.getInstance();
    final cachedData = prefs.getString('all_countries_data');
    final cacheTime = prefs.getInt('countries_cache_time');

    if (cachedData != null && cacheTime != null) {
      final currentTime = DateTime.now().millisecondsSinceEpoch;
      if (currentTime - cacheTime < oneMonthInMilliseconds) {
        try {
          final jsonData = json.decode(cachedData) as Map<String, dynamic>;
          final entity = CountriesEntity.fromJson(jsonData);
          print('jsonData: $jsonData');
          return entity.countries.map((country) => country.toJson()).toList();
        } catch (e) {
          print('Failed to parse cached countries data: $e');
        }
      }
    }

    final params = <String, dynamic>{
      if (continentId != null) 'continentId': continentId,
      if (size != null) 'size': size,
      if (page != null) 'page': page,
    };

    HttpResultBean result = await Http.instance.get(
      GeoApi.getCountries,
      params: params,
      isShowLoading: false,
    );
    if (result.isSuccess()) {
      final data = result.data as Map<String, dynamic>;
      final entity = CountriesEntity.fromJson(data);
      final countryList =
          entity.countries.map((country) => country.toJson()).toList();
      print('countryList: $countryList');
      prefs.setString('all_countries_data', json.encode(entity.toJson()));
      prefs.setInt(
          'countries_cache_time', DateTime.now().millisecondsSinceEpoch);

      return countryList;
    }
    return [];
  }

  static Future<List<Map<String, dynamic>>> getContinentList() async {
    return [
      {
        'id': 0,
        'name': '热门',
        'nameEn': 'hot',
        'nameCn': '热门',
        'nameLocal': '热门',
        'code': 'hot'
      },
      {
        'id': 1,
        'name': '亚洲',
        'nameEn': 'Asia',
        'nameCn': '亚洲',
        'nameLocal': '亚洲',
        'code': 'AS'
      },
      {
        'id': 2,
        'name': '欧洲',
        'nameEn': 'Europe',
        'nameCn': '欧洲',
        'nameLocal': '欧洲',
        'code': 'EU'
      },
      {
        'id': 4,
        'name': '北美洲',
        'nameEn': 'North America',
        'nameCn': '北美洲',
        'nameLocal': '北美洲',
        'code': 'NA'
      },
      {
        'id': 5,
        'name': '南美洲',
        'nameEn': 'South America',
        'nameCn': '南美洲',
        'nameLocal': '南美洲',
        'code': 'SA'
      },
      {
        'id': 3,
        'name': '非洲',
        'nameEn': 'Africa',
        'nameCn': '非洲',
        'nameLocal': '非洲',
        'code': 'AF'
      },
      {
        'id': 6,
        'name': '大洋洲',
        'nameEn': 'Oceania',
        'nameCn': '大洋洲',
        'nameLocal': '大洋洲',
        'code': '0C'
      },
      {
        'id': 7,
        'name': '南极洲',
        'nameEn': 'Antarctica',
        'nameCn': '南极洲',
        'nameLocal': '南极洲',
        'code': 'AQ'
      },
    ];
  }

  static Future<List<Map<String, dynamic>>> getAllCountryList() async {
    return [
      {
        'id': 1,
        'imageUrl': 'img_american.png',
        'name': '美国',
        'nameEn': 'American',
        'nameCn': '美国',
        'code': 'American',
        'nameLocal': '美国',
        'continentId': 4,
        'continent': {
          'id': 4,
          'name': '北美洲',
          'nameEn': 'North America',
          'nameCn': '北美洲',
          'nameLocal': '北美洲',
          'code': 'NA'
        },
      },
      {
        'id': 2,
        'imageUrl': 'img_denmark.png',
        'name': '丹麦',
        'nameEn': 'Denmark',
        'nameCn': '丹麦',
        'code': 'Denmark',
        'nameLocal': '丹麦',
        'continentId': 2,
        'continent': {
          'id': 2,
          'name': '欧洲',
          'nameEn': 'Europe',
          'nameCn': '欧洲',
          'nameLocal': '欧洲',
          'code': 'EU'
        },
        'isHot': true
      },
      {
        'id': 3,
        'imageUrl': 'img_australia.png',
        'name': '澳大利亚',
        'nameEn': 'Australia',
        'nameCn': '澳大利亚',
        'code': 'Australia',
        'nameLocal': '澳大利亚',
        'continentId': 6,
        'continent': {
          'id': 6,
          'name': '大洋洲',
          'nameEn': 'Oceania',
          'nameCn': '大洋洲',
          'nameLocal': '大洋洲',
          'code': '0C'
        },
        'isHot': true
      },
      {
        'id': 4,
        'imageUrl': 'img_china.png',
        'name': '中国',
        'nameEn': 'China',
        'nameCn': '中国',
        'code': 'China',
        'nameLocal': '中国',
        'continentId': 1,
        'continent': {
          'id': 1,
          'name': '亚洲',
          'nameEn': 'Asia',
          'nameCn': '亚洲',
          'nameLocal': '亚洲',
          'code': 'AS'
        },
        'isHot': true
      },
      {
        'id': 5,
        'imageUrl': 'img_finland.png',
        'name': '芬兰',
        'nameEn': 'Finland',
        'nameCn': '芬兰',
        'code': 'Finland',
        'nameLocal': '芬兰',
        'continentId': 2,
        'continent': {
          'id': 2,
          'name': '欧洲',
          'nameEn': 'Europe',
          'nameCn': '欧洲',
          'nameLocal': '欧洲',
          'code': 'EU'
        },
        'isHot': true
      },
      {
        'id': 6,
        'imageUrl': 'img_uk.png',
        'name': '英国',
        'nameEn': 'UK',
        'nameCn': '英国',
        'code': 'UK',
        'nameLocal': '英国',
        'continentId': 2,
        'continent': {
          'id': 2,
          'name': '欧洲',
          'nameEn': 'Europe',
          'nameCn': '欧洲',
          'nameLocal': '欧洲',
          'code': 'EU'
        },
        'isHot': true
      },
      {
        'id': 7,
        'imageUrl': 'img_france.png',
        'name': '法国',
        'nameEn': 'France',
        'nameCn': '法国',
        'code': 'France',
        'nameLocal': '法国',
        'continentId': 2,
        'continent': {
          'id': 2,
          'name': '欧洲',
          'nameEn': 'Europe',
          'nameCn': '欧洲',
          'nameLocal': '欧洲',
          'code': 'EU'
        },
        'isHot': true
      },
      {
        'id': 8,
        'imageUrl': 'img_japan.png',
        'name': '日本',
        'nameEn': 'Japan',
        'nameCn': '日本',
        'code': 'Japan',
        'nameLocal': '日本',
        'continentId': 1,
        'continent': {
          'id': 1,
          'name': '亚洲',
          'nameEn': 'Asia',
          'nameCn': '亚洲',
          'nameLocal': '亚洲',
          'code': 'AS'
        },
        'isHot': true
      },
      {
        'id': 9,
        'imageUrl': 'img_italy.png',
        'name': '意大利',
        'nameEn': 'Italy',
        'nameCn': '意大利',
        'code': 'Italy',
        'nameLocal': '意大利',
        'continentId': 2,
        'continent': {
          'id': 2,
          'name': '欧洲',
          'nameEn': 'Europe',
          'nameCn': '欧洲',
          'nameLocal': '欧洲',
          'code': 'EU'
        },
        'isHot': true
      },
      {
        'id': 10,
        'imageUrl': 'img_thailand.png',
        'name': '泰国',
        'nameEn': 'Thailand',
        'nameCn': '泰国',
        'code': 'Thailand',
        'nameLocal': '泰国',
        'continentId': 1,
        'continent': {
          'id': 1,
          'name': '亚洲',
          'nameEn': 'Asia',
          'nameCn': '亚洲',
          'nameLocal': '亚洲',
          'code': 'AS'
        },
        'isHot': true
      },
      {
        'id': 11,
        'imageUrl': 'img_iceland.png',
        'name': '冰岛',
        'nameEn': 'Iceland',
        'nameCn': '冰岛',
        'code': 'Iceland',
        'nameLocal': '冰岛',
        'continentId': 2,
        'continent': {
          'id': 2,
          'name': '欧洲',
          'nameEn': 'Europe',
          'nameCn': '欧洲',
          'nameLocal': '欧洲',
          'code': 'EU'
        },
        'isHot': true
      },
      {
        'id': 12,
        'imageUrl': 'img_spain.png',
        'name': '西班牙',
        'nameEn': 'Spain',
        'nameCn': '西班牙',
        'code': 'Spain',
        'nameLocal': '西班牙',
        'continentId': 2,
        'continent': {
          'id': 2,
          'name': '欧洲',
          'nameEn': 'Europe',
          'nameCn': '欧洲',
          'nameLocal': '欧洲',
          'code': 'EU'
        },
        'isHot': true
      }
    ];
  }

  static Future<List<Map<String, dynamic>>> getAllCityList() async {
    return [
      {
        'id': 1,
        'imageUrl': 'img_beijing.png',
        'name': '北京',
        'countryId': 4,
      },
      {
        'id': 2,
        'imageUrl': 'img_shanghai.png',
        'name': '上海',
        'countryId': 4,
      },
      {
        'id': 3,
        'imageUrl': 'img_guangzhou.png',
        'name': '广州',
        'countryId': 4,
      },
      {
        'id': 4,
        'imageUrl': 'img_chengdu.png',
        'name': '成都',
        'countryId': 4,
      }
    ];
  }
}
