/*
 * @Description: Card selection
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-18 00:30:03
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-04-03 15:41:39
 */

import 'package:client/app/constants.dart';
import 'package:client/app/routes.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/iter_text.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter/services.dart';
import 'package:client/common/widgets/clickable_button.dart';
import 'package:gap/gap.dart';
import '../service/card_selection_service.dart';
import '../entity/geo_entity.dart';

class CardSelectionPage extends StatefulWidget {
  const CardSelectionPage({super.key});

  @override
  State<CardSelectionPage> createState() => _CardSelectionPageState();
}

class _CardSelectionPageState extends BaseState<CardSelectionPage> {
  int selectionLevel = 0; // 0:countries 1:cities
  int _selectedContinentId = 0;
  final Set<int> _selectedCities = {};
  late List<Continent> _continentList;
  late int _total;
  late List<Map<String, dynamic>> _countriesList;
  late List<Map<String, dynamic>> _citiesList;
  late List<Map<String, dynamic>> _continentCountList;
  bool _isLoading = true;

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
    _countriesList = await CardSelectionService.getCountriesData();
    _citiesList = await CardSelectionService.getAllCityList();
    _continentCountList = _selectedContinentId == 0
        ? _countriesList.take(30).toList()
        : _countriesList
            .where((country) => country['continentId'] == _selectedContinentId)
            .toList();
    setState(() {
      _isLoading = false;
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
                ? _countriesList.take(30).toList()
                : _countriesList
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
                final selectItem = _continentCountList[index];
                final isSelected = _selectedCities.contains(selectItem['id']);
                return GestureDetector(
                  onTap: () {
                    setState(() {
                      print(selectItem);
                      print(isSelected);
                      if (selectionLevel == 0) {
                        _continentCountList = _citiesList
                            .where(
                                (city) => city['countryId'] == selectItem['id'])
                            .toList();
                        selectionLevel = 1;
                      } else {
                        if (isSelected) {
                          _selectedCities.remove(selectItem['id']);
                        } else {
                          _selectedCities.add(selectItem['id']);
                        }
                      }
                    });
                  },
                  child: Container(
                    child: Stack(
                      children: [
                        BaseImage.asset(
                          name: selectItem['imageUrl'],
                          width: 142.w,
                          fit: BoxFit.cover,
                        ),
                        Positioned(
                          bottom: 0,
                          right: 8.w,
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.end,
                            children: [
                              IterText(selectItem['name'],
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
                              if (selectItem['nameEn'] != null &&
                                  selectItem['nameEn'].isNotEmpty)
                                IterText(selectItem['nameEn'],
                                    textAlign: TextAlign.end,
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
              // continentTab
              Container(
                width: double.infinity,
                child: Wrap(
                  alignment: WrapAlignment.start,
                  spacing: 10,
                  runSpacing: 10,
                  children: [
                    ..._continentList
                        .map(
                          (tag) => IntrinsicWidth(
                            child: GestureDetector(
                              onTap: () => {
                                setState(() {
                                  print(_selectedContinentId);
                                  print(_countriesList[0]['continentId']);
                                  _selectedContinentId = tag.id;
                                  _continentCountList =
                                      _selectedContinentId == 0
                                          ? _countriesList.take(30).toList()
                                          : _countriesList
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
              // cityScrollView
              Container(
                width: double.infinity,
                child: SingleChildScrollView(
                  scrollDirection: Axis.horizontal,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.start,
                    children: _selectedCities
                        .map((id) =>
                            _citiesList.firstWhere((item) => item['id'] == id))
                        .map((city) => ClickableButton(
                              text: city['name'],
                              onTapIcon: () => {
                                setState(() {
                                  _selectedCities.remove(city['id']);
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
                onTap: () => {go(Routes.poiSearch)},
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
                        "已选${_selectedCities.length}",
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
