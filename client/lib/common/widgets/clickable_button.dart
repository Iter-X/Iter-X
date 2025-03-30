import 'package:client/common/utils/app_config.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class ClickableButton extends StatelessWidget {
  final String text;
  final VoidCallback onTapIcon;
  final VoidCallback? onTapText;
  final double? height;
  final double? width;
  final IconData? icon;
  final Color? iconColor;
  final double rotationAngle;
  final double? gap;

  const ClickableButton({
    super.key,
    required this.text,
    required this.onTapIcon,
    this.onTapText,
    this.height,
    this.width,
    this.icon,
    this.iconColor,
    this.rotationAngle = 0.0,
    this.gap,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTapText,
      child: Container(
        height: height,
        width: width,
        padding: EdgeInsets.all(9.w),
        decoration: BoxDecoration(
          color: BaseColor.secondary,
          borderRadius: BorderRadius.circular(AppConfig.cornerRadius),
        ),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            ConstrainedBox(
              constraints: BoxConstraints(maxWidth: 120.w),
              child: Text(
                text,
                style: TextStyle(
                  color: BaseColor.primary,
                  fontSize: 16.sp,
                ),
                overflow: TextOverflow.ellipsis,
                maxLines: 1,
              ),
            ),
            SizedBox(width: gap ?? 5.w),
            GestureDetector(
              onTap: onTapIcon,
              child: Transform.rotate(
                angle: rotationAngle,
                child: Icon(
                  icon ?? Icons.cancel,
                  color: iconColor ?? BaseColor.primary,
                  size: 22.sp,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
