# =============================================================================
#  Docker Container for JSON Schema Validator (v1.0.0+202112022356)
# =============================================================================

# -----------------------------------------------------------------------------
#  Build Stage
# -----------------------------------------------------------------------------
FROM golang:alpine AS builder

COPY ./validator /workspaces/validator

WORKDIR /workspaces/validator

RUN \
    apk add --no-cache build-base && \
    go mod download && \
    go test ./... && \
    go install . && \
    # Smoke test
    validator 2>&1 | grep "missing arg"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM alpine

COPY --from=builder /go/bin/validator /usr/local/bin/validator

# Smoke test
RUN validator 2>&1 | grep "missing arg"

WORKDIR /workspaces
