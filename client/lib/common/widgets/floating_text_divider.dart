import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class FloatingTextDivider extends StatelessWidget {
  final Widget child;
  final Color color;
  final double thickness;
  final double spacing;
  final double floatingOffset;

  const FloatingTextDivider({
    super.key,
    required this.child,
    this.color = Colors.grey,
    this.thickness = 1,
    this.spacing = 8,
    this.floatingOffset = -8,
  });

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: Alignment.center,
      clipBehavior: Clip.none,
      children: [
        Row(
          children: [
            Expanded(
              child: Divider(
                color: color,
                thickness: thickness,
                endIndent: spacing.w,
                height: thickness,
              ),
            ),
            const SizedBox(),
            Expanded(
              child: Divider(
                color: color,
                thickness: thickness,
                indent: spacing.w,
                height: thickness,
              ),
            ),
          ],
        ),
        Positioned(
          top: floatingOffset.sp,
          child: Container(
            padding: EdgeInsets.symmetric(
              horizontal: 10.w,
            ),
            decoration: BoxDecoration(
              color: Colors.white,
            ),
            child: child,
          ),
        )
      ],
    );
  }
}
