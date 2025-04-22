import 'dart:convert';
import 'dart:ui' as ui;

import 'package:amap_flutter_base/amap_flutter_base.dart';
import 'package:amap_flutter_map/amap_flutter_map.dart';
import 'package:client/app/constants.dart';
import 'package:client/business/trip/entity/trip.dart';
import 'package:client/business/trip/service/trip_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
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
  late final AMapApiKey amapApiKey;

  CustomStyleOptions? _customStyleOptions;

  final List<Color> _dailyColors = [
    AppColor.highlight,
    // AppColor.primary,
    // AppColor.secondary,
    // AppColor.bg,
  ];

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

  Future<BitmapDescriptor> _createNumberedMarker(
      int dayNumber, int sequenceNumber) async {
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
        text: '$dayNumber.$sequenceNumber',
        style: const TextStyle(
          color: AppColor.secondary,
          fontSize: 32,
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

  Future<List<LatLng>> _getRoutePoints(LatLng start, LatLng end) async {
    try {
      final response = await http.get(
        Uri.parse(
          'https://restapi.amap.com/v3/direction/driving?origin=${start.longitude},${start.latitude}&destination=${end.longitude},${end.latitude}&key=${dotenv.env['AMAP_WEB_KEY']}',
        ),
      );

      if (response.statusCode == 200) {
        final data = json.decode(response.body);
        if (data['status'] == '1' && data['route'] != null) {
          final path = data['route']['paths'][0];
          final points = path['steps'].map<LatLng>((step) {
            final location = step['polyline'].split(';')[0].split(',');
            return LatLng(
              double.parse(location[1]),
              double.parse(location[0]),
            );
          }).toList();
          return points;
        }
      }
    } catch (e) {
      debugPrint('Failed to get route points: $e');
    }
    // Fallback to direct line if route planning fails
    return [start, end];
  }

  Future<Set<Polyline>> _buildPolylines(Trip trip) async {
    final polylines = <Polyline>{};

    // Create polylines for each day's trip
    for (var dayIndex = 0; dayIndex < trip.dailyTrips.length; dayIndex++) {
      final dailyTrip = trip.dailyTrips[dayIndex];
      if (dailyTrip.dailyItineraries.length < 2) continue;

      final points = <LatLng>[];
      for (var i = 0; i < dailyTrip.dailyItineraries.length - 1; i++) {
        final start = LatLng(
          dailyTrip.dailyItineraries[i].poi.latitude,
          dailyTrip.dailyItineraries[i].poi.longitude,
        );
        final end = LatLng(
          dailyTrip.dailyItineraries[i + 1].poi.latitude,
          dailyTrip.dailyItineraries[i + 1].poi.longitude,
        );
        final routePoints = await _getRoutePoints(start, end);
        points.addAll(routePoints);
      }

      polylines.add(
        Polyline(
          points: points,
          color: _dailyColors[dayIndex % _dailyColors.length],
          width: 5,
          capType: CapType.arrow,
          alpha: 0.8,
        ),
      );
    }

    return polylines;
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
            trip.dailyTrips.asMap().entries.expand((entry) {
              final dayNumber = entry.key + 1;
              return List.generate(
                entry.value.dailyItineraries.length,
                (index) async {
                  final marker =
                      await _createNumberedMarker(dayNumber, index + 1);
                  return Marker(
                    position: LatLng(
                      entry.value.dailyItineraries[index].poi.latitude,
                      entry.value.dailyItineraries[index].poi.longitude,
                    ),
                    icon: marker,
                    anchor: const Offset(0.5, 0.5),
                  );
                },
              );
            }).toList(),
          ),
          builder: (context, snapshot) {
            if (!snapshot.hasData) {
              return const Center(child: CircularProgressIndicator());
            }

            return FutureBuilder<Set<Polyline>>(
              future: _buildPolylines(trip),
              builder: (context, polylineSnapshot) {
                if (!polylineSnapshot.hasData) {
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
                  polylines: polylineSnapshot.data!,
                  mapType: MapType.normal,
                  customStyleOptions: _customStyleOptions,
                );
              },
            );
          },
        );
      },
    );
  }
}
