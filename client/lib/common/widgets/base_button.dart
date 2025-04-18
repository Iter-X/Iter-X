import 'package:client/common/material/image.dart';
import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

class BaseButton extends StatelessWidget {
  final String? iconName;
  final double iconSize;
  final String text;
  final double textSize;
  final Color textColor;
  final double width;
  final double height;
  final Color? backgroundColor;
  final bool hasBorder;
  final Color borderColor;
  final VoidCallback? onTap;
  final double borderRadius;
  final double gap;
  final bool isLoading;
  final Widget? loadingWidget;

  BaseButton({
    super.key,
    this.iconName,
    double? iconSize,
    required this.text,
    double? textSize,
    Color? textColor,
    double? width,
    double? height,
    this.backgroundColor,
    this.hasBorder = false,
    Color? borderColor,
    this.onTap,
    double? borderRadius,
    double? gap,
    this.isLoading = false,
    this.loadingWidget,
  })  : iconSize = iconSize?.w ?? 26.w,
        textSize = textSize?.sp ?? 18.sp,
        textColor = textColor ?? AppColor.primary,
        width = width?.w ?? double.infinity,
        height = height?.h ?? 52.h,
        borderColor = borderColor ?? AppColor.primary,
        borderRadius = borderRadius?.r ?? 24.r,
        gap = gap?.w ?? 20.w;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        width: width,
        height: height,
        decoration: BoxDecoration(
          color: backgroundColor ??
              (hasBorder ? Colors.transparent : AppColor.primary),
          borderRadius: BorderRadius.circular(borderRadius),
          border: hasBorder
              ? Border.all(
                  width: 1.w,
                  color: borderColor,
                )
              : null,
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            if (iconName != null) ...[
              BaseImage.asset(
                name: iconName!,
                size: iconSize,
              ),
              Gap(gap),
            ],
            isLoading
                ? loadingWidget ?? CircularProgressIndicator()
                : Text(
                    text,
                    style: TextStyle(
                      color: textColor,
                      fontSize: textSize,
                    ),
                  )
          ],
        ),
      ),
    );
  }
}
