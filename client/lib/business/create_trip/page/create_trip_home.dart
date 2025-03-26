import 'package:client/app/routes.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/text_field.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter/services.dart';
import 'package:gap/gap.dart';

import '../../../common/material/app.dart';
import '../../../common/utils/color.dart';
import '../../common/widgets/buttom_widgets.dart';
import '../../../common/material/state.dart';
import '../widgets/bottom_create_widgets.dart';
import '../widgets/widgets.dart';

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

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: BaseColor.c_f2f2f2,
      resizeToAvoidBottomInset: true,
      appBar: AppBar(
        backgroundColor: Colors.transparent,
        elevation: 0,
        systemOverlayStyle: SystemUiOverlayStyle(
          statusBarColor: Colors.transparent,
          statusBarIconBrightness: Brightness.dark,
          statusBarBrightness: Brightness.light,
        ),
        leading: ButtonBackWidget(),
      ),
      body: SafeAreaX(
        child: Column(
          children: [
            Expanded(
              child: SingleChildScrollView(
                child: Container(
                  margin: EdgeInsets.only(
                    left: 35.w,
                    right: 35.w,
                  ),
                  padding: EdgeInsets.only(top: _hasFocus ? 20.h : 100.h),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        _selectOneHasFocus ? '你想去哪？' : 'Hi, Leo',
                        style: TextStyle(
                          fontSize: 30.sp,
                          color: BaseColor.c_1D1F1E,
                          fontWeight: FontWeight.w600,
                        ),
                      ),
                      Text(
                        _selectOneHasFocus
                            ? '输入目的地'
                            : '世界是一本书，那些不旅行的人只读了其中的一页\n——Danny Kaye',
                        style: TextStyle(
                          fontSize: 18.sp,
                          color: const Color(0xFF888888),
                        ),
                      ),
                      Gap(70.h),
                      _selectOneHasFocus
                          ? const SizedBox()
                          : Row(
                              children: [
                                ItemCreateWidget(
                                  isSelected: selectIndex == 0,
                                  img: selectIndex == 0
                                      ? 'ic_create_manually.png'
                                      : 'ic_create_manually_black.png',
                                  text: '手动创建',
                                  onTap: () {
                                    setState(() {
                                      selectIndex = 0;
                                    });
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
                                    setState(() {
                                      selectIndex = 1;
                                    });
                                  },
                                ),
                              ],
                            ),
                      _selectOneHasFocus
                          ? const SizedBox()
                          : Container(
                              margin: EdgeInsets.only(top: 15.h),
                              child: Row(
                                children: [
                                  ItemCreateWidget(
                                    isSelected: selectIndex == 2,
                                    img: selectIndex == 2
                                        ? 'ic_create_link.png'
                                        : 'ic_create_link_black.png',
                                    text: '外部链接',
                                    onTap: () {
                                      setState(() {
                                        selectIndex = 2;
                                      });
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
                                      setState(() {
                                        selectIndex = 3;
                                      });
                                    },
                                  ),
                                ],
                              ),
                            ),
                      Gap(138.h),
                    ],
                  ),
                ),
              ),
            ),
            if (selectIndex == 0 || selectIndex == 2)
              CreateManuallyWidget(
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
              )
            else if (selectIndex == 3)
              CreatePhotoWidget(
                onTap: () {},
              )
            else
              CreateVoiceWidget(onTextRecognized: onTextRecognized)
          ],
        ),
      ),
    );
  }

  void _onFocusChange() {
    setState(() {
      _hasFocus = _focusNode.hasFocus;
      _selectOneHasFocus = selectIndex == 0 && _hasFocus;
    });
  }
}
