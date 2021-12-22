# Numary Machine [![test](https://github.com/numary/machine/actions/workflows/main.yml/badge.svg)](https://github.com/numary/machine/actions/workflows/main.yml)

A virtual machine for moving money.

This repo bundles:
* The Numary Machine VM
* A Numscript parser
* A Numscript compiler

# Example

```
send [USD/2 1099] (
  source = {
    @users:001:wallet
    @users:001:credit
  }
  destination = {
    85% to @drivers:033
    15% to @platform:fees
  }
)
```

# Documentation

You can find the complete Numary documentation at [docs.numary.com](https://docs.numary.com)
