import 'dart:ui';

import 'package:client/app/constants.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:client/business/trip/widgets/edit_title_widget.dart';
import 'package:client/common/utils/date_util.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class TripOverviewView extends StatefulWidget {
  final TripService service;

  const TripOverviewView({
    super.key,
    required this.service,
  });

  @override
  State<TripOverviewView> createState() => _TripOverviewViewState();
}

class _TripOverviewViewState extends State<TripOverviewView> {
  bool _isShowEditTitle = false;

  Widget _buildDaySection(BuildContext context, DailyTrip dailyTrip) {
    return GestureDetector(
      onTap: () {},
      child: Container(
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
                children:
                    dailyTrip.dailyItineraries.asMap().entries.map((entry) {
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

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        SingleChildScrollView(
          child: Padding(
            padding: EdgeInsets.only(left: 20.w, right: 20.w),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                ...widget.service.trip!.dailyTrips.map((dailyTrip) {
                  return Column(
                    children: [
                      _buildDaySection(context, dailyTrip),
                      SizedBox(height: 15.h),
                    ],
                  );
                }),
                _buildAddDayButton(),
                SizedBox(height: 20.h),
              ],
            ),
          ),
        ),
        if (_isShowEditTitle) ...[
          Positioned.fill(
            child: GestureDetector(
              onTap: () {
                setState(() {
                  _isShowEditTitle = false;
                });
              },
              child: ClipRect(
                child: BackdropFilter(
                  filter: ImageFilter.blur(sigmaX: 1, sigmaY: 1),
                  child: Container(
                    color: AppColor.primary.withOpacity(0.5),
                  ),
                ),
              ),
            ),
          ),
          Positioned(
            left: 0,
            right: 0,
            bottom: 0,
            child: Column(
              children: [
                EditTitleWidget(
                  initialTitle: widget.service.trip?.title ?? '',
                  onSave: (newTitle) async {
                    if (widget.service.trip != null) {
                      await widget.service.updateTrip(
                        tripId: widget.service.trip!.id,
                        title: newTitle,
                      );
                      setState(() {
                        _isShowEditTitle = false;
                      });
                    }
                  },
                ),
                Container(
                  height: MediaQuery.of(context).padding.bottom,
                  color: AppColor.bottomBar,
                ),
              ],
            ),
          ),
        ],
      ],
    );
  }
}
