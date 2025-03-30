import 'package:client/common/material/image.dart';
import 'package:client/common/utils/app_config.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

// 创建行程按钮样式
class ItemCreateWidget extends StatelessWidget {
  final bool isSelected;
  final Function onTap;
  final String img;
  final String text;

  const ItemCreateWidget({
    super.key,
    required this.isSelected,
    required this.onTap,
    required this.img,
    required this.text,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => onTap.call(),
      child: Container(
        width: 170.w,
        height: 70.h,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(AppConfig.borderRadius),
          color: isSelected ? BaseColor.primary : BaseColor.buttonGrayBG,
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            BaseImage.asset(
              name: img,
              size: 30.w,
            ),
            Gap(10.w),
            Text(
              text,
              style: TextStyle(
                fontSize: 18.sp,
                color: isSelected
                    ? BaseColor.secondaryFont
                    : BaseColor.primaryFont,
              ),
            )
          ],
        ),
      ),
    );
  }
}
