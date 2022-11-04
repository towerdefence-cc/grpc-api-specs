package cc.towerdefence.api.utils.spring.exception;

import io.grpc.Status;
import org.jetbrains.annotations.NotNull;

public class ResourceAlreadyExistsException extends GenericGrpcException {

    public ResourceAlreadyExistsException(@NotNull String description) {
        super(Status.ALREADY_EXISTS, description);
    }
}
