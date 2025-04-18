import 'package:client/app/constants.dart';
import 'package:client/business/trip/widgets/select_days_widget.dart';
import 'package:client/common/utils/date_util.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class EditTitleWidget extends StatefulWidget {
  final String initialTitle;
  final String initialDescription;
  final DateTime? initialStartTs;
  final DateTime? initialEndTs;
  final int initialDuration;
  final Function(String, String, DateTime?, DateTime?, int) onSave;
  static final ValueNotifier<bool> isExpandedNotifier =
      ValueNotifier<bool>(false);

  const EditTitleWidget({
    super.key,
    required this.initialTitle,
    required this.initialDescription,
    required this.initialStartTs,
    required this.initialEndTs,
    required this.initialDuration,
    required this.onSave,
  });

  @override
  State<EditTitleWidget> createState() => _EditTitleWidgetState();
}

class _EditTitleWidgetState extends State<EditTitleWidget> {
  late TextEditingController _titleController;
  late TextEditingController _descriptionController;
  late DateTime? _startTs;
  late DateTime? _endTs;
  late int _duration;
  bool _showSelectDays = false;
  static const int _maxTitleLength = 50;
  static const int _maxDescriptionLength = 250;

  @override
  void initState() {
    super.initState();
    _titleController = TextEditingController(text: widget.initialTitle);
    _descriptionController =
        TextEditingController(text: widget.initialDescription);
    _startTs = widget.initialStartTs;
    _endTs = widget.initialEndTs;
    _duration = widget.initialDuration;
    EditTitleWidget.isExpandedNotifier.value = true;
  }

  @override
  void dispose() {
    EditTitleWidget.isExpandedNotifier.value = false;
    _titleController.dispose();
    _descriptionController.dispose();
    super.dispose();
  }

  Future<void> _selectDate(BuildContext context) async {
    final DateTime firstDate = DateTime(2000);
    final DateTime lastDate = DateTime(2100);
    final DateTime now = DateTime.now();

    // 处理1970/01/01的情况，使用当前日期
    DateTime startDate = _startTs ?? now;
    if (startDate.year <= 1970 || !startDate.isAfter(firstDate)) {
      startDate = now;
    }

    DateTime endDate = _endTs ?? startDate.add(const Duration(days: 1));
    if (endDate.year <= 1970 || !endDate.isBefore(lastDate)) {
      endDate = startDate.add(const Duration(days: 1));
    }

    // 确保开始日期不晚于结束日期
    if (startDate.isAfter(endDate)) {
      endDate = startDate.add(const Duration(days: 1));
    }

    final DateTimeRange? picked = await showDateRangePicker(
      context: context,
      firstDate: firstDate,
      lastDate: lastDate,
      initialDateRange: DateTimeRange(start: startDate, end: endDate),
    );
    if (picked != null) {
      setState(() {
        _startTs = picked.start;
        _endTs = picked.end;
        _duration = picked.end.difference(picked.start).inDays + 1;
      });
    }
  }

  void _showSelectDaysWidget() {
    setState(() {
      _showSelectDays = true;
    });
  }

