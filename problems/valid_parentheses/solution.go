func isValid(s string) bool {
    state := []string{}
    for _, item := range s {
        switch string(item) {
            case "{", "(", "[":
            state = append(state,string(item))
            case "}":
            if len(state) == 0 || state[len(state)-1] != "{"{
                return false
            }
            state = state[:len(state) -1]
            case ")":
            if len(state) == 0 || state[len(state)-1] != "("{
                return false
            }
            state = state[:len(state) -1]
            case "]":
            if len(state) == 0 || state[len(state)-1] != "["{
                return false
            }
            state = state[:len(state) -1]
        }
        
    }
    return len(state) == 0
}