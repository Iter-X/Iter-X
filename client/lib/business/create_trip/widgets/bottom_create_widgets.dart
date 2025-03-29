import 'package:client/common/material/image.dart';
import 'package:client/common/material/text_field.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter_speech/flutter_speech.dart';
import 'package:gap/gap.dart';

// 手动创建和外部链接 底部样式
class CreateManuallyWidget extends StatelessWidget {
  final TextEditingController controller;
  final FocusNode focusNode;
  final bool hasFocus;
  final int selectIndex;
  final Function onTap;

  const CreateManuallyWidget({
    super.key,
    required this.controller,
    required this.focusNode,
    required this.hasFocus,
    required this.selectIndex,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    final double bottomPadding = MediaQuery.of(context).viewInsets.bottom;

    return Container(
      width: double.infinity,
      height: 56.h,
      margin: EdgeInsets.only(
        bottom: bottomPadding + (bottomPadding > 0 ? 10.h : 90.h),
        left: 35.w,
        right: 35.w,
      ),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(28.w),
        color: BaseColor.c_E3E3E3,
      ),
      child: Row(
        children: [
          Expanded(
            child: BaseTextFieldWidget(
              controller: controller,
              focusNode: focusNode,
              hintText: selectIndex == 0 ? '输入目的地开始规划行程' : '粘贴外部链接，智能识别规划',
              hintStyle: TextStyle(
                fontSize: 15.sp,
                color: BaseColor.c_999999,
              ),
              style: TextStyle(
                fontSize: 15.sp,
                color: BaseColor.c_1D1F1E,
              ),
              contentPadding: EdgeInsets.only(
                left: 24.w,
                right: 24.w,
              ),
            ),
          ),
          GestureDetector(
            onTap: () => onTap.call(),
            child: Container(
              width: 107.w,
              height: 46.h,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(24.w),
                color: BaseColor.c_1D1F1E,
              ),
              margin: EdgeInsets.only(right: 6.w),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  BaseImage.asset(
                    name: 'ic_create_picard.png',
                    size: 18.w,
                  ),
                  Gap(5.w),
                  Text(
                    selectIndex == 0 ? '图卡选择' : '粘贴',
                    style: TextStyle(
                      fontSize: 14.sp,
                      color: BaseColor.c_f2f2f2,
                    ),
                  )
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}

// 截图创建 底部样式
class CreatePhotoWidget extends StatelessWidget {
  final Function onTap;

  const CreatePhotoWidget({
    super.key,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => onTap.call(),
      child: Container(
        width: double.infinity,
        height: 52.h,
        margin: EdgeInsets.only(
          bottom: 90.h,
          left: 35.w,
          right: 35.w,
        ),
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(28.w),
          color: BaseColor.c_1D1F1E,
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            BaseImage.asset(
              name: 'ic_select_photo.png',
              size: 24.w,
            ),
            Gap(10.w),
            Text(
              '选择照片上传',
              style: TextStyle(
                fontSize: 18.sp,
                color: BaseColor.c_f2f2f2,
              ),
            )
          ],
        ),
      ),
    );
  }
}

// 语音创建 底部样式
class CreateVoiceWidget extends StatefulWidget {
  final Function(String) onTextRecognized;

  const CreateVoiceWidget({
    super.key,
    required this.onTextRecognized,
  });

  @override
  _CreateVoiceWidgetState createState() => _CreateVoiceWidgetState();
}

class _CreateVoiceWidgetState extends State<CreateVoiceWidget> {
  double _scale = 1.0;
  String _recognizedText = '';
  late SpeechRecognition _speechRecognition;
  bool _isListening = false;
  final String _locale = 'zh-CN';
  final List<Map<String, dynamic>> _languages = [
    {
      'name': 'Chinese',
      'code': 'zh-CN',
    },
    {
      'name': 'English',
      'code': 'en-US',
    },
  ];

  @override
  void initState() {
    super.initState();
    _speechRecognition = SpeechRecognition();
    _speechRecognition.setAvailabilityHandler((bool result) {
      setState(() {
        // 处理语音识别可用性
      });
    });
    _speechRecognition.setRecognitionStartedHandler(() {
      setState(() {
        _isListening = true;
        print('开始识别');
      });
    });
    _speechRecognition.setRecognitionResultHandler((String text) {
      setState(() {
        _recognizedText = text;
        widget.onTextRecognized(text); // 调用回调函数传递识别结果
        _isListening = false;
        print('识别结果：$text');
      });
    });
    _speechRecognition.setRecognitionCompleteHandler((String text) {
      setState(() {
        _isListening = false;
        print('识别完成：$text');
      });
    });
    _speechRecognition.activate(_languages[0]['code']).then((result) {
      setState(() {
        // 处理激活结果
        print('激活结果：$result');
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return Column(mainAxisAlignment: MainAxisAlignment.end, children: [
      if (_isListening)
        Container(
            margin: EdgeInsets.only(bottom: 20.h),
            child: BaseImage.asset(
              name: 'ic_create_voice_black.png',
              width: 50.w,
              height: 50.h,
              fit: BoxFit.cover,
            )),
      if (_recognizedText != '')
        Text(
          _recognizedText,
          style: TextStyle(
            fontSize: 16.sp,
            color: BaseColor.c_1D1F1E,
          ),
        ),
      Container(
        margin: EdgeInsets.only(bottom: 90.h, top: 20.h),
        child: GestureDetector(
          onLongPress: () {
            _recognizedText = '';
            _scale = 1.2;
            _speechRecognition.listen();
          },
          onLongPressEnd: (_) {
            setState(() {
              _scale = 1.0;
            });
            _speechRecognition.stop();
          },
          child: Transform.scale(
            scale: _scale,
            child: BaseImage.asset(
              name: 'btn_create_voice.png',
              width: 100.w,
              height: 100.h,
              fit: BoxFit.cover,
            ),
          ),
        ),
      )
    ]);
  }
}
