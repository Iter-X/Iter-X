// 资源路径工具类
class AssetUtil {
  AssetUtil._();

  static const String base = 'assets/images';

  static String getAsset(String name) {
    return '$base/$name';
  }
}
