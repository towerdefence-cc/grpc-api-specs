package cc.towerdefence.api.utils.spring.exception;

import io.grpc.Status;
import org.jetbrains.annotations.NotNull;

public class ResourceNotFoundException extends GenericGrpcException {

    public ResourceNotFoundException(@NotNull String description) {
        super(Status.NOT_FOUND, description);
    }
}
