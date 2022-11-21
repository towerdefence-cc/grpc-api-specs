package cc.towerdefence.api.utils;

import cc.towerdefence.api.service.FriendGrpc;
import cc.towerdefence.api.service.McPlayerGrpc;
import cc.towerdefence.api.service.McPlayerSecurityGrpc;
import cc.towerdefence.api.service.PermissionServiceGrpc;
import cc.towerdefence.api.service.PlayerTrackerGrpc;
import cc.towerdefence.api.service.PrivateMessageGrpc;
import cc.towerdefence.api.service.ServerDiscoveryGrpc;
import cc.towerdefence.api.service.velocity.VelocityFriendGrpc;
import cc.towerdefence.api.service.velocity.VelocityPrivateMessageGrpc;
import cc.towerdefence.api.service.velocity.VelocityServerGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import lombok.Getter;
import org.jetbrains.annotations.NotNull;

import java.util.Optional;
import java.util.function.Function;

// We directly return the optionals here, so it's functionally identical
@SuppressWarnings("OptionalUsedAsFieldOrParameterType")
public class GrpcStubCollection {
    private static final boolean DEVELOPMENT = System.getenv("KUBERNETES_SERVICE_HOST") == null;

    private static final Function<String, String> HOST = string -> {
        if (DEVELOPMENT) return "localhost";
        return string + ".towerdefence.svc";
    };

    private static final @NotNull @Getter Optional<FriendGrpc.FriendFutureStub> friendService;
    private static final @NotNull @Getter Optional<McPlayerGrpc.McPlayerFutureStub> playerService;
    private static final @NotNull @Getter Optional<McPlayerSecurityGrpc.McPlayerSecurityFutureStub> playerSecurityService;
    private static final @NotNull @Getter Optional<PermissionServiceGrpc.PermissionServiceFutureStub> permissionService;
    private static final @NotNull @Getter Optional<PlayerTrackerGrpc.PlayerTrackerFutureStub> playerTrackerService;
    private static final @NotNull @Getter Optional<PrivateMessageGrpc.PrivateMessageFutureStub> privateMessageService;
    private static final @NotNull @Getter Optional<ServerDiscoveryGrpc.ServerDiscoveryFutureStub> serverDiscoveryService;

    static {
        friendService = createChannel("friend-manager").map(FriendGrpc::newFutureStub);
        playerService = createChannel("mc-player").map(McPlayerGrpc::newFutureStub);
        playerSecurityService = createChannel("mc-player-security").map(McPlayerSecurityGrpc::newFutureStub);
        permissionService = createChannel("permission").map(PermissionServiceGrpc::newFutureStub);
        playerTrackerService = createChannel("player-tracker").map(PlayerTrackerGrpc::newFutureStub);
        privateMessageService = createChannel("private-message").map(PrivateMessageGrpc::newFutureStub);
        serverDiscoveryService = createChannel("server-discovery").map(ServerDiscoveryGrpc::newFutureStub);
    }

    /**
     * @param name Service name in the format "svc-name"
     * @return Optional of a ManagedChannel, empty if service is not enabled.
     */
    private static Optional<ManagedChannel> createChannel(String name) {
        String portEnvVarName = name.toUpperCase().replace('-', '_') + "_PORT";
        String portEnvVarValue = System.getenv(portEnvVarName);

        if (portEnvVarValue == null) return Optional.empty();
        int port = Integer.parseInt(System.getenv(portEnvVarName));

        return Optional.of(ManagedChannelBuilder.forAddress(HOST.apply(name), port)
                .defaultLoadBalancingPolicy("round_robin")
                .usePlaintext()
                .build());
    }
}
