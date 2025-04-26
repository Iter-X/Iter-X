import 'package:client/app/constants.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:client/common/utils/date_util.dart';
import 'package:client/common/utils/logger.dart';
import 'package:client/common/utils/toast.dart';
import 'package:client/common/widgets/text_divider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:provider/provider.dart';

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
  final int day;
  final DailyItinerary itinerary;
  final VoidCallback? onDragStarted;
  final VoidCallback? onDragEnded;
  final Function(DailyItinerary)? onItemDropped;

  const POIItem({
    required this.key,
    required this.poi,
    required this.isLast,
    required this.tripId,
    required this.day,
    required this.itinerary,
    this.onDragStarted,
    this.onDragEnded,
    this.onItemDropped,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return DragTarget<DailyItinerary>(
      onWillAccept: (data) => data != null && data.id != itinerary.id,
      onAccept: (data) {
        if (onItemDropped != null) {
          onItemDropped!(data);
        }
      },
      builder: (context, candidateData, rejectedData) {
        return LongPressDraggable<DailyItinerary>(
          data: itinerary,
          feedback: Container(
            width: 260.w,
            padding: EdgeInsets.all(8.w),
            decoration: BoxDecoration(
              color: AppColor.white.withOpacity(0.95),
              borderRadius: BorderRadius.circular(AppConfig.imageRadius),
              boxShadow: [
                BoxShadow(
                  color: Colors.black.withOpacity(0.1),
                  blurRadius: 4,
                  offset: const Offset(0, 2),
                ),
              ],
            ),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Day $day',
                  style: TextStyle(
                    fontSize: 12.sp,
                    fontWeight: AppFontWeight.medium,
                    color: AppColor.grayFont,
                  ),
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                ),
                SizedBox(height: 4.h),
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
                  SizedBox(height: 4.h),
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
          onDragStarted: onDragStarted,
          onDragCompleted: onDragEnded,
          child: Container(
            decoration: BoxDecoration(
              color: candidateData.isNotEmpty
                  ? AppColor.primary.withOpacity(0.1)
                  : Colors.transparent,
              borderRadius: BorderRadius.circular(AppConfig.imageRadius),
            ),
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
                        borderRadius:
                            BorderRadius.circular(AppConfig.imageRadius),
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
          ),
        );
      },
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
  final int day;

  const POIList({
    super.key,
    required this.itineraries,
    required this.tripId,
    required this.day,
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
            onAcceptWithDetails: (data) async {
              // Move item before the current item
              final service = Provider.of<TripService>(context, listen: false);
              await service.moveItineraryItem(
                tripId: tripId,
                dailyTripId: data.data.dailyTripId,
                itineraryId: data.data.id,
                newDay: day,
                newIndex: index,
              );
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
                        day: day,
                        itinerary: itinerary,
                        onDragStarted: () {
                          print('Drag started for ${itinerary.poi.nameCn}');
                        },
                        onDragEnded: () {
                          print('Drag ended for ${itinerary.poi.nameCn}');
                        },
                        onItemDropped: (data) {
                          // Handle item dropped
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
                      onAcceptWithDetails: (data) async {
                        // Move item before the current item
                        final service =
                            Provider.of<TripService>(context, listen: false);
                        await service.moveItineraryItem(
                          tripId: tripId,
                          dailyTripId: data.data.dailyTripId,
                          itineraryId: data.data.id,
                          newDay: day,
                          newIndex: index,
                        );
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
                      onAcceptWithDetails: (data) async {
                        // Move item after the current item
                        final service =
                            Provider.of<TripService>(context, listen: false);
                        await service.moveItineraryItem(
                          tripId: tripId,
                          dailyTripId: data.data.dailyTripId,
                          itineraryId: data.data.id,
                          newDay: day,
                          newIndex: index + 1,
                        );
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
  final TripService service;
  final DailyTrip dailyTrip;
  final String tripId;

  const DayContent({
    super.key,
    required this.service,
    required this.dailyTrip,
    required this.tripId,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        SizedBox(height: 10.h),
        if (dailyTrip.dailyItineraries.isNotEmpty) ...[
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
                  day: dailyTrip.day,
                ),
              ),
            ],
          ),
        ],
        Divider(color: AppColor.bg),
        SizedBox(height: 15.h),
        Center(
          child: GestureDetector(
            onTap: () async {
              try {
                final newDailyTrip = await service.addDay(
                  tripId: service.trip!.id,
                  afterDay: dailyTrip.day,
                  notes: '',
                );
                if (newDailyTrip != null) {
                  // 获取从当前天数开始的新列表
                  final updatedDays = await service.fetchDailyTripsFromDay(
                    tripId: service.trip!.id,
                    fromDay: dailyTrip.day,
                  );
                  if (updatedDays.isNotEmpty && service.trip != null) {
                    // 找到当前天数在列表中的位置
                    final index = service.trip!.dailyTrips.indexWhere(
                      (day) => day.day == dailyTrip.day,
                    );
                    if (index != -1) {
                      // 使用service方法更新数据
                      service.updateDailyTripsFromIndex(
                        index: index,
                        updatedDays: updatedDays,
                      );
                    }
                  }
                } else {
                  ToastX.show('添加失败');
                }
              } catch (e) {
                BaseLogger.e('Error adding day: $e');
                ToastX.show('添加失败');
              }
            },
            child: Row(
              mainAxisSize: MainAxisSize.min,
              children: [
                Icon(
                  Icons.add_circle,
                  size: 22.w,
                  color: AppColor.primary.withOpacity(0.8),
                ),
                SizedBox(width: 10.w),
                Text(
                  '添加',
                  style: TextStyle(
                    fontSize: 16.sp,
                    fontWeight: AppFontWeight.regular,
                    color: AppColor.primaryFont,
                  ),
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}

class TripDetailView extends StatelessWidget {
  final String tripId;
  final TripService service;

  const TripDetailView({
    super.key,
    required this.tripId,
    required this.service,
  });

  Widget _buildDaySection(BuildContext context, DailyTrip dailyTrip) {
    return GestureDetector(
      onTap: () {},
      child: DragTarget<DailyItinerary>(
        onWillAccept: (data) => data != null,
        onAccept: (data) async {
          try {
            await service.moveItineraryItem(
              tripId: tripId,
              dailyTripId: data.dailyTripId,
              itineraryId: data.id,
              newDay: dailyTrip.day,
              newIndex: dailyTrip.dailyItineraries.isEmpty
                  ? 0
                  : dailyTrip.dailyItineraries.length,
            );
          } catch (e) {
            print('Error moving itinerary item: $e');
            ToastX.show('Failed to move item');
          }
        },
        builder: (context, candidateData, rejectedData) {
          return Container(
            padding: EdgeInsets.only(
              left: 15.w,
              right: 15.w,
              top: 15.h,
              bottom: 15.h,
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
                  service: service,
                  dailyTrip: dailyTrip,
                  tripId: tripId,
                ),
              ],
            ),
          );
        },
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
            SizedBox(height: 20.h),
          ],
        ),
      ),
    );
  }
}
