import 'package:client/business/create_trip/service/poi_search_service.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/utils/app_config.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/widgets/base_button.dart';
import 'package:client/common/widgets/clickable_button.dart';
import 'package:client/common/widgets/preference_button.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';
import 'package:provider/provider.dart';

class PoiSearchPage extends StatefulWidget {
  const PoiSearchPage({super.key});

  @override
  State<PoiSearchPage> createState() => _PoiSearchPageState();
}

class _PoiSearchPageState extends State<PoiSearchPage> {
  final TextEditingController _searchController = TextEditingController();
  final FocusNode _focusNode = FocusNode();
  bool _hasFocus = false;

  @override
  void initState() {
    super.initState();
    _focusNode.addListener(_onFocusChange);
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<PoiSearchService>().searchPoi('');
    });
  }

  @override
  void dispose() {
    _focusNode.removeListener(_onFocusChange);
    _focusNode.dispose();
    _searchController.dispose();
    super.dispose();
  }

  void _onFocusChange() {
    setState(() {
      _hasFocus = _focusNode.hasFocus;
    });
  }

  Widget _buildPoiCard(PoiItem poi) {
    return Consumer<PoiSearchService>(
      builder: (context, service, child) {
        final isSelected =
            service.selectedPois.any((item) => item.name == poi.name);

        return GestureDetector(
          onTap: () => service.togglePoiSelection(poi),
          child: Container(
            padding: EdgeInsets.only(top: 2.h),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                ClipRRect(
                  child: Stack(
                    children: [
                      Image.network(
                        poi.imageUrl,
                        width: 142.w,
                        height: 142.w,
                        fit: BoxFit.cover,
                      ),
                      if (isSelected)
                        Positioned(
                          right: 2.w,
                          bottom: 2.w,
                          child: Icon(
                            Icons.verified,
                            color: BaseColor.secondary,
                            size: 33.sp,
                          ),
                        ),
                    ],
                  ),
                ),
                Expanded(
                  child: Container(
                    padding: EdgeInsets.only(left: 10.w, right: 10.w),
                    decoration: BoxDecoration(
                      color: isSelected
                          ? BaseColor.selectedItem
                          : Colors.transparent,
                    ),
                    child: SizedBox(
                      height: 142.w,
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Text(
                            poi.name,
                            style: TextStyle(
                              color: BaseColor.primaryFont,
                              fontSize: 18.sp,
                              fontWeight: AppFontWeight.medium,
                            ),
                            maxLines: 2,
                            overflow: TextOverflow.ellipsis,
                          ),
                          Text(
                            poi.englishName,
                            style: TextStyle(
                              color: BaseColor.primaryFont,
                              fontSize: 14.sp,
                            ),
                            maxLines: 2,
                            overflow: TextOverflow.ellipsis,
                          ),
                          SizedBox(height: 5.h),
                          Row(
                            children: [
                              ...List.generate(5, (index) {
                                if (index < poi.rating.floor()) {
                                  return Icon(Icons.star,
                                      color: BaseColor.highlight, size: 18.sp);
                                } else if (index == poi.rating.floor() &&
                                    poi.rating % 1 > 0) {
                                  return Icon(Icons.star_half,
                                      color: BaseColor.highlight, size: 18.sp);
                                }
                                return Icon(Icons.star_border,
                                    color: BaseColor.highlight, size: 18.sp);
                              }),
                              Gap(8.w),
                              Text(
                                poi.reviews,
                                style: TextStyle(
                                  color: BaseColor.grayFont,
                                  fontSize: 14.sp,
                                ),
                              ),
                            ],
                          ),
                          SizedBox(height: 5.h),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                poi.duration,
                                style: TextStyle(
                                  color: BaseColor.grayFont,
                                  fontSize: 14.sp,
                                ),
                              ),
                              Text(
                                poi.popularity,
                                style: TextStyle(
                                  color: BaseColor.grayFont,
                                  fontSize: 14.sp,
                                ),
                              ),
                            ],
                          ),
                        ],
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        );
      },
    );
  }

  Widget _buildSelectedPois() {
    return Consumer<PoiSearchService>(
      builder: (context, service, child) {
        if (service.selectedPois.isEmpty) return SizedBox.shrink();

        final selectedPois = service.selectedPois.toList();
        final maxVisiblePois = 4; // 最多显示3个POI
        final isExpanded = ValueNotifier<bool>(false);

        return Container(
          padding: EdgeInsets.only(bottom: 15.h),
          child: ValueListenableBuilder<bool>(
            valueListenable: isExpanded,
            builder: (context, expanded, child) {
              final visiblePois = expanded
                  ? selectedPois
                  : selectedPois.take(maxVisiblePois).toList();
              final hasMore = selectedPois.length > maxVisiblePois;

              return Wrap(
                spacing: 10.w,
                runSpacing: 10.h,
                children: [
                  ...visiblePois.map((poi) {
                    return ClickableButton(
                      text: poi.name,
                      onTapIcon: () => service.removeSelectedPoi(poi),
                      icon: Icons.cancel,
                      iconColor: BaseColor.closeButton,
                    );
                  }),
                  if (hasMore)
                    ClickableButton(
                      text: expanded
                          ? ''
                          : '+ ${selectedPois.length - maxVisiblePois}',
                      onTapIcon: () => isExpanded.value = !isExpanded.value,
                      icon: Icons.expand_circle_down,
                      iconColor: BaseColor.highlight,
                      gap: expanded ? 0 : 5.w,
                      rotationAngle: expanded ? 0 : 1 * 3.14,
                    ),
                ],
              );
            },
          ),
        );
      },
    );
  }

  Widget _buildBottomBar() {
    return Container(
      padding: EdgeInsets.symmetric(horizontal: 15.w, vertical: 15.h),
      decoration: BoxDecoration(
        color: Colors.white,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          _buildSelectedPois(),
          Row(
            children: [
              Expanded(
                child: Container(
                  decoration: BoxDecoration(
                    color: BaseColor.bg,
                    borderRadius: BorderRadius.circular(AppConfig.cornerRadius),
                  ),
                  child: TextField(
                    controller: _searchController,
                    focusNode: _focusNode,
                    onChanged: (value) {
                      context.read<PoiSearchService>().searchPoi(value);
                    },
                    onSubmitted: (value) {
                      _focusNode.unfocus();
                      context.read<PoiSearchService>().searchPoi(value);
                    },
                    decoration: InputDecoration(
                      hintText: '输入搜索',
                      hintStyle: TextStyle(
                        color: BaseColor.inputPlaceholder,
                        fontSize: 16.sp,
                      ),
                      prefixIcon: Icon(
                        Icons.travel_explore,
                        color: BaseColor.primaryFont,
                        size: 28.sp,
                      ),
                      border: InputBorder.none,
                      contentPadding: EdgeInsets.symmetric(
                          horizontal: 16.w, vertical: 12.h),
                    ),
                  ),
                ),
              ),
              Gap(10.w),
              Consumer<PoiSearchService>(
                builder: (context, service, child) {
                  final selectedCount = service.selectedPois.length;
                  return BaseButton(
                    text: _hasFocus
                        ? '搜索'
                        : selectedCount > 0
                            ? '已选$selectedCount 确认增加'
                            : '暂不添加',
                    textColor: BaseColor.secondaryFont,
                    width: 146.w,
                    textSize: 16.sp,
                    onTap: () {
                      if (_hasFocus) {
                        _focusNode.unfocus();
                        context
                            .read<PoiSearchService>()
                            .searchPoi(_searchController.text);
                      } else {
                        if (selectedCount > 0) {
                          // TODO: add selected pois to trip
                        }
                        Navigator.pop(context);
                      }
                    },
                  );
                },
              ),
            ],
          )
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      backgroundColor: BaseColor.bg,
      bottomColor: BaseColor.bottomBar,
      leading: ReturnButton(),
      title: '洛杉矶Los Angeles',
      actions: [PreferenceButton()],
      child: Column(
        children: [
          Expanded(
            child: Consumer<PoiSearchService>(
              builder: (context, service, child) {
                if (service.isLoading) {
                  return const Center(
                      child: CircularProgressIndicator()); // TODO: skeleton
                }
                return ListView(
                  children:
                      service.poiList.map((poi) => _buildPoiCard(poi)).toList(),
                );
              },
            ),
          ),
          _buildBottomBar(),
        ],
      ),
    );
  }
}
