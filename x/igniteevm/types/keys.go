package types

const (
	// ModuleName defines the module name
	ModuleName = "igniteevm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_igniteevm"
)

var (
	ParamsKey = []byte("p_igniteevm")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
