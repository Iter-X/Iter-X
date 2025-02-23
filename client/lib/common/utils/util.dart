import 'dart:convert' as convert;

import 'package:decimal/decimal.dart';
import 'package:flutter/material.dart';

class BaseUtil {
  BaseUtil._();

  /// 对象是否为空
  static bool isEmpty(dynamic o) {
    if (o == null) {
      return true;
    }
    if (o is String) {
      return o == '' || o == 'null';
    } else if (o is List) {
      return o.isEmpty;
    }
    return false;
  }

  static bool isNotEmpty(dynamic o) {
    return !isEmpty(o);
  }

  /// 判断手机号是否正确
  static bool isMobile(String str) {
    return str.length == 11;
  }

  static String getStrWithMaxLen(String? str, {int len = 5}) {
    if (str == null || str.isEmpty == true) {
      return '';
    }
    if (str.length <= len) {
      return str;
    }
    return '${str.substring(0, len)}...';
  }

  /*
  * Base64加密
  */
  static String base64Encode(String data) {
    var content = convert.utf8.encode(data);
    return convert.base64Encode(content);
  }

  /*
  * Base64解密
  */
  static String base64Decode(String data) {
    List<int> bytes = convert.base64Decode(data);
    // 网上找的很多都是String.fromCharCodes，这个中文会乱码
    // String txt1 = String.fromCharCodes(bytes);
    return convert.utf8.decode(bytes);
  }

  ///超过四位数的数字转化为w格式,如：38128 => 3.8w，381285 => 38.1w
  static String formatCharCount(int count) {
    if (count <= 0 || count.isNaN) {
      return '0';
    }
    String strCount = count.toString();
    if (strCount.length >= 5) {
      String prefix = strCount.substring(0, strCount.length - 4);
      if (strCount.length == 5) {
        prefix += '.${strCount[1]}';
      }
      if (strCount.length == 6) {
        prefix += '.${strCount[2]}';
      }
      return '${prefix}w';
    }
    return strCount;
  }

  ///验证是否是中文  ,正则中的^表示开头,$表示结束
  static bool isChinese(String value) {
    return RegExp(r"^[\u4e00-\u9fa5]+$").hasMatch(value);
  }

  static String num2Str(v) {
    if (v == null) {
      return '0';
    }
    if (v is! num) {
      return '?';
    }
    // 去掉无效0
    return Decimal.parse(v.toString()).toString();
  }

  static String? cutZero(String? str) {
    if (str == null || str.isEmpty) {
      return str;
    }
    if (str.endsWith('.0') == true) {
      return str.substring(0, str.length - 2);
    }
    return str;
  }

  static String desensitizeMobile(String? mobile) {
    if (mobile?.length != 11) {
      return '****';
    }
    return mobile!.replaceRange(3, 7, '****');
  }

  static String getWaterMobile(String phone) {
    return phone.length == 11 ? phone.substring(7, 11) : phone;
  }

  static bool intListEquals(List<int>? l1, List<int>? l2) {
    if ((l1?.length ?? 0) != (l2?.length ?? 0)) {
      return false;
    }
    //
    if (l1?.isNotEmpty == true && l2?.isNotEmpty == true) {
      l1!.sort();
      l2!.sort();
      for (int i = 0; i < l1.length; i++) {
        if (l1[i] != l2[i]) {
          return false;
        }
      }
    }
    //
    return true;
  }

  static String getPrintSize(limit){
    String  size = "";
    //内存转换
    if(limit == 0) {
      size = '0KB';
    } else if(limit < 0.1 * 1024){                            //小于0.1KB，则转化成B
      size = limit.toString();
      size = "${size.substring(0,size.indexOf(".")+3)}B";
    }else if(limit < 0.1 * 1024 * 1024){            //小于0.1MB，则转化成KB
      size = (limit/1024).toString();
      size = "${size.substring(0,size.indexOf(".")+3)}KB";
    }else if(limit < 0.1 * 1024 * 1024 * 1024){        //小于0.1GB，则转化成MB
      size = (limit/(1024 * 1024)).toString() ;
      size = "${size.substring(0,size.indexOf(".")+3)}MB";
    }else{                                            //其他转化成GB
      size = (limit/(1024 * 1024 * 1024)).toString();
      size = "${size.substring(0,size.indexOf(".")+3)}GB";
    }
    return size;
  }

  // 判断文件是否是图片
  static bool isImageFile(String filePath) {
    final imageExtensions = ['jpg', 'jpeg', 'png', 'gif', 'bmp', 'svg']; // 常见图片文件扩展名

    final fileExtension = filePath.split('.').last.toLowerCase();
    return imageExtensions.contains(fileExtension);
  }

  static String stringSub(String str, int length) {
    String result = '';
    if (isEmpty(str)) {
      return result;
    }
    if (str.length < length) {
      return str;
    } else {
      result = '${str.substring(0, length)}...';
    }
    return result;
  }

  // 判断是否是物理设备
  // static Future<bool> getIsPhysicalDevice() async {
  //   final DeviceInfoPlugin deviceInfo = DeviceInfoPlugin();
  //   if (PlatformUtils().isAndroid) {
  //     AndroidDeviceInfo androidInfo = await deviceInfo.androidInfo;
  //     return androidInfo.isPhysicalDevice;
  //   } else if (PlatformUtils().isIOS) {
  //     IosDeviceInfo iosInfo = await deviceInfo.iosInfo;
  //     return iosInfo.isPhysicalDevice;
  //   }
  //   return false;
  // }

  // 10000->1w
  static String formatNumberW(int number) {
    if (number >= 10000) {
      double w = number / 10000;
      return '${w.toStringAsFixed(1)}w';
    } else {
      return number.toString();
    }
  }
}

extension FixAutoLines on String {
  String fixAutoLines() {
    return Characters(this).join('\u{200B}');
  }
}
