#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(xcod status 2>&1)
    if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(xcod status 2>&1)
if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
  wait
fi

PASSWORD="12345678"
GAS_PRICES="0.025uxco"
CHAIN_ID="pandora-4"
FEE=$(yes $PASSWORD | xcod keys show fee -a)
RESERVE_OUT=$(yes $PASSWORD | xcod keys show reserveOut -a)

xcod_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  xcod tx $cmd \
    --gas-prices="$GAS_PRICES" \
    --chain-id="$CHAIN_ID" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

xcod_q() {
  xcod q "$@" --output=json | jq .
}

BOND_DID="did:xco:U7GK8p8rVhJMKhBVRCJJ8c"
#BOND_DID_FULL='{
#  "did":"did:xco:U7GK8p8rVhJMKhBVRCJJ8c",
#  "verifyKey":"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW",
#  "encryptionPublicKey":"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m",
#  "secret":{
#    "seed":"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053",
#    "signKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC",
#    "encryptionPrivateKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC"
#  }
#}'

MIGUEL_ADDR="xco1acltgu0kwgnuqdgewracms3nhz8c6n2grk0uz0"
FRANCESCO_ADDR="xco1zyaz6rkpxa9mdlzazc9uuch4hqc7l5eatsunes"
SHAUN_ADDR="xco10uqnjz60h3lkxxsgnlvxql3yfpkgv6gc6wuz4c"
MIGUEL_DID_FULL='{
  "did":"did:xco:4XJLBfGtWSGKSz4BeRxdun",
  "verifyKey":"A84TX+JIkb1s1WUj/Uxt4P7N9K/yWg4eUQIzjMmdgOO0",
  "encryptionPublicKey":"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS",
  "secret":{
    "seed":"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180",
    "signKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh",
    "encryptionPrivateKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh"
  }
}'
FRANCESCO_DID="did:xco:UKzkhVSHc3qEFva5EY2XHt"
FRANCESCO_DID_FULL='{
  "did":"did:xco:UKzkhVSHc3qEFva5EY2XHt",
  "verifyKey":"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej",
  "encryptionPublicKey":"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si",
  "secret":{
    "seed":"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de",
    "signKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM",
    "encryptionPrivateKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM"
  }
}'
SHAUN_DID_FULL='{
  "did":"did:xco:U4tSpzzv91HHqWW1YmFkHJ",
  "verifyKey":"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG",
  "encryptionPublicKey":"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i",
  "secret":{
    "seed":"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e",
    "signKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR",
    "encryptionPrivateKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR"
  }
}'

# Ledger DIDs
echo "Ledgering DID 1/3..."
xcod_tx did add-did-doc "$MIGUEL_DID_FULL" 
echo "Ledgering DID 2/3..."
xcod_tx did add-did-doc "$FRANCESCO_DID_FULL" 
echo "Ledgering DID 3/3..."
xcod_tx did add-did-doc "$SHAUN_DID_FULL" 

# d0 := 500.0   // initial raise (reserve)
# p0 := 0.01    // initial price (reserve per token)
# theta := 0.4  // initial allocation (percentage)
# kappa := 3.0  // degrees of polynomial (i.e. x^2, x^4, x^6)

# R0 = 300              // initial reserve (1-theta)*d0
# S0 = 50000            // initial supply
# V0 = 416666666666.667 // invariant

echo "Creating bond..."
xcod_tx bonds create-bond \
  --token=abc \
  --name="A B C" \
  --description="Description about A B C" \
  --function-type=augmented_function \
  --function-parameters="d0:500.0,p0:0.01,theta:0.4,kappa:3.0" \
  --reserve-tokens=res \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE" \
  --reserve-withdrawal-address="$RESERVE_OUT" \
  --max-supply=1000000abc \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --allow-sells \
  --alpha-bond \
  --batch-blocks=1 \
  --outcome-payment="100000" \
  --bond-did="$BOND_DID" \
  --creator-did="$MIGUEL_DID_FULL" \
  --controller-did="$FRANCESCO_DID"
