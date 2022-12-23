package cc.towerdefence.api.utils.callback;

import io.grpc.stub.StreamObserver;
import org.jetbrains.annotations.Nullable;

import java.util.function.Consumer;

public class FunctionalStreamObserver {

    public static <T> StreamObserver<T> create(Consumer<T> onNext, Consumer<Throwable> onError, @Nullable Runnable onCompleted) {
        return new StreamObserver<>() {
            @Override
            public void onNext(T value) {
                onNext.accept(value);
            }

            @Override
            public void onError(Throwable t) {
                onError.accept(t);
            }

            @Override
            public void onCompleted() {
                if (onCompleted != null) onCompleted.run();
            }
        };
    }
}
