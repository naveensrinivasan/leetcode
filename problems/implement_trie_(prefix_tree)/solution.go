type Trie struct {
    root *Node
}

type Node struct{
    C [26]*Node
    EndWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{root:&Node{}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    curr := this.root
    for i:= range word{
        if curr.C[word[i] - 'a'] == nil{
            curr.C[word[i] - 'a'] = &Node{}
        }
        curr = curr.C[word[i] - 'a']
    }
    curr.EndWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    curr := this.root
    for i:= range word{
        if curr.C[word[i] - 'a'] == nil{
            return false
        }
        curr = curr.C[word[i] - 'a']
    }
    if curr.EndWord{
        return true
    }
    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(word string) bool {
    curr := this.root
    for i:= range word{
        if curr.C[word[i] - 'a'] == nil{
            return false
        }
        curr = curr.C[word[i] - 'a']
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */