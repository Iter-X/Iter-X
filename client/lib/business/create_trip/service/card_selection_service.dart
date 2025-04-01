import 'package:client/app/apis/geo_api.dart';
import 'package:client/common/dio/http.dart';
import 'package:client/common/dio/http_result_bean.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../entity/contient_entity.dart';
import 'dart:convert';

class CardSelectionService {
  // 获取大洲数据
  static Future<ContientEntity?> getContinentsData() async {
    final prefs = await SharedPreferences.getInstance();
    final cachedData = prefs.getString('continents_data');
    if (cachedData != null) {
      try {
        // 使用 json.decode 解析 JSON 字符串
        final jsonData = json.decode(cachedData) as Map<String, dynamic>;
        return ContientEntity.fromJson(jsonData);
      } catch (e) {
        print('Failed to parse cached continents data: $e');
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
        'code': 'hot'
      };
      final newContinents = [
        Continent.fromJson(newContinent),
        ...entity.continents
      ];
      // 转换 total 为数字类型
      final total = int.tryParse(data['total'].toString()) ?? 0;
      final newEntity =
          ContientEntity(continents: newContinents, total: total + 1);
      prefs.setString('continents_data', json.encode(newEntity.toJson()));
      return newEntity;
    }
    return null;
  }

  // 获取国家数据
  static Future<HttpResultBean> getCountriesData({
    String? continentId,
    int size = 30,
    int page = 1,
  }) async {
    final params = <String, dynamic>{
      if (continentId != null) 'continentId': continentId,
      'size': size,
      'page': page,
    };

    return await Http.instance.get(
      GeoApi.getCountries,
      // 传递查询参数
      params: params,
      isShowLoading: false,
    );
  }

  // 虚拟请求：获取大洲列表
  static Future<List<Map<String, dynamic>>> getContinentList() async {
    return [
      {
        'continentId': '',
        'name': '热门',
        'nameEn': 'hot',
        'nameCn': '热门',
        'code': 'hot'
      },
      {'continentId': '1', 'name': '亚洲', 'englishName': 'Asia'},
      {'continentId': '2', 'name': '欧洲', 'englishName': 'Europe'},
      {'continentId': '3', 'name': '北美洲', 'englishName': 'North America'},
      {'continentId': '4', 'name': '南美洲', 'englishName': 'South America'},
      {'continentId': '5', 'name': '非洲', 'englishName': 'Africa'},
      {'continentId': '6', 'name': '大洋洲', 'englishName': 'Oceania'},
    ];
  }

  // 虚拟请求：获取所有国家列表
  static Future<List<Map<String, dynamic>>> getAllCountryList() async {
    return [
      {
        'countryId': '1',
        'image': 'img_american.png',
        'name': '美国',
        'englishName': 'American',
        'continentId': 4,
        'isHot': true
      },
      {
        'countryId': '2',
        'image': 'img_denmark.png',
        'name': '丹麦',
        'englishName': 'Denmark',
        'continentId': 2,
        'isHot': true
      },
      {
        'countryId': '3',
        'image': 'img_australia.png',
        'name': '澳大利亚',
        'englishName': 'Australia',
        'continentId': 6,
        'isHot': true
      },
      {
        'countryId': '4',
        'image': 'img_china.png',
        'name': '中国',
        'englishName': 'China',
        'continentId': 1,
        'isHot': true
      },
      {
        'countryId': '5',
        'image': 'img_finland.png',
        'name': '芬兰',
        'englishName': 'Finland',
        'continentId': 2,
        'isHot': true
      },
      {
        'countryId': '6',
        'image': 'img_uk.png',
        'name': '英国',
        'englishName': 'UK',
        'continentId': 2,
        'isHot': true
      },
      {
        'countryId': '7',
        'image': 'img_france.png',
        'name': '法国',
        'englishName': 'France',
        'continentId': 2,
        'isHot': true
      },
      {
        'countryId': '8',
        'image': 'img_japan.png',
        'name': '日本',
        'englishName': 'Japan',
        'continentId': 1,
        'isHot': true
      },
      {
        'countryId': '9',
        'image': 'img_italy.png',
        'name': '意大利',
        'englishName': 'Italy',
        'continentId': 2,
        'isHot': true
      },
      {
        'countryId': '10',
        'image': 'img_thailand.png',
        'name': '泰国',
        'englishName': 'Thailand',
        'continentId': 1,
        'isHot': true
      },
      {
        'countryId': '11',
        'image': 'img_iceland.png',
        'name': '冰岛',
        'englishName': 'Iceland',
        'continentId': 2,
        'isHot': true
      },
      {
        'countryId': '12',
        'image': 'img_spain.png',
        'name': '西班牙',
        'englishName': 'Spain',
        'continentId': 2,
        'isHot': true
      }
    ];
  }

  // 虚拟请求：获取所有城市列表
  static Future<List<Map<String, dynamic>>> getAllCityList() async {
    return [
      {
        'cityId': '1',
        'image': 'img_beijing.png',
        'name': '北京',
        'countryId': '4',
      },
      {
        'cityId': '2',
        'image': 'img_shanghai.png',
        'name': '上海',
        'countryId': '4',
      },
      {
        'cityId': '3',
        'image': 'img_guangzhou.png',
        'name': '广州',
        'countryId': '4',
      },
      {
        'cityId': '4',
        'image': 'img_chengdu.png',
        'name': '成都',
        'countryId': '4',
      }
    ];
  }
}
