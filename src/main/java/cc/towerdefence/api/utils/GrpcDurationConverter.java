package cc.towerdefence.api.utils;

import com.google.protobuf.Duration;

public class GrpcDurationConverter {

    public static Duration convert(java.time.Duration duration) {
        return Duration.newBuilder()
                .setSeconds(duration.getSeconds())
                .setNanos(duration.getNano())
                .build();
    }

    public static Duration convertMillis(long millis) {
        return Duration.newBuilder()
                .setSeconds(millis / 1000)
                .setNanos((int) ((millis % 1000) * 1000000))
                .build();
    }

    public static Duration convertNanos(long nanos) {
        return Duration.newBuilder()
                .setSeconds(nanos / 1000000000)
                .setNanos((int) (nanos % 1000000000))
                .build();
    }
}
