package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	path := os.Args[1]

	statfs := &syscall.Statfs_t{}
	if err := syscall.Statfs(path, statfs); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("fs type %d \n", statfs.Type)
	fmt.Printf("fs block size %d\n", statfs.Bsize)
	fmt.Printf("total blocks %d\n", statfs.Blocks)
	fmt.Printf("free blocks in fs %d\n", statfs.Bfree)
	fmt.Printf("free blocks available to unprivileged user %d\n", statfs.Bavail)

	reservedBlocks := statfs.Bfree - statfs.Bavail
	fmt.Printf("reserved blocks %d\n", reservedBlocks)

	reservedBytes := reservedBlocks * uint64(statfs.Bsize)
	reservedGiB := bytesToGiB(reservedBytes)
	fmt.Printf("reserved %d bytes = %f GiB\n", reservedBytes, reservedGiB)
}

func bytesToGiB(bytes uint64) float64 {
	return float64(bytes) / (1024.0 * 1024.0 * 1024.0)
}
