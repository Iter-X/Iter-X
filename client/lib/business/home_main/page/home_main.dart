import 'package:client/app/constants.dart';
import 'package:client/app/routes.dart';
import 'package:client/business/home/page/home.dart';
import 'package:client/business/mine/page/mine_page.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/bottom_nav_bar.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/state.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class HomeMainPage extends StatefulWidget {
  const HomeMainPage({super.key});

  @override
  State<HomeMainPage> createState() => _HomeMainPageState();
}

class _HomeMainPageState extends BaseState<HomeMainPage> {
  int _selectedIndex = 0;

  final List<Widget> _pages = [
    const HomePage(),
    const MinePage(),
  ];

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      hasAppBar: false,
      top: false,
      backgroundColor: AppColor.bg,
      bottomColor: AppColor.bottomBar,
      bottomBarCfg: BottomNavConfig(
        items: [
          BottomBarItem(
            iconImage: 'ic_home.png',
            selectedIconImage: 'ic_home_selected.png',
          ),
          BottomBarItem(
            iconImage: 'ic_mine.png',
            selectedIconImage: 'ic_mine_selected.png',
          ),
        ],
        selectedIndex: _selectedIndex,
        onIndexChanged: (index) => setState(() => _selectedIndex = index),
        centerButton: BaseImage.asset(
          name: 'ic_add.png',
          size: 38.w,
        ),
        onCenterButtonTap: () => go(Routes.createTripHome),
      ),
      child: _pages[_selectedIndex],
    );
  }
}
