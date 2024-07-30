package types

const (
	// ModuleName defines the module name
	ModuleName = "blogchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blogchain"
)

var (
	ParamsKey = []byte("p_blogchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
