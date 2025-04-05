# Median_of_Two_Sorted_Arrays Benchmarks

**Last run:** 2025-04-05 16:04:59  
**Go version:** go1.24.1  
**CPU:** AMD EPYC 7763 64-Core Processor

## Large Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 204 | 5872161 | 802879 | 3 |
| V2 | 5,492 | 210832 | 802819 | 1 |
| V3 | 167554,323 | 7.176 | 0 | 0 |

## Medium Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 35,340 | 33632 | 8248 | 3 |
| V2 | 469,830 | 2376 | 8192 | 1 |
| V3 | 167226,840 | 7.193 | 0 | 0 |

## Small Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 6977,835 | 172.5 | 136 | 3 |
| V2 | 26532,405 | 43.83 | 80 | 1 |
| V3 | 100000,000 | 10.91 | 0 | 0 |

