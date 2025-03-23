import 'package:client/app/routes.dart';
import 'package:client/business/home/page/home.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import '../../../common/material/state.dart';
import '../../../common/utils/color.dart';
import '../../mine/page/mine.dart';
import '../widgets/widgets.dart';

class HomeMainPage extends StatefulWidget {
  const HomeMainPage({super.key});

  @override
  State<HomeMainPage> createState() => _HomeMainPageState();
}

class _HomeMainPageState extends BaseState<HomeMainPage> {
  int selectIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: Column(
        children: [
          Container(
            color: BaseColor.c_f2f2f2,
            child: SafeArea(
              bottom: false,
              child: Container(height: 0),
            ),
          ),
          Expanded(
            child: Container(
              color: BaseColor.c_f2f2f2,
              child: selectIndex == 0 ? HomePage() : MinePage(),
            ),
          ),
          Container(
            color: Colors.white,
            child: SafeArea(
              top: false,
              child: Container(
                color: BaseColor.bottomBar,
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Container(
                      height: 72.h,
                      width: double.infinity,
                      decoration: BoxDecoration(
                        border: Border(
                          top: BorderSide(
                            color: BaseColor.bottomBarLine,
                            width: 1,
                          ),
                        ),
                      ),
                      padding: EdgeInsets.symmetric(vertical: 15.h),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Expanded(
                            child: Center(
                              child: ItemHomeBottomWidget(
                                img: selectIndex == 0 ? 'ic_home_selected.png' : 'ic_home.png',
                                onTap: () {
                                  setState(() {
                                    selectIndex = 0;
                                  });
                                },
                              ),
                            ),
                          ),
                          Expanded(
                            child: Center(
                              child: ItemHomeBottomWidget(
                                img: 'ic_add.png',
                                onTap: () {
                                  go(Routes.createTripHome);
                                },
                              ),
                            ),
                          ),
                          Expanded(
                            child: Center(
                              child: ItemHomeBottomWidget(
                                img: selectIndex == 1 ? 'ic_mine_selected.png' : 'ic_mine.png',
                                onTap: () {
                                  setState(() {
                                    selectIndex = 1;
                                  });
                                },
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ],
                ),
              ),
            ),
          )
        ],
      ),
    );
  }
}
