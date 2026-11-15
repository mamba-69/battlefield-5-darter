package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.2" )

func u7i1Yo0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aI4t7ParWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SR9IJppBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00fK31IsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q0mqdIQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhZ0gwwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5c9v4Rj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lu9f2XHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3OtcDYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5Fc6vTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXx1yBIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btA05PHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMYDd3L0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmPJ3XCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjVRsmwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5rh4eaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzjPe9GnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KP68B4tBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XN2YSKjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kq2VIajOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ccX6QsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHerqg2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msWAQO0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4P9EL8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05pYjTbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0mTH7RCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSxkdYDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7xcPs8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NhuWsCX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sozWlpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HiKWPxpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPP4DZD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qEILV7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrOvMUMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lBVU4k7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4s4jmDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqd2TGo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3iyEVFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zgqd3VMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gu4rwwIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDFXFiulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXUFmREaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fz8yHDmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VoW3eILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaAnNXtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6h5bC44XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpY1UzUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZuaWbWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRCJ29pUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NANhq6BBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSOudYcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzrEqnlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sL7gntbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OpdM1AfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xq5DLbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5oABj6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUpw3jxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESmtbbWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phDapH07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4sE3RngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGmKg55BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faK0dR9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5wUpeTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O15lZV5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uUGwJfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REeY9y08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYB4zNzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9tavWQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5CpqqG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uh9lCfXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXGmwqr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XBZpO1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nWeMhs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXCExiRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOac9W7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYo03wsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mI10jFhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAofEFCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHLXfaBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5yuBtk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwiQhjgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnWdfcGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5iAv6fcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdoAhh5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24LbIcCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33IvBAErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyzF7Qs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEHVRkJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwuWDwuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tO8i3wyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbUrz7UuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3Pahfr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7HKsFRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUX0ipc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIPEyOO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thuYkhX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LY0bKXp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFpnfvkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nt11hTeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oynRqgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLIuNnzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5RUjvcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hxsg9azhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aRIuWkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyiDKR2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvGycELQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8asfAeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tWvWKoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lNLaDVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4ar9Ij7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGwPbkXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egMHKWDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGfwIaF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func db301UIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azgICOZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJVAi95qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FHMvIfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFFGYXPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qE2T5wxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5z0QknZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HMq9ckeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Wd1cSl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgIWZKZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHmJ90i7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZYWqokvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7btvnLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDX74rb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnMTDN0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ileq0qgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiY7AbMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMr8RrjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6l10J4yXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2Ek0HoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCQAKxtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTd5eCgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBx421qbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vN7rLbthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAQfCK0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ak7JytxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pHN4vKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWR3cIvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUepNg8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1z795zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wygLwIaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6P7tcxYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDP2U01SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddaXqP51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func geZXxDI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRGrkGYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfjgZRM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFKm7kRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTBT2b5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0H5y4JaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9LAzj7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdYaY7MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFTn9pX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmKKQUWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9Et19atWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4umFvwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WNTFn2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38ZPgJtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njl4dNaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULU9Fpm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GE5UI6lrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GknE6JrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLxSvJkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bWjviOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADkgBowdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpmvZOjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuqU8tBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bparr12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ss7htt5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jX2k2dnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yt7BdqDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SLrIkVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIUgKs28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwbD6ZM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFhNqzOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nP70s1BLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzUKEKP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJbvDgyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4M81njwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsqTqe6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXLduMwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvtwEu9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1J8aTqpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iE8lxaGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i915qbhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOBHBgxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOP8tMslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZzPNHGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbhYXiHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAhrkelpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XL76IpZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaXc8lC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfgWZ68NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQdCSyjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wajgJPBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVlV6RqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q81PLULsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7k2AZbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpFtlulgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxYJ6MSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9jr0WxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Whw6iznVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmqDpAFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ljjhgZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqwSLetdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6TRhpkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Es4hpqNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIBYRV6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXg4QLOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AVg4TOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4D0QoBhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAbcb2PPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jthaWo46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uF85pST3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47CJKf1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ieYGtxJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0i7eEXqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XClip9yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HF2fX1i0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCXMJ0lsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmmOqjZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8M8rmq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yU782hyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjJkMqa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XOhGOBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTDP5xnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AweBAraVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IMod3uFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXhf9ChXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCg16sgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyv94kR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCPsOHopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqFHttS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kavcn4ggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uP3pIoKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqjIYtckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjEZVd1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jazZ8p8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTpuhSl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUye2C9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjIndN3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnkV2L5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qtm9MQETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FbQyTX3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zl2uU2ZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5UQ7R4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rvmc2CmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duaI0OttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5W5Z5m67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odRxThbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYs6ZB5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNRM7YpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1T5nHhQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fEEkc78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwojOLiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8tXfH52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCZaGrdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbRDgEssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaxBKc53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0ckIvHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22AAaJntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYM9GE5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXekFh8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qe0Vx6HcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxPGvYCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyAhcwWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIdgoZVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fF8oOohuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OtzxVEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cHuJHYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gr3IBK0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88jGWNc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9uu7gKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEWeQBnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCa6vqiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxvq9vDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqViQj6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWmf7fvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz7BEiLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhcJuf0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROZ9gBcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBPqOqnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqJzUKwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUB41lfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifE2BRHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoMav7MsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYPo2XdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mzs2qGHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GzXUpu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wG83IjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDHn8ApiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3GIo9rmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhedgXeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 496w1XilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDutPqLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZleV4mnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cte31MUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsscEECgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxF1lsiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkbQL9iBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jf1frEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZHls5sRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DcuegVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3gcopGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHoL9kQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtpYH278Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQ6YmI84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qheYxLCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXwWf5NIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFasXtX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgJHVSzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhOZtRquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9E0hctxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5eGTktiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLBT9M6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjtIUZnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRvzkTYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjSg9XNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpp9sXvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2YwGSqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7Df2tIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3oLkYfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWKKuNjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddlhZywAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vt5jszARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoOdmHpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLyfc9arWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aK6pYRnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CC2knMXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRn0iPAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzH5kAfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGFjpxo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mlb6gv83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2LWMWOSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huVUR6T2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAO7EKpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgjXWl7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIAUYKXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjTrCTVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjP9PI4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7vr612qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aghVH5SyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEbzpAQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THKnRvVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdq3lwD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJEFDZP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jz0B53EUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnM6F9ZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58IY5eD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJG7OnzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q465KUNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uojVJmTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQodm1hWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXvQuvT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSrD8Qt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RA1c0toWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEm3t8yaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFyJ1eBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAh2i1PNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func givQnBGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgwLR9wOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMtWHLkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQangOgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukdZhpshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZYEhbz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjeHuqUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycbRL4PCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQQlGzIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRrvNlb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpyr6ZzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKaw7dNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdPb7MfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOZFY96rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPCOivKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTG23hd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwfA7CApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JD0fQ1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiOHC3yDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEHUur4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efe8gMsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqXtr1ZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5Houc8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTE68g2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoSEYwSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdgZ4213Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGBL68UPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hahDPVlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSJkDBbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpYvmJ94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ii2ZuQfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rb3QCCT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDpBere7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9oSlgzEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wj68c1IqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXeAHLzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JJyBSgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taHaMEjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7bvZWuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OS7K2hVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1a4Oh076Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38cEMVVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7O16AnSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVoYRqPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JibaOjDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOTQ8YisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boV4Z1o5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ck0xDSngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFLwtWUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25R4p4bzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jqkqJXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJKy6xomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ore34bxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhGoVdBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4MJB15xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LG9FXr6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEdXNJOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esEMokqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVmuMqY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTrYbIDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Od6ePK4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyZXb9n3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KcksJuXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ph0rxTz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYiVoFefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vuZ1hwDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRriiV0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6IQajDRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TD2PuUcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WE2QMepPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApoqA4zrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1KyOVGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDsSr72BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waUouOH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yul5wrC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7S3qxRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTF3yseTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTmCLMTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bZZL35ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ri8WanwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tp7s85DuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xncOJvkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjTPRUUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2r585lwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4T8yBVteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAnZIdXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avqFkDrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6UIg1gPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7lPMN8p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgMlDPIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4MGj94uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhb1YSAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67dV70nOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4IT5AdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inY1ZSvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88i7xQvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpbhm25RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3EJ4A8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UZLQtiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8vSblWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zI9fdOLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSxnn9N7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIpjczxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjT7sdTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zd8GYrvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcb7ytlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RPwF8yyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UbPTWXZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEwszbRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMmBAOY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DX2C1bY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzdFQ1SGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmdavD8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivvgdn8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlLZCCI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJ1wltGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5NWHV7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQN0yWJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BvFyewDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tT1CTg2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ljneGL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4SxIuN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkVREsoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyxYVtvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04dhb142Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2IJEbl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCZYUPZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfgjjkqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlqE5QhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5WAwOXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPo9dm2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrikuTWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0D1CblXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WvQkQVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1S0zVEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVePyCtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYodumXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYUgS3T4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xb6Tt0WCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypmOsGoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EUpqUU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2pR11gZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7KQoJMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmYWNE2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Dyxk1jHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J51lTjZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSMLcZ0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIq78jqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6QHHnnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eu0l7N3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CL7j5J6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9y82vLV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMiv9P5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3txFWigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qgw13CD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hM4yGH8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QuCOfaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyovAZaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZy5FUPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BekAy4g6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0O2DFYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlkGtJb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P5mf3Ms0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLTd5Yx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3FOYsD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBWkgla7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WtSNGs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z21OtGfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fi3kfQ5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0lkfpYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0HZqjsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onLQj4CMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjzUZSo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ry3DafdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3gfAQI2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaAJC9enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8hLvLqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wcSehRpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRfMflq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfet3oF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDyMKezHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QAWEEPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97G4oFYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AWLkl12UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjsfzE9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRmd2RfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2cV4GFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5icSORHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKfj9UvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rfBZjDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTRUgLKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgtpGmiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CG7nCCthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13M1fLyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KklIywswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4Kk4eJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tby8YCJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gq5ekyIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0gc2OVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLkMCoS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5C0WdcTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXEHQzDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQkcGVqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ak32OAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8ZIMvegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONuC5Yu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqcnlWd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYDM47xSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wm9VfSoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7NEgSl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOW1myvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDHqGv6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhQDoAr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPeMNzLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhb85KLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKaWx1G4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZX3PClECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1aZEV6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyC072OeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYzB5pQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPU7w3n2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNoFT2slWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmoFncuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJCXid5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cklT4KcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YqWwN4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFmfVfGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtnGGQ1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dioKe7MgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEq5TQt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4k0KiXg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taJ7Xhm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7Gk2urpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhRQodw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUWWBfxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDJ2DMxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22G9EQUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGEq0MplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bk2Gy6QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKEnyM00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTBkgQIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93RuQuy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrh74qyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHqIplIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUXVGe1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrKyVGttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59RHs6tHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bme6iQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ov4Bnx5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMiZCI33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rylO0tX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBqcL9NyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFIJYMliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rndqMcihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 736mNg1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0Y81V91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlF5sLNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiRHStR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRGduqYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAje4LpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o6KENOgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SktXfKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4T0nYjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9e47xG4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZKU7WyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajeu3qmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZm87vTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HYXcmq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnigHJnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tHNFl1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQrrFYe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6RRgLdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1MFGM3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TJwuBTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLPMKEC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNJShjk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbCDQLzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAkHO83lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJeOJEHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqFl1zizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meziDVpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQrfaRZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKGIV0cBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D87sGxiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bt9IsDsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPCwx0UiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccJ3kG7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYSGQfuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3PN4aU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYZX8GLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12x0cOstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXENuIXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TGZIb3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhixhSArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4zA2l3FOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oUbA5ycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nP7nwtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXtcuk1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRHLUYrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3gw1cJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZtJgZRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVqgoPBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjk4jitXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoAWgiINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oc7B30khWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlo46hhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvGazCqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZ0pMX6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func di8LH6JcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOuVYd6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2Cf7JnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FwAoPiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZmiY0AGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LgAOTAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqzol0xzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQvod7RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWxXiFJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GDRC3WgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2Zhfj93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boKY1VWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNAjGugQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JXvyhD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kd2QIRoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaOvpUmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xlodZdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func joWzrCDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jx5GqzKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGSbk82oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoKP3gUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKYnXNydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNis28bTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFJoFLUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTxyyNS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5hewF8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7k5vlmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27zcRBRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SbjNu0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAG58iNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCaPykd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSqZJyH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdAZVTjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iDczfKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCaDJeLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fBaDz4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cs440jVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrEYeQW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EEg6MpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLxaU64uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUaACZyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLHqRiHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVrw6ACEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7azdFddCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJ2ENJbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TULDRHL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGUSqdZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHGrt9t5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4ia0HwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHJ2JfAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOrar16MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYiHpH8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVuhiG0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YL7kYNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v90ggbkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvcShGatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14JmuF8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRww9QxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oca7DhEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nVPBoFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEFh5UAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8w7FQN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5k28bw8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeU27lHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dlrcK3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPt8gqCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qupTElcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPZQkDFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3j3Dl5TqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeXO9db2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XG8m3m8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HJSj5epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTv7MaSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgkrcFMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSLpyxrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFDooYNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjYMO9ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qngzwnKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMxTmWajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnKoCHz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkwKLOT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYle7uZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkhTdz78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xG1MhXV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7NG3cnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSuXKw3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66zS80OUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFxt5i9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeSORPetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Xxj9pR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tqa7vthzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsTHq7ZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gsCpqvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5GzBR6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fU4YZ3QWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DI3ESnDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oH05VirhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoKD5hESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJ3VkEhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5TqkbR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0l5TjYeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1B6zDgBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 953nIzJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qiE0Ua2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgzwRdyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFO5OmfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrcV5JVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MqBLQOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNvVk2yUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JceXKIJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCIUQVnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVLgRFeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQqjLpezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4BZD1KcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZiZdHNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Adfy6FvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzxmNRA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVElX9qJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5v1lws8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AzqIWlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82rsnIjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nl0vMIliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZijtgLLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQqKcmhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YWxcI2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDUYhxnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SY1KfguBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQyzIoeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wyS26rNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8Wc7MidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4KXJ4YkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25vtpTCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxW2uAUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qw5if7sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFtFQtgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIzi7QjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5w75k56uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9f3X9LoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P46kNWQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2dVHERmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7HAr1D1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yK5c7sHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31f5Ol72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSEAEtlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFY444LhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8EusxUwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYjjsRS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McVyKC6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8tGMWayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8HflYBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0H01TXKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTORtimnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 680EcFyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bR2WH9rCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymF0cqbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jekKyQWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOkEII9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etIXMJTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9z1zcACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIX9VD2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGdh5LatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCPBqnlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57GNnw5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3vbT8jLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51X0DuofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGae2DnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNnvif04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oD1KnkjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLoKEXpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHmlzip0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUaZuZE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gk3uc1ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIFlsB03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func abex0JOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PS7N0qt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zk7vXAVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nKEcxmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfM1RrJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKpKfcccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ww88Opb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNsn9CkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjV6DZ71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAaR1tWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVftbK0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwZJI0ZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9JlmaSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nz9rdGrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUIle3LaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMm3WSYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xmn0QC7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sbijN08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOdVubqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjI21XX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5b2KReFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6P2SmgSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uP5DwHDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiR5uXZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Oq8TZSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJQ1qVw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ab22YhKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bi5ygdIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNqY8CuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2giaVqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sc7hyviEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdZUaC4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uaMSweuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zh9muw6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agabMQWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdVpG4ajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0MRllQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Akz6GPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N069PMSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOfmtVHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36g40R1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJjqaJnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3avik2xhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67HU9VZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGxJJuiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhqKWhIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UIqce9s4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHFcdEGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J92qu5DmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUxBWL6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrHdAqWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmGLC8fGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAu92Yl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TACFahNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FspVFsCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xy0PeyWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDdK4fOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U95bUirbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhLJXlP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uz5qEdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMClRmHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tzoARdKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKGodNDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L25BMtukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 745jAhwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmqv0wHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCgQ0fiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOhDGT9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sn9xZz3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlRFHjGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqngxertWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBJacN6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rQMFVVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ib5wEuxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9FmSD1t9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGb4vb6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fEZTU51eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gKrsLlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZ1CeJlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

