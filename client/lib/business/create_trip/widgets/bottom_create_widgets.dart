import 'package:client/app/constants.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/text_field.dart';
import 'package:client/common/widgets/base_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter_speech/flutter_speech.dart';

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
    return Container(
      width: double.infinity,
      height: 56.h,
      margin: EdgeInsets.only(
        top: 20.h,
        bottom: 20.h,
        left: 35.w,
        right: 35.w,
      ),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(AppConfig.cornerRadius),
        color: AppColor.inputGrayBG,
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
                color: AppColor.c_999999,
              ),
              style: TextStyle(
                fontSize: 15.sp,
                color: AppColor.c_1D1F1E,
              ),
              contentPadding: EdgeInsets.only(
                left: 24.w,
                right: 24.w,
              ),
            ),
          ),
          Container(
            margin: EdgeInsets.only(right: 6.w),
            child: BaseButton(
              width: 107.w,
              height: 46.h,
              iconName: 'ic_create_picard.png',
              iconSize: 18.w,
              text: selectIndex == 0 ? '图卡选择' : '粘贴',
              textSize: 14.sp,
              textColor: AppColor.secondaryFont,
              backgroundColor: AppColor.primary,
              borderRadius: AppConfig.cornerRadius,
              gap: 5.w,
              onTap: () => onTap.call(),
            ),
          ),
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
    return Container(
      margin: EdgeInsets.only(
        top: 20.h,
        bottom: 20.h,
        left: 35.w,
        right: 35.w,
      ),
      child: BaseButton(
        height: 52.h,
        iconName: 'ic_select_photo.png',
        iconSize: 24.w,
        text: '选择照片上传',
        textSize: 18.sp,
        textColor: AppColor.secondaryFont,
        backgroundColor: AppColor.primary,
        borderRadius: AppConfig.cornerRadius,
        gap: 10.w,
        onTap: () => onTap.call(),
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
    return Container(
      margin: EdgeInsets.symmetric(horizontal: 35.w),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          if (_isListening)
            Container(
              margin: EdgeInsets.only(bottom: 10.h),
              child: BaseImage.asset(
                name: 'ic_create_voice_black.png',
                width: 50.w,
                height: 50.h,
                fit: BoxFit.cover,
              ),
            ),
          if (_recognizedText.isNotEmpty)
            Container(
              margin: EdgeInsets.only(bottom: 10.h),
              child: Text(
                _recognizedText,
                style: TextStyle(
                  fontSize: 16.sp,
                  color: AppColor.c_1D1F1E,
                ),
                textAlign: TextAlign.center,
              ),
            ),
          Container(
            margin: EdgeInsets.only(bottom: 20.h),
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
          ),
        ],
      ),
    );
  }
}
