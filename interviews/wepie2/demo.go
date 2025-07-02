package wepie2

func schedule(tasks []string, cooldown int) int {
	cnt := make(map[string]int)
	for _, t := range tasks {
		cnt[t]++
	}

	var totalTime int
	var seqLen int
	for len(cnt) > 0 {
		seqLen = len(cnt)
		totalTime += len(cnt)
		for t := range cnt {
			cnt[t]--
			if cnt[t] == 0 {
				delete(cnt, t)
			}
		}
		if len(cnt) != 0 && seqLen < cooldown+1 {
			totalTime += cooldown + 1 - seqLen
		}
	}

	return totalTime
}
