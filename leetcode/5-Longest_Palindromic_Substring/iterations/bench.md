# Longest_Palindromic_Substring Benchmarks

**Last run:** 2025-04-09 15:10:58  
**Go version:** go1.24.1  
**CPU:** AMD EPYC 7763 64-Core Processor

## Large Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 1 | 7765562766 | 0 | 0 |
| V2 | 1,178 | 1027283 | 1810438 | 2 |
| V3 | 3,472 | 343464 | 0 | 0 |

## Medium Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 1,566 | 770809 | 0 | 0 |
| V2 | 111,434 | 10495 | 18432 | 2 |
| V3 | 348,600 | 3437 | 0 | 0 |

## Small Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 10490,440 | 111.8 | 0 | 0 |
| V2 | 8523,796 | 138.8 | 216 | 2 |
| V3 | 29744,331 | 41.34 | 0 | 0 |

