package cc.towerdefence.api.utils;

import com.google.protobuf.Timestamp;

import java.time.Instant;

public class GrpcTimestampConverter {

    public static Timestamp convert(Instant instant) {
        return Timestamp.newBuilder()
                .setSeconds(instant.getEpochSecond())
                .setNanos(instant.getNano())
                .build();
    }

    public static Instant reverse(Timestamp timestamp) {
        return Instant.ofEpochSecond(timestamp.getSeconds(), timestamp.getNanos());
    }

    public static Timestamp now() {
        return convert(Instant.now());
    }

    public static Timestamp convertMillis(long millis) {
        return Timestamp.newBuilder()
                .setSeconds(millis / 1000)
                .setNanos((int) ((millis % 1000) * 1000000))
                .build();
    }

    public static Timestamp convertNanos(long nanos) {
        return Timestamp.newBuilder()
                .setSeconds(nanos / 1000000000)
                .setNanos((int) (nanos % 1000000000))
                .build();
    }
}
