import 'dart:ui' as ui;

import 'package:amap_flutter_base/amap_flutter_base.dart';
import 'package:amap_flutter_map/amap_flutter_map.dart';
import 'package:client/app/constants.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:provider/provider.dart';

class TripMapPage extends StatefulWidget {
  final String tripId;

  const TripMapPage({
    super.key,
    required this.tripId,
  });

  @override
  State<TripMapPage> createState() => _TripMapPageState();
}

class _TripMapPageState extends State<TripMapPage> {
  // 高德地图配置
  late final AMapApiKey amapApiKey;

  CustomStyleOptions? _customStyleOptions;

  @override
  void initState() {
    super.initState();
    amapApiKey = AMapApiKey(
      androidKey: dotenv.env['AMAP_ANDROID_KEY'] ?? '',
      iosKey: dotenv.env['AMAP_IOS_KEY'] ?? '',
    );
    _loadCustomStyle();
  }

  Future<void> _loadCustomStyle() async {
    _customStyleOptions ??= CustomStyleOptions(false);

    try {
      // 加载自定义地图样式文件
      final styleByteData =
          await rootBundle.load('assets/maps/amap/style.data');
      final styleExtraByteData =
          await rootBundle.load('assets/maps/amap/style_extra.data');

      setState(() {
        _customStyleOptions!.styleData = styleByteData.buffer.asUint8List();
        _customStyleOptions!.styleExtraData =
            styleExtraByteData.buffer.asUint8List();
        _customStyleOptions!.enabled = true;
      });
    } catch (e) {
      debugPrint('Failed to load custom map style: $e');
    }
  }

  Future<BitmapDescriptor> _createNumberedMarker(int number) async {
    final ui.PictureRecorder recorder = ui.PictureRecorder();
    final Canvas canvas = Canvas(recorder);
    final Size size = const Size(80, 80);

    // Draw circle background
    final paint = Paint()
      ..color = AppColor.highlight
      ..style = PaintingStyle.fill;
    canvas.drawCircle(
        Offset(size.width / 2, size.height / 2), size.width / 2, paint);

    // Draw number
    final textPainter = TextPainter(
      text: TextSpan(
        text: number.toString(),
        style: const TextStyle(
          color: AppColor.secondary,
          fontSize: 42,
          fontWeight: FontWeight.bold,
        ),
      ),
      textDirection: TextDirection.ltr,
    );
    textPainter.layout();
    textPainter.paint(
      canvas,
      Offset(
        (size.width - textPainter.width) / 2,
        (size.height - textPainter.height) / 2,
      ),
    );

    final ui.Image image = await recorder.endRecording().toImage(
          size.width.toInt(),
          size.height.toInt(),
        );
    final ByteData? byteData =
        await image.toByteData(format: ui.ImageByteFormat.png);
    final Uint8List uint8List = byteData!.buffer.asUint8List();

    return BitmapDescriptor.fromBytes(uint8List);
  }

  @override
  Widget build(BuildContext context) {
    return Consumer<TripService>(
      builder: (context, service, child) {
        if (service.isLoading) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        }

        final trip = service.trip;
        if (trip == null) {
          return const Center(
            child: Text('暂无行程数据'),
          );
        }

        // 获取所有POI点的位置
        final locations = trip.dailyTrips.expand((dailyTrip) {
          return dailyTrip.dailyItineraries.map((itinerary) {
            return LatLng(
              itinerary.poi.latitude,
              itinerary.poi.longitude,
            );
          });
        }).toList();

        return FutureBuilder<List<Marker>>(
          future: Future.wait(
            List.generate(
              locations.length,
              (index) async {
                final marker = await _createNumberedMarker(index + 1);
                return Marker(
                  position: locations[index],
                  icon: marker,
                  anchor: const Offset(0.5, 0.5),
                );
              },
            ),
          ),
          builder: (context, snapshot) {
            if (!snapshot.hasData) {
              return const Center(child: CircularProgressIndicator());
            }

            return AMapWidget(
              apiKey: amapApiKey,
              privacyStatement: AMapConfig.privacyStatement,
              initialCameraPosition: CameraPosition(
                target: locations.isNotEmpty
                    ? locations.first
                    : const LatLng(39.909187, 116.397451),
                zoom: 10,
              ),
              markers: snapshot.data!.toSet(),
              polylines: _buildPolylines(trip),
              mapType: MapType.normal,
              customStyleOptions: _customStyleOptions,
            );
          },
        );
      },
    );
  }

  Set<Polyline> _buildPolylines(Trip trip) {
    final polylines = <Polyline>{};

    // 为每一天的行程创建折线
    for (final dailyTrip in trip.dailyTrips) {
      if (dailyTrip.dailyItineraries.length < 2) continue;

      final points = dailyTrip.dailyItineraries.map((itinerary) {
        return LatLng(
          itinerary.poi.latitude,
          itinerary.poi.longitude,
        );
      }).toList();

      polylines.add(
        Polyline(
          points: points,
          color: AppColor.primary,
          width: 3,
        ),
      );
    }

    return polylines;
  }
}
