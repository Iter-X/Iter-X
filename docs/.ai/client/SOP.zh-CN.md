# Standard Operating Procedures (SOP) Manual
1. 客户端语言是Flutter，代码在 `client` 目录下，主要代码都在 `client/lib` 目录下
2. 所有页面位于 `client/lib/bussiness` 目录下，如 `client/lib/bussiness/auth`
3. 通常每个页面下会包含这么几个目录（当不存在时你可以自己根据情况选择创建）：
   - `entity`：实体类
   - `page`：页面文件
   - `service`：网络请求
   - `widgets`：封装组件
4. 所有的公共工具配置都在 `client/lib/common` 目录下，几个重要的：
   - `dio`：网络请求配置
   - `material`：封装的基础组件
   - `utils`：工具、配置类，如全局配置、颜色都在这里
   - `widgets`：多个页面可用的通用组件


## 路由
路由在 `client/lib/app/routes.dart` 中定义，使用 `fluro` 进行路由管理，有新页面需要到这里进行注册

## 登录状态 & 用户信息管理
登录状态使用 `provider` 进行管理，相关代码在 `client/lib/app/notifier/user.dart` 中，在页面中可以从这里获取对应的用户信息。在API请求的时候，可以直接从这里获取token信息，我们会用refresh token来获取新的token，确保用户在APP中的持续登录状态

## 网络请求
网络请求使用 `dio` 进行管理，相关代码在 `client/lib/common/dio` 中，URI的常量定义在 `client/lib/app/apis` 目录下，有新的接口需要到这边去增加，这里是根据模块管理的，比如 `auth` 模块的接口就在 `client/lib/app/apis/auth.dart` 中定义，如果发现不存在文件的时候可以自己创建文件

对应的请求一般是定义在对应的页面的 `service` 目录下，比如 `client/lib/bussiness/auth/service` 目录下

## 本地存储
本地存储使用 `shared_preferences` 进行管理，相关代码在 `client/lib/common/utils/shared_preference_util.dart` 中，可以直接调用这里的方法进行存储和读取（只能存不敏感的信息）

## 全局配置
全局配置在 `client/lib/app/constants.dart` 中，可以在这里定义一些全局的常量，比如颜色、字体等：
- 字号可以直接AppFontWeight.normal进行调用，不要直接写w300这种数字
- 颜色可以直接AppColor.primary进行调用，不要直接写颜色值
- 如果有圆角间距之类直接从AppConfig中调用，比如AppConfig.boxRadius就是通用圆角，而AppConfig.cornerRadius就是按钮的圆角（全圆）

## SafeArea & AppBar
我们自己定义了 `AppBarWithSafeArea` 在 `lib/common/material/app_bar_with_safe_area.dart` 中，这个是带有安全区域的AppBar，使用的时候直接调用这个组件即可，具体用法参考代码或者其他页面的使用，这个是可以单独开启AppBar，单独设置上下安全距离是否开启，上下背景之类的

## Toast
有Toast的统一调用 `lib/common/utils/toast.dart` 中，直接调用`ToastX.show('提示内容');`即可

## flutter_screenutil
我们使用 `flutter_screenutil` 进行屏幕适配，在设置字体大小的时候，使用`number.sp`，在设置宽高间距的时候，使用`number.w`和`number.h`，正方形的情况下，使用同一个单位，如`width: 100.w, height: 100.w`或`width: 100.h, height: 100.h`
