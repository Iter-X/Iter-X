import 'dart:convert';

import 'package:client/app/apis/geo_api.dart';
import 'package:client/business/create_trip/entity/geo_entity.dart';
import 'package:client/common/dio/http.dart';
import 'package:client/common/dio/http_result_bean.dart';
import 'package:client/common/utils/logger.dart';
import 'package:shared_preferences/shared_preferences.dart';

// The data cache time for continents and countries is one month, in milliseconds. Here it is set to 30 days.
// const oneMonthInMilliseconds = 30 * 24 * 60 * 60 * 1000;
const oneMonthInMilliseconds = 0;

class CardSelectionService {
  static Future<ContinentEntity?> getContinentsData() async {
    final prefs = await SharedPreferences.getInstance();
    final cachedData = prefs.getString('continents_data');
    final cacheTime = prefs.getInt('continents_cache_time');

    if (cachedData != null && cacheTime != null) {
      final currentTime = DateTime.now().millisecondsSinceEpoch;
      if (currentTime - cacheTime < oneMonthInMilliseconds) {
        try {
          final jsonData = json.decode(cachedData) as Map<String, dynamic>;
          return ContinentEntity.fromJson(jsonData);
        } catch (e) {
          BaseLogger.e('Failed to parse cached continents data: $e');
        }
      }
    }

    HttpResultBean result = await Http.instance.get(
      GeoApi.getContinents,
      isShowLoading: false,
    );
    if (result.isSuccess()) {
      final data = result.data as Map<String, dynamic>;
      final entity = ContinentEntity.fromJson(data);
      final newContinent = {
        'id': 0,
        'name': '热门',
        'nameEn': 'hot',
        'nameCn': '热门',
        'nameLocal': '热门',
        'code': 'hot',
        'isHot': true
      };
      final newContinents = [
        Continent.fromJson(newContinent),
        ...entity.continents
      ];
      final total = int.tryParse(data['total'].toString()) ?? 0;
      final newEntity =
          ContinentEntity(continents: newContinents, total: total + 1);
      prefs.setString('continents_data', json.encode(newEntity.toJson()));
      prefs.setInt(
          'continents_cache_time', DateTime.now().millisecondsSinceEpoch);

      return newEntity;
    }
    return null;
  }

  static Future<CountriesEntity?> getCountriesData({
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
          if (continentId != null && continentId != 0) {
            final filteredCountries = entity.countries
                .where((country) => country.continentId == continentId)
                .toList();
            return CountriesEntity(
                countries: filteredCountries, total: filteredCountries.length);
          }
          if (continentId == 0) {
            final hotCountries =
                entity.countries.where((country) => country.isHot).toList();
            return CountriesEntity(
                countries: hotCountries, total: hotCountries.length);
          }
          return entity;
        } catch (e) {
          BaseLogger.e('Failed to parse cached countries data: $e');
        }
      }
    }

    final params = <String, dynamic>{
      if (continentId != null && continentId != 0) 'continentId': continentId,
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
      prefs.setString('all_countries_data', json.encode(entity.toJson()));
      prefs.setInt(
          'countries_cache_time', DateTime.now().millisecondsSinceEpoch);

      return entity;
    }
    return null;
  }

  static Future<StatesEntity?> getStatesData({
    required int countryId,
    int? size = 100,
    int? page = 1,
  }) async {
    final params = <String, dynamic>{
      'countryId': countryId,
      if (size != null) 'size': size,
      if (page != null) 'page': page,
    };

    HttpResultBean result = await Http.instance.get(
      GeoApi.getStates,
      params: params,
      isShowLoading: false,
    );
    if (result.isSuccess()) {
      final data = result.data as Map<String, dynamic>;
      return StatesEntity.fromJson(data);
    }
    return null;
  }

  static Future<CitiesEntity?> getCitiesData({
    required int stateId,
    int? size = 100,
    int? page = 1,
  }) async {
    final params = <String, dynamic>{
      'stateId': stateId,
      if (size != null) 'size': size,
      if (page != null) 'page': page,
    };

    HttpResultBean result = await Http.instance.get(
      GeoApi.getCities,
      params: params,
      isShowLoading: false,
    );
    if (result.isSuccess()) {
      final data = result.data as Map<String, dynamic>;
      return CitiesEntity.fromJson(data);
    }
    return null;
  }
}
