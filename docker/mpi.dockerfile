FROM debian:10.9-slim
RUN apt-get update && \
    apt-get install -q -y --no-install-recommends openmpi-bin && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*
