import 'package:client/common/material/image.dart';
import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class PreferenceButton extends StatelessWidget {
  final VoidCallback? onTap;
  final Color? color;

  const PreferenceButton({super.key, this.onTap, this.color});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        margin: EdgeInsets.only(right: 20.w),
        width: 28.w,
        height: 28.w,
        alignment: Alignment.center,
        child: BaseImage.asset(
          name: 'setting.svg',
          size: 28.w,
          color: color ?? AppColor.primaryFont,
        ),
      ),
    );
  }
}
