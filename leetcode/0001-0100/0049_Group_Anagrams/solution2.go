package leetcode0049

import "math/big"

func groupAnagrams2(strs []string) [][]string {
	primes := [26]int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101,
	}
	primesBig := [26]*big.Int{}
	for i, pri := range primes {
		primesBig[i] = big.NewInt(int64(pri))
	}
	hashFunc := func(s string) string {
		res := big.NewInt(1)
		for _, c := range s {
			res.Mul(res, primesBig[c-'a'])
		}
		return res.String()
	}
	anagrams := make(map[string][]string)
	for _, str := range strs {
		key := hashFunc(str)
		anagrams[key] = append(anagrams[key], str)
	}
	res := make([][]string, 0, len(anagrams))
	for _, lst := range anagrams {
		res = append(res, lst)
	}
	return res
}
