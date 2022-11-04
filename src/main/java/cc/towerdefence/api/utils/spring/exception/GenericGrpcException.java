package cc.towerdefence.api.utils.spring.exception;

import io.grpc.Metadata;
import io.grpc.Status;
import io.grpc.StatusException;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

public abstract class GenericGrpcException extends RuntimeException {
    private final @NotNull Status status;
    private final @NotNull String description;
    private @Nullable Metadata metadata;

    protected GenericGrpcException(@NotNull Status status, @NotNull String description) {
        this.status = status;
        this.description = description;
    }

    public @NotNull Status getStatus() {
        return status;
    }

    public @NotNull String getDescription() {
        return description;
    }

    public @NotNull GenericGrpcException withMetadata(@Nullable Metadata metadata) {
        this.metadata = metadata;
        return this;
    }

    public StatusException toStatusException() {
        return this.status.withDescription(this.description).withCause(this).asException(this.metadata);
    }
}
