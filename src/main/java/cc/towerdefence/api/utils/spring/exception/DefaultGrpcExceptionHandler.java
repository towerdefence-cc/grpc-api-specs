package cc.towerdefence.api.utils.spring.exception;

import io.grpc.StatusException;
import net.devh.boot.grpc.server.advice.GrpcAdvice;
import net.devh.boot.grpc.server.advice.GrpcExceptionHandler;

@GrpcAdvice
public class DefaultGrpcExceptionHandler {

    @GrpcExceptionHandler
    public StatusException handleGenericGrpcException(GenericGrpcException ex) {
        return ex.toStatusException();
    }
}
