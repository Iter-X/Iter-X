/*
 * @Description: 图卡选择页
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-18 00:30:03
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-03-18 00:36:13
 */

import 'package:flutter/material.dart';
import 'package:client/common/utils/color.dart';

class CardSelectionPage extends StatefulWidget {
  const CardSelectionPage({super.key});

  @override
  State<CardSelectionPage> createState() => _CardSelectionPageState();
}

class _CardSelectionPageState extends State<CardSelectionPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('选择目的国家'),
      ),
      body: const Center(
        child: Text('图卡选择'),
      ),
    );
  }
}
