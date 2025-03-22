import 'package:client/common/material/image.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

// 创建行程按钮样式
class ItemCreateWidget extends StatelessWidget {
  bool isSelected;
  Function onTap;
  String img;
  String text;

  ItemCreateWidget({
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
          borderRadius: BorderRadius.circular(12.w),
          color: isSelected ? BaseColor.c_1D1F1E : BaseColor.c_E3E3E3,
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
                color: isSelected ? Colors.white : BaseColor.c_1D1F1E,
              ),
            )
          ],
        ),
      ),
    );
  }
}
