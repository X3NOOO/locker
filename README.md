# locker

```
  /&&&&&\    
 &&/   \&&   .__                       __
 &&.. ..&&   |  |     ____     ____   |  | __   ____   _______
&&&&&&&&&&&  |  |    /  _ \  _/ ___\  |  |/ / _/ __ \  \_  __ \
&&&&' '&&&&  |  |__ (  <_> ) \  \___  |    <  \  ___/   |  | \/
&&&&& &&&&&  |____/  \____/   \___  > |__|_ \  \___  >  |__|    1.1
&&&&&&&&&&&                       \/       \/      \/           
```

locker is cli program for locking files and folders.

https://user-images.githubusercontent.com/48159366/151880905-26bed182-4fb2-494b-ba1b-340bbe856632.mov

## features

- Every installation have randomized master password. Even if someone can copy your files they will not be able to unlock them
- Files are tared before encryption so you can lock the folder as well

## installation

1. `git clone https://github.com/X3NOOO/locker`
2. `cd locker`
3. `make release`
4. `make install`

## usage

- `locker`          - Display basic information
- `locker help`     - Display help message
- `locker lock`     - Lock directory/file
- `locker unlock`   - Unlock file/directory
- `locker license`  - Display GNU GPL v3 license

## donation

- XMR: `49F3GknYgs7cRfMJghrd9dHZKe63Z6Y3aJKPecDKqLRje5YebzWvz3VWsTa8e8Sk92G7WJEsyp8L1VEeNxmdj2vZNJSACo1`
- DOGE: `DFYc29EsSuSbyLndGrKBGoC2usHRUqiiXb`
- BTC: `bc1q08p6wd86806uf2cj95j4pcgl584jvaqkhs37pp`
- ETH: `0x84FfD8524a66505344A1cbfC3212392Db5b2474d`
- LTC: `Lew3VmzbkaxzoYG3jNHf263oEDMrQ3ecN1`

## todo

- [X] changing file permission
- [X] encryption
- [X] dir encryption
- [X] `--debug` flag
- [X] support of larger files
- [ ] windows support
