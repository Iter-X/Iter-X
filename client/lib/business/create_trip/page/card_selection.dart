/*
 * @Description: 
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-23 20:35:12
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-03-27 17:54:06
 */
/*
 * @Description: 图卡选择页
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-18 00:30:03
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-03-23 22:32:11
 */

import 'package:flutter/material.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/material/image.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter/services.dart';
import 'package:gap/gap.dart';

import '../../../common/material/app.dart';
import '../../../common/material/iter_text.dart';
import '../../common/widgets/buttom_widgets.dart';

class CardSelectionPage extends StatefulWidget {
  const CardSelectionPage({super.key});

  @override
  State<CardSelectionPage> createState() => _CardSelectionPageState();
}

class _CardSelectionPageState extends State<CardSelectionPage> {
  int selectionLevel = 0; // 0:国家 1:城市 2:景点
  var _selectedContinentId = '0'; // 选择的洲id
  final Set<String> _selectedCountries = {};
  List<Map<String, dynamic>> _continentList = [
    {'continentId': '0', 'name': '热门', 'englishName': 'hot'},
    {'continentId': '1', 'name': '亚洲', 'englishName': 'Asia'},
    {'continentId': '2', 'name': '欧洲', 'englishName': 'Europe'},
    {'continentId': '3', 'name': '北美洲', 'englishName': 'North America'},
    {'continentId': '4', 'name': '南美洲', 'englishName': 'South America'},
    {'continentId': '5', 'name': '非洲', 'englishName': 'Africa'},
    {'continentId': '6', 'name': '大洋洲', 'englishName': 'Oceania'},
  ];
  final List<Map<String, dynamic>> _allCountryList = [
    {
      'countryId': '1',
      'image': 'img_american.png',
      'name': '美国',
      'englishName': 'American',
      'continentId': '3',
      'isHot': true
    },
    {
      'countryId': '2',
      'image': 'img_denmark.png',
      'name': '丹麦',
      'englishName': 'Denmark',
      'continentId': '2',
      'isHot': true
    },
    {
      'countryId': '3',
      'image': 'img_australia.png',
      'name': '澳大利亚',
      'englishName': 'Australia',
      'continentId': '4',
      'isHot': true
    },
    {
      'countryId': '4',
      'image': 'img_china.png',
      'name': '中国',
      'englishName': 'China',
      'continentId': '1',
      'isHot': true
    },
    {
      'countryId': '5',
      'image': 'img_finland.png',
      'name': '芬兰',
      'englishName': 'Finland',
      'continentId': '2',
      'isHot': true
    },
    {
      'countryId': '6',
      'image': 'img_uk.png',
      'name': '英国',
      'englishName': 'UK',
      'continentId': '2',
      'isHot': true
    },
    {
      'countryId': '7',
      'image': 'img_france.png',
      'name': '法国',
      'englishName': 'France',
      'continentId': '2',
      'isHot': true
    },
    {
      'countryId': '8',
      'image': 'img_japan.png',
      'name': '日本',
      'englishName': 'Japan',
      'continentId': '1',
      'isHot': true
    },
    {
      'countryId': '9',
      'image': 'img_italy.png',
      'name': '意大利',
      'englishName': 'Italy',
      'continentId': '2',
      'isHot': true
    },
    {
      'countryId': '10',
      'image': 'img_thailand.png',
      'name': '泰国',
      'englishName': 'Thailand',
      'continentId': '1',
      'isHot': true
    },
    {
      'countryId': '11',
      'image': 'img_iceland.png',
      'name': '冰岛',
      'englishName': 'Iceland',
      'continentId': '2',
      'isHot': true
    },
    {
      'countryId': '12',
      'image': 'img_spain.png',
      'name': '西班牙',
      'englishName': 'Spain',
      'continentId': '2',
      'isHot': true
    }
  ];
  final List<Map<String, dynamic>> _allCityList = [
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
  late List<Map<String, dynamic>> _continentCountList;

  @override
  void initState() {
    super.initState();
    _continentCountList =
        _allCountryList.where((country) => country['isHot'] == true).toList();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: BaseColor.c_f2f2f2,
        appBar: AppBar(
            title: selectionLevel == 0 ? Text('选择目的国家') : Text('选择目的城市'),
            systemOverlayStyle: SystemUiOverlayStyle(
              statusBarColor: Colors.transparent,
              statusBarIconBrightness: Brightness.dark,
              statusBarBrightness: Brightness.light,
            ),
            backgroundColor: Colors.transparent,
            leading: ButtonBackWidget(onTap: () {
              if (selectionLevel == 1) {
                setState(() {
                  selectionLevel = 0;
                  _continentCountList = _selectedContinentId == '0'
                      ? _allCountryList
                          .where((country) => country['isHot'] == true)
                          .toList()
                      : _allCountryList
                          .where((country) =>
                              country['continentId'] == _selectedContinentId)
                          .toList();
                });
              } else {
                Navigator.pop(context);
              }
            })),
        body: SafeAreaX(
          child: Column(children: [
            Expanded(
                child: SingleChildScrollView(
                    child: GridView.builder(
              shrinkWrap: true,
              physics: NeverScrollableScrollPhysics(),
              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 3,
                crossAxisSpacing: 4,
                mainAxisSpacing: 4,
              ),
              itemCount: _continentCountList.length,
              itemBuilder: (context, index) {
                final country = _continentCountList[index];
                final isSelected =
                    _selectedCountries.contains(country['cityId']);
                return GestureDetector(
                  onTap: () {
                    setState(() {
                      if (selectionLevel == 0) {
                        _continentCountList = _allCityList
                            .where((city) =>
                                city['countryId'] == country['countryId'])
                            .toList();
                        selectionLevel = 1;
                      } else {
                        if (isSelected) {
                          _selectedCountries.remove(country['cityId']);
                        } else {
                          _selectedCountries.add(country['cityId']);
                        }
                      }
                    });
                  },
                  child: Container(
                    child: Stack(
                      children: [
                        BaseImage.asset(
                          name: country['image'],
                          width: 142.w,
                          fit: BoxFit.cover,
                        ),
                        Positioned(
                          bottom: 0,
                          right: 8.w,
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.end,
                            children: [
                              IterText(country['name'],
                                  style: TextStyle(
                                    fontSize: 22.sp,
                                    color: isSelected
                                        ? BaseColor.c_1D1F1E
                                        : BaseColor.c_f2f2f2,
                                    fontWeight: FontWeight.w900,
                                  ),
                                  borders: BorderProperties(
                                    width: 2,
                                    color: isSelected
                                        ? BaseColor.c_f2f2f2
                                        : BaseColor.c_1D1F1E,
                                  )),
                              if (country['englishName'] != null &&
                                  country['englishName'].isNotEmpty)
                                IterText(country['englishName'],
                                    textAlign: TextAlign.end, // 文字居中
                                    style: TextStyle(
                                      fontSize: 22.sp,
                                      color: isSelected
                                          ? BaseColor.c_1D1F1E
                                          : BaseColor.c_f2f2f2,
                                      fontWeight: FontWeight.w900,
                                    ),
                                    borders: BorderProperties(
                                      width: 2,
                                      color: isSelected
                                          ? BaseColor.c_f2f2f2
                                          : BaseColor.c_1D1F1E,
                                    )),
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
                );
              },
            ))),
            // 大洲Tab
            Container(
              padding: EdgeInsets.all(20),
              color: BaseColor.c_F2F2F2,
              child: Column(
                children: [
                  Container(
                    width: double.infinity,
                    child: Wrap(
                      alignment: WrapAlignment.start,
                      spacing: 10, // 水平间距
                      runSpacing: 10, // 垂直间距
                      children: [
                        ..._continentList
                            .map((tag) => IntrinsicWidth(
                                    child: GestureDetector(
                                  onTap: () => {
                                    setState(() {
                                      _selectedContinentId = tag['continentId'];
                                      _continentCountList =
                                          _selectedContinentId == '0'
                                              ? _allCountryList
                                                  .where((country) =>
                                                      country['isHot'] == true)
                                                  .toList()
                                              : _allCountryList
                                                  .where((country) =>
                                                      country['continentId'] ==
                                                      _selectedContinentId)
                                                  .toList();
                                      selectionLevel = 0;
                                    })
                                  },
                                  child: Container(
                                      alignment: Alignment.center,
                                      padding: EdgeInsets.symmetric(
                                          horizontal: 15), // 内间距
                                      height: 38.h,
                                      decoration: BoxDecoration(
                                        borderRadius:
                                            BorderRadius.circular(28.w),
                                        color: selectionLevel == 0 &&
                                                _selectedContinentId ==
                                                    tag['continentId']
                                            ? BaseColor.c_1D1F1E
                                            : BaseColor.c_E3E3E3,
                                      ),
                                      child: Row(
                                          mainAxisAlignment:
                                              MainAxisAlignment.center,
                                          children: [
                                            if (selectionLevel == 0 &&
                                                _selectedContinentId ==
                                                    tag['continentId'])
                                              BaseImage.asset(
                                                name: 'ic_create_picard.png',
                                                size: 18.w,
                                              ),
                                            Gap(5.w),
                                            Text(
                                              tag['name'],
                                              style: TextStyle(
                                                fontSize: 14.sp,
                                                color: selectionLevel == 0 &&
                                                        _selectedContinentId ==
                                                            tag['continentId']
                                                    ? BaseColor.c_F2F2F2
                                                    : BaseColor.c_1D1F1E,
                                              ),
                                            ),
                                          ])),
                                )))
                            .toList()
                      ],
                    ),
                  ),
                  SizedBox(height: 20),
                  // 已选城市滚动区
                  Container(
                    width: double.infinity,
                    child: SingleChildScrollView(
                      scrollDirection: Axis.horizontal,
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.start,
                        children: _selectedCountries
                            .map((id) => _allCityList
                                .firstWhere((item) => item['cityId'] == id))
                            .map((country) => Row(children: [
                                  IntrinsicWidth(
                                      child: Container(
                                    height: 42.w,
                                    alignment: Alignment.center,
                                    padding:
                                        EdgeInsets.symmetric(horizontal: 15),
                                    decoration: BoxDecoration(
                                      color: BaseColor.c_E3E3E3,
                                      borderRadius: BorderRadius.only(
                                        topLeft: Radius.circular(5.sp),
                                        bottomLeft: Radius.circular(5.sp),
                                      ),
                                    ),
                                    child: Text(
                                      country['name'],
                                      style: TextStyle(
                                        fontSize: 16.sp,
                                        color: BaseColor.c_1D1F1E,
                                      ),
                                    ),
                                  )),
                                  GestureDetector(
                                      onTap: () => {
                                            setState(() {
                                              _selectedCountries
                                                  .remove(country['cityId']);
                                            })
                                          },
                                      child: Container(
                                        alignment: Alignment.center,
                                        decoration: BoxDecoration(
                                          color: BaseColor.c_1D1F1E,
                                          borderRadius: BorderRadius.only(
                                            topRight: Radius.circular(5.sp),
                                            bottomRight: Radius.circular(5.sp),
                                          ),
                                        ),
                                        margin: EdgeInsets.only(right: 10.w),
                                        width: 38.w,
                                        height: 42.w,
                                        child: BaseImage.asset(
                                          name: 'ic_card_cancel.png',
                                          size: 26.w,
                                        ),
                                      ))
                                ]))
                            .toList(),
                      ),
                    ),
                  ),
                  SizedBox(height: 20),
                  GestureDetector(
                    onTap: () => {print('点击')},
                    child: Container(
                      width: 390.w,
                      height: 52.h,
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(28.w),
                        color: BaseColor.c_1D1F1E,
                      ),
                      margin: EdgeInsets.only(right: 6.w),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          BaseImage.asset(
                            name: 'ic_card_location.png',
                            size: 24.w,
                          ),
                          Gap(5.w),
                          Text(
                            "已选${_selectedCountries.length}",
                            style: TextStyle(
                              fontSize: 18.sp,
                              color: BaseColor.c_f2f2f2,
                            ),
                          ),
                          Gap(10.w),
                          BaseImage.asset(
                            name: 'ic_card_arrow_right.png',
                            size: 20.w,
                          ),
                          Gap(10.w),
                          BaseImage.asset(
                            name: 'ic_card_generate.png',
                            size: 24.w,
                          ),
                          Gap(5.w),
                          Text(
                            "生成行程",
                            style: TextStyle(
                              fontSize: 18.sp,
                              color: BaseColor.c_f2f2f2,
                            ),
                          )
                        ],
                      ),
                    ),
                  )
                ],
              ),
            )
          ]),
        ));
  }
}
