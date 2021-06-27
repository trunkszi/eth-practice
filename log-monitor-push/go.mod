module practices/log-monitor-push

go 1.16

require (
	ethtool v0.0.0
	github.com/ethereum/go-ethereum v1.10.4
	practices/contract/warpper/eztoken v0.0.0

)

replace (
	ethtool v0.0.0 => ../../ethtool
	practices/contract/warpper/eztoken v0.0.0 => ../../practices/contract/warpper/eztoken
)
