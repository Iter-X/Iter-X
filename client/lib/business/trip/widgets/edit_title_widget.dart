import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class EditTitleWidget extends StatefulWidget {
  final String initialTitle;
  final Function(String) onSave;
  static final ValueNotifier<bool> isExpandedNotifier =
      ValueNotifier<bool>(false);

  const EditTitleWidget({
    super.key,
    required this.initialTitle,
    required this.onSave,
  });

  @override
  State<EditTitleWidget> createState() => _EditTitleWidgetState();
}

class _EditTitleWidgetState extends State<EditTitleWidget> {
  late TextEditingController _titleController;

  @override
  void initState() {
    super.initState();
    _titleController = TextEditingController(text: widget.initialTitle);
    EditTitleWidget.isExpandedNotifier.value = true;
  }

  @override
  void dispose() {
    EditTitleWidget.isExpandedNotifier.value = false;
    _titleController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: 180.h,
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
                    widget.onSave(_titleController.text);
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
          Padding(
            padding: EdgeInsets.all(20.w),
            child: TextField(
              controller: _titleController,
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
          ),
        ],
      ),
    );
  }
}
