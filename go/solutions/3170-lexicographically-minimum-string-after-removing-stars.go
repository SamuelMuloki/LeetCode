package solutions

import "container/heap"

func ClearStars(s string) string {
    h := &MinHeap{}
	heap.Init(h)

    st := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            num := heap.Pop(h).(int)
            j := len(st)-1
            for ; j > 0 && int(st[j]-'a') != num; j-- {}
            st1 := st[j+1:]
            st = st[:j]
            st = append(st, st1...)
        } else {
            heap.Push(h, int(s[i]-'a'))
            st = append(st, s[i])
        }
    }

    return string(st)
}
