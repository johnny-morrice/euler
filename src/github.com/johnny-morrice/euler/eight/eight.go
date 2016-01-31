package main

import (
    "fmt"
    "strconv"
    "flag"
)

type args struct {
    n int
    all bool
    digits string
}

func main() {
    args := readArgs()
    products := factorMatrix(args.digits, args.n)
    if args.all {
        for i := 0; i < args.n; i++ {
            fmt.Println(max(products, i))
        }
    } else {
        fmt.Println(max(products, args.n - 1))
    }
}

func readArgs() args {
    params := args{}
    flag.IntVar(&params.n, "factcount", 13, "Count of factors")
    flag.StringVar(&params.digits, "digits", problem, "Digit string")
    flag.BoolVar(&params.all, "all", false, "Print max factors for all n < factcount")
    flag.Parse()
    return params
}

func factorMatrix(digits string, factcount int) [][]int64 {
    // Prepare matrix
    digcount := len(digits)
    products := make([][]int64, digcount)
    for i := 0; i < digcount; i++ {
        products[i] = make([]int64, factcount)
    }
    for i := 0; i < digcount; i++ {
        products[i][0] = chrnum(digits, i)
    }

    // Dynamic programming
    maxi := digcount - factcount
    for i := 0; i < maxi; i++ {
        for j := 1; j < factcount; j++ {
            products[i][j] = products[i][j - 1] * products[i + j][0]
        }
    }

    return products
}

func max(products [][]int64, factcount int) int64 {
    // Find max
    maxi := len(products) - factcount
    max := int64(0)
    for i := 0; i < maxi; i++ {
        prod := products[i][factcount]
        if prod > max {
            max = prod
        }
    }

    return max
}

func chrnum(digits string, index int) int64 {
    ch := digits[index:index+1]
    n, err := strconv.Atoi(ch)
    if err != nil {
        panic(err)
    }
    return int64(n)
}

const problem = "73167176531330624919225119674426574742355349194934" +
"96983520312774506326239578318016984801869478851843" +
"85861560789112949495459501737958331952853208805511" +
"12540698747158523863050715693290963295227443043557" +
"66896648950445244523161731856403098711121722383113" +
"62229893423380308135336276614282806444486645238749" +
"30358907296290491560440772390713810515859307960866" +
"70172427121883998797908792274921901699720888093776" +
"65727333001053367881220235421809751254540594752243" +
"52584907711670556013604839586446706324415722155397" +
"53697817977846174064955149290862569321978468622482" +
"83972241375657056057490261407972968652414535100474" +
"82166370484403199890008895243450658541227588666881" +
"16427171479924442928230863465674813919123162824586" +
"17866458359124566529476545682848912883142607690042" +
"24219022671055626321111109370544217506941658960408" +
"07198403850962455444362981230987879927244284909188" +
"84580156166097919133875499200524063689912560717606" +
"05886116467109405077541002256983155200055935729725" +
"71636269561882670428252483600823257530420752963450"