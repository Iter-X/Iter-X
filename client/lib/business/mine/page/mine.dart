import 'package:client/app/notifier/user.dart';
import 'package:client/app/routes.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:provider/provider.dart';


class MinePage extends StatefulWidget {
  const MinePage({super.key});

  @override
  State<MinePage> createState() => _MinePageState();
}

class _MinePageState extends BaseState<MinePage> {
  @override
  void initState() {
    super.initState();
    // 进入页面时刷新用户信息
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<UserNotifier>().refreshUserInfo();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Consumer<UserNotifier>(
      builder: (context, userNotifier, child) {
        final user = userNotifier.user;
        
        return Scaffold(
          appBar: AppBar(
            title: Text('个人中心'),
            backgroundColor: Colors.transparent,
            elevation: 0,
          ),
          body: SingleChildScrollView(
            child: Column(
              children: [
                // 用户头像和基本信息
                Container(
                  padding: EdgeInsets.all(20.w),
                  child: Row(
                    children: [
                      CircleAvatar(
                        radius: 40.w,
                        backgroundImage: user?.hasValidAvatar == true
                            ? NetworkImage(user!.avatarUrl)
                            : null,
                        child: user?.hasValidAvatar != true
                            ? Icon(Icons.person, size: 40.w)
                            : null,
                      ),
                      SizedBox(width: 20.w),
                      Expanded(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              user?.nickname ?? user?.username ?? '未设置昵称',
                              style: TextStyle(
                                fontSize: 20.sp,
                                fontWeight: FontWeight.bold,
                                color: BaseColor.c_1D1F1E,
                              ),
                            ),
                            if (user?.email != null) ...[
                              SizedBox(height: 5.h),
                              Text(
                                user!.email!,
                                style: TextStyle(
                                  fontSize: 14.sp,
                                  color: Colors.grey,
                                ),
                              ),
                            ],
                          ],
                        ),
                      ),
                    ],
                  ),
                ),
                
                // 分割线
                Divider(height: 1.h),
                
                // 功能列表
                ListTile(
                  leading: Icon(Icons.edit),
                  title: Text('编辑资料'),
                  trailing: Icon(Icons.chevron_right),
                  onTap: () {
                    // TODO: 跳转到编辑资料页面
                  },
                ),
                
                ListTile(
                  leading: Icon(Icons.settings),
                  title: Text('设置'),
                  trailing: Icon(Icons.chevron_right),
                  onTap: () {
                    // TODO: 跳转到设置页面
                  },
                ),
                
                // 退出登录按钮
                Container(
                  width: double.infinity,
                  padding: EdgeInsets.symmetric(horizontal: 20.w, vertical: 20.h),
                  child: ElevatedButton(
                    onPressed: () async {
                      await userNotifier.logout();
                      if (mounted) {
                        go(Routes.login, clearStack: true);
                      }
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: BaseColor.c_1D1F1E,
                      padding: EdgeInsets.symmetric(vertical: 12.h),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(24.w),
                      ),
                    ),
                    child: Text(
                      '退出登录',
                      style: TextStyle(
                        fontSize: 16.sp,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        );
      },
    );
  }
}
