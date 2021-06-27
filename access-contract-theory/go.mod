module practices/access-contract-theory

go 1.16
require (
	ethtool v0.0.0
    github.com/ethereum/go-ethereum v1.10.4
    practices/contract/warpper/eztoken v0.0.0
)
replace (
	practices/contract/warpper/eztoken v0.0.0 => ../../practices/contract/warpper/eztoken
	ethtool v0.0.0 => ../../ethtool
)