func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
    m := make(map[string]map[string]string)
	for _, words := range similarPairs {
        if oldMap,ok:= m[words[0]];!ok{
            newMap := make(map[string]string)
            newMap[words[1]] = words[1]
            m[words[0]] = newMap
        }else{
            oldMap[words[1]] = words[1]
        }
		 if oldMap,ok:= m[words[1]];!ok{
            newMap := make(map[string]string)
            newMap[words[0]] = words[0]
            m[words[1]] = newMap
        }else{
            oldMap[words[0]] = words[0]
        }
	}
	for i := 0; i < len(sentence1); i++ {
       if sentence1[i] != sentence2[i] && m[sentence1[i]][sentence2[i]] != sentence2[i] {
           return false
       }
    }
	return true
}


