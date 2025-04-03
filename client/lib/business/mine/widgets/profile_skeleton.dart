import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:shimmer_animation/shimmer_animation.dart';

class ProfileSkeleton extends StatelessWidget {
  const ProfileSkeleton({super.key});

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: const EdgeInsets.only(bottom: 15),
      children: [
        _buildHeaderSkeleton(),
        SizedBox(height: 20.h),
        _buildStatsRowSkeleton(),
        SizedBox(height: 10.h),
        _buildStatsRowSkeleton(),
        SizedBox(height: 20.h),
        _buildSectionSkeleton(),
        SizedBox(height: 20.h),
        _buildTripsSkeleton(),
      ],
    );
  }

  Widget _buildHeaderSkeleton() {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 20.h, horizontal: 20.w),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Shimmer(
                  duration: const Duration(seconds: 2),
                  child: Container(
                    width: 150.w,
                    height: 35.h,
                    decoration: BoxDecoration(
                      color: Colors.grey[200],
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
                SizedBox(height: 8.h),
                Shimmer(
                  duration: const Duration(seconds: 2),
                  child: Container(
                    width: 120.w,
                    height: 20.h,
                    decoration: BoxDecoration(
                      color: Colors.grey[200],
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
              ],
            ),
          ),
          Shimmer(
            duration: const Duration(seconds: 2),
            child: Container(
              width: 70.w,
              height: 70.w,
              decoration: BoxDecoration(
                color: Colors.grey[200],
                shape: BoxShape.circle,
                border: Border.all(
                  color: AppColor.primary,
                  width: 1,
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildStatsRowSkeleton() {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 15.w),
      child: Row(
        children: [
          Expanded(child: _buildStatCardSkeleton()),
          SizedBox(width: 10.w),
          Expanded(child: _buildStatCardSkeleton()),
        ],
      ),
    );
  }

  Widget _buildStatCardSkeleton() {
    return Container(
      padding: EdgeInsets.all(15.w),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(AppConfig.boxRadius),
      ),
      child: Row(
        children: [
          Shimmer(
            duration: const Duration(seconds: 2),
            child: Container(
              width: 40.w,
              height: 40.w,
              decoration: BoxDecoration(
                color: Colors.grey[200],
                borderRadius: BorderRadius.circular(8),
              ),
            ),
          ),
          SizedBox(width: 10.w),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
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
                SizedBox(height: 4.h),
                Shimmer(
                  duration: const Duration(seconds: 2),
                  child: Container(
                    width: 40.w,
                    height: 18.h,
                    decoration: BoxDecoration(
                      color: Colors.grey[200],
                      borderRadius: BorderRadius.circular(4),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildSectionSkeleton() {
    return Container(
      margin: EdgeInsets.symmetric(horizontal: 15.w),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(AppConfig.boxRadius),
      ),
      child: Padding(
        padding: EdgeInsets.all(15.w),
        child: Row(
          children: [
            Shimmer(
              duration: const Duration(seconds: 2),
              child: Container(
                width: 22.w,
                height: 22.w,
                decoration: BoxDecoration(
                  color: Colors.grey[200],
                  borderRadius: BorderRadius.circular(4),
                ),
              ),
            ),
            SizedBox(width: 5.w),
            Shimmer(
              duration: const Duration(seconds: 2),
              child: Container(
                width: 80.w,
                height: 16.h,
                decoration: BoxDecoration(
                  color: Colors.grey[200],
                  borderRadius: BorderRadius.circular(4),
                ),
              ),
            ),
            const Spacer(),
            Shimmer(
              duration: const Duration(seconds: 2),
              child: Container(
                width: 16.w,
                height: 16.w,
                decoration: BoxDecoration(
                  color: Colors.grey[200],
                  borderRadius: BorderRadius.circular(4),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildTripsSkeleton() {
    return Container(
      margin: EdgeInsets.symmetric(horizontal: 15.w),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(AppConfig.boxRadius),
      ),
      child: Column(
        children: [
          _buildSectionSkeleton(),
          ...List.generate(2, (index) => _buildTripCardSkeleton(index == 0)),
        ],
      ),
    );
  }

  Widget _buildTripCardSkeleton(bool isFirst) {
    return Shimmer(
      duration: const Duration(seconds: 2),
      child: Container(
        height: 160.h,
        margin: EdgeInsets.fromLTRB(15.w, isFirst ? 0 : 10.w, 15.w, 15.w),
        decoration: BoxDecoration(
          color: Colors.grey[200],
          borderRadius: BorderRadius.circular(12),
        ),
      ),
    );
  }
}
