package Model

type AVLTreeNode struct{
	data interface{}
	lChild,
	rChild *AVLTreeNode
	height int
}

func NewAVLTreeNode(data interface{}) *AVLTreeNode {
	return &AVLTreeNode{
			data: data,
			lChild: nil,
			rChild: nil,
			height: 0,
		}
}

func (this *AVLTreeNode) SetData(data interface{})  {
	this.data = data
}

func (this *AVLTreeNode) GetData() interface{} {
	return this.data
}


func (this *AVLTreeNode) HasLChild() bool {
	if this.lChild != nil {
		return true
	}
	return false
}

func (this *AVLTreeNode) SetLChild(node *AVLTreeNode)  {
	this.lChild = node
}

func (this *AVLTreeNode) GetLChild() *AVLTreeNode {
	return this.lChild
}

func (this *AVLTreeNode) HasRChild() bool {
	if this.rChild != nil {
		return true
	}
	return false
}

func (this *AVLTreeNode) SetRChild(node *AVLTreeNode)  {
	this.rChild = node
}

func (this *AVLTreeNode) GetRChild() *AVLTreeNode {
	return this.rChild
}