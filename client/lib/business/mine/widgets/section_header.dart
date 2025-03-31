import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';

class SectionHeader extends StatelessWidget {
  final String title;
  final String emoji;
  final VoidCallback? onTap;

  const SectionHeader({
    super.key,
    required this.title,
    required this.emoji,
    this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(AppConfig.boxRadius),
        ),
        child: Padding(
          padding: const EdgeInsets.all(15),
          child: Row(
            children: [
              Text(
                emoji,
                style: const TextStyle(
                  fontSize: 22,
                  fontWeight: AppFontWeight.semiBold,
                ),
              ),
              const SizedBox(width: 5),
              Text(
                title,
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: AppFontWeight.semiBold,
                  color: AppColor.primaryFont,
                  letterSpacing: 0.5,
                ),
              ),
              const Spacer(),
              const Icon(
                Icons.arrow_forward_ios,
                size: 16,
                color: AppColor.grayFont,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
