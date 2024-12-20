import 'dart:math';

import 'package:flutter/material.dart';
import 'package:duo_client/widgets/playingcard.dart' as duo;
import 'package:flutter_riverpod/flutter_riverpod.dart';

class DUOCardStack extends ConsumerStatefulWidget {
  const DUOCardStack(
      {super.key, required this.cards, this.randomAngles = false, this.onTap});

  final List<duo.PlayingCard> cards;
  final bool randomAngles;
  final Function? onTap;

  @override
  ConsumerState<DUOCardStack> createState() => _DUOCardStackState();
}

class _DUOCardStackState extends ConsumerState<DUOCardStack> {
  double randomOffset() {
    return Random().nextDouble() * 10 * (Random().nextBool() ? 1 : -1);
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: 400,
      height: 400,
      child: Stack(
        children: [
          for (int i = 0; i < widget.cards.length; i++)
            Positioned(
              left: randomOffset(),
              child: Transform.rotate(
                angle: widget.randomAngles ? randomOffset() * 0.03 : 0,
                child: SizedBox(
                  height: MediaQuery.of(context).size.height * 0.4,
                  width: MediaQuery.of(context).size.width * 0.2,
                  child: InkWell(
                    onTap: widget.onTap != null
                        ? () => widget.onTap!(widget.cards[i])
                        : null,
                    child: widget.cards[i],
                  ),
                ),
              ),
            ),
        ],
      ),
    );
  }
}
