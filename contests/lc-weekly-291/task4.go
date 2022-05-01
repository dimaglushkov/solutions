package main

/*func appealSum(s string) int64 {
	var res, i, j int64
	var N = int64(len(s))
	var letters [26]int

	for i < N {
		if j < N && letters[s[j]-'a'] == 0 {
			letters[s[j]-'a']++
			j++
		} else {
			letters[s[i]-'a']--
			i++
		}
	}
	return res
}*/

// TLE
// At first I tried to solve it with straight-forward idea of counting appeal for every possible substring
// but later realized that I could use longer string appeal to get next strings appeal faster
// Unfortunately, didn't have enough time to implement this idea

func appealSum(s string) int64 {
	var res int64 = 0
	var letters [26]int

	var words = make(map[string]int64)

	findAppeal := func(str string, appeal int64) int64 {
		if str == "" {
			return 0
		}
		if words[str] != 0 {
			return words[str]
		}

		words[str] = appeal
		if letters[str[0]-'a'] == 1 {
			words[str[1:]] = appeal - 1
		} else {
			words[str[1:]] = appeal
		}

		if letters[str[len(str)-1]-'a'] == 1 {
			words[str[:len(str)-1]] = appeal - 1
		} else {
			words[str[:len(str)-1]] = appeal
		}

		return appeal
	}

	N := len(s)
	var appeal int64
	for _, r := range s {
		if letters[r-'a'] == 0 {
			appeal++
		}
		letters[r-'a']++
	}
	findAppeal(s, appeal)
	for i := N - 1; i >= 0; i-- {
		for j := 0; j < N-i+1; j++ {
			res += words[s[j:j+i]]
		}
	}
	return res
}

func main() {
	println(appealSum("abbca"))
	println(appealSum("code"))
	println(appealSum("a"))
}
