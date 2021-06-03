# FS reserved space checker

The code in this repo uses the `Statfs` syscall to generate output similar to `df`. It is a PoC for finding the reserved disk space of particular fs mounts on particular hosts.