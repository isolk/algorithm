package tree

type SearchTree struct {
	root *Node
	count int
}

func (s *SearchTree) Add(val int)  {
	insertNode := &Node{Value: val}
	if s.root == nil{
		s.root = insertNode
	}else {
		k := s.root
		for k!= nil{
			if val < k.Value{
				if k.Left == nil{
					k.Left = insertNode
					break
				}
				k = k.Left
			}else {
				if k.Right == nil{
					k.Right =insertNode
					break
				}
				k = k.Right
			}
		}
	}
}

func (s *SearchTree) Find(val int)bool  {
	for k:= s.root;k != nil; {
		if val == k.Value{
			return true
		}else if val < k.Value{
			k = k.Left
		}else {
			k = k.Right
		}
	}
	return false
}

func (s *SearchTree) Delete(val int)bool  {
	father := &Node{Left: s.root}
	l := true
	now := s.root
	for now != nil{
		if now.Value == val{
			if now.Left == nil && now.Right == nil{
				if l{
					father.Left = nil
				}else{
					father.Right = nil
				}
			}else if now.Left == nil {
				if l{
					father.Left = now.Right
				}else {
					father.Right = now.Right
				}
			}else if now.Right == nil{
				if l{
					father.Left = now.Left
				}else {
					father.Right = now.Left
				}
			}else {
				fa := now
				min := now.Right
				for min.Left != nil{
					fa = min
					min = min.Left
				}
				now.Value = min.Value
				fa.Left = nil
			}
			return true
		}else if val < now.Value{
			father=now
			now = now.Left
			l = true
		}else {
			father=now
			l = false
			now = now.Right
		}
	}
	return false
}

func (s *SearchTree) PrintAll()  {
	MidTraverse(s.root)
}
