import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class TextDivider extends StatelessWidget {
  final Widget child;
  final Color color;
  final double thickness;
  final double spacing;
  final double? height;

  const TextDivider({
    super.key,
    required this.child,
    this.color = Colors.grey,
    this.thickness = 1,
    this.spacing = 8,
    this.height,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: height,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Expanded(
            child: Divider(
              color: color,
              thickness: thickness,
              endIndent: spacing.w,
            ),
          ),
          child,
          Expanded(
            child: Divider(
              color: color,
              thickness: thickness,
              indent: spacing.w,
            ),
          ),
        ],
      ),
    );
  }
}
