/*
 * @Description: 图卡选择页
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-18 00:30:03
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-04-01 18:49:46
 */

import 'package:client/app/constants.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/iter_text.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter/services.dart';
import 'package:client/common/widgets/clickable_button.dart';
import 'package:gap/gap.dart';
import '../service/card_selection_service.dart';
import '../entity/contient_entity.dart';

class CardSelectionPage extends StatefulWidget {
  const CardSelectionPage({super.key});

  @override
  State<CardSelectionPage> createState() => _CardSelectionPageState();
}

class _CardSelectionPageState extends State<CardSelectionPage> {
  int selectionLevel = 0; // 0:国家 1:城市 2:景点
  // 修改为 String 类型
  int _selectedContinentId = 0;
  final Set<String> _selectedCountries = {};
  late List<Continent> _continentList;
  late int _total;
  late List<Map<String, dynamic>> _allCountryList;
  late List<Map<String, dynamic>> _allCityList;
  late List<Map<String, dynamic>> _continentCountList;
  bool _isLoading = true; // 添加加载状态标志

  @override
  void initState() {
    super.initState();
    _fetchData();
  }

  Future<void> _fetchData() async {
    final continentEntity = await CardSelectionService.getContinentsData();
    if (continentEntity != null) {
      _continentList = continentEntity.continents;
      _total = continentEntity.total;
    }
    _allCountryList = await CardSelectionService.getAllCountryList();
    _allCityList = await CardSelectionService.getAllCityList();
    _continentCountList =
        _allCountryList.where((country) => country['isHot'] == true).toList();
    setState(() {
      _isLoading = false; // 数据加载完成，更新加载状态
    });
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return AppBarWithSafeArea(
        backgroundColor: AppColor.bg,
        hasAppBar: true,
        title: '加载中...',
        leading: ReturnButton(onTap: () {
          Navigator.pop(context);
        }),
        child: const Center(
          child: CircularProgressIndicator(),
        ),
      );
    }

    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      hasAppBar: true,
      title: selectionLevel == 0 ? '选择目的国家' : '选择目的城市',
      leading: ReturnButton(onTap: () {
        if (selectionLevel == 1) {
          setState(() {
            selectionLevel = 0;
            _continentCountList = _selectedContinentId == 0
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
      }),
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
                                        ? AppColor.c_1D1F1E
                                        : AppColor.c_f2f2f2,
                                    fontWeight: AppFontWeight.black,
                                  ),
                                  borders: BorderProperties(
                                    width: 2,
                                    color: isSelected
                                        ? AppColor.c_f2f2f2
                                        : AppColor.c_1D1F1E,
                                  )),
                              if (country['englishName'] != null &&
                                  country['englishName'].isNotEmpty)
                                IterText(country['englishName'],
                                    textAlign: TextAlign.end, // 文字居中
                                    style: TextStyle(
                                      fontSize: 22.sp,
                                      color: isSelected
                                          ? AppColor.c_1D1F1E
                                          : AppColor.c_f2f2f2,
                                      fontWeight: AppFontWeight.black,
                                    ),
                                    borders: BorderProperties(
                                      width: 2,
                                      color: isSelected
                                          ? AppColor.c_f2f2f2
                                          : AppColor.c_1D1F1E,
                                    )),
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
                );
              },
            ),
          ),
        ),
        Container(
          padding: EdgeInsets.all(20),
          decoration: BoxDecoration(
            color: AppColor.bottomBar,
            border: Border(
              top: BorderSide(
                color: AppColor.bottomBarLine,
                width: 1,
              ),
            ),
          ),
          child: Column(
            children: [
              // 大洲Tab
              Container(
                width: double.infinity,
                child: Wrap(
                  alignment: WrapAlignment.start,
                  spacing: 10, // 水平间距
                  runSpacing: 10, // 垂直间距
                  children: [
                    ..._continentList
                        .map(
                          (tag) => IntrinsicWidth(
                            child: GestureDetector(
                              onTap: () => {
                                setState(() {
                                  print(_selectedContinentId);
                                  print(_allCountryList[0]['continentId']);
                                  _selectedContinentId = tag.id;
                                  _continentCountList =
                                      _selectedContinentId == 0
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
                                padding: EdgeInsets.symmetric(horizontal: 15),
                                // 内间距
                                height: 38.h,
                                decoration: BoxDecoration(
                                  borderRadius: BorderRadius.circular(28.w),
                                  color: selectionLevel == 0 &&
                                          _selectedContinentId == tag.id
                                      ? AppColor.c_1D1F1E
                                      : AppColor.c_E3E3E3,
                                ),
                                child: Row(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                      if (selectionLevel == 0 &&
                                          _selectedContinentId == tag.id)
                                        BaseImage.asset(
                                          name: 'ic_create_picard.png',
                                          size: 18.w,
                                        ),
                                      Gap(5.w),
                                      Text(
                                        tag.nameCn,
                                        style: TextStyle(
                                          fontSize: 14.sp,
                                          color: selectionLevel == 0 &&
                                                  _selectedContinentId == tag.id
                                              ? AppColor.c_F2F2F2
                                              : AppColor.c_1D1F1E,
                                        ),
                                      ),
                                    ]),
                              ),
                            ),
                          ),
                        )
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
                        .map((country) => ClickableButton(
                              text: country['name'],
                              onTapIcon: () => {
                                setState(() {
                                  _selectedCountries.remove(country['cityId']);
                                })
                              },
                              icon: Icons.cancel,
                              iconColor: AppColor.closeButton,
                              margin: EdgeInsets.only(right: 10),
                            ))
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
                    color: AppColor.c_1D1F1E,
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
                          color: AppColor.c_f2f2f2,
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
                          color: AppColor.c_f2f2f2,
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
    );
  }
}
