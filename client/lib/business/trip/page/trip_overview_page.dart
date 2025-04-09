import 'package:client/app/constants.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/utils/date_util.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:provider/provider.dart';

class TripOverviewPage extends StatefulWidget {
  final String tripId;

  const TripOverviewPage({
    super.key,
    required this.tripId,
  });

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
      service.fetchTripData(tripId: widget.tripId);
    });
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      hasAppBar: true,
      bottom: false,
      leading: ReturnButton(),
      child: Consumer<TripService>(
        builder: (context, service, child) {
          if (service.isLoading) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          }

          final trip = service.trip;
          if (trip == null) {
            return const Center(
              child: Text('暂无行程数据'),
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
                      trip.title,
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
                        '${DateUtil.formatDate(trip.startTs)} | ${trip.dailyTrips.length}天',
                        style: TextStyle(
                          fontSize: 20.sp,
                          fontWeight: AppFontWeight.regular,
                          color: AppColor.primaryFont,
                        ),
                      ),
                    ],
                  ),
                  SizedBox(height: 10.h),
                  Text(
                    trip.description,
                    style: TextStyle(
                      fontSize: 16.sp,
                      fontWeight: AppFontWeight.regular,
                      color: AppColor.secondaryFont,
                    ),
                  ),
                  SizedBox(height: 30.h),
                  ...trip.dailyTrips.map((dailyTrip) {
                    return Column(
                      children: [
                        _buildDaySection(dailyTrip),
                        SizedBox(height: 15.h),
                      ],
                    );
                  }),
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

  Widget _buildDaySection(DailyTrip dailyTrip) {
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
                'Day ${dailyTrip.day}',
                style: TextStyle(
                  fontSize: 18.sp,
                  fontWeight: AppFontWeight.semiBold,
                  color: AppColor.primaryFont,
                ),
              ),
              Text(
                DateUtil.formatDate(dailyTrip.date),
                style: TextStyle(
                  fontSize: 14.sp,
                  fontWeight: AppFontWeight.regular,
                  color: AppColor.primaryFont,
                ),
              ),
            ],
          ),
          SizedBox(height: 10.h),
          Text(
            dailyTrip.notes,
            style: TextStyle(
              fontSize: 16.sp,
              fontWeight: AppFontWeight.regular,
              color: AppColor.primaryFont,
            ),
          ),
          if (dailyTrip.dailyItineraries.isNotEmpty) ...[
            SizedBox(height: 10.h),
            Divider(color: AppColor.bg),
            SizedBox(height: 10.h),
            Wrap(
              runSpacing: 5.h,
              children: dailyTrip.dailyItineraries.asMap().entries.map((entry) {
                final isLast =
                    entry.key == dailyTrip.dailyItineraries.length - 1;
                return Row(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Text(
                      entry.value.poi.nameCn,
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
