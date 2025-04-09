# Two_Sum Benchmarks

**Last run:** 2025-04-09 15:10:12  
**Go version:** go1.24.1  
**CPU:** AMD EPYC 7763 64-Core Processor

## Large Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| BruteForce | 1 | 3138888457 | 16 | 1 |
| Optimized | 240 | 5089855 | 2364583 | 258 |

## Medium Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| BruteForce | 3,775 | 316972 | 16 | 1 |
| Optimized | 38,190 | 31606 | 36960 | 6 |

## Small Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| BruteForce | 22092,264 | 54.71 | 16 | 1 |
| Optimized | 3255,355 | 367.8 | 344 | 4 |

