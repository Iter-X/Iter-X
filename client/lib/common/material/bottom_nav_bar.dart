import 'package:client/app/constants.dart';
import 'package:client/common/material/image.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class BottomNavConfig {
  final List<BottomBarItem> items;
  final int selectedIndex;
  final ValueChanged<int> onIndexChanged;
  final Color? backgroundColor;
  final Color? borderColor;
  final Widget? centerButton;
  final VoidCallback? onCenterButtonTap;

  const BottomNavConfig({
    required this.items,
    required this.selectedIndex,
    required this.onIndexChanged,
    this.backgroundColor,
    this.borderColor,
    this.centerButton,
    this.onCenterButtonTap,
  });
}

class BottomNavBar extends StatelessWidget {
  final BottomNavConfig config;

  const BottomNavBar({
    super.key,
    required this.config,
  });

  @override
  Widget build(BuildContext context) {
    final itemCount = config.items.length;
    final halfItems = itemCount ~/ 2;

    List<Widget> rowChildren = [];

    // Add first half of items
    for (var i = 0; i < halfItems; i++) {
      rowChildren.add(
        Expanded(
          child: _buildBottomBarItem(
            item: config.items[i],
            isSelected: config.selectedIndex == i,
            onTap: () => config.onIndexChanged(i),
          ),
        ),
      );
    }

    // Add center button if provided
    if (config.centerButton != null) {
      rowChildren.add(
        Expanded(
          child: Center(
            child: GestureDetector(
              onTap: config.onCenterButtonTap,
              child: config.centerButton!,
            ),
          ),
        ),
      );
    }

    // Add second half of items
    for (var i = halfItems; i < itemCount; i++) {
      rowChildren.add(
        Expanded(
          child: _buildBottomBarItem(
            item: config.items[i],
            isSelected: config.selectedIndex == i,
            onTap: () => config.onIndexChanged(i),
          ),
        ),
      );
    }

    return Container(
      decoration: BoxDecoration(
        color: config.backgroundColor ?? AppColor.bottomBar,
        border: Border(
          top: BorderSide(
            color: config.borderColor ?? AppColor.bottomBarLine,
            width: 1.w,
          ),
        ),
      ),
      padding: EdgeInsets.symmetric(vertical: 15.h),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: rowChildren,
      ),
    );
  }

  Widget _buildBottomBarItem({
    required BottomBarItem item,
    required bool isSelected,
    required VoidCallback onTap,
  }) {
    return GestureDetector(
      onTap: onTap,
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          if (item.iconImage != null)
            BaseImage.asset(
              name: isSelected ? item.selectedIconImage! : item.iconImage!,
              size: 38.w,
            )
          else
            Icon(
              isSelected ? item.selectedIcon! : item.icon!,
              size: 30.w,
              color: isSelected ? AppColor.highlight : AppColor.primaryFont,
            ),
          if (item.label != null)
            Text(
              item.label!,
              style: TextStyle(
                fontSize: 12.sp,
                fontWeight: AppFontWeight.regular,
                color: isSelected ? AppColor.highlight : AppColor.primaryFont,
              ),
            ),
        ],
      ),
    );
  }
}

class BottomBarItem {
  final IconData? icon;
  final IconData? selectedIcon;
  final String? iconImage;
  final String? selectedIconImage;
  final String? label;

  const BottomBarItem({
    this.icon,
    this.selectedIcon,
    this.iconImage,
    this.selectedIconImage,
    this.label,
  }) : assert(
          (icon != null || (iconImage != null && selectedIconImage != null)),
          'Either provide icon or both iconImage and selectedIconImage',
        );
}
