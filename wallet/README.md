Host Wallet
===========

## API

1. Get Address

```bash
curl -XPOST http://localhost:8080/address -H "Content-Type: application/json" -d '{"userID":"1", "sign":"XDQD7yPJ0o0rfzFwuW+C554FJpKcI2dqjS6PXbzkZYo/PBrWR8YIy4qHGn8jUnqAOok61uYgroG39+DBb61aMl9bM8OdPMHOcrJ5fKkpQIfk3hedeQf3JSVgC/AZw6tmEJh7FsNZIID4XDLdjzkyCueVIYzlo7akZSLT6EAXxXijRvYLLx3N4w4VTYLSd48VctU9uhdosqRZlkGBKZ38sD6HPjiqW1NT3AiL7jKYT3UoKTeI6YWYz8QhdXveK+RUswH8XitetXBrB7EBxSLWr7JatOT3vq9NJhlpMRmRoC141L1cPQCSNCov8qvnLQ/TEuzkT6NcJe/C4WLAeqx9ww=="}'
```

2. Sign

```bash
curl -XPOST http://localhost:8080/sign -H "Content-Type: application/json" -d '{"userID":"1", "data":"48656c6c6f20476f7068657221", "sign":"bK2G/ta+90PeHat7IUaYfxRr2m/e1uwfJwxjNH4bOGuJoa8Echlok7WUmPXLLiKTO0SPUDCk7tHbSFYH5V/KLtzAbF/QuwURsRxNPdW4UPi3g2pDLpERfwSPmwCLm6Jy7v/LdQghxMbYSGSiR1Ahd7Bq4Ygi248G5amUBeGMw4SX4W1JF8LIQPtVVW1aoyxIk5ctIohCR5pDKTeVtb/ebkdA51y10jnKdgtFwJcOO9VdfaaZed+REKDMyJCLZMVAMMUdv2ILoaOHhJ49h73/0aPCkRegaYmRXpw9giJ0T7qYYqOsFJhVdQ6mNYjGYUEcez1j56GWf8vgDMGJWYAuBw=="}'
```