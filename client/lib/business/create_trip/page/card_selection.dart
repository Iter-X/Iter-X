/*
 * @Description: 目的地选择页面
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-18 00:30:03
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-04-03 15:41:39
 */

import 'package:client/app/constants.dart';
import 'package:client/app/routes.dart';
import 'package:client/business/create_trip/entity/geo_entity.dart';
import 'package:client/business/create_trip/service/card_selection_service.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/iter_text.dart';
import 'package:client/common/widgets/clickable_button.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

/// 选择层级枚举
enum SelectionLevel { country, city }

class CardSelectionPage extends StatefulWidget {
  const CardSelectionPage({super.key});

  @override
  State<CardSelectionPage> createState() => _CardSelectionPageState();
}

class _CardSelectionPageState extends State<CardSelectionPage> {
  SelectionLevel _selectionLevel = SelectionLevel.country;
  int _selectedContinentId = 0;
  int? _selectedCountryId;
  final Set<int> _selectedCities = {};
  bool _isLoading = true;

  ContinentEntity? _continentList;
  CountriesEntity? _countriesList;
  CitiesEntity? _citiesList;
  late List<dynamic> _showItems = [];

  @override
  void initState() {
    super.initState();
    _loadData();
  }

  Future<void> _loadData() async {
    final continents = await CardSelectionService.getContinentsData();
    final countries = await CardSelectionService.getCountriesData();

    if (mounted) {
      setState(() {
        _continentList = continents;
        _countriesList = countries;
        _showItems = countries?.countries.take(30).toList() ?? [];
        _isLoading = false;
      });
    }
  }

  Future<void> _loadCitiesByCountry(int countryId) async {
    final cities =
        await CardSelectionService.getCitiesData(countryId: countryId);
    if (mounted) {
      setState(() {
        _citiesList = cities;
        _showItems = cities?.cities ?? [];
        _selectionLevel = SelectionLevel.city;
        _selectedCountryId = countryId;
      });
    }
  }

  // 处理返回按钮点击
  void _handleReturn() {
    if (_selectionLevel == SelectionLevel.city) {
      setState(() {
        _selectionLevel = SelectionLevel.country;
        _showItems = _selectedContinentId == 0
            ? _countriesList!.countries.take(30).toList()
            : _countriesList!.countries
                .where((country) => country.continentId == _selectedContinentId)
                .toList();
      });
    } else {
      Navigator.pop(context);
    }
  }

  void _handleItemTap(dynamic item) {
    if (_selectionLevel == SelectionLevel.country) {
      _loadCitiesByCountry((item as Country).id);
    } else {
      setState(() {
        final id = (item as City).id;
        if (_selectedCities.contains(id)) {
          _selectedCities.remove(id);
        } else {
          _selectedCities.add(id);
        }
      });
    }
  }

  // 处理大洲标签点击
  void _handleContinentTap(Continent tag) {
    setState(() {
      _selectedContinentId = tag.id;
      _showItems = _selectedContinentId == 0
          ? _countriesList!.countries.take(30).toList()
          : _countriesList!.countries
              .where((country) => country.continentId == _selectedContinentId)
              .toList();
      _selectionLevel = SelectionLevel.country;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return _buildLoadingView();
    }

    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      bottomColor: AppColor.bottomBar,
      hasAppBar: true,
      title: _getPageTitle(),
      leading: ReturnButton(onTap: _handleReturn),
      child: Column(
        children: [
          Expanded(
            child: _buildGridView(),
          ),
          _buildBottomBar(),
        ],
      ),
    );
  }

  Widget _buildLoadingView() {
    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      hasAppBar: true,
      title: '加载中...',
      leading: ReturnButton(onTap: () => Navigator.pop(context)),
      child: const Center(
        child: CircularProgressIndicator(),
      ),
    );
  }

  // 获取页面标题
  String _getPageTitle() {
    switch (_selectionLevel) {
      case SelectionLevel.country:
        return '选择目的国家';
      case SelectionLevel.city:
        return '选择目的城市';
    }
  }

  Widget _buildGridView() {
    return SingleChildScrollView(
      child: GridView.builder(
        padding: EdgeInsets.only(bottom: 2),
        shrinkWrap: true,
        physics: const NeverScrollableScrollPhysics(),
        gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: 3,
          crossAxisSpacing: 2,
          mainAxisSpacing: 2,
        ),
        itemCount: _showItems.length,
        itemBuilder: (context, index) => _buildGridItem(_showItems[index]),
      ),
    );
  }

  Widget _buildGridItem(dynamic item) {
    final ItemData itemData = _getItemData(item);
    final isSelected =
        item is City ? _selectedCities.contains(itemData.id) : false;

    return GestureDetector(
      onTap: () => _handleItemTap(item),
      child: Stack(
        children: [
          _buildImage(
            imageUrl: itemData.imageUrl,
            width: 142.w,
          ),
          Positioned(
            bottom: 2.w,
            right: 6.w,
            child: _buildItemText(itemData, isSelected),
          ),
        ],
      ),
    );
  }

  Widget _buildImage({required String imageUrl, required double width}) {
    if (imageUrl.isEmpty) {
      return BaseImage.asset(
        name: 'placeholder.png',
        width: width,
        fit: BoxFit.cover,
      );
    }

    return BaseImage.net(
      imageUrl,
      width: width,
      fit: BoxFit.cover,
    );
  }

  Widget _buildBottomBar() {
    return Container(
      padding: const EdgeInsets.all(20),
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
          if (_selectionLevel == SelectionLevel.country) _buildContinentTabs(),
          if (_selectionLevel == SelectionLevel.country)
            const SizedBox(height: 20),
          _buildSelectedCities(),
          if (_selectedCities.isNotEmpty) const SizedBox(height: 20),
          _buildGenerateButton(),
        ],
      ),
    );
  }

  Widget _buildContinentTabs() {
    return SizedBox(
      width: double.infinity,
      child: Wrap(
        alignment: WrapAlignment.start,
        spacing: 10,
        runSpacing: 10,
        children: _continentList!.continents
            .map((tag) => _buildContinentTab(tag))
            .toList(),
      ),
    );
  }

  Widget _buildContinentTab(Continent tag) {
    final bool isSelected = _selectedContinentId == tag.id;

    return IntrinsicWidth(
      child: GestureDetector(
        onTap: () => _handleContinentTap(tag),
        child: Container(
          alignment: Alignment.center,
          padding: EdgeInsets.symmetric(horizontal: 15),
          height: 38.h,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(28.w),
            color: isSelected ? AppColor.c_1D1F1E : AppColor.c_E3E3E3,
          ),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              if (isSelected)
                BaseImage.asset(
                  name: 'ic_create_picard.png',
                  width: 18.w,
                  fit: BoxFit.cover,
                ),
              Gap(5.w),
              Text(
                tag.nameCn,
                style: TextStyle(
                  fontSize: 14.sp,
                  color: isSelected ? AppColor.c_F2F2F2 : AppColor.c_1D1F1E,
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildSelectedCities() {
    return SizedBox(
      width: double.infinity,
      child: SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.start,
          children: _selectedCities
              .map((id) =>
                  _citiesList!.cities.firstWhere((item) => item.id == id))
              .map((city) => ClickableButton(
                    text: city.name,
                    onTapIcon: () =>
                        setState(() => _selectedCities.remove(city.id)),
                    icon: Icons.cancel,
                    iconColor: AppColor.closeButton,
                    margin: const EdgeInsets.only(right: 10),
                  ))
              .toList(),
        ),
      ),
    );
  }

  Widget _buildGenerateButton() {
    return GestureDetector(
      onTap: () => Navigator.pushNamed(context, Routes.poiSearch),
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
              width: 24.w,
              fit: BoxFit.cover,
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
              width: 24.w,
              fit: BoxFit.cover,
            ),
            Gap(10.w),
            BaseImage.asset(
              name: 'ic_card_generate.png',
              width: 24.w,
              fit: BoxFit.cover,
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
    );
  }
}

// 列表项数据模型
class ItemData {
  final int id;
  final String name;
  final String nameEn;
  final String imageUrl;

  ItemData({
    required this.id,
    required this.name,
    required this.nameEn,
    required this.imageUrl,
  });
}

extension _CardSelectionStateExt on _CardSelectionPageState {
  ItemData _getItemData(dynamic item) {
    if (item is Country) {
      return ItemData(
        id: item.id,
        name: item.name,
        nameEn: item.nameEn,
        imageUrl: item.imageUrl,
      );
    } else if (item is GeoState) {
      return ItemData(
        id: item.id,
        name: item.name,
        nameEn: item.nameEn,
        imageUrl: item.imageUrl,
      );
    } else {
      final city = item as City;
      return ItemData(
        id: city.id,
        name: city.name,
        nameEn: city.nameEn,
        imageUrl: city.imageUrl,
      );
    }
  }

  Widget _buildItemText(ItemData itemData, bool isSelected) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.end,
      children: [
        SizedBox(
          width: 120.w,
          child: IterText(
            itemData.name,
            alignment: AlignmentDirectional.bottomEnd,
            textAlign: TextAlign.end,
            style: TextStyle(
              fontSize: 22.sp,
              color: isSelected ? AppColor.primaryFont : AppColor.secondaryFont,
              fontWeight: AppFontWeight.black,
            ),
            maxLines: 2,
            borders: BorderProperties(
              width: 2,
              color: isSelected ? AppColor.secondaryFont : AppColor.primaryFont,
            ),
          ),
        ),
        if (itemData.nameEn.isNotEmpty)
          SizedBox(
            width: 120.w,
            child: FittedBox(
              fit: BoxFit.scaleDown,
              alignment: Alignment.centerRight,
              child: IterText(
                itemData.nameEn,
                textAlign: TextAlign.end,
                style: TextStyle(
                  fontSize: 22.sp,
                  color: isSelected
                      ? AppColor.primaryFont
                      : AppColor.secondaryFont,
                  fontWeight: AppFontWeight.black,
                ),
                maxLines: 2,
                borders: BorderProperties(
                  width: 2,
                  color: isSelected
                      ? AppColor.secondaryFont
                      : AppColor.primaryFont,
                ),
              ),
            ),
          ),
      ],
    );
  }
}
