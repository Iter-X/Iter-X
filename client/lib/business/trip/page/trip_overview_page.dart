import 'dart:ui';

import 'package:client/app/constants.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:client/business/trip/widgets/edit_title_widget.dart';
import 'package:client/business/trip/widgets/trip_collaborators_section.dart';
import 'package:client/business/trip/widgets/trip_detail_view.dart';
import 'package:client/business/trip/widgets/trip_overview_view.dart';
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
  bool _isDetailView = true;
  bool _isShowEditTitle = false;

  @override
  void initState() {
    super.initState();
    // Initialize data
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final service = context.read<TripService>();
      service.fetchTripData(tripId: widget.tripId);
    });
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
          GestureDetector(
            onTap: () {
              setState(() {
                _isShowEditTitle = true;
              });
            },
            child: SizedBox(
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

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        AppBarWithSafeArea(
          backgroundColor: AppColor.bg,
          hasAppBar: true,
          bottom: false,
          surfaceTintColor: AppColor.bg,
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
                        ? TripDetailView(
                            tripId: widget.tripId,
                            service: service,
                          )
                        : TripOverviewView(service: service),
                  ),
                ],
              );
            },
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
                  initialTitle: context.read<TripService>().trip?.title ?? '',
                  initialDescription: context.read<TripService>().trip?.description ?? '',
                  initialStartTs: context.read<TripService>().trip?.startTs,
                  initialEndTs: context.read<TripService>().trip?.endTs,
                  initialDuration: context.read<TripService>().trip?.dailyTrips.length ?? 1,
                  onSave: (newTitle, newDescription, newStartTs, newEndTs, newDuration) async {
                    final service = context.read<TripService>();
                    if (service.trip != null) {
                      await service.updateTrip(
                        tripId: service.trip!.id,
                        title: newTitle,
                        description: newDescription,
                        startTs: newStartTs,
                        endTs: newEndTs,
                        duration: newDuration,
                      );
                    }
                    setState(() {
                      _isShowEditTitle = false;
                    });
                  },
                ),
              ],
            ),
          ),
        ],
      ],
    );
  }
}