echo "Created bond..."
xcod_q bonds bond "$BOND_DID"

echo "Miguel buys 20000abc..."
xcod_tx bonds buy 20000abc 100000res "$BOND_DID" "$MIGUEL_DID_FULL" 
echo "Miguel's account..."
xcod_q bank balances "$MIGUEL_ADDR"

echo "Francesco buys 20000abc..."
xcod_tx bonds buy 20000abc 100000res "$BOND_DID" "$FRANCESCO_DID_FULL" 
echo "Francesco's account..."
xcod_q bank balances "$FRANCESCO_ADDR"

echo "Shaun cannot buy 10001abc..."
xcod_tx bonds buy 10001abc 100000res "$BOND_DID" "$SHAUN_DID_FULL" 
echo "Shaun cannot sell anything..."
xcod_tx bonds sell 10000abc "$BOND_DID" "$SHAUN_DID_FULL" 
echo "Shaun can buy 10000abc..."
xcod_tx bonds buy 10000abc 100000res "$BOND_DID" "$SHAUN_DID_FULL" 
echo "Shaun's account..."
xcod_q bank balances "$SHAUN_ADDR"

echo "Bond state is now open..."  # since 50000 (S0) reached
xcod_q bonds bond "$BOND_DID"

echo "Current price is 0.018..."
xcod_q bonds current-price "$BOND_DID"

echo "Changing public alpha 0.5->0.51..."
NEW_ALPHA="0.51"
xcod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL" 
echo "Current price is now approx 0.011..."
xcod_q bonds current-price "$BOND_DID"

echo "Changing public alpha 0.51->0.4..."
NEW_ALPHA="0.4"
xcod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL" 
echo "Current price is still approx 0.011..."
xcod_q bonds current-price "$BOND_DID"

echo "Cannot change public alpha 0.5->0.6..."
NEW_ALPHA="0.6"
xcod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL" 

echo "Miguel sells 20000abc..."
xcod_tx bonds sell 20000abc "$BOND_DID" "$MIGUEL_DID_FULL" 
echo "Miguel's account..."
xcod_q bank balances "$MIGUEL_ADDR"

echo "Francesco makes outcome payment of 50000000 [1]..."
xcod_tx bonds make-outcome-payment "$BOND_DID" "50000000" "$FRANCESCO_DID_FULL" 
echo "Francesco makes outcome payment of 100000000 [2]..."
xcod_tx bonds make-outcome-payment "$BOND_DID" "100000000" "$FRANCESCO_DID_FULL" 
echo "Francesco makes outcome payment of 150000000 [3]..."
xcod_tx bonds make-outcome-payment "$BOND_DID" "150000000" "$FRANCESCO_DID_FULL" 
echo "Francesco's account..."
xcod_q bank balances "$FRANCESCO_ADDR"
echo "Bond outcome payment reserve is now 300000000..."
xcod_q bonds bond "$BOND_DID"

echo "Francesco updates the bond state to SETTLE"
xcod_tx bonds update-bond-state "SETTLE" "$BOND_DID" "$FRANCESCO_DID_FULL"
echo "Bond outcome payment reserve is now empty (moved to main reserve)..."
xcod_q bonds bond "$BOND_DID"

echo "Francesco withdraws share..."
xcod_tx bonds withdraw-share "$BOND_DID" "$FRANCESCO_DID_FULL" 
echo "Francesco's account..."
xcod_q bank balances "$FRANCESCO_ADDR"

echo "Shaun withdraws share..."
xcod_tx bonds withdraw-share "$BOND_DID" "$SHAUN_DID_FULL" 
echo "Shaun's account..."
xcod_q bank balances "$SHAUN_ADDR"

echo "Bond reserve is now empty and supply is 0..."
xcod_q bonds bond "$BOND_DID"