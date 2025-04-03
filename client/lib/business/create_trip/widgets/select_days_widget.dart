import 'package:client/common/utils/color.dart';
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
      height: 250.h,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.only(
          topLeft: Radius.circular(10.w),
          topRight: Radius.circular(10.w),
        ),
        color: BaseColor.c_F2F2F2,
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
                color: selectDays == i + 1 ? BaseColor.c_375F77 : Colors.transparent,
              ),
              child: Text(
                '${i + 1}天',
                style: TextStyle(
                  color: selectDays == i + 1 ? Colors.white : BaseColor.c_1D1F1E,
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
