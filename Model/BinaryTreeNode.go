package Model

type BinaryTreeNode struct {
	data interface{}
	parent,
	lChild,
	rChild *BinaryTreeNode
	height int
}

func NewBinaryTreeNode(e interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{data: e}
}

func (this *BinaryTreeNode) GetData() interface{} {
	return this.data
}

func (this *BinaryTreeNode) SetData(e interface{})  {
	this.data = e
}

func (this *BinaryTreeNode) HasLChild() bool {
	return this.lChild != nil
}

func (this *BinaryTreeNode) HasRChild() bool {
	return this.rChild != nil
}

func (this *BinaryTreeNode) HasParent() bool {
	return this.parent != nil
}

func (this *BinaryTreeNode) GetLChild() *BinaryTreeNode {
	if !this.HasLChild() {
		return nil
	}
	return this.lChild
}

func (this *BinaryTreeNode) GetRChild() *BinaryTreeNode {
	if !this.HasRChild() {
		return nil
	}
	return this.rChild
}

func (this *BinaryTreeNode) GetParent() *BinaryTreeNode {
	if this.HasParent() {
		return this.parent
	}
	return nil
}

func (this *BinaryTreeNode) SetParent(node *BinaryTreeNode)  {
	this.parent = node
}

func (this *BinaryTreeNode) CutOffParent()  {
	if this.HasParent() {
		parent := this.GetParent()
		this.SetParent(nil)
		this.lChild = nil
		this.rChild = nil
		if parent.GetLChild() == this {
			parent.lChild = nil
		} else {
			parent.rChild = nil
		}
	}
}

func (this *BinaryTreeNode) SetLChild(node *BinaryTreeNode) *BinaryTreeNode {
	oldChild := this.lChild
	if this.HasLChild() {
		oldChild.CutOffParent()
		this.SetLChild(node)
	} else {
		this.lChild = node
	}
	return oldChild
}

func (this *BinaryTreeNode) SetRChild(node *BinaryTreeNode) *BinaryTreeNode {
	oldChild := this.rChild
	if this.HasRChild() {
		oldChild.CutOffParent()
		this.SetRChild(node)
	} else {
		this.rChild = node
	}
	return oldChild
}