package jumpList

type level struct {
	next *node
	span int
}

type node struct {
	backward *node
	score int
	levels []level
}

type JumpList struct {
	head *node
	tail *node
	maxLevel int
	length int
}

func newList()*JumpList{
	j :=JumpList{
		head: &node{
			levels: make([]level,32),
		},
		maxLevel: 0,
		length: 0,
	}
	return &j
}

func (j *JumpList) Insert(score int,val interface{})  {
	cur := j.head.levels[j.maxLevel].next
	for cur != nil{
		if cur.score > score{
			cur = cur.
		}
	}
}
