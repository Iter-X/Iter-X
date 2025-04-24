import 'package:client/app/constants.dart';
import 'package:client/business/trip/entity/collaborator.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class TripCollaboratorsSection extends StatelessWidget {
  final TripService service;

  const TripCollaboratorsSection({
    super.key,
    required this.service,
  });

  Widget _buildAvatar(String avatarUrl) {
    if (avatarUrl.startsWith('assets/')) {
      return Image.asset(
        avatarUrl,
        width: 30.w,
        height: 30.w,
        fit: BoxFit.cover,
      );
    }
    return Image.network(
      avatarUrl,
      width: 30.w,
      height: 30.w,
      fit: BoxFit.cover,
      errorBuilder: (context, error, stackTrace) {
        return Image.asset(
          Collaborator.defaultAvatar,
          width: 30.w,
          height: 30.w,
          fit: BoxFit.cover,
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    // Layout parameters
    final double avatarSize = 30.w;
    final double avatarGap = 10.w;
    final double iconGap = 5.w;

    // Filter only accepted collaborators
    final acceptedCollaborators =
        service.collaborators.where((c) => c.status == 'Accepted').toList();

    final maxVisibleAvatars = 4;
    final displayedCollaborators =
        acceptedCollaborators.take(maxVisibleAvatars).toList();
    final remainingCount = acceptedCollaborators.length - maxVisibleAvatars;

    final stackWidth = avatarSize *
            (displayedCollaborators.length + (remainingCount > 0 ? 1 : 0)) -
        (displayedCollaborators.length - 1 + (remainingCount > 0 ? 1 : 0)) *
            avatarGap;

    if (service.loadingCollaborators) {
      return SizedBox(
        width: avatarSize,
        height: avatarSize,
        child: CircularProgressIndicator(
          strokeWidth: 2.w,
          color: AppColor.primary,
        ),
      );
    }

    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        SizedBox(
          height: avatarSize,
          width: stackWidth,
          child: Stack(
            clipBehavior: Clip.none,
            children: [
              ...List.generate(displayedCollaborators.length, (index) {
                final collaborator = displayedCollaborators[index];
                return Positioned(
                  left: index * (avatarSize - avatarGap),
                  child: Container(
                    width: avatarSize,
                    height: avatarSize,
                    decoration: BoxDecoration(
                      shape: BoxShape.circle,
                      border: Border.all(
                        color: AppColor.highlight,
                        width: 1.w,
                      ),
                      color: AppColor.bg,
                    ),
                    child: ClipOval(
                      child: _buildAvatar(collaborator.avatarUrl),
                    ),
                  ),
                );
              }),
              if (remainingCount > 0)
                Positioned(
                  left:
                      displayedCollaborators.length * (avatarSize - avatarGap),
                  child: Container(
                    width: avatarSize,
                    height: avatarSize,
                    decoration: BoxDecoration(
                      shape: BoxShape.circle,
                      color: AppColor.buttonGrayBG,
                    ),
                    child: Center(
                      child: Text(
                        '+$remainingCount',
                        style: TextStyle(
                          fontSize: 12.sp,
                          fontWeight: AppFontWeight.medium,
                          color: AppColor.grayFont,
                        ),
                      ),
                    ),
                  ),
                ),
            ],
          ),
        ),
        SizedBox(width: iconGap),
        Container(
          width: avatarSize,
          height: avatarSize,
          decoration: BoxDecoration(
            shape: BoxShape.circle,
            color: AppColor.buttonGrayBG,
          ),
          child: IconButton(
            padding: EdgeInsets.zero,
            icon: Icon(
              Icons.group_add,
              size: 18.sp,
              color: AppColor.highlight,
            ),
            onPressed: () {},
          ),
        ),
      ],
    );
  }
}
