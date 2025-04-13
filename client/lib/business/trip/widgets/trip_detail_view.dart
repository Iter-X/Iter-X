import 'package:client/app/constants.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:client/common/utils/date_util.dart';
import 'package:client/common/widgets/text_divider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class Timeline extends StatelessWidget {
  final IconData icon;
  final bool isFirst;
  final bool isLast;

  const Timeline({
    super.key,
    required this.icon,
    this.isFirst = false,
    this.isLast = false,
  });

  @override
  Widget build(BuildContext context) {
    final double iconHeight = 32.w;
    final double iconWidth = 32.w;
    // Calculate height based on POIItem structure
    // POIItem height = top padding (15.h) + image height (56.w) + bottom padding (15.h)
    final double itemHeight = 15.h + 56.w + 15.h;
    final double textDividerHeight = 16.h;
    // Add TextDivider height if not last item
    final double totalHeight =
        isLast ? itemHeight : itemHeight + textDividerHeight / 2;

    return SizedBox(
      width: iconWidth,
      height: totalHeight,
      child: Stack(
        alignment: Alignment.center,
        children: [
          if (!isFirst)
            Positioned(
              top: 0,
              bottom: (itemHeight - iconHeight) / 2,
              child: Container(
                width: 2.w,
                color: AppColor.bg,
              ),
            ),
          if (!isLast)
            Positioned(
              top: iconHeight + (itemHeight - iconHeight) / 2,
              bottom: 0,
              child: Container(
                width: 2.w,
                color: AppColor.bg,
              ),
            ),
          // Icon container
          Positioned(
            top: (itemHeight - iconHeight) / 2,
            child: Container(
              width: iconWidth,
              height: iconHeight,
              decoration: BoxDecoration(
                color: AppColor.bg,
                shape: BoxShape.circle,
              ),
              child: Icon(
                icon,
                size: 26.w,
                color: AppColor.primaryFont,
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class POIItem extends StatelessWidget {
  final POI poi;
  final bool isLast;

  const POIItem({
    super.key,
    required this.poi,
    required this.isLast,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: EdgeInsets.symmetric(vertical: 15.h),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              // POI image
              ClipRRect(
                borderRadius: BorderRadius.circular(AppConfig.imageRadius),
                child: Container(
                  width: 56.w,
                  height: 56.w,
                  color: AppColor.bg,
                  child: Icon(
                    Icons.image_not_supported,
                    size: 24.w,
                    color: AppColor.primaryFont.withOpacity(.8),
                  ),
                ),
              ),
              SizedBox(width: 10.w),
              // POI name and description
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      poi.nameCn,
                      style: TextStyle(
                        fontSize: 16.sp,
                        fontWeight: AppFontWeight.medium,
                        color: AppColor.primaryFont,
                      ),
                      maxLines: 1,
                      overflow: TextOverflow.ellipsis,
                    ),
                    if (poi.nameEn.isNotEmpty) ...[
                      Text(
                        poi.nameEn,
                        style: TextStyle(
                          fontSize: 14.sp,
                          fontWeight: AppFontWeight.regular,
                          color: AppColor.primaryFont,
                        ),
                        maxLines: 1,
                        overflow: TextOverflow.ellipsis,
                      ),
                    ],
                  ],
                ),
              ),
            ],
          ),
        ),
        if (!isLast) ...[
          TextDivider(
            color: AppColor.bg,
            height: 16.h,
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Icon(
                  Icons.directions_car,
                  size: 16.w,
                  color: AppColor.grayFont,
                ),
                SizedBox(width: 5.w),
                Text(
                  '385km | 3h47m',
                  style: TextStyle(
                    fontSize: 10.sp,
                    fontWeight: AppFontWeight.medium,
                    color: AppColor.grayFont,
                  ),
                ),
              ],
            ),
          ),
        ],
      ],
    );
  }
}

class TripDetailView extends StatelessWidget {
  final TripService service;

  const TripDetailView({
    super.key,
    required this.service,
  });

  Widget _buildDaySection(BuildContext context, DailyTrip dailyTrip) {
    return GestureDetector(
      onTap: () {},
      child: Container(
        padding: EdgeInsets.only(
          left: 15.w,
          right: 15.w,
          top: 15.w,
          bottom: dailyTrip.dailyItineraries.isEmpty ? 15.w : 0,
        ),
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
              Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  // Timeline column
                  Column(
                    children: dailyTrip.dailyItineraries.map((itinerary) {
                      final poi = itinerary.poi;
                      IconData poiIcon;
                      switch (poi.type.toLowerCase()) {
                        case 'city':
                          poiIcon = Icons.location_city;
                          break;
                        case 'airport':
                          poiIcon = Icons.flight;
                          break;
                        default:
                          poiIcon = Icons.landscape;
                      }
                      return Timeline(
                        icon: poiIcon,
                        isFirst:
                            dailyTrip.dailyItineraries.indexOf(itinerary) == 0,
                        isLast: dailyTrip.dailyItineraries.indexOf(itinerary) ==
                            dailyTrip.dailyItineraries.length - 1,
                      );
                    }).toList(),
                  ),
                  SizedBox(width: 20.w),
                  // POI items column
                  Expanded(
                    child: Column(
                      children: dailyTrip.dailyItineraries.map((itinerary) {
                        return POIItem(
                          poi: itinerary.poi,
                          isLast:
                              dailyTrip.dailyItineraries.indexOf(itinerary) ==
                                  dailyTrip.dailyItineraries.length - 1,
                        );
                      }).toList(),
                    ),
                  ),
                ],
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
    return SingleChildScrollView(
      child: Padding(
        padding: EdgeInsets.only(left: 20.w, right: 20.w, top: 60.h),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            ...service.trip!.dailyTrips.map((dailyTrip) {
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
    );
  }
}
