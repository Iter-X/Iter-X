import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class SelectDaysWidget extends StatefulWidget {
  final Function(int) onTap;
  final int? selectDays;
  static final ValueNotifier<bool> isExpandedNotifier =
      ValueNotifier<bool>(false);

  const SelectDaysWidget({
    super.key,
    required this.onTap,
    this.selectDays,
  });

  @override
  State<SelectDaysWidget> createState() => _SelectDaysWidgetState();
}

class _SelectDaysWidgetState extends State<SelectDaysWidget> {
  @override
  void initState() {
    super.initState();
    SelectDaysWidget.isExpandedNotifier.value = true;
  }

  @override
  void dispose() {
    SelectDaysWidget.isExpandedNotifier.value = false;
    super.dispose();
  }

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
      ),
      child: GridView.builder(
        padding: EdgeInsets.only(top: 17.h),
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: 5,
          childAspectRatio: 1.5,
        ),
        itemBuilder: (c, i) {
          return GestureDetector(
            onTap: () => widget.onTap.call(i + 1),
            child: Container(
              width: 56.w,
              height: 56.w,
              alignment: Alignment.center,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: widget.selectDays == i + 1
                    ? AppColor.highlight
                    : Colors.transparent,
              ),
              child: Text(
                '${i + 1}天',
                style: TextStyle(
                  color: widget.selectDays == i + 1
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
