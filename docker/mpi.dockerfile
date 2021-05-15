FROM python:3.8-slim
RUN apt-get update && \
    apt-get install -q -y openmpi-bin libopenmpi-dev && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/* && \
    pip install mpi4py
