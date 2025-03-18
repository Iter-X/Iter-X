import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

import '../../../common/material/image.dart';
import '../../../common/material/text_field.dart';
import '../../../common/utils/color.dart';

// 手动创建和外部链接 底部样式
class CreateManuallyWidget extends StatelessWidget {
  final TextEditingController controller;
  final FocusNode focusNode;
  final bool hasFocus;
  final int selectIndex;
  final Function onTap;

  const CreateManuallyWidget({
    super.key,
    required this.controller,
    required this.focusNode,
    required this.hasFocus,
    required this.selectIndex,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    final double bottomPadding = MediaQuery.of(context).viewInsets.bottom;

    return Container(
      width: double.infinity,
      height: 56.h,
      margin: EdgeInsets.only(
        bottom: bottomPadding + (bottomPadding > 0 ? 10.h : 90.h),
        left: 35.w,
        right: 35.w,
      ),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(28.w),
        color: BaseColor.c_E3E3E3,
      ),
      child: Row(
        children: [
          Expanded(
            child: BaseTextFieldWidget(
              controller: controller,
              focusNode: focusNode,
              hintText: selectIndex == 0 ? '输入目的地开始规划行程' : '粘贴外部链接，智能识别规划',
              hintStyle: TextStyle(
                fontSize: 15.sp,
                color: BaseColor.c_999999,
              ),
              style: TextStyle(
                fontSize: 15.sp,
                color: BaseColor.c_1D1F1E,
              ),
              contentPadding: EdgeInsets.only(
                left: 24.w,
                right: 24.w,
              ),
            ),
          ),
          GestureDetector(
            onTap: () => onTap.call(),
            child: Container(
              width: 107.w,
              height: 46.h,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(24.w),
                color: BaseColor.c_1D1F1E,
              ),
              margin: EdgeInsets.only(right: 6.w),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  BaseImage.asset(
                    name: 'ic_create_picard.png',
                    size: 18.w,
                  ),
                  Gap(5.w),
                  Text(
                    selectIndex == 0 ? '图卡选择' : '粘贴',
                    style: TextStyle(
                      fontSize: 14.sp,
                      color: BaseColor.c_f2f2f2,
                    ),
                  )
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}

// 截图创建 底部样式
class CreatePhotoWidget extends StatelessWidget {
  final Function onTap;

  const CreatePhotoWidget({
    super.key,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => onTap.call(),
      child: Container(
        width: double.infinity,
        height: 52.h,
        margin: EdgeInsets.only(
          bottom: 90.h,
          left: 35.w,
          right: 35.w,
        ),
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(28.w),
          color: BaseColor.c_1D1F1E,
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            BaseImage.asset(
              name: 'ic_select_photo.png',
              size: 24.w,
            ),
            Gap(10.w),
            Text(
              '选择照片上传',
              style: TextStyle(
                fontSize: 18.sp,
                color: BaseColor.c_f2f2f2,
              ),
            )
          ],
        ),
      ),
    );
  }
}
