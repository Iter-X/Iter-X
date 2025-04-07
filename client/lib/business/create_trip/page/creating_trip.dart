import 'package:client/app/constants.dart';
import 'package:client/app/routes.dart';
import 'package:client/common/dio/http.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/toast.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:lottie/lottie.dart';

class CreatingTripPage extends StatefulWidget {
  final Map<String, dynamic> params;

  const CreatingTripPage({super.key, required this.params});

  @override
  State<CreatingTripPage> createState() => _CreatingTripPageState();
}

class _CreatingTripPageState extends BaseState<CreatingTripPage> {
  final List<String> _tips = [
    '正在为您规划行程...',
    '正在计算最佳路线...',
    '正在整理景点信息...',
    '马上就好...'
  ];
  int _currentTipIndex = 0;
  bool _isCreating = true;

  @override
  void initState() {
    super.initState();
    // _createTrip();
    _startTipsAnimation();
  }

  void _startTipsAnimation() {
    Future.delayed(const Duration(seconds: 3), () {
      if (!_isCreating && mounted) return;
      setState(() {
        _currentTipIndex = (_currentTipIndex + 1) % _tips.length;
      });
      _startTipsAnimation();
    });
  }

  Future<void> _createTrip() async {
    final result = await Http.instance.post(
      '/api/v1/trips/card',
      data: widget.params,
      isShowLoading: false,
    );

    _isCreating = false;
    if (result.isSuccess()) {
      final tripId = result.data['trip']['id'];
      if (mounted) {
        go(Routes.tripOverview, arguments: {'tripId': tripId});
      }
    } else {
      ToastX.show(result.msg ?? '创建行程失败');
      if (mounted) {
        Navigator.pop(context);
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      hasAppBar: false,
      backgroundColor: AppColor.bg,
      child: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Lottie.asset(
              'assets/lottie/earth.json',
              width: 200.w,
              height: 200.w,
            ),
            SizedBox(height: 40.h),
            AnimatedSwitcher(
              duration: const Duration(milliseconds: 500),
              child: Text(
                _tips[_currentTipIndex],
                key: ValueKey(_currentTipIndex),
                style: TextStyle(
                  fontSize: 16.sp,
                  color: AppColor.primaryFont,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
