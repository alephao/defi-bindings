//go:generate abigen --abi=lib/mev-inspect-rs/abi/aavecore.json 			--pkg=aave 			--out=aave/AaveCore.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/aavepool.json 			--pkg=aave 			--out=aave/AavePool.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/bpool.json 					--pkg=balancer 	--out=balancer/BPool.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/bproxy.json 				--pkg=balancer 	--out=balancer/BProxy.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/cether.json 				--pkg=compound 	--out=compound/CEther.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/comptroller.json 		--pkg=compound 	--out=compound/Comptroller.go --alias _borrowGuardianPaused=PBorrowGuardianPaused,_mintGuardianPaused=PMintGuardianPaused
//go:generate abigen --abi=lib/mev-inspect-rs/abi/ctoken.json 				--pkg=compound 	--out=compound/CToken.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/curvepool.json 			--pkg=curve 		--out=curve/CurvePool.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/curveregistry.json 	--pkg=curve 		--out=curve/CurveRegistry.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/unipair.json 				--pkg=uniswapv2 --out=uniswapv2/UniPair.go
//go:generate abigen --abi=lib/mev-inspect-rs/abi/unirouterv2.json 		--pkg=uniswapv2 --out=uniswapv2/UniRouterV2.go
package main

func main() {

}
