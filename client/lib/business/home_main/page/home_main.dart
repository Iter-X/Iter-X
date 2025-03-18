import 'package:client/app/routes.dart';
import 'package:client/business/home/page/home.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import '../../../common/material/state.dart';
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
      body: Column(
        children: [
          Expanded(
            child: selectIndex == 0 ? HomePage() : MinePage(),
          ),
          Container(
            height: 102.h,
            width: double.infinity,
            margin: EdgeInsets.only(
              left: 50.w,
              right: 50.w,
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                ItemHomeBottomWidget(
                  img: selectIndex == 0 ? 'ic_home_selected.png' : 'ic_home.png',
                  onTap: () {
                    setState(() {
                      selectIndex = 0;
                    });
                  },
                ),
                ItemHomeBottomWidget(
                  img: 'ic_add.png',
                  onTap: () {
                    go(Routes.createTripHome);
                  },
                ),
                ItemHomeBottomWidget(
                  img: selectIndex == 1 ? 'ic_mine_selected.png' : 'ic_mine.png',
                  onTap: () {
                    setState(() {
                      selectIndex = 1;
                    });
                  },
                ),
              ],
            ),
          )
        ],
      ),
    );
  }
}
