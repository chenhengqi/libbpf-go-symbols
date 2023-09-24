package main

import (
	"log"

	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/rlimit"
)

func init() {
	err := rlimit.RemoveMemlock()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	vmlinuxBTF, err := btf.LoadKernelSpec()
	if err != nil {
		log.Fatal(err)
	}

	var types []btf.Type
	iter := vmlinuxBTF.Iterate()
	for iter.Next() {
		intType, ok := iter.Type.(*btf.Int)
		if ok {
			types = append(types, intType)
		}
	}

	intTypes, err := btf.NewBuilder(types)
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	result, err := intTypes.Marshal(data, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(result))
}
