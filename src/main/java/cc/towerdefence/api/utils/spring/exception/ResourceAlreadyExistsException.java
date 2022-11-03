package cc.towerdefence.api.utils.spring.exception;

import org.jetbrains.annotations.NotNull;

public class ResourceAlreadyExistsException extends RuntimeException {
    private final @NotNull String description;

    public ResourceAlreadyExistsException(@NotNull String description) {
        this.description = description;
    }

    public @NotNull String getDescription() {
        return this.description;
    }

    @Override
    public String getMessage() {
        return this.getDescription();
    }
}
