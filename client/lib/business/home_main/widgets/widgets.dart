import 'package:client/common/material/image.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

// 首页底部按钮样式
class ItemHomeBottomWidget extends StatelessWidget {
  String img;
  Function onTap;

  ItemHomeBottomWidget({
    super.key,
    required this.img,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => onTap.call(),
      child: BaseImage.asset(
        name: img,
        size: 38.w,
      ),
    );
  }
}
