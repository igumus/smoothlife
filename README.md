# Smooth Life Implementation in Go

Smooth Life (according to paper); 
```
is the generic generalization of Conway's "Game of Life" to a continuous domain. 
```

## Quick Start

```console
$ make build
$ ./bin/smoothlife
```

## Arguments

| Argument         | Description     | Default Value|
|--------------|-----------|------------|
| ra | Paper's outer circle radius | 11.0 |
| dt | Infinitesimal time to give ability to obtain arbitrary small time steps | 0.05 |
| b1 | Paper's b1 param | 0.278 |
| b2 | Paper's b2 param | 0.365 |
| d1 | Paper's d1 param | 0.267 |
| d2 | Paper's d2 param | 0.445 |
| alpha-n | Paper's alpha(n) param | 0.028 |
| alpha-m | Paper's alpha(m) param | 0.147 |
| interval | Duration between each step | 300ms |
| step | Execution step count | 200 |
| with-paper-diff | Switch between diff algorithms<sup>1</sup> | true |

### Argument notes

1. `with-paper-diff` command line argument is used to switch between paper's diff proposal and tsoding's diff implementation. `true` means `use paper's diff proposal`.

## Todos

- [ ] enhance terminal rendering
- [ ] tweak automaton parameters to explore new life forms :)
- [ ] check anti-aliasing method for neighbour value selection which described in paper (page:3, paragrah:2)

## References

- [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
- [Paper](https://arxiv.org/abs/1111.1567)
- [tsoding](https://github.com/tsoding):
    - [Source in C](https://github.com/tsoding/SmoothLife):
    - [YT Video: Simulating Life on my Computer like a Scientist](https://www.youtube.com/watch?v=L68_BBiuHUw)
    - [YT Video: My GPU Almost Died Doing This Life Simulation](https://www.youtube.com/watch?v=9s8vjf_vLaA)
