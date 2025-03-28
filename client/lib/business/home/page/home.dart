import 'package:client/business/auth/page/input_code.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/utils/asset.dart';
import 'package:client/common/utils/color.dart';
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
    return Column(
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
                      color: BaseColor.c_1D1F1E,
                      fontWeight: FontWeight.w600,
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
              BaseImage.net(
                null,
                size: 70.w,
                assetName: AssetUtil.getAsset('ic_default_head.png'),
              ),
            ],
          ),
        ),
      ],
    );
  }
}
