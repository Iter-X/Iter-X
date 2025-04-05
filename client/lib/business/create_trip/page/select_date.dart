import 'dart:ui';

import 'package:client/app/constants.dart';
import 'package:client/business/create_trip/widgets/select_days_widget.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/date_time_util.dart';
import 'package:client/common/utils/toast.dart';
import 'package:client/common/widgets/base_button.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:intl/intl.dart';
import 'package:lunar/calendar/Lunar.dart';
import 'package:scrollable_clean_calendar/controllers/clean_calendar_controller.dart';
import 'package:scrollable_clean_calendar/scrollable_clean_calendar.dart';
import 'package:scrollable_clean_calendar/utils/enums.dart';

class SelectDatePage extends StatefulWidget {
  const SelectDatePage({super.key});

  @override
  State<SelectDatePage> createState() => _SelectDatePageState();
}

class _SelectDatePageState extends BaseState<SelectDatePage> {
  String startTime = ''; // 开始时间
  DateTime? _startTime; // 开始时间
  DateTime? _endTime; // 结束时间
  int? selectDays; // 灵活天数
  bool isShowSelectDays = false;
  int selectRangeDays = 0;
  final GlobalKey _bottomBarKey = GlobalKey();

  late CleanCalendarController _calendarController;

  @override
  void initState() {
    _calendarController = CleanCalendarController(
      minDate: DateTime(1900, 1, 1),
      maxDate: DateTime(2200, 12, 31),
      initialFocusDate: DateTime.now(),
      weekdayStart: 7,
      onRangeSelected: (startDate, endDate) {
        selectDays = null;
        startTime = DateFormat('MM-dd').format(startDate);
        selectRangeDays = 0;
        if (endDate != null) {
          if (DateTimeUtil.isSameDay(startDate, endDate)) {
            selectRangeDays = 1;
          } else {
            selectRangeDays = endDate.difference(startDate).inDays + 1;
          }
        }
        _startTime = startDate;
        _endTime = endDate;
        setState(() {});
      },
    );
    super.initState();
  }

