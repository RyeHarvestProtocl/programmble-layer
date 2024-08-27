# RPC Methods Specification

## Send transaction

### simple transfer

transfer_sig is used when native btc transfers / native runes token transfers are triggered. Since theres no contract interaction, "calldata_sig" should be leave empty. user should sign the transfer transaction using their private key, and if the transfer is considered valid by RyeHarvest's programmable layer's node, it will be fully signed and send to the BTC network.

"id" uniquely identify each request when doing batch requests.

```json
{
  "jsonrpc": "2.0",
  "method": "rhp_sendTransaction",
  "params": [
    {
      "transfer_sig": "IAhhvyviQFFVR28YYdbeQIviVve9emkxb1zwRbpEnnBzBLqiNLMaZ2Jz8h2jAwxWGlm5wn7AjRrvGBICRtq/UOE=",
      "calldata_sig": ""
    }
  ],
  "id": 1
}
```

### contract call

If this request does not include any asset ownership transfer on the BTC network, then a transfer signature is not required. If it does, the user needs to sign twice: once for the asset ownership transfer of the account on the BTC network, and once for the contract call on the programmable layer.

user sign the following data:

```json
{
  "from": "tb1paqws24hmyjuxqxcy00guz27w3s8ecs7tzd4w5d7uy0cjvxhtyluqfc2qk7",
  "to": "tb1pkpy90pkwnu88tvshtx5xmzghzw9f8g9zdqx86n78huk3rtr3krrqr7nukv",
  "value": "0x9184e72a000", // btc value to send in sat
  "calldata": "" // follow the same rules as ethereum.
  "nonce": "0x0"
}
```

And if there's asset ownership transfer along with this contract call, user should sign the transfer transaction also and put them into params together. The request should look like:

```json
{
  "jsonrpc": "2.0",
  "method": "rhp_sendTransaction",
  "params": [
    {
      "transfer_sig": "IAhhvyviQFFVR28YYdbeQIviVve9emkxb1zwRbpEnnBzBLqiNLMaZ2Jz8h2jAwxWGlm5wn7AjRrvGBICRtq/UOE=",
      "calldata_sig": "IAhhvyviQFFVR28YYdbeQIviVve9emkxb1zwRbpEnnBzBLqiNLMaZ2Jz8h2jAwxWGlm5wn7AjRrvGBICRtq/UOE="
    }
  ],
  "id": 1
}
```

## examples

### fund request

```sh
grpcurl -plaintext -d '{"txHash": "your_tx_hash_here", "btcAmount": "0.01", "runeId": "your_rune_id_here", "runeAmount": "100"}' localhost:50051 programmableLayer.FundTxHandler/SubmitFundRequest

{
  "success": true
}

```

### mint request

```sh
grpcurl -plaintext -d '{"transferSignature": "your_transfer_signature_here", "calldataSignature": "your_calldata_signature_here", "mintAmount": "0.01", "publicKey": "your_public_key_here", "userAddress": "your_user_address_here", "expiredAt": 1234567890, "nonce": 1}' localhost:50051 programmableLayer.MintTxHandler/SubmitMintRequest


{
  "success": true
}

```

### claim request

```sh
grpcurl -plaintext -d '{"calldataSignature": "your_calldata_signature_here", "mintAmount": "0.01", "publicKey": "your_public_key_here", "userAddress": "your_user_address_here", "expiredAt": "2023-12-31T23:59:59Z", "nonce": 1}' localhost:50051 programmableLayer.ClaimTxHandler/SubmitClaimRequest


{
  "success": true
}

```
