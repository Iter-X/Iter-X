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
    // POIItem height = top padding (15.h) + image height (56.h) + bottom padding (15.h)
    final double itemHeight = 15.h + 56.h + 15.h;
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
  final Key key;
  final String tripId;
  final VoidCallback? onDragStarted;
  final VoidCallback? onDragEnded;

  const POIItem({
    required this.key,
    required this.poi,
    required this.isLast,
    required this.tripId,
    this.onDragStarted,
    this.onDragEnded,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return LongPressDraggable<DailyItinerary>(
      data: DailyItinerary(
        id: 'temp_${DateTime.now().millisecondsSinceEpoch}',
        tripId: tripId,
        dailyTripId: 'temp_daily_${DateTime.now().millisecondsSinceEpoch}',
        poiId: poi.id,
        poi: poi,
        notes: '',
        createdAt: DateTime.now(),
        updatedAt: DateTime.now(),
      ),
      feedback: Material(
        elevation: 4,
        child: Container(
          width: 200.w,
          padding: EdgeInsets.all(8.w),
          decoration: BoxDecoration(
            color: AppColor.white,
            borderRadius: BorderRadius.circular(8.r),
          ),
          child: Text(
            poi.nameCn,
            style: TextStyle(
              fontSize: 16.sp,
              fontWeight: AppFontWeight.medium,
              color: AppColor.primaryFont,
            ),
          ),
        ),
      ),
      onDragStarted: onDragStarted,
      onDragCompleted: onDragEnded,
      child: Column(
        key: key,
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
                    width: 56.h,
                    height: 56.h,
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
      ),
    );
  }
}

class DayHeader extends StatelessWidget {
  final DailyTrip dailyTrip;

  const DayHeader({
    super.key,
    required this.dailyTrip,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
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
      ],
    );
  }
}

class DragIndicator extends StatelessWidget {
  final bool isTop;
  final bool isVisible;

  const DragIndicator({
    super.key,
    this.isTop = false,
    required this.isVisible,
  });

  @override
  Widget build(BuildContext context) {
    if (!isVisible) return const SizedBox.shrink();

    return Stack(
      children: [
        Positioned(
          left: 0,
          top: isTop ? 4.h : null,
          bottom: isTop ? null : 4.h,
          child: Container(
            width: 8.w,
            height: 8.w,
            decoration: BoxDecoration(
              color: Colors.transparent,
              shape: BoxShape.circle,
              border: Border.all(
                color: AppColor.primary.withOpacity(.6),
                width: 2.h,
              ),
            ),
          ),
        ),
        Positioned(
          left: 8.w,
          top: isTop ? 7.h : null,
          bottom: isTop ? null : 7.h,
          right: 0,
          child: Container(
            height: 2.h,
            color: AppColor.primary.withOpacity(.6),
          ),
        ),
      ],
    );
  }
}

class POIList extends StatelessWidget {
  final List<DailyItinerary> itineraries;
  final String tripId;

  const POIList({
    super.key,
    required this.itineraries,
    required this.tripId,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        ...itineraries.asMap().entries.map((entry) {
          final index = entry.key;
          final itinerary = entry.value;
          final isLast = index == itineraries.length - 1;

          return DragTarget<DailyItinerary>(
            onWillAcceptWithDetails: (data) => true,
            onAcceptWithDetails: (data) {
              print('Dropped before item ${index + 1}');
            },
            builder: (context, candidateData, rejectedData) {
              return Stack(
                children: [
                  Column(
                    children: [
                      POIItem(
                        key: ValueKey(itinerary.id),
                        poi: itinerary.poi,
                        tripId: tripId,
                        isLast: isLast,
                        onDragStarted: () {
                          print('Drag started for ${itinerary.poi.nameCn}');
                        },
                        onDragEnded: () {
                          print('Drag ended for ${itinerary.poi.nameCn}');
                        },
                      ),
                    ],
                  ),
                  Positioned(
                    bottom: !isLast ? 43.h + 16.h : 43.h,
                    height: 43.h,
                    left: 0,
                    right: 0,
                    child: DragTarget<DailyItinerary>(
                      onWillAcceptWithDetails: (data) => true,
                      onAcceptWithDetails: (data) {
                        print('Dropped before item ${index + 1}');
                      },
                      builder: (context, candidateData, rejectedData) {
                        return Container(
                          child: candidateData.isNotEmpty
                              ? DragIndicator(isTop: true, isVisible: true)
                              : null,
                        );
                      },
                    ),
                  ),
                  Positioned(
                    top: 43.h,
                    height: 43.h,
                    left: 0,
                    right: 0,
                    child: DragTarget<DailyItinerary>(
                      onWillAcceptWithDetails: (data) => true,
                      onAcceptWithDetails: (data) {
                        print('Dropped after item ${index + 1}');
                      },
                      builder: (context, candidateData, rejectedData) {
                        return Container(
                          child: candidateData.isNotEmpty
                              ? DragIndicator(isTop: false, isVisible: true)
                              : null,
                        );
                      },
                    ),
                  ),
                ],
              );
            },
          );
        }),
      ],
    );
  }
}

class DayContent extends StatelessWidget {
  final DailyTrip dailyTrip;
  final String tripId;

  const DayContent({
    super.key,
    required this.dailyTrip,
    required this.tripId,
  });

  @override
  Widget build(BuildContext context) {
    if (dailyTrip.dailyItineraries.isEmpty) {
      return const SizedBox.shrink();
    }

    return Column(
      children: [
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
                  isFirst: dailyTrip.dailyItineraries.indexOf(itinerary) == 0,
                  isLast: dailyTrip.dailyItineraries.indexOf(itinerary) ==
                      dailyTrip.dailyItineraries.length - 1,
                );
              }).toList(),
            ),
            SizedBox(width: 20.w),
            // POI items column
            Expanded(
              child: POIList(
                itineraries: dailyTrip.dailyItineraries,
                tripId: tripId,
              ),
            ),
          ],
        ),
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
      child: DragTarget<DailyItinerary>(
        onWillAcceptWithDetails: (data) => true,
        onAcceptWithDetails: (data) {
          print('Dropped in day ${dailyTrip.day}');
        },
        builder: (context, candidateData, rejectedData) {
          return Container(
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
                DayHeader(dailyTrip: dailyTrip),
                DayContent(
                  dailyTrip: dailyTrip,
                  tripId: service.trip!.id,
                ),
              ],
            ),
          );
        },
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
