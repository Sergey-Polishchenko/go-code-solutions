# Median_of_Two_Sorted_Arrays Benchmarks

**Last run:** 2025-04-09 15:10:39  
**Go version:** go1.24.1  
**CPU:** AMD EPYC 7763 64-Core Processor

## Large Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 195 | 5911636 | 802884 | 3 |
| V2 | 5,374 | 216898 | 802820 | 1 |
| V3 | 167315,563 | 7.168 | 0 | 0 |

## Medium Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 35,094 | 34295 | 8248 | 3 |
| V2 | 460,131 | 2417 | 8192 | 1 |
| V3 | 167199,706 | 7.181 | 0 | 0 |

## Small Dataset
| Algorithm | Runs | ns/op | Alloc (B/op) | Allocs (op) |
|-----------|-----:|------:|-------------:|------------:|
| V1 | 6845,824 | 173.1 | 136 | 3 |
| V2 | 27194,976 | 43.12 | 80 | 1 |
| V3 | 100000,000 | 11.04 | 0 | 0 |

