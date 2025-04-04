import 'dart:io';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:shimmer_animation/shimmer_animation.dart';

// Loading type enum
enum ImageLoadingType {
  circular,
  shimmer,
}

class BaseImage {
  BaseImage._();

  static Widget net(
    String url, {
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
    ImageLoadingType loadingType = ImageLoadingType.shimmer,
  }) {
    width ??= size;
    height ??= size;

    Widget getLoadingWidget() {
      switch (loadingType) {
        case ImageLoadingType.circular:
          return Center(
            child: placeholderChild ??
                CircularProgressIndicator(
                  valueColor: AlwaysStoppedAnimation<Color>(
                    progressColor ?? AppColor.bg,
                  ),
                ),
          );
        case ImageLoadingType.shimmer:
          return Shimmer(
            colorOpacity: 0.3,
            child: Container(
              width: width,
              height: height,
              decoration: BoxDecoration(
                color: Colors.grey[300],
                borderRadius: BorderRadius.circular(circular ?? 0),
              ),
            ),
          );
      }
    }

    return create(
      url.isNotEmpty == true
          ? CachedNetworkImage(
              fit: fit ?? getBoxFit(width: width, height: height),
              imageUrl: url,
              memCacheWidth: (1080 * 0.8).toInt(),
              memCacheHeight: (1920 * 0.8).toInt(),
              maxWidthDiskCache: (1080 * 0.8).toInt(),
              maxHeightDiskCache: (1920 * 0.8).toInt(),
              placeholder: (context, url) => getLoadingWidget(),
              errorWidget: (context, url, error) =>
                  placeholderChild ??
                  Icon(
                    Icons.error,
                    color: progressColor ?? AppColor.scaffoldBackgroundColor,
                    size: 15.w,
                  ),
            )
          : asset(
              name: assetName ?? 'placeholder.png',
              fit: getBoxFit(width: size, height: height),
            ),
      aspectRatio: aspectRatio,
      width: width,
      height: height,
      circular: circular,
      borderWidth: borderWidth,
      borderColor: borderColor,
    );
  }

  static Widget asset({
    String base = AppConfig.assetBaseDir,
    required String name,
    double? aspectRatio,
    double? size,
    double? width,
    double? height,
    double? circular,
    double? borderWidth,
    Color? borderColor,
    Color? color,
    Rect? centerSlice,
    BoxFit? fit,
  }) {
    width ??= size;
    height ??= size;
    bool isSvg = name.toLowerCase().endsWith('.svg');

    return create(
      isSvg
          ? SvgPicture.asset(
              '$base/$name',
              width: double.infinity,
              height: double.infinity,
              fit: fit ?? getBoxFit(width: width, height: height),
              colorFilter: color != null
                  ? ColorFilter.mode(color, BlendMode.srcIn)
                  : null,
              allowDrawingOutsideViewBox: true,
            )
          : Image.asset(
              '$base/$name',
              fit: fit ?? getBoxFit(width: width, height: height),
              color: color,
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
    String base = AppConfig.assetBaseDir,
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
        child: getBoxFit(width: width, height: height) == BoxFit.cover
            ? AspectRatio(aspectRatio: aspectRatio ?? 1, child: child)
            : child,
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
