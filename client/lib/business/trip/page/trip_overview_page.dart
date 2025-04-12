import 'package:client/app/constants.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:client/business/trip/widgets/trip_collaborators_section.dart';
import 'package:client/business/trip/widgets/trip_day_section.dart';
import 'package:client/business/trip/widgets/trip_settings_drawer.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/utils/date_util.dart';
import 'package:client/common/widgets/preference_button.dart';
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
  bool _isDetailView = false;

  @override
  void initState() {
    super.initState();
    // Initialize data
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final service = context.read<TripService>();
      service.fetchTripData(tripId: widget.tripId);
    });
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

  void _showSettingsMenu(BuildContext context) {
    showGeneralDialog(
      context: context,
      barrierDismissible: true,
      barrierLabel: '',
      barrierColor: Colors.black.withOpacity(0.5),
      transitionDuration: const Duration(milliseconds: 300),
      pageBuilder: (context, animation, secondaryAnimation) {
        return Container();
      },
      transitionBuilder: (context, animation, secondaryAnimation, child) {
        return SlideTransition(
          position: Tween<Offset>(
            begin: const Offset(1.0, 0.0),
            end: Offset.zero,
          ).animate(CurvedAnimation(
            parent: animation,
            curve: Curves.easeOutCubic,
          )),
          child: TripSettingsDrawer(
            isDetailView: _isDetailView,
            onViewModeChanged: (value) {
              setState(() {
                _isDetailView = value;
              });
            },
          ),
        );
      },
    );
  }

  Widget _buildHeader(TripService service) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 20.w),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(height: 20.h),
          SizedBox(
            width: double.infinity,
            child: Text(
              service.trip!.title,
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
                '${DateUtil.formatDate(service.trip!.startTs)} | ${service.trip!.dailyTrips.length}天',
                style: TextStyle(
                  fontSize: 20.sp,
                  fontWeight: AppFontWeight.regular,
                  color: AppColor.primaryFont,
                ),
              ),
              TripCollaboratorsSection(service: service),
            ],
          ),
          SizedBox(height: 20.h),
        ],
      ),
    );
  }

  Widget _buildDetailView(TripService service) {
    return SingleChildScrollView(
      child: Padding(
        padding: EdgeInsets.only(left: 20.w, right: 20.w, top: 60.h),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            ...service.trip!.dailyTrips.map((dailyTrip) {
              return Column(
                children: [
                  TripDaySection(dailyTrip: dailyTrip),
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

  Widget _buildOverviewView(TripService service) {
    return SingleChildScrollView(
      child: Padding(
        padding: EdgeInsets.only(left: 20.w, right: 20.w, top: 60.h),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            ...service.trip!.dailyTrips.map((dailyTrip) {
              return Column(
                children: [
                  TripDaySection(dailyTrip: dailyTrip),
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

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      hasAppBar: true,
      bottom: false,
      elevation: 0,
      leading: ReturnButton(),
      actions: [
        PreferenceButton(
          onTap: () => _showSettingsMenu(context),
        ),
      ],
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

          return Column(
            children: [
              _buildHeader(service),
              Expanded(
                child: _isDetailView
                    ? _buildDetailView(service)
                    : _buildOverviewView(service),
              ),
            ],
          );
        },
      ),
    );
  }
}
