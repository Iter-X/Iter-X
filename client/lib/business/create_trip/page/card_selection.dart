/*
 * @Description: 图卡选择页
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-18 00:30:03
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-03-20 01:36:28
 */

import 'package:flutter/material.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/material/image.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

import '../../common/widgets/buttom_widgets.dart';

class CardSelectionPage extends StatefulWidget {
  const CardSelectionPage({super.key});

  @override
  State<CardSelectionPage> createState() => _CardSelectionPageState();
}

class _CardSelectionPageState extends State<CardSelectionPage> {
  int selectionLevel = 0; // 0:国家 1:城市 2:景点
  final Set<String> _selectedCountries = {};
  final List<Map<String, dynamic>> _countryList = [
    {
      'id': '1',
      'image': 'img_american.png',
      'name': '美国',
      'englishName': 'American'
    },
    {
      'id': '2',
      'image': 'img_denmark.png',
      'name': '丹麦',
      'englishName': 'Denmark'
    },
    {
      'id': '3',
      'image': 'img_australia.png',
      'name': '澳大利亚',
      'englishName': 'Australia'
    },
  ];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: BaseColor.c_f2f2f2,
      appBar: AppBar(
          title: selectionLevel == 0 ? Text('选择目的国家') : Text('选择目的城市'),
          backgroundColor: Colors.transparent,
          leading: ButtonBackWidget()),
      body: Column(children: [
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
          itemCount: _countryList.length,
          itemBuilder: (context, index) {
            final country = _countryList[index];
            final isSelected = _selectedCountries.contains(country['id']);
            return GestureDetector(
              onTap: () {
                setState(() {
                  if (isSelected) {
                    _selectedCountries.remove(country['id']);
                  } else {
                    _selectedCountries.add(country['id']);
                  }
                });
              },
              child: Container(
                decoration: BoxDecoration(
                  border: Border.all(
                    color: isSelected ? Colors.black : Colors.transparent,
                    width: 2,
                  ),
                ),
                child: Column(
                  children: [
                    BaseImage.asset(
                      name: country['image'],
                      size: 142.w,
                      fit: BoxFit.cover,
                    ),
                    Text(country['name']),
                  ],
                ),
              ),
            );
          },
        ))),
        Container(
          padding: EdgeInsets.all(16),
          color: BaseColor.c_F2F2F2,
          child: Column(
            children: [
              Wrap(
                spacing: 10, // 水平间距
                runSpacing: 10, // 垂直间距
                children: [
                  GestureDetector(
                    onTap: () => {print('目的地')},
                    child: Container(
                      width: 81.w,
                      height: 38.h,
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(28.w),
                        color: BaseColor.c_1D1F1E,
                      ),
                      child: Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            BaseImage.asset(
                              name: 'ic_create_picard.png',
                              size: 18.w,
                            ),
                            Gap(5.w),
                            Text(
                              "目的地",
                              style: TextStyle(
                                fontSize: 14.sp,
                                color: BaseColor.c_f2f2f2,
                              ),
                            ),
                          ]),
                    ),
                  ),
                  ...['亚洲', '欧洲', '北美', '南美', '非洲', '大洋洲', '南极洲']
                      .map((tag) => IntrinsicWidth(
                              child: GestureDetector(
                            onTap: () => {print('$tag')},
                            child: Container(
                                alignment: Alignment.center,
                                padding:
                                    EdgeInsets.symmetric(horizontal: 15), // 内间距
                                height: 38.h,
                                decoration: BoxDecoration(
                                  borderRadius: BorderRadius.circular(28.w),
                                  color: BaseColor.c_E3E3E3,
                                ),
                                child: Text(
                                  "$tag",
                                  style: TextStyle(
                                    fontSize: 14.sp,
                                    color: BaseColor.c_1D1F1E,
                                  ),
                                )),
                          )))
                      .toList()
                ],
              ),
              SizedBox(height: 20),
              // 已选城市滚动区
              SingleChildScrollView(
                scrollDirection: Axis.horizontal,
                child: Row(
                  children: _selectedCountries
                      .map((id) =>
                          _countryList.firstWhere((item) => item['id'] == id))
                      .map((country) => ChoiceChip(
                            label: Text(country['name']),
                            onSelected: (value) {
                              setState(() {
                                _selectedCountries.remove(country['id']);
                              });
                            },
                            selected:
                                _selectedCountries.contains(country['id']),
                          ))
                      .toList(),
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
    );
  }
}
