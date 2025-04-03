import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class StatsCard extends StatelessWidget {
  final String value;
  final String label;
  final String emoji;

  const StatsCard({
    super.key,
    required this.value,
    required this.label,
    required this.emoji,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(15.w),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(AppConfig.boxRadius),
      ),
      child: Row(
        children: [
          Text(
            emoji,
            style: TextStyle(
              fontSize: 35.sp,
              fontWeight: AppFontWeight.semiBold,
            ),
          ),
          SizedBox(width: 10.w),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  label,
                  style: TextStyle(
                    fontSize: 14.sp,
                    fontWeight: AppFontWeight.semiBold,
                    color: AppColor.grayFont,
                    letterSpacing: 0.5.sp,
                  ),
                ),
                Text(
                  value,
                  style: TextStyle(
                    fontSize: 18.sp,
                    fontWeight: AppFontWeight.extraBold,
                    color: AppColor.highlight,
                    letterSpacing: 0.5.sp,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
