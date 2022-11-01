package cc.towerdefence.api.utils.resolvers;

import cc.towerdefence.api.service.McPlayerGrpc;
import cc.towerdefence.api.service.McPlayerProto;
import cc.towerdefence.api.utils.utils.FunctionalFutureCallback;
import com.github.benmanes.caffeine.cache.Cache;
import com.github.benmanes.caffeine.cache.Caffeine;
import com.google.common.util.concurrent.Futures;
import com.google.common.util.concurrent.ListenableFuture;
import io.grpc.Status;

import java.time.Duration;
import java.util.UUID;
import java.util.concurrent.ForkJoinPool;
import java.util.function.Consumer;

public class PlayerResolver {
    private static final Cache<String, CachedMcPlayer> USERNAME_TO_PLAYER_CACHE = Caffeine.newBuilder()
            .maximumSize(1000)
            .expireAfterWrite(Duration.ofMinutes(5))
            .build();

    private static McPlayerGrpc.McPlayerFutureStub playerService;

    public static void retrievePlayerUuid(String username, Consumer<CachedMcPlayer> callback, Consumer<Status> errorCallback) {
        CachedMcPlayer cacheResult = USERNAME_TO_PLAYER_CACHE.getIfPresent(username);
        if (cacheResult != null) callback.accept(cacheResult);

        requestMcPlayer(username, uuid -> {
            USERNAME_TO_PLAYER_CACHE.put(username, uuid);
            callback.accept(uuid);
        }, errorCallback);
    }

    private static void requestMcPlayer(String username, Consumer<CachedMcPlayer> callback, Consumer<Status> errorCallback) {
        ListenableFuture<McPlayerProto.PlayerResponse> playerResponseFuture = playerService
                .getPlayerByUsername(McPlayerProto.PlayerUsernameRequest.newBuilder().setUsername(username).build());

        Futures.addCallback(playerResponseFuture, FunctionalFutureCallback.create(
                playerResponse -> callback.accept(new CachedMcPlayer(UUID.fromString(playerResponse.getId()), playerResponse.getCurrentUsername())),
                throwable -> errorCallback.accept(Status.fromThrowable(throwable))
        ), ForkJoinPool.commonPool());
    }

    public record CachedMcPlayer(UUID uuid, String username) {}
}
