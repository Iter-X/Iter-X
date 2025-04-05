import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class SelectDaysWidget extends StatelessWidget {
  final Function(int) onTap;
  final int? selectDays;

  const SelectDaysWidget({
    super.key,
    required this.onTap,
    this.selectDays,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: 350.h,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.only(
          topLeft: Radius.circular(AppConfig.boxRadius),
          topRight: Radius.circular(AppConfig.boxRadius),
        ),
        color: AppColor.bottomBar,
        border: Border(
          top: BorderSide(
            color: AppColor.bottomBarLine,
            width: 1.w,
          ),
        ),
      ),
      child: GridView.builder(
        padding: EdgeInsets.only(top: 17.h),
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: 5,
          childAspectRatio: 1.5,
        ),
        itemBuilder: (c, i) {
          return GestureDetector(
            onTap: () => onTap.call(i + 1),
            child: Container(
              width: 56.w,
              height: 56.w,
              alignment: Alignment.center,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: selectDays == i + 1
                    ? AppColor.highlight
                    : Colors.transparent,
              ),
              child: Text(
                '${i + 1}天',
                style: TextStyle(
                  color: selectDays == i + 1
                      ? AppColor.white
                      : AppColor.primaryFont,
                  fontSize: 18.sp,
                ),
              ),
            ),
          );
        },
        itemCount: 99,
      ),
    );
  }
}
