# Proof-of-Work (PoW)

The challenger will receive a prefix and number n representing difficulty. He has to compute an answer such that `sha256(prefix + answer)` has n leading zero bits.

This PoW is used in Balsn CTF 2019.

## Pow solver 

We support various languages to help you solve the obnoxious PoW.

- [X] Python 2
- [X] Python 3
- [X] NodeJs
- [X] Browser-based
- [X] Ruby
- [X] Go (contributed by [Lee Xun](https://github.com/LeeXun), optimized by @RobinJadoul)
- [ ] Rust
- [ ] C/C++
- [ ] Java

## Modules

- `powser.py`: A IP-based PoW server module implemented in Python 3.8. Suitable for web challenges.
- `nc_powser.py`: A simple PoW server module implemented in Python 3.8. Suitable for nc-based challenges.
