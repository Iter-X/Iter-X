import 'package:client/app/constants.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/widgets/preference_button.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:provider/provider.dart';

class TripOverviewPage extends StatefulWidget {
  const TripOverviewPage({super.key});

  @override
  State<TripOverviewPage> createState() => _TripOverviewPageState();
}

class _TripOverviewPageState extends State<TripOverviewPage> {
  @override
  void initState() {
    super.initState();
    // 初始化数据
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final service = context.read<TripService>();
      service.fetchTripData();
    });
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      hasAppBar: true,
      bottom: false,
      leading: ReturnButton(),
      actions: [PreferenceButton()],
      child: Consumer<TripService>(
        builder: (context, service, child) {
          if (service.isLoading) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          }

          return SingleChildScrollView(
            child: Padding(
              padding: EdgeInsets.symmetric(horizontal: 20.w),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SizedBox(height: 20.h),
                  SizedBox(
                    width: double.infinity,
                    child: Text(
                      service.title,
                      style: TextStyle(
                        fontSize: 30.sp,
                        fontWeight: AppFontWeight.bold,
                        color: AppColor.primaryFont,
                      ),
                      maxLines: 2,
                      overflow: TextOverflow.ellipsis,
                    ),
                  ),
                  SizedBox(height: 5.h),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      Text(
                        '2025/01/25 | 10天',
                        style: TextStyle(
                          fontSize: 20.sp,
                          fontWeight: AppFontWeight.regular,
                          color: AppColor.primaryFont,
                        ),
                      ),
                      _buildParticipantsSection(service),
                    ],
                  ),
                  SizedBox(height: 70.h),
                  ...service.days.map((day) {
                    return Column(
                      children: [
                        _buildDaySection(
                          day: day.day,
                          date: day.date,
                          cities: day.cities,
                          spots: day.spots,
                        ),
                        SizedBox(height: 15.h),
                      ],
                    );
                  }),
                  _buildUnplannedSection(service),
                  SizedBox(height: 15.h),
                  _buildAddDayButton(),
                  SizedBox(height: 20.h),
                ],
              ),
            ),
          );
        },
      ),
    );
  }

  Widget _buildParticipantsSection(TripService service) {
    // 布局参数
    final double avatarSize = 30.w; // 头像大小
    final double avatarGap = 10.w; // 头像间距（负值表示重叠）
    final double iconGap = 5.w; // 图标与头像组的间距

    // 最多显示4个头像，第5个位置显示剩余数量
    final maxVisibleAvatars = 4;
    final displayedParticipants =
        service.participants.take(maxVisibleAvatars).toList();
    final remainingCount = service.participants.length - maxVisibleAvatars;

    // 计算Stack的总宽度：头像宽度 * (显示的头像数 + 剩余数量显示) - 重叠部分
    final stackWidth = avatarSize *
            (displayedParticipants.length + (remainingCount > 0 ? 1 : 0)) -
        (displayedParticipants.length - 1 + (remainingCount > 0 ? 1 : 0)) *
            avatarGap;

    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        // 参与者头像列表
        SizedBox(
          height: avatarSize,
          width: stackWidth,
          child: Stack(
            clipBehavior: Clip.none,
            children: [
              ...List.generate(displayedParticipants.length, (index) {
                final participant = displayedParticipants[index];
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
                      child: Image.network(
                        participant.avatar,
                        width: avatarSize,
                        height: avatarSize,
                        fit: BoxFit.cover,
                      ),
                    ),
                  ),
                );
              }),
              if (remainingCount > 0)
                Positioned(
                  left: displayedParticipants.length * (avatarSize - avatarGap),
                  child: Container(
                    width: avatarSize,
                    height: avatarSize,
                    decoration: BoxDecoration(
                      shape: BoxShape.circle,
                      color: AppColor.buttonGrayBG,
                    ),
                    child: Center(
                      child: Text(
                        '+1$remainingCount',
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

  Widget _buildUnplannedSection(TripService service) {
    return Container(
      padding: EdgeInsets.all(15.w),
      decoration: BoxDecoration(
        color: AppColor.white,
        borderRadius: BorderRadius.circular(AppConfig.boxRadius),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Text(
                '待规划',
                style: TextStyle(
                  fontSize: 18.sp,
                  fontWeight: AppFontWeight.semiBold,
                  color: AppColor.primaryFont,
                ),
              ),
              Text(
                '${service.unplannedSpots.length}个兴趣点位',
                style: TextStyle(
                  fontSize: 16.sp,
                  fontWeight: AppFontWeight.regular,
                  color: AppColor.primaryFont,
                ),
              ),
            ],
          ),
          SizedBox(height: 10.h),
          Divider(color: AppColor.bg),
          SizedBox(height: 10.h),
          Wrap(
            spacing: 20.w,
            runSpacing: 15.h,
            children: service.unplannedSpots
                .map((spot) => Text(
                      spot.name,
                      style: TextStyle(
                        fontSize: 16.sp,
                        fontWeight: AppFontWeight.regular,
                        color: AppColor.primaryFont,
                      ),
                    ))
                .toList(),
          ),
        ],
      ),
    );
  }

  Widget _buildDaySection({
    required String day,
    required String date,
    required List<String> cities,
    List<String>? spots,
  }) {
    return Container(
      padding: EdgeInsets.all(15.w),
      decoration: BoxDecoration(
        color: AppColor.white,
        borderRadius: BorderRadius.circular(AppConfig.boxRadius),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Text(
                day,
                style: TextStyle(
                  fontSize: 18.sp,
                  fontWeight: AppFontWeight.semiBold,
                  color: AppColor.primaryFont,
                ),
              ),
              Text(
                date,
                style: TextStyle(
                  fontSize: 14.sp,
                  fontWeight: AppFontWeight.regular,
                  color: AppColor.primaryFont,
                ),
              ),
            ],
          ),
          SizedBox(height: 10.h),
          Divider(color: AppColor.bg),
          SizedBox(height: 10.h),
          Wrap(
            runSpacing: 5.h,
            children: cities.asMap().entries.map((entry) {
              final isLast = entry.key == cities.length - 1;
              return Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text(
                    entry.value,
                    style: TextStyle(
                      fontSize: 18.sp,
                      fontWeight: AppFontWeight.medium,
                      color: AppColor.primaryFont,
                    ),
                  ),
                  if (!isLast)
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 8.w),
                      child: Icon(
                        Icons.arrow_forward_ios,
                        size: 15.w,
                        color: AppColor.primaryFont,
                      ),
                    ),
                ],
              );
            }).toList(),
          ),
          if (spots != null) ...[
            SizedBox(height: 10.h),
            Divider(color: AppColor.bg),
            SizedBox(height: 10.h),
            Wrap(
              runSpacing: 5.h,
              children: spots.asMap().entries.map((entry) {
                final isLast = entry.key == spots.length - 1;
                return Row(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Text(
                      entry.value,
                      style: TextStyle(
                        fontSize: 16.sp,
                        fontWeight: AppFontWeight.regular,
                        color: AppColor.primaryFont,
                      ),
                    ),
                    if (!isLast)
                      Padding(
                        padding: EdgeInsets.symmetric(horizontal: 8.w),
                        child: Icon(
                          Icons.arrow_forward_ios,
                          size: 12.w,
                          color: AppColor.primaryFont,
                        ),
                      ),
                  ],
                );
              }).toList(),
            ),
          ],
        ],
      ),
    );
  }

  Widget _buildAddDayButton() {
    return GestureDetector(
      onTap: () {},
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Icon(
            Icons.add_circle,
            size: 20.w,
            color: AppColor.primary.withOpacity(0.8),
          ),
          SizedBox(width: 10.w),
          Text(
            '添加一天',
            style: TextStyle(
              fontSize: 16.sp,
              fontWeight: AppFontWeight.regular,
              color: AppColor.primaryFont,
            ),
          ),
        ],
      ),
    );
  }
}
