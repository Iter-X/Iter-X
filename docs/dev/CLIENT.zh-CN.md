<div align="center">
  <img src="../logo.png" alt="Logo" width="290" height="251" />
</div>

<div align="center">

| [English](CLIENT.md) | [中文简体](CLIENT.zh-CN.md) |

</div>

# 客户端开发规范

## 代码风格规范

### 命名规范
- 变量/参数：小驼峰命名法（如 userName，isLoading）
- 文件夹名/dart文件名：全小写加下划线（如 home_main）
- 常量：全大写加下划线（如 MAX_ITEMS）
- 资源文件：类型_图片名（如 ic_name.png）
- 类名：大驼峰命名法
  - 页面文件：后缀加 Page（如 UserInfoDetailPage）
  - 弹窗文件：后缀加 Dialog
  - 其余组件：基本用 Widget 即可
- 工具类：后缀加 Utils
- 方法名：

  | 方法 | 说明 |
  |------|------|
  | initXXX() | 初始化 |
  | isXX() | 返回 bool |
  | getXX() | 返回某个值 |
  | setXXX() | 设置某个值 |
  | updateXXX() | 更新数据 |
  其他方法名尽量保证方法名能直观表达方法用途

### 代码格式
- 缩进：tab 缩进
- 空行：
  - 类与类之间空一行
  - 方法之间空一行
  - 逻辑块内无空行（可以通过使用注释 // 代替空行）
- 导入语句：可以使用 as 别名（如 import 'package:flutter/material.dart' as m;）

### Dart 语言规范
- 字符串插值：优先使用 String interpolation，而非 + 拼接
  ```dart
  // 推荐
  Text('Hello, ${userName}!');
  ```
- 条件语句：简单的条件语句使用三元运算符简化代码
  ```dart
  bool isValid = length >= 10 ? true : false;
  ```
- 集合操作：使用 isEmpty 代替 length == 0

## 二、项目结构规范

### 基础目录结构
```
assets/  
├── images/         # 存放图片
├── voice/          # 存放语音
lib/
├── app/            # 常量配置（api路径、路由路径等）
│   ├── apis/       # api路径
│   ├── events/     # event事件
│   ├── foundation/ # mixin文件
│   ├── notifier/   # 状态管理
│   ├── constants.dart   # 常量
│   ├── routes.dart     # 路由路径
├── business/       # 功能模块（按业务划分）
│   ├── auth/       # 认证模块
│   │   ├── dialog/    # 弹窗组件
│   │   ├── entity/    # 实体类
│   │   ├── page/      # 页面文件
│   │   ├── service/   # 网络请求
│   │   └── widgets/   # 封装组件
│   ├── common/     # 公共模块
├── common/         # 公共配置
│   ├── dio/        # 网络请求配置
│   ├── material/    # 封装的基础组件
│   ├── utils/       # 工具类
├── main.dart       # 主入口
```

### 模块化设计
- 根据功能模块划分目录
- 每个功能模块独立开发
- 遇到两次以上的组件/方法就进行封装并注释

## 性能优化规范

### 减少 Widget 重建
- 使用 const 构建函数标记不可变组件
- 避免在 build 方法中调用 setState 或耗时操作
- 可以使用 RepaintBoundary 包裹频繁更新的局部区域

### 状态管理优化
- 选择高效的状态管理方案（如 provider、bloc 等）
- 避免过度细分状态，按需更新

### 异步操作
- 使用 async/await 或者 FutureBuilder 处理异步任务
- 避免在主线程执行耗时操作

### 资源管理
- 图片资源使用封装好的图片组件，可以缓存常用资源
- 避免内存泄露，及时关闭 stream 或 timer 等

## 协作与版本控制

### Git 规范
- 提交的文案参考 git 项目中的说明规范

### 代码审查
- 提交代码时检查未使用或未引用的代码，若真的没有使用则及时删除
- 在复杂逻辑中及时增加注释
- 实体类中表明字段含义 