module practices/access-contract-bind

go 1.16

require (
	ethtool v0.0.0
	github.com/ethereum/go-ethereum v1.10.4
	contract/warpper/eztoken v0.0.0
)

replace (
	contract/warpper/eztoken v0.0.0 => ../contract/warpper/eztoken
	ethtool v0.0.0 => ../../ethtool
)
