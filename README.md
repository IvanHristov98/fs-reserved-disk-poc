# FS reserved disk size

The code in this repo uses the [Statfs](https://man7.org/linux/man-pages/man2/statfs.2.html) syscall to generate output similar to `df`. It is a PoC for finding the reserved disk space of a mounted disk regardless of the FS. It applies only for *nix OS.

It can be used by running:

```bash
go build main.go

# Example applications within diego-cell instances
# For ext4
./main /var/vcap/data # the arg is a pathname of any file within the mounted fs
fs type 61267 # 61267 is 0xef53 in hex which is ext4
fs block size 4096
total blocks 23607165
free blocks in fs 19499293
free blocks available to unprivileged user 18289361
reserved blocks 1209932
reserved 4955881472 bytes = 4.615524 GiB # 

# For xfs
./main /var/vcap/data/grootfs/store/unprivileged
fs type 1481003842 # 1481003842 is 0x58465342 in hex which is xfs
fs block size 4096
total blocks 19665399
free blocks in fs 18881796
free blocks available to unprivileged user 18881796
reserved blocks 0
reserved 0 bytes = 0.000000 GiB
```