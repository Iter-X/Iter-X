import 'dart:io';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:client/common/utils/asset.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class BaseImage {
  BaseImage._();

  static Widget net(
    String? url, {
    progressColor,
    double? aspectRatio,
    double? size,
    double? width,
    double? height,
    double? circular,
    double? borderWidth,
    Color? borderColor,
    String? assetName,
    Widget? placeholderChild,
    BoxFit? fit,
  }) {
    width ??= size;
    height ??= size;
    return create(
      url?.isNotEmpty == true
          ? CachedNetworkImage(
              fit: fit ?? getBoxFit(width: width, height: height),
              imageUrl: url!,
              memCacheWidth: (1080 * 0.8).toInt(),
              memCacheHeight: (1920 * 0.8).toInt(),
              maxWidthDiskCache: (1080 * 0.8).toInt(),
              maxHeightDiskCache: (1920 * 0.8).toInt(),
              placeholder: (context, url) => Center(
                child: placeholderChild ??
                    CircularProgressIndicator(
                      valueColor: AlwaysStoppedAnimation<Color>(
                        progressColor ?? BaseColor.scaffoldBackgroundColor,
                      ),
                    ),
              ),
              errorWidget: (context, url, error) =>
                  placeholderChild ??
                  Icon(
                    Icons.error,
                    color: progressColor ?? BaseColor.scaffoldBackgroundColor,
                    size: 15.w,
                  ),
            )
          : assetName?.isNotEmpty == true
              ? Image.asset(
                  assetName!,
                  fit: getBoxFit(width: size, height: height),
                )
              : Container(),
      aspectRatio: aspectRatio,
      width: width,
      height: height,
      circular: circular,
      borderWidth: borderWidth,
      borderColor: borderColor,
    );
  }

  static Widget asset({
    String base = AssetUtil.base,
    required String name,
    double? aspectRatio,
    double? size,
    double? width,
    double? height,
    double? circular,
    double? borderWidth,
    Color? borderColor,
    Color? imageColor,
    Rect? centerSlice,
    BoxFit? fit,
  }) {
    width ??= size;
    height ??= size;
    return create(
      Image.asset(
        '$base/$name',
        fit: fit ?? getBoxFit(width: size, height: height),
        color: imageColor,
        centerSlice: centerSlice,
      ),
      aspectRatio: aspectRatio,
      width: width,
      height: height,
      circular: circular,
      borderWidth: borderWidth,
      borderColor: borderColor,
    );
  }

  static Widget file({
    String base = AssetUtil.base,
    required String name,
    double? aspectRatio,
    double? size,
    double? width,
    double? height,
    double? circular,
    double? borderWidth,
    Color? borderColor,
    Color? imageColor,
  }) {
    width ??= size;
    height ??= size;
    return create(
      Image.file(
        File(name),
        fit: getBoxFit(width: size, height: height),
        color: imageColor,
      ),
      aspectRatio: aspectRatio,
      width: width,
      height: height,
      circular: circular,
      borderWidth: borderWidth,
      borderColor: borderColor,
    );
  }

  static Widget create(
    Widget child, {
    double? aspectRatio,
    double? size,
    double? width,
    double? height,
    double? circular,
    double? borderWidth,
    Color? borderColor,
  }) {
    width ??= size;
    height ??= size;
    return Container(
      width: width,
      height: height,
      decoration: BoxDecoration(
        border: Border.all(
          width: borderWidth ?? 0,
          color: borderColor ?? Colors.transparent,
        ),
        borderRadius: BorderRadius.all(Radius.circular(circular ?? 0)),
      ),
      child: ClipRRect(
        borderRadius: BorderRadius.circular(circular ?? 0),
        child: getBoxFit(width: width, height: height) == BoxFit.cover ? AspectRatio(aspectRatio: aspectRatio ?? 1, child: child) : child,
      ),
    );
  }

  static BoxFit getBoxFit({double? width, double? height}) {
    if (width != null && height == null) {
      return BoxFit.fitWidth;
    }
    if (width == null && height != null) {
      return BoxFit.fitHeight;
    }
    if (width != null && height != null && width == height) {
      return BoxFit.fill;
    }
    return BoxFit.cover;
  }
}