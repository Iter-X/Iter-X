import 'package:client/app/constants.dart';
import 'package:client/business/create_trip/entity/geo_entity.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

class CityDropdown extends StatefulWidget {
  final List<City> cities;
  final int selectedIndex;
  final Function(int) onCityChanged;

  const CityDropdown({
    super.key,
    required this.cities,
    required this.selectedIndex,
    required this.onCityChanged,
  });

  @override
  State<CityDropdown> createState() => _CityDropdownState();
}

class _CityDropdownState extends State<CityDropdown> {
  bool _isExpanded = false;
  final LayerLink _layerLink = LayerLink();
  OverlayEntry? _overlayEntry;

  @override
  void dispose() {
    _overlayEntry?.remove();
    _overlayEntry = null;
    super.dispose();
  }

  void _removeOverlay() {
    _overlayEntry?.remove();
    _overlayEntry = null;
    if (mounted) {
      setState(() {
        _isExpanded = false;
      });
    }
  }

  void _toggleDropdown() {
    setState(() {
      _isExpanded = !_isExpanded;
      if (_isExpanded) {
        _overlayEntry = _createOverlayEntry();
        Overlay.of(context).insert(_overlayEntry!);
      } else {
        _removeOverlay();
      }
    });
  }

  OverlayEntry _createOverlayEntry() {
    RenderBox renderBox = context.findRenderObject() as RenderBox;
    var size = renderBox.size;
    final offset = renderBox.localToGlobal(Offset.zero);

    return OverlayEntry(
      builder: (context) => Stack(
        children: [
          Positioned.fill(
            child: GestureDetector(
              onTap: _removeOverlay,
              behavior: HitTestBehavior.opaque,
              child: Container(
                color: Colors.transparent,
              ),
            ),
          ),
          Positioned(
            left: 0,
            right: 0,
            top: offset.dy + size.height + 5.0,
            child: Material(
              elevation: 4,
              child: Container(
                constraints: BoxConstraints(maxHeight: 200.h),
                decoration: BoxDecoration(
                  color: AppColor.bg,
                ),
                child: ListView.builder(
                  padding: EdgeInsets.zero,
                  shrinkWrap: true,
                  itemCount: widget.cities.length,
                  itemBuilder: (context, index) {
                    final city = widget.cities[index];
                    return InkWell(
                      onTap: () {
                        widget.onCityChanged(index);
                        _removeOverlay();
                      },
                      child: Container(
                        padding: EdgeInsets.symmetric(
                          vertical: 12.h,
                          horizontal: 16.w,
                        ),
                        alignment: Alignment.center,
                        decoration: BoxDecoration(
                          color: index == widget.selectedIndex
                              ? AppColor.selectedItem
                              : Colors.transparent,
                        ),
                        child: Text(
                          city.name,
                          style: TextStyle(
                            color: AppColor.primaryFont,
                            fontSize: 16.sp,
                            fontWeight: index == widget.selectedIndex
                                ? AppFontWeight.medium
                                : AppFontWeight.regular,
                          ),
                          textAlign: TextAlign.center,
                        ),
                      ),
                    );
                  },
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return CompositedTransformTarget(
      link: _layerLink,
      child: GestureDetector(
        onTap: _toggleDropdown,
        child: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              widget.cities[widget.selectedIndex].name,
              style: TextStyle(
                color: AppColor.primaryFont,
                fontSize: 18.sp,
                fontWeight: AppFontWeight.medium,
              ),
            ),
            Gap(5.w),
            Icon(
              _isExpanded ? Icons.keyboard_arrow_up : Icons.keyboard_arrow_down,
              color: AppColor.primaryFont,
              size: 24.sp,
            ),
          ],
        ),
      ),
    );
  }
} 