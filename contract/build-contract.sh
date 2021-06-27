#!/usr/bin/env bash

set -e
set -o pipefail

srcDir=contract
abiDir=$srcDir/build/

if [ ! -e abiDir ];then
  mkdir -p $abiDir
fi

# shellcheck disable=SC2045
for file in `ls $srcDir/*.sol`; do
  target=$(basename $file .sol)
  pkgDir=$srcDir/wrapper/${target,,}

  if [ ! -e pkgDir ];then
    mkdir -p $pkgDir
  fi

  echo "Compiling Solidity file ${target}.sol"
  solc --bin --abi --optimize --overwrite \
          --allow-paths "$(pwd)" \
          $file -o $abiDir

  echo "Generating Go file ${target}.go"
  abigen --bin=${abiDir}/${target}.bin \
  --abi=${abiDir}/${target}.abi \
  --pkg=${target,,} \
  --out=${pkgDir}/${target}.go

  echo "Complete"
done
