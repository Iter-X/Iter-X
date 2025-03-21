import 'package:client/common/material/image.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

// 返回按钮
class ButtonBackWidget extends StatelessWidget {
  const ButtonBackWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Navigator.pop(context);
      },
      child: Container(
        padding: EdgeInsets.all(20.w),
        child: BaseImage.asset(
          name: 'return_btn.png',
          width: 28.w,
          fit: BoxFit.fitWidth,
        ),
      ),
    );
  }
}
