import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class TripSettingsDrawer extends StatefulWidget {
  final bool isDetailView;
  final Function(bool) onViewModeChanged;

  const TripSettingsDrawer({
    super.key,
    required this.isDetailView,
    required this.onViewModeChanged,
  });

  @override
  State<TripSettingsDrawer> createState() => _TripSettingsDrawerState();
}

class _TripSettingsDrawerState extends State<TripSettingsDrawer> {
  bool _isDetailView = false;

  @override
  void initState() {
    super.initState();
    _isDetailView = widget.isDetailView;
  }

  @override
  void didUpdateWidget(TripSettingsDrawer oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.isDetailView != widget.isDetailView) {
      setState(() {
        _isDetailView = widget.isDetailView;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.centerRight,
      child: Container(
        width: MediaQuery.of(context).size.width * 4 / 5,
        height: MediaQuery.of(context).size.height,
        decoration: BoxDecoration(
          color: AppColor.bg,
          borderRadius: BorderRadius.only(
            topLeft: Radius.circular(AppConfig.boxRadius),
            bottomLeft: Radius.circular(AppConfig.boxRadius),
          ),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withOpacity(0.1),
              blurRadius: 10,
              offset: const Offset(0, 0),
            ),
          ],
        ),
        child: Column(
          children: [
            // Top padding for safe area
            SizedBox(height: MediaQuery.of(context).padding.top),

            // Header
            Container(
              padding: EdgeInsets.symmetric(horizontal: 20.w, vertical: 20.h),
              decoration: BoxDecoration(
                border: Border(
                  bottom: BorderSide(
                    color: Colors.grey.withOpacity(0.2),
                    width: 1,
                  ),
                ),
              ),
              child: Row(
                children: [
                  Expanded(
                    child: Text(
                      '设置',
                      style: TextStyle(
                        fontSize: 24.sp,
                        fontWeight: AppFontWeight.bold,
                        color: AppColor.primaryFont,
                      ),
                    ),
                  ),
                  IconButton(
                    icon: Icon(Icons.close, size: 24.w),
                    onPressed: () => Navigator.pop(context),
                  ),
                ],
              ),
            ),

            // Content
            Expanded(
              child: SingleChildScrollView(
                padding: EdgeInsets.symmetric(horizontal: 20.w),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    SizedBox(height: 20.h),
                    _buildSection(
                      title: '视图模式',
                      children: [
                        _buildToggleItem(
                          title: _isDetailView ? '详情模式' : '总览模式',
                          value: _isDetailView,
                          onChanged: (value) {
                            setState(() {
                              _isDetailView = value;
                            });
                            widget.onViewModeChanged(value);
                          },
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildSection({
    required String title,
    required List<Widget> children,
  }) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          title,
          style: TextStyle(
            fontSize: 18.sp,
            fontWeight: AppFontWeight.semiBold,
            color: AppColor.primaryFont,
          ),
        ),
        SizedBox(height: 10.h),
        ...children,
        SizedBox(height: 20.h),
      ],
    );
  }

  Widget _buildToggleItem({
    required String title,
    required bool value,
    required Function(bool) onChanged,
  }) {
    return Container(
      margin: EdgeInsets.only(bottom: 10.h),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(12.r),
      ),
      child: ListTile(
        title: Text(
          title,
          style: TextStyle(
            fontSize: 16.sp,
            fontWeight: AppFontWeight.medium,
            color: AppColor.primaryFont,
          ),
        ),
        trailing: Container(
          width: 40.w,
          height: 24.h,
          decoration: BoxDecoration(
            color: value ? AppColor.primary : Colors.grey[300],
            borderRadius: BorderRadius.circular(12.r),
          ),
          child: AnimatedAlign(
            duration: const Duration(milliseconds: 200),
            alignment: value ? Alignment.centerRight : Alignment.centerLeft,
            child: Container(
              width: 20.w,
              height: 20.h,
              margin: EdgeInsets.all(2.w),
              decoration: BoxDecoration(
                color: Colors.white,
                shape: BoxShape.circle,
              ),
            ),
          ),
        ),
        onTap: () => onChanged(!value),
      ),
    );
  }
}