  Widget _buildCalendar() {
    return ScrollableCleanCalendar(
      calendarController: _calendarController,
      layout: Layout.BEAUTY,
      calendarCrossAxisSpacing: 0,
      locale: 'zh',
      weekdayBuilder: (context, day) {
        return Center(
          child: Text(
            day.substring(1, 2),
            style: TextStyle(
              color: AppColor.primaryFont,
              fontSize: 18.sp,
            ),
          ),
        );
      },
      monthBuilder: (context, month) {
        String yearStr = month.substring(month.length - 4, month.length);
        String monthStr = month.substring(0, month.length - 5);
        return Container(
          margin: EdgeInsets.only(left: 10.w),
          child: Text(
            '$yearStr-${getChangeMonth(monthStr)}',
            style: TextStyle(
              color: AppColor.primaryFont,
              fontSize: 30.sp,
              fontWeight: AppFontWeight.bold,
            ),
          ),
        );
      },
      dayBuilder: (context, day) {
        var time = Lunar.fromDate(day.day);
        String chineseDay;
        if (time.getDayInChinese() == '初一') {
          chineseDay = '${time.getMonthInChinese()}月';
        } else {
          chineseDay = time.getDayInChinese();
        }
        return Container(
          decoration: getBoxDecoration(
              day.selectedMinDate, day.selectedMaxDate, day.day),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                day.text,
                style: TextStyle(
                  color: isWhiteText(
                          day.selectedMinDate, day.selectedMaxDate, day.day)
                      ? AppColor.secondaryFont
                      : AppColor.primaryFont,
                  fontSize: 18.sp,
                ),
              ),
              Text(
                chineseDay,
                style: TextStyle(
                  color: isWhiteText(
                          day.selectedMinDate, day.selectedMaxDate, day.day)
                      ? AppColor.secondaryFont
                      : AppColor.primaryFont,
                  fontSize: 13.sp,
                ),
              )
            ],
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        AppBarWithSafeArea(
          backgroundColor: AppColor.bg,
          bottomColor: AppColor.bottomBar,
          hasAppBar: true,
          leading: ReturnButton(),
          title: "选择出行时间",
          child: Column(
            children: [
              Expanded(
                child: _buildCalendar(),
              ),
              _buildBottomContent(),
            ],
          ),
        ),
        if (isShowSelectDays)
          Positioned(
            top: 0,
            left: 0,
            right: 0,
            bottom: 0,
            child: GestureDetector(
              onTap: () {
                setState(() {
                  isShowSelectDays = false;
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
        if (isShowSelectDays)
          Column(
            children: [
              const Spacer(),
              SelectDaysWidget(
                onTap: (days) {
                  _calendarController.clearSelectedDates();
                  selectDays = days;
                  isShowSelectDays = false;
                  setState(() {});
                },
                selectDays: selectDays,
              ),
              _buildBottomContent(),
              Container(
                height: MediaQuery.of(context).padding.bottom,
                color: AppColor.bottomBar,
              )
            ],
          ),
      ],
    );
  }

  Widget _buildBottomContent() {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 30.h),
      decoration: BoxDecoration(
        color: AppColor.bottomBar,
        border: Border(
          top: BorderSide(
            color: AppColor.bottomBarLine,
            width: 1.w,
          ),
        ),
      ),
      child: Column(
        children: [
          Container(
            margin: EdgeInsets.only(
              left: 30.w,
              right: 30.w,
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Row(
                  children: [
                    if (selectDays != null || _startTime != null)
                      Row(
                        children: [
                          if (selectDays == null)
                            RichText(
                              text: TextSpan(
                                children: [
                                  TextSpan(
                                    text: startTime,
                                    style: TextStyle(
                                      fontSize: 18.sp,
                                      color: AppColor.primaryFont,
                                      fontWeight: AppFontWeight.bold,
                                    ),
                                  ),
                                  TextSpan(
                                    text: '出发 ',
                                    style: TextStyle(
                                      fontSize: 18.sp,
                                      color: AppColor.primaryFont,
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          RichText(
                            text: TextSpan(
                              children: [
                                TextSpan(
                                  text: '出行',
                                  style: TextStyle(
                                    fontSize: 18.sp,
                                    color: AppColor.primaryFont,
                                  ),
                                ),
                                TextSpan(
                                  text: '${selectDays ?? selectRangeDays}',
                                  style: TextStyle(
                                    fontSize: 18.sp,
                                    color: AppColor.primaryFont,
                                    fontWeight: AppFontWeight.bold,
                                  ),
                                ),
                                TextSpan(
                                  text: '天',
                                  style: TextStyle(
                                    fontSize: 18.sp,
                                    color: AppColor.primaryFont,
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ],
                      )
                    else
                      Text(
                        '选择出行时间',
                        style: TextStyle(
                          fontSize: 18.sp,
                          color: AppColor.primaryFont,
                        ),
                      ),
                  ],
                ),
                GestureDetector(
                  onTap: () {
                    setState(() {
                      isShowSelectDays = !isShowSelectDays;
                    });
                  },
                  child: Row(
                    children: [
                      Text(
                        '灵活时间',
                        style: TextStyle(
                          color: AppColor.primaryFont,
                          fontSize: 16.sp,
                        ),
                      ),
                      Icon(
                        isShowSelectDays
                            ? Icons.keyboard_arrow_up
                            : Icons.keyboard_arrow_down,
                        color: AppColor.primaryFont,
                        size: 24.sp,
                      ),
                    ],
                  ),
                )
              ],
            ),
          ),
          Container(
            margin: EdgeInsets.only(
              top: 30.h,
              left: 15.w,
              right: 15.w,
            ),
            child: BaseButton(
              text: '下一步',
              textColor: AppColor.white,
              onTap: next,
            ),
          ),
        ],
      ),
    );
  }

  String getChangeMonth(String month) {
    switch (month) {
      case '一月':
        return '01';
      case '二月':
        return '02';
      case '三月':
        return '03';
      case '四月':
        return '04';
      case '五月':
        return '05';
      case '六月':
        return '06';
      case '七月':
        return '07';
      case '八月':
        return '08';
      case '九月':
        return '09';
      case '十月':
        return '10';
      case '十一月':
        return '11';
      case '十二月':
        return '12';
      default:
        return '';
    }
  }

  // 获取日期的背景样式
  BoxDecoration? getBoxDecoration(
      DateTime? day1, DateTime? day2, DateTime day3) {
    if (day1 == null && day2 == null) {
      if (DateTimeUtil.isSameDay(day3, DateTime.now())) {
        return BoxDecoration(
          color: AppColor.primary,
          shape: BoxShape.circle,
        );
      }
      return null;
    }
    // 选中的日期范围是同一天
    if (DateTimeUtil.isSameDay(day1, day3) &&
        DateTimeUtil.isSameDay(day2, day3)) {
      return BoxDecoration(
        color: AppColor.highlight,
        shape: BoxShape.circle,
      );
    }
    // 开始时间是当前选中的日期
    if (DateTimeUtil.isSameDay(day1, day3)) {
      return BoxDecoration(
        color: AppColor.highlight,
        borderRadius: BorderRadius.only(
          topLeft: Radius.circular(56.w),
          bottomLeft: Radius.circular(56.w),
        ),
      );
    }
    // 结束时间是当前选中的日期
    if (DateTimeUtil.isSameDay(day2, day3)) {
      return BoxDecoration(
        color: AppColor.highlight,
        borderRadius: BorderRadius.only(
          topRight: Radius.circular(56.w),
          bottomRight: Radius.circular(56.w),
        ),
      );
    }
    // 日期范围是选中的日期
    if (day1 != null &&
        day2 != null &&
        day3.isAfter(day1) &&
        day3.isBefore(day2)) {
      return BoxDecoration(
        color: AppColor.highlight,
      );
    }
    if (!DateTimeUtil.isSameDay(day1, day3) &&
        !DateTimeUtil.isSameDay(day2, day3) &&
        DateTimeUtil.isSameDay(day3, DateTime.now())) {
      return BoxDecoration(
        color: AppColor.primary,
        shape: BoxShape.circle,
      );
    }
    return null;
  }

  bool isWhiteText(DateTime? day1, DateTime? day2, DateTime day3) {
    return DateTimeUtil.isSameDay(day3, DateTime.now()) ||
        DateTimeUtil.isSameDay(day1, day3) ||
        DateTimeUtil.isSameDay(day2, day3) ||
        (day1 != null &&
            day2 != null &&
            day3.isAfter(day1) &&
            day3.isBefore(day2));
  }

  void next() {
    if (selectDays == null) {
      if (_startTime == null) {
        ToastX.show('请选择出行时间');
        return;
      }
      if (_endTime == null) {
        ToastX.show('请选择返程时间');
        return;
      }
      ToastX.show("请选择出行时间");
    }
  }
}
