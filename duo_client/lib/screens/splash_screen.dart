import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';

class SplashScreen<T> extends StatefulWidget {
  static const route = '/splash_in_your_face';
  final Future<T> Function() onLoading;
  final void Function(T value) onLoadingComplete;
  const SplashScreen({
    required this.onLoading,
    required this.onLoadingComplete,
    super.key,
  });

  @override
  State<SplashScreen> createState() => _SplasScreenState<T>();
}

class _SplasScreenState<T> extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) async {
      T value = await widget.onLoading();
      widget.onLoadingComplete(value);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Image.asset(
              'assets/logo.png',
              width: 100.0,
              height: 100.0,
            ),
            const SizedBox(height: 40.0),
            SpinKitFadingCube(
              color: Theme.of(context).colorScheme.primary,
              size: 50.0,
            ),
          ],
        ),
      ),
    );
  }
}
