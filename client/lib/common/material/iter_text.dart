/*
 * @Description: 支持文字描边的组件，后续有其他文字需要的功能也考虑在这里拓展
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-27 17:17:58
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-03-27 17:51:36
 */
import 'package:flutter/material.dart';

class IterText extends StatelessWidget {
  final String data;
  final TextStyle? style;
  final TextAlign? textAlign;
  final TextDirection? textDirection;
  final Locale? locale;
  final bool? softWrap;
  final TextOverflow? overflow;
  final TextScaler? textScaler;
  final int? maxLines;
  final String? semanticsLabel;
  final TextWidthBasis? textWidthBasis;
  final TextHeightBehavior? textHeightBehavior;
  final BorderProperties? borders;
  final AlignmentDirectional? alignment;

  const IterText(
    this.data, {
    super.key,
    this.style,
    this.textAlign,
    this.textDirection,
    this.locale,
    this.softWrap,
    this.overflow,
    this.textScaler,
    this.maxLines,
    this.semanticsLabel,
    this.textWidthBasis,
    this.textHeightBehavior,
    this.borders,
    this.alignment,
  });

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: alignment ?? AlignmentDirectional.center,
      children: [
        if (borders != null)
          Text(
            data,
            style: (style ?? const TextStyle()).copyWith(
              foreground: Paint()
                ..style = PaintingStyle.stroke
                ..strokeWidth = borders!.width
                ..color = borders!.color,
            ),
            textAlign: textAlign,
            textDirection: textDirection,
            locale: locale,
            softWrap: softWrap,
            overflow: overflow,
            textScaler: textScaler,
            maxLines: maxLines,
            semanticsLabel: semanticsLabel,
            textWidthBasis: textWidthBasis,
            textHeightBehavior: textHeightBehavior,
          ),
        Text(
          data,
          style: style,
          textAlign: textAlign,
          textDirection: textDirection,
          locale: locale,
          softWrap: softWrap,
          overflow: overflow,
          textScaler: textScaler,
          maxLines: maxLines,
          semanticsLabel: semanticsLabel,
          textWidthBasis: textWidthBasis,
          textHeightBehavior: textHeightBehavior,
        ),
      ],
    );
  }
}

class BorderProperties {
  final double width;
  final Color color;

  const BorderProperties({
    required this.width,
    required this.color,
  });
}
