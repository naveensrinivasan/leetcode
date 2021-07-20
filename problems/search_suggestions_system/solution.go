func suggestedProducts(products []string, searchWord string) [][]string {
    root := NewTrie()
    for _, w := range products{
        root.insert(w)
    }
    return root.search(searchWord)
}

type Trie struct{
    C  []*Trie
    S  []string
}

func NewTrie()*Trie{
    return &Trie{
        C : make([]*Trie,26),
        S : []string{},
    }
}

func (t *Trie)insert(w string){
    for i:= range w{
        if t.C[w[i] - 'a'] == nil{
            t.C[w[i] - 'a'] = NewTrie()
        }
        t = t.C[w[i] - 'a']
        t.S = append(t.S,w)
        sort.Strings(t.S)
        if len(t.S) > 3{
            t.S = t.S[:3]
        }
    }   
}

func (t *Trie)search(w string)[][]string{
    var res [][]string
    for i:= range w{
        if t != nil{
            t = t.C[w[i] - 'a']
        }
        if t == nil{
            res = append(res,[]string{})
        }else{
            res = append(res,t.S)
        }
    }
    return res
}