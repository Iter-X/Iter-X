import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';

class StatsCard extends StatelessWidget {
  final String value;
  final String label;
  final String emoji;

  const StatsCard({
    super.key,
    required this.value,
    required this.label,
    required this.emoji,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(15),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(AppConfig.boxRadius),
      ),
      child: Row(
        children: [
          Text(
            emoji,
            style: const TextStyle(
              fontSize: 35,
              fontWeight: AppFontWeight.semiBold,
            ),
          ),
          const SizedBox(width: 10),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  label,
                  style: const TextStyle(
                    fontSize: 14,
                    fontWeight: AppFontWeight.semiBold,
                    color: AppColor.grayFont,
                    letterSpacing: 0.5,
                  ),
                ),
                Text(
                  value,
                  style: const TextStyle(
                    fontSize: 18,
                    fontWeight: AppFontWeight.extraBold,
                    color: AppColor.highlight,
                    letterSpacing: 0.5,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
