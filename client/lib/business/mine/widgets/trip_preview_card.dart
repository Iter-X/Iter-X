import 'package:cached_network_image/cached_network_image.dart';
import 'package:client/app/constants.dart';
import 'package:client/common/material/iter_text.dart';
import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import '../entity/user_profile.dart';

class TripPreviewCard extends StatelessWidget {
  final TripPreview trip;
  final bool isFirst;
  final bool isLast;

  const TripPreviewCard({
    super.key,
    required this.trip,
    this.isFirst = false,
    this.isLast = false,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 160,
      margin: EdgeInsets.fromLTRB(15, isFirst ? 0 : 10, 15, isLast ? 15 : 0),
      child: ClipRRect(
        borderRadius: BorderRadius.circular(12),
        child: Stack(
          fit: StackFit.expand,
          children: [
            CachedNetworkImage(
              imageUrl: trip.imageUrl,
              fit: BoxFit.cover,
              placeholder: (context, url) => Container(
                color: Colors.grey[200],
                child: const Center(
                  child: CircularProgressIndicator(),
                ),
              ),
              errorWidget: (context, url, error) {
                return Container(
                  color: Colors.grey[200],
                  child: const Center(
                    child: Icon(Icons.error_outline, size: 40),
                  ),
                );
              },
            ),
            Container(
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  begin: Alignment.topCenter,
                  end: Alignment.bottomCenter,
                  colors: [
                    Colors.transparent,
                    Colors.black.withOpacity(0.6),
                  ],
                ),
              ),
            ),
            Positioned(
              bottom: 15,
              right: 15,
              child: IterText(
                trip.title,
                style: TextStyle(
                  fontSize: 22.sp,
                  color: AppColor.secondary,
                  fontWeight: AppFontWeight.black,
                ),
                borders: BorderProperties(
                  width: 2.sp,
                  color: AppColor.primary,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
