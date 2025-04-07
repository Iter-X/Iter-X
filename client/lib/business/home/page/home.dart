import 'package:client/app/constants.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/image.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      top: true,
      bottom: false,
      hasAppBar: false,
      child: Column(
        children: [
          Gap(48.h),
          Container(
            margin: EdgeInsets.only(
              left: 35.w,
              right: 35.w,
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Hi, Leo',
                      style: TextStyle(
                        fontSize: 30.sp,
                        color: AppColor.c_1D1F1E,
                        fontWeight: AppFontWeight.semiBold,
                      ),
                    ),
                    Text(
                      'Explore the world',
                      style: TextStyle(
                        fontSize: 18.sp,
                        color: const Color(0xFF888888),
                      ),
                    ),
                  ],
                ),
                BaseImage.asset(
                  name: 'ic_default_head.png',
                  width: 70.w,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
