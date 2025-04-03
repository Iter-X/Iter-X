import 'package:client/app/constants.dart';
import 'package:client/app/routes.dart';
import 'package:client/business/create_trip/widgets/bottom_create_widgets.dart';
import 'package:client/business/create_trip/widgets/widgets.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

class CreateTripHomePage extends StatefulWidget {
  const CreateTripHomePage({super.key});

  @override
  State<CreateTripHomePage> createState() => _CreateTripHomePageState();
}

class _CreateTripHomePageState extends BaseState<CreateTripHomePage> {
  int selectIndex = 0;
  final _controller = TextEditingController();
  final FocusNode _focusNode = FocusNode();
  bool _hasFocus = false;
  bool _selectOneHasFocus = false;
  String recognizedText = ''; // 用于存储识别到的文本

  @override
  void initState() {
    super.initState();
    _focusNode.addListener(_onFocusChange);
  }

  @override
  void dispose() {
    _focusNode.removeListener(_onFocusChange);
    _focusNode.dispose();
    _controller.dispose();
    super.dispose();
  }

  void onTextRecognized(String text) {
    setState(() {
      recognizedText = text;
      print('语音向外传输:$text');
    });
  }

  Widget _buildButtonGroup() {
    return Column(
      children: [
        Row(
          children: [
            ItemCreateWidget(
              isSelected: selectIndex == 0,
              img: selectIndex == 0
                  ? 'ic_create_manually.png'
                  : 'ic_create_manually_black.png',
              text: '手动创建',
              onTap: () {
                _switchTab(0);
              },
            ),
            Gap(20.w),
            ItemCreateWidget(
              isSelected: selectIndex == 1,
              img: selectIndex == 1
                  ? 'ic_create_voice.png'
                  : 'ic_create_voice_black.png',
              text: '语音创建',
              onTap: () {
                _switchTab(1);
              },
            ),
          ],
        ),
        Gap(20.w),
        Row(
          children: [
            ItemCreateWidget(
              isSelected: selectIndex == 2,
              img: selectIndex == 2
                  ? 'ic_create_link.png'
                  : 'ic_create_link_black.png',
              text: '外部链接',
              onTap: () {
                _switchTab(2);
              },
            ),
            Gap(20.w),
            ItemCreateWidget(
              isSelected: selectIndex == 3,
              img: selectIndex == 3
                  ? 'ic_create_photo.png'
                  : 'ic_create_photo_black.png',
              text: '截图创建',
              onTap: () {
                _switchTab(3);
              },
            ),
          ],
        ),
      ],
    );
  }

  Widget _buildHeaderSection() {
    return Container(
      width: double.infinity,
      margin: EdgeInsets.symmetric(horizontal: 35.w),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            _selectOneHasFocus ? '你想去哪？' : 'Hi, Leo',
            style: TextStyle(
              fontSize: 30.sp,
              color: AppColor.primaryFont,
              fontWeight: AppFontWeight.bold,
            ),
          ),
          Text(
            _selectOneHasFocus
                ? '输入目的地'
                : '世界是一本书，那些不旅行的人只读了其中的一页\n——Danny Kaye',
            style: TextStyle(
              fontSize: 18.sp,
              color: AppColor.grayFont,
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildBottomSection() {
    if (selectIndex == 0 || selectIndex == 2) {
      return CreateManuallyWidget(
        controller: _controller,
        focusNode: _focusNode,
        hasFocus: _hasFocus,
        selectIndex: selectIndex,
        onTap: () async {
          if (selectIndex == 0) {
            go(Routes.cardSelection);
          } else {
            final ClipboardData? clipboardData =
                await Clipboard.getData(Clipboard.kTextPlain);
            if (clipboardData != null && clipboardData.text != null) {
              setState(() {
                _controller.text = clipboardData.text!;
              });
            }
          }
        },
      );
    } else if (selectIndex == 3) {
      return CreatePhotoWidget(
        onTap: () {},
      );
    } else {
      return CreateVoiceWidget(onTextRecognized: onTextRecognized);
    }
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      hasAppBar: true,
      leading: ReturnButton(
        onTap: () {
          if (_hasFocus) {
            _focusNode.unfocus();
          } else {
            Navigator.of(context).pop();
          }
        },
      ),
      child: Column(
        children: [
          Flexible(
            flex: 2,
            child: _buildHeaderSection(),
          ),
          Gap(20.h),
          if (!_selectOneHasFocus)
            Flexible(
              flex: 3,
              child: Container(
                margin: EdgeInsets.symmetric(horizontal: 35.w),
                child: _buildButtonGroup(),
              ),
            ),
          if (!_selectOneHasFocus) Gap(20.h),
          Flexible(
            flex: 2,
            child: _buildBottomSection(),
          ),
        ],
      ),
    );
  }

  void _switchTab(int index) {
    setState(() {
      selectIndex = index;
    });
    _onFocusChange();
  }

  void _onFocusChange() {
    setState(() {
      _hasFocus = _focusNode.hasFocus;
      _selectOneHasFocus = selectIndex == 0 && _hasFocus;
    });
  }
}
