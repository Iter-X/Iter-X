import 'dart:math';

import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class BaseTextField {
  BaseTextField._();

  static InputDecoration defaultDecoration = InputDecoration(
    isDense: false,
    hintText: '请输入',
    hintStyle: TextStyle(
      color: BaseColor.hint,
      fontSize: 16.sp,
    ),
    border: InputBorder.none,
    isCollapsed: true,
    // 重点，相当于高度包裹的意思，必须设置为true，不然有默认奇妙的最小高度
    fillColor: Colors.transparent,
    // 背景颜色，必须结合filled: true,才有效
    filled: true,
    // 重点，必须设置为true，fillColor才有效
    contentPadding: EdgeInsets.zero,
    // 内容内边距，影响高度
  );
}

class BaseTextFieldWidget extends StatefulWidget {
  final FocusNode? focusNode;
  final TextEditingController? controller;
  final TextInputType? keyboardType;
  final TextStyle? style;
  final TextStyle? hintStyle;
  final bool? enabled;
  final List<TextInputFormatter>? inputFormatters;
  final bool? obscureText;

  //

  final int? lengthLimit;

  //
  final String? hintText;
  final Color? fillColor;
  final Widget? suffixIcon;

  //
  final ValueChanged<String>? onChanged;
  final EdgeInsetsGeometry? contentPadding;

  //
  final TextAlign textAlign;
  final double? circular;

  //
  final bool linesNull;
  final int? maxLines;
  final int? minLines;
  final int? maxLength;

  //
  final bool autofocus;
  final bool canClear;
  final Widget? clearWidget;
  final Color? clearColor;

  //
  final TextInputAction? textInputAction;
  final VoidCallback? onEditingComplete;
  final ValueChanged<String>? onSubmitted;

  const BaseTextFieldWidget({
    Key? key,
    this.focusNode,
    this.controller,
    this.keyboardType,
    this.style,
    this.hintStyle,
    this.enabled = true,
    this.inputFormatters,
    this.obscureText,
    //

    this.lengthLimit,
    //
    this.hintText,
    this.fillColor,
    this.suffixIcon,
    //
    this.onChanged,
    this.contentPadding,
    //
    this.textAlign = TextAlign.start,
    this.circular,
    //
    this.linesNull = false,
    this.maxLines,
    this.minLines,
    this.maxLength,
    //
    this.autofocus = false,
    //
    this.canClear = false,
    this.clearWidget,
    this.clearColor,
    //
    this.textInputAction,
    this.onEditingComplete,
    this.onSubmitted,
  }) : super(key: key);

  @override
  State<BaseTextFieldWidget> createState() => _BaseTextFieldWidgetState();
}

class _BaseTextFieldWidgetState extends State<BaseTextFieldWidget> {
  //
  bool _showClear = false;

  //
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = widget.controller ?? TextEditingController();
  }

  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: BorderRadius.circular(widget.circular ?? 0),
      child: Row(
        children: [
          Expanded(
            child: TextField(
              focusNode: widget.focusNode,
              controller: _controller,
              decoration: BaseTextField.defaultDecoration.copyWith(
                hintText: widget.hintText,
                hintStyle: widget.hintStyle,
                fillColor: widget.fillColor,
                suffixIcon: widget.suffixIcon,
                contentPadding: widget.contentPadding,
              ),
              keyboardType: widget.keyboardType,
              style: widget.style ??
                  TextStyle(
                    color: BaseColor.c_1D1F1E,
                    fontSize: 16.sp,
                  ),
              enabled: widget.enabled,
              maxLines: widget.maxLines is int
                  ? widget.maxLines
                  : widget.linesNull
                      ? null
                      : widget.enabled == true
                          ? widget.minLines != null
                              ? max(3, widget.minLines ?? 1)
                              : widget.obscureText == true
                                  ? 1
                                  : 3
                          : widget.obscureText == true
                              ? 1
                              : null,
              // 实现高度随内容自适应
              minLines: widget.linesNull ? null : widget.minLines ?? 1,
              maxLength: widget.maxLength,
              inputFormatters: [
                LengthLimitingTextInputFormatter(widget.lengthLimit ?? 200),
                ...?widget.inputFormatters,
              ],
              obscureText: widget.obscureText ?? false,
              onChanged: _onChanged,
              textAlign: widget.textAlign,
              autofocus: widget.autofocus,
              textInputAction: widget.textInputAction,
              onEditingComplete: widget.onEditingComplete,
              onSubmitted: widget.onSubmitted,
            ),
          ),
          widget.canClear && _showClear
              ? GestureDetector(
                  onTap: () {
                    setState(() {
                      _controller.text = '';
                    });
                    _onChanged('');
                  },
                  child: Container(
                    color: Colors.transparent,
                    padding: const EdgeInsets.all(2),
                    child: widget.clearWidget ??
                        Icon(
                          Icons.close,
                          color: widget.clearColor ?? BaseColor.hint,
                          size: 18,
                        ),
                  ),
                )
              : const SizedBox.shrink(),
        ],
      ),
    );
  }

  void _onChanged(value) {
    setState(
      () => _showClear = _controller.text.isNotEmpty == true,
    );
    widget.onChanged?.call(value);
  }
}
