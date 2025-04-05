import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:shimmer_animation/shimmer_animation.dart';

class PoiSkeleton extends StatelessWidget {
  const PoiSkeleton({super.key});

  Widget _buildPoiItemSkeleton() {
    return Container(
      padding: EdgeInsets.only(top: 2.h),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Shimmer(
            duration: const Duration(seconds: 2),
            child: Container(
              width: 142.w,
              height: 142.w,
              decoration: BoxDecoration(
                color: Colors.grey[200],
                borderRadius: BorderRadius.circular(8),
              ),
            ),
          ),
          Expanded(
            child: Container(
              padding: EdgeInsets.only(left: 10.w, right: 10.w),
              height: 142.w,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Shimmer(
                    duration: const Duration(seconds: 2),
                    child: Container(
                      width: 200.w,
                      height: 20.h,
                      decoration: BoxDecoration(
                        color: Colors.grey[200],
                        borderRadius: BorderRadius.circular(4),
                      ),
                    ),
                  ),
                  SizedBox(height: 8.h),
                  Shimmer(
                    duration: const Duration(seconds: 2),
                    child: Container(
                      width: 150.w,
                      height: 16.h,
                      decoration: BoxDecoration(
                        color: Colors.grey[200],
                        borderRadius: BorderRadius.circular(4),
                      ),
                    ),
                  ),
                  SizedBox(height: 12.h),
                  Row(
                    children: List.generate(5, (index) {
                      return Padding(
                        padding: EdgeInsets.only(right: 4.w),
                        child: Shimmer(
                          duration: const Duration(seconds: 2),
                          child: Container(
                            width: 18.w,
                            height: 18.w,
                            decoration: BoxDecoration(
                              color: Colors.grey[200],
                              borderRadius: BorderRadius.circular(4),
                            ),
                          ),
                        ),
                      );
                    }),
                  ),
                  SizedBox(height: 12.h),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Shimmer(
                        duration: const Duration(seconds: 2),
                        child: Container(
                          width: 80.w,
                          height: 14.h,
                          decoration: BoxDecoration(
                            color: Colors.grey[200],
                            borderRadius: BorderRadius.circular(4),
                          ),
                        ),
                      ),
                      Shimmer(
                        duration: const Duration(seconds: 2),
                        child: Container(
                          width: 60.w,
                          height: 14.h,
                          decoration: BoxDecoration(
                            color: Colors.grey[200],
                            borderRadius: BorderRadius.circular(4),
                          ),
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      padding: EdgeInsets.all(0),
      itemCount: 3,
      itemBuilder: (context, index) => _buildPoiItemSkeleton(),
    );
  }
} 