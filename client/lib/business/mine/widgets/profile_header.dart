import 'package:cached_network_image/cached_network_image.dart';
import 'package:client/common/utils/app_config.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';

class ProfileHeader extends StatelessWidget {
  final String name;
  final String avatarUrl;

  const ProfileHeader({
    super.key,
    required this.name,
    required this.avatarUrl,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(20),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Hi, $name 👋',
                  style: const TextStyle(
                    fontSize: 30,
                    fontWeight: AppFontWeight.semiBold,
                    color: BaseColor.primaryFont,
                  ),
                ),
                const SizedBox(height: 8),
                const Text(
                  'Explore the world',
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: AppFontWeight.medium,
                    color: BaseColor.grayFont,
                    letterSpacing: 0.5,
                  ),
                ),
              ],
            ),
          ),
          Container(
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              border: Border.all(
                color: BaseColor.highlight,
                width: 1,
              ),
            ),
            child: CircleAvatar(
              radius: 40,
              backgroundColor: Colors.white,
              child: ClipOval(
                child: CachedNetworkImage(
                  imageUrl: avatarUrl,
                  width: 70,
                  height: 70,
                  fit: BoxFit.cover,
                  placeholder: (context, url) =>
                      const CircularProgressIndicator(),
                  errorWidget: (context, url, error) {
                    return const Icon(Icons.person,
                        size: 40); // TODO: change to a default avatar
                  },
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
