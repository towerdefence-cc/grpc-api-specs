package cc.towerdefence.api.utils.resolvers;

import cc.towerdefence.api.model.PlayerProto;
import cc.towerdefence.api.service.McPlayerGrpc;
import cc.towerdefence.api.service.McPlayerProto;
import cc.towerdefence.api.utils.utils.FunctionalFutureCallback;
import com.github.benmanes.caffeine.cache.Cache;
import com.github.benmanes.caffeine.cache.Caffeine;
import com.google.common.util.concurrent.Futures;
import com.google.common.util.concurrent.ListenableFuture;
import io.grpc.Status;
import org.jetbrains.annotations.NotNull;

import java.time.Duration;
import java.util.UUID;
import java.util.concurrent.ForkJoinPool;
import java.util.function.Consumer;
import java.util.function.Function;

@SuppressWarnings("unused")
public class PlayerResolver {
    private static final Cache<String, CachedMcPlayer> USERNAME_TO_PLAYER_CACHE = Caffeine.newBuilder()
            .maximumSize(1000)
            .expireAfterWrite(Duration.ofMinutes(5))
            .build();

    private static McPlayerGrpc.McPlayerFutureStub playerService;
    // The alternative should be used as a server software specific option to avoid calling the mc-player service if they're on the same server.
    private static Function<String, CachedMcPlayer> alternativeUsernameResolver;

    public static void setPlayerService(@NotNull McPlayerGrpc.McPlayerFutureStub playerService, Function<String, CachedMcPlayer> alternativeUsernameResolver) {
        if (PlayerResolver.playerService != null) throw new IllegalStateException("Player service already set");
        PlayerResolver.playerService = playerService;
        PlayerResolver.alternativeUsernameResolver = alternativeUsernameResolver;
    }

    public static void retrievePlayerData(String username, Consumer<CachedMcPlayer> callback, Consumer<Status> errorCallback) {
        String usernameLowercase = username.toLowerCase();

        CachedMcPlayer alternativeOption = alternativeUsernameResolver.apply(usernameLowercase);
        if (alternativeOption != null) {
            callback.accept(alternativeOption);
            return;
        }

        CachedMcPlayer cacheResult = USERNAME_TO_PLAYER_CACHE.getIfPresent(usernameLowercase);
        if (cacheResult != null) callback.accept(cacheResult);

        requestMcPlayer(usernameLowercase, uuid -> {
            USERNAME_TO_PLAYER_CACHE.put(usernameLowercase, uuid);
            callback.accept(uuid);
        }, errorCallback);
    }

    private static void requestMcPlayer(String username, Consumer<CachedMcPlayer> callback, Consumer<Status> errorCallback) {
        ListenableFuture<McPlayerProto.PlayerResponse> playerResponseFuture = playerService
                .getPlayerByUsername(PlayerProto.PlayerUsernameRequest.newBuilder().setUsername(username).build());

        Futures.addCallback(playerResponseFuture, FunctionalFutureCallback.create(
                playerResponse -> callback.accept(new CachedMcPlayer(UUID.fromString(playerResponse.getId()), playerResponse.getCurrentUsername())),
                throwable -> errorCallback.accept(Status.fromThrowable(throwable))
        ), ForkJoinPool.commonPool());
    }

    public record CachedMcPlayer(UUID uuid, String username) {}
}
