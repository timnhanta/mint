package types

const (
	// ModuleName defines the module name
	ModuleName = "custommint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_custommint"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
