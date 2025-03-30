import 'package:client/app/constants.dart';

// 资源路径工具类
class AssetUtil {
  AssetUtil._();

  static String getAsset(String name) {
    // return '$base/$name';
    return '${AppConfig.assetBaseDir}/$name';
  }
}
