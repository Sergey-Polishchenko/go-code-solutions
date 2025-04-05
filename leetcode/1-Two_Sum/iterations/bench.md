# Two_Sum Benchmarks

**Last run:** 2025-04-05 16:04:31  
**Go version:** go1.24.1  
**CPU:** AMD EPYC 7763 64-Core Processor

## Large Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| BruteForce | 1 | 3139642667 | 16 | 1 |
| Optimized | 241 | 4962350 | 2364743 | 258 |

## Medium Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| BruteForce | 3,780 | 317607 | 16 | 1 |
| Optimized | 38,070 | 31492 | 36960 | 6 |

## Small Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| BruteForce | 21755,402 | 54.19 | 16 | 1 |
| Optimized | 3289,644 | 367.1 | 344 | 4 |

