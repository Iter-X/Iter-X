import 'package:client/common/material/image.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class ReturnButton extends StatelessWidget {
  final VoidCallback? onTap;
  final Color? color;

  const ReturnButton({super.key, this.onTap, this.color});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        if (onTap != null) {
          onTap!();
        } else {
          Navigator.pop(context);
        }
      },
      child: Container(
        margin: EdgeInsets.only(left: 20.w),
        width: 28.w,
        height: 28.w,
        alignment: Alignment.center,
        child: BaseImage.asset(
          name: 'return_btn.svg',
          size: 28.w,
          color: color ?? BaseColor.primaryFont,
        ),
      ),
    );
  }
}
