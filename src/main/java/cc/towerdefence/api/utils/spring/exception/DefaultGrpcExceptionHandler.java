package cc.towerdefence.api.utils.spring.exception;

import io.grpc.Status;
import io.grpc.StatusException;
import net.devh.boot.grpc.server.advice.GrpcAdvice;
import net.devh.boot.grpc.server.advice.GrpcExceptionHandler;

@GrpcAdvice
public class DefaultGrpcExceptionHandler {

    @GrpcExceptionHandler
    public StatusException handleResourceAlreadyExistsException(ResourceAlreadyExistsException e) {
        return Status.ALREADY_EXISTS.withDescription(e.getDescription()).withCause(e).asException();
    }
}
