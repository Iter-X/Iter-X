import 'package:client/app/constants.dart';
import 'package:client/business/trip/page/trip_overview_page.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/bottom_nav_bar.dart';
import 'package:flutter/material.dart';

class TripMainPage extends StatefulWidget {
  const TripMainPage({super.key});

  @override
  State<TripMainPage> createState() => _TripMainPageState();
}

class _TripMainPageState extends State<TripMainPage> {
  int _selectedIndex = 0;

  final List<Widget> _pages = [
    const TripOverviewPage(),
    const Placeholder(), // 地图页面
    const Placeholder(), // 图册页面
    const Placeholder(), // 智能页面
    const Placeholder(), // 记账页面
  ];

  final List<BottomBarItem> _bottomBarItems = [
    BottomBarItem(
      icon: Icons.view_timeline_outlined,
      selectedIcon: Icons.view_timeline,
      label: '总览',
    ),
    BottomBarItem(
      icon: Icons.paid_outlined,
      selectedIcon: Icons.paid,
      label: '记账',
    ),
    BottomBarItem(
      icon: Icons.auto_awesome_outlined,
      selectedIcon: Icons.auto_awesome,
      label: '智能',
    ),
    BottomBarItem(
      icon: Icons.image_outlined,
      selectedIcon: Icons.image,
      label: '图册',
    ),
    BottomBarItem(
      icon: Icons.near_me_outlined,
      selectedIcon: Icons.near_me,
      label: '地图',
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      hasAppBar: false,
      backgroundColor: AppColor.bg,
      bottomColor: AppColor.bottomBar,
      top: false,
      bottomBarCfg: BottomNavConfig(
        items: _bottomBarItems,
        selectedIndex: _selectedIndex,
        onIndexChanged: (index) => setState(() => _selectedIndex = index),
      ),
      child: _pages[_selectedIndex],
    );
  }
}
