import 'package:duo_client/pb/friend.pb.dart';
import 'package:duo_client/provider/friend_provider.dart';
import 'package:duo_client/provider/storage_provider.dart';
import 'package:duo_client/utils/constants.dart';
import 'package:duo_client/widgets/leaderboard_list_tile.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';

class LeaderboardScreen extends ConsumerStatefulWidget {
  const LeaderboardScreen({super.key});

  @override
  ConsumerState<LeaderboardScreen> createState() => _LeaderboardScreenState();
}

class _LeaderboardScreenState extends ConsumerState<LeaderboardScreen> {
  @override
  Widget build(BuildContext context) {
    StorageProvider _storageProvider = ref.read(storageProvider);
    var self_user = Friend(
        name: _storageProvider.username, score: 100, state: FriendState.online);

    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: SvgPicture.asset(
            'res/icons/menu.svg',
            colorFilter:
                const ColorFilter.mode(Colors.white54, BlendMode.srcIn),
          ),
          onPressed: () {},
        ),
        title: const Text('Leaderboard'),
        actions: [
          IconButton(
            icon: SvgPicture.asset(
              'res/icons/refresh.svg',
              colorFilter: const ColorFilter.mode(
                Colors.white54,
                BlendMode.srcIn,
              ),
            ),
            onPressed: () {},
          ),
          const SizedBox(width: Constants.defaultPadding / 2),
          IconButton(
            icon: SvgPicture.asset(
              'res/icons/bell.svg',
              colorFilter: const ColorFilter.mode(
                Colors.white54,
                BlendMode.srcIn,
              ),
            ),
            onPressed: () {},
          ),
          const SizedBox(width: Constants.defaultPadding / 2),
        ],
        centerTitle: true,
      ),
      body: SingleChildScrollView(
        child: Consumer(
          builder: (context, ref, child) {
            var friends = [...ref.watch(friendProvider).friends, self_user]
              ..sort(
                (a, b) => b.score.compareTo(a.score),
              );

            return Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const SizedBox(height: Constants.defaultPadding),
                ListView.builder(
                  shrinkWrap: true,
                  itemCount: friends.length,
                  itemBuilder: (context, index) {
                    return LeaderboardListTile(
                      boardPlace: index + 1,
                      friend: friends[index],
                      ownScore: self_user.score,
                      isBetterThanPlayer:
                          friends[index].score > self_user.score,
                      isPlayer: friends[index].name == self_user.name,
                    );
                  },
                ),
              ],
            );
          },
        ),
      ),
    );
  }
}