  void _onDaysSelected(int days) {
    setState(() {
      _duration = days;
      if (_startTs != null) {
        _endTs = _startTs!.add(Duration(days: days - 1));
      }
      _showSelectDays = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Container(
          width: double.infinity,
          height: 400.h,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.only(
              topLeft: Radius.circular(AppConfig.boxRadius),
              topRight: Radius.circular(AppConfig.boxRadius),
            ),
            color: AppColor.bottomBar,
          ),
          child: Column(
            children: [
              Container(
                padding: EdgeInsets.symmetric(horizontal: 20.w, vertical: 15.h),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      '编辑行程',
                      style: TextStyle(
                        fontSize: 18.sp,
                        fontWeight: AppFontWeight.medium,
                        color: AppColor.primaryFont,
                      ),
                    ),
                    GestureDetector(
                      onTap: () {
                        widget.onSave(
                          _titleController.text,
                          _descriptionController.text,
                          _startTs,
                          _endTs,
                          _duration,
                        );
                        setState(() {
                          _showSelectDays = false;
                        });
                      },
                      child: Text(
                        '保存',
                        style: TextStyle(
                          fontSize: 16.sp,
                          fontWeight: AppFontWeight.medium,
                          color: AppColor.primary,
                        ),
                      ),
                    ),
                  ],
                ),
              ),
              Divider(color: AppColor.bottomBarLine, height: 1.h),
              Expanded(
                child: SingleChildScrollView(
                  padding: EdgeInsets.all(20.w),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        '标题',
                        style: TextStyle(
                          fontSize: 14.sp,
                          color: AppColor.grayFont,
                        ),
                      ),
                      SizedBox(height: 8.h),
                      TextField(
                        controller: _titleController,
                        maxLength: _maxTitleLength,
                        style: TextStyle(
                          fontSize: 16.sp,
                          color: AppColor.primaryFont,
                        ),
                        decoration: InputDecoration(
                          hintText: '请输入标题',
                          hintStyle: TextStyle(
                            fontSize: 16.sp,
                            color: AppColor.grayFont,
                          ),
                          counterText:
                              '${_titleController.text.length}/$_maxTitleLength',
                          counterStyle: TextStyle(
                            fontSize: 12.sp,
                            color:
                                _titleController.text.length > _maxTitleLength
                                    ? AppColor.primary
                                    : AppColor.grayFont,
                          ),
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.w),
                            borderSide: BorderSide(
                              color: AppColor.borderLine,
                              width: 1.w,
                            ),
                          ),
                          enabledBorder: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.w),
                            borderSide: BorderSide(
                              color: AppColor.borderLine,
                              width: 1.w,
                            ),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.w),
                            borderSide: BorderSide(
                              color: AppColor.primary,
                              width: 1.w,
                            ),
                          ),
                        ),
                      ),
                      SizedBox(height: 20.h),
                      Text(
                        '描述',
                        style: TextStyle(
                          fontSize: 14.sp,
                          color: AppColor.grayFont,
                        ),
                      ),
                      SizedBox(height: 8.h),
                      TextField(
                        controller: _descriptionController,
                        maxLines: 3,
                        maxLength: _maxDescriptionLength,
                        style: TextStyle(
                          fontSize: 16.sp,
                          color: AppColor.primaryFont,
                        ),
                        decoration: InputDecoration(
                          hintText: '请输入描述',
                          hintStyle: TextStyle(
                            fontSize: 16.sp,
                            color: AppColor.grayFont,
                          ),
                          counterText:
                              '${_descriptionController.text.length}/$_maxDescriptionLength',
                          counterStyle: TextStyle(
                            fontSize: 12.sp,
                            color: _descriptionController.text.length >
                                    _maxDescriptionLength
                                ? AppColor.primary
                                : AppColor.grayFont,
                          ),
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.w),
                            borderSide: BorderSide(
                              color: AppColor.borderLine,
                              width: 1.w,
                            ),
                          ),
                          enabledBorder: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.w),
                            borderSide: BorderSide(
                              color: AppColor.borderLine,
                              width: 1.w,
                            ),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.w),
                            borderSide: BorderSide(
                              color: AppColor.primary,
                              width: 1.w,
                            ),
                          ),
                        ),
                      ),
                      SizedBox(height: 20.h),
                      Text(
                        '时间范围',
                        style: TextStyle(
                          fontSize: 14.sp,
                          color: AppColor.grayFont,
                        ),
                      ),
                      SizedBox(height: 8.h),
                      GestureDetector(
                        onTap: () => _selectDate(context),
                        child: Container(
                          padding: EdgeInsets.symmetric(
                              horizontal: 15.w, vertical: 12.h),
                          decoration: BoxDecoration(
                            border: Border.all(
                              color: AppColor.borderLine,
                              width: 1.w,
                            ),
                            borderRadius: BorderRadius.circular(8.w),
                          ),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                _startTs != null && _endTs != null
                                    ? '${DateUtil.formatDate(_startTs!)} 至 ${DateUtil.formatDate(_endTs!)}'
                                    : '请选择时间范围',
                                style: TextStyle(
                                  fontSize: 16.sp,
                                  color: _startTs != null && _endTs != null
                                      ? AppColor.primaryFont
                                      : AppColor.grayFont,
                                ),
                              ),
                              Icon(
                                Icons.calendar_today,
                                size: 20.w,
                                color: AppColor.primaryFont,
                              ),
                            ],
                          ),
                        ),
                      ),
                      SizedBox(height: 20.h),
                      Text(
                        '天数',
                        style: TextStyle(
                          fontSize: 14.sp,
                          color: AppColor.grayFont,
                        ),
                      ),
                      SizedBox(height: 8.h),
                      GestureDetector(
                        onTap: _showSelectDaysWidget,
                        child: Container(
                          padding: EdgeInsets.symmetric(
                              horizontal: 15.w, vertical: 12.h),
                          decoration: BoxDecoration(
                            border: Border.all(
                              color: AppColor.borderLine,
                              width: 1.w,
                            ),
                            borderRadius: BorderRadius.circular(8.w),
                          ),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                '$_duration天',
                                style: TextStyle(
                                  fontSize: 16.sp,
                                  color: AppColor.primaryFont,
                                ),
                              ),
                              Icon(
                                Icons.arrow_forward_ios,
                                size: 16.w,
                                color: AppColor.primaryFont,
                              ),
                            ],
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
        if (_showSelectDays)
          Positioned(
            left: 0,
            right: 0,
            bottom: 0,
            child: SelectDaysWidget(
              selectDays: _duration,
              onTap: _onDaysSelected,
            ),
          ),
      ],
    );
  }
}
