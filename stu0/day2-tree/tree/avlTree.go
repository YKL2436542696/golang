package main

import "fmt"

// 二叉树结构定义

type AVLNode struct {
	Left, Right *AVLNode // 表示指向左孩子和右孩子
	Data        int      // 结点存储数据
	Height      int      // 每个节点的平衡因子，若该节点的平衡因子绝对值大于1则不平衡，需要调整这个子树结构
}

// 搜索

func (avlNode *AVLNode) Get(key int) *AVLNode {
	q := avlNode
	if avlNode != nil {
		if key > avlNode.Data {
			// 大于当前节点值，往右走
			q = avlNode.Right.Get(key)
		} else if key < avlNode.Data {
			// 小于当前节点值，往左走
			q = avlNode.Left.Get(key)
		} else if key == avlNode.Data {
			q = avlNode
		}
	}
	return q
}

// 插入
// 实现步骤:
// 1、使用二叉搜索树的插入逻辑把新元素插进去
// 2、插入后，检查每个节点的平衡因子
// 3、如果每个节点的平衡因子为0或1或-1，则进行下一个操作
// 4、如果任何节点的平衡因子绝对值超过1，则称该树不平衡。在这种情况下，进行适当的旋转，使其平衡并进行下一步操作。
// 总的来说就是，插入，遍历节点查看平衡因子，如果平衡因子绝对值不超过2，则继续插入；如果超过，则执行平衡操作

func (avlNode *AVLNode) Insert(Value int) *AVLNode {
	// 加锁
	// treeNode.lock.Lock()
	// defer treeNode.lock.Unlock()
	// 创建叶子节点
	if avlNode == nil {
		return &AVLNode{
			Left:   nil,
			Right:  nil,
			Height: 0,
			Data:   Value,
		}
	}
	// 从上往下插入节点
	if Value > avlNode.Data { // go right
		avlNode.Right = avlNode.Right.Insert(Value)
	} else if Value < avlNode.Data { // go left
		avlNode.Left = avlNode.Left.Insert(Value)
	} else {
		fmt.Println("值已存在")
	}
	// 在插入路径中，从下往上递归更新平衡因子
	avlNode.Height = avlNode.Left.GetHeight() - avlNode.Right.GetHeight()
	//进行旋转调整
	avlNode = avlNode.adjust()
	return avlNode
}

// 逆时针旋转，左旋

func (avlNode *AVLNode) LeftRotate() *AVLNode {
	//headNode:支点，左旋找传入节点的右子节点作为支点
	headNode := avlNode.Right
	//把支点的左子节点作为传入节点的新右子节点
	avlNode.Right = headNode.Left
	//再把传入节点作为支点的新左子节点
	headNode.Left = avlNode
	//把支点作为首节点
	avlNode = headNode
	// 更新旋转结点的高度
	// 先更新avlNode的高度，因为headNode结点在avlNode结点的上面，headNode计算高度的时候要根据avlNode的高度来计算
	avlNode.Height = max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.Height = max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1
	return headNode
}

// 顺时针旋转，右旋

func (avlNode *AVLNode) RightRotate() *AVLNode {
	//headNode:支点，右旋找传入节点的左子节点作为支点
	headNode := avlNode.Left
	//把支点的右子节点作为传入节点的新左子节点
	avlNode.Left = headNode.Right
	//再把传入节点作为支点的新右子节点
	headNode.Right = avlNode
	//把支点作为首节点
	avlNode = headNode
	// 更新旋转后结点的高度
	avlNode.Height = max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.Height = max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1
	return headNode

}

// 若传入节点的左比右高2，并且该节点的左子节点的右比左大
// 则先让该节点的左子节点左旋，再让该节点自己右旋

func (avlNode *AVLNode) LR_Rotate() *AVLNode {
	// 先把左孩子结点进行左旋
	avlNode.Left = avlNode.Left.LeftRotate()
	// 然后把自己右旋
	return avlNode.RightRotate()
}

// 若传入节点的右比左高2，并且该节点的右节点的左比右大
// 则先让该节点的右子节点右旋，再让该节点自己左旋

func (avlNode *AVLNode) RL_Rotate() *AVLNode {
	// 先把右孩子进行右旋
	avlNode.Right = avlNode.Right.RightRotate()
	// 然后把自己右旋
	return avlNode.LeftRotate()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//递归计算子树的高度在其左和右子树高度中取最大值+1

func (avlNode *AVLNode) GetHeight() int {
	if avlNode == nil {
		return 0
	}
	return max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
}

// 遍历删除节点
// 查找节点的直接前驱或直接后继，前驱必定在左子树，后继必定在右子树

func (avlNode *AVLNode) Remove(Value int) *AVLNode {
	if avlNode == nil { // 插入时找不到节点
		return nil
	}
	// 删除操作
	if avlNode.Data > Value {
		avlNode.Left = avlNode.Left.Remove(Value)
	} else if avlNode.Data < Value {
		avlNode.Right = avlNode.Right.Remove(Value)
	} else {
		// 找到了要删除节点，找到之后有三种情况：
		return avlNode.remove()
	}
	// 接着更新平衡因子，剩下的和插入一样，调整树
	avlNode.Height = avlNode.Left.GetHeight() - avlNode.Right.GetHeight()
	//调整
	avlNode = avlNode.adjust()
	return avlNode
}

func (avlNode *AVLNode) adjust() *AVLNode {
	if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == 2 {
		//如果传入节点的右子树高度超了左字数高度2
		if avlNode.Right.Right.GetHeight() > avlNode.Right.Left.GetHeight() {
			// 判断传入节点的右子树，把该右子树的节点作为支点
			// 右边比左边高则
			// 进行左旋
			return avlNode.LeftRotate()
		} else {
			// 左边比右边高则
			// 先右旋
			// 再把传入节点作为支点进行左旋
			return avlNode.RL_Rotate()
		}
	} else if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == -2 {
		//如果传入节点的左子树高度超了右子数高度2
		if avlNode.Left.Right.GetHeight() < avlNode.Left.Left.GetHeight() {
			// 判断传入节点的左子树，把该左子树的节点作为支点
			// 左边比右边高则
			// 进行右旋
			return avlNode.RightRotate()
		} else {
			//右边比左边高则
			//先左旋
			//在把传入节点作为支点进行右旋
			return avlNode.LR_Rotate()
		}
	}
	// 如果都不符合平衡因子绝对值超过1则原路返回
	return avlNode
}

// 1、节点是叶子节点，直接置为nil
// 2、节点有一个孩子，把孩子当作当前节点
// 3、选取前驱或者后继。删除后直接返回即可。

func (avlNode *AVLNode) remove() *AVLNode {
	if avlNode.Left == nil && avlNode.Right == nil {
		avlNode = nil
	} else if avlNode.Left == nil {
		// 删除节点的左子树为空，则把右子树当作当前节点
		avlNode = avlNode.Right
	} else if avlNode.Right == nil {
		// 删除节点的右子树为空，则把左子树当作当前节点
		avlNode = avlNode.Left
	} else {
		// 左右子树都不为空，需要找到顶替元素。则需要找到当前节点中序遍历的直接前驱或直接后继，这里使用后继来写。
		// 找出直接后继节点
		min := avlNode.Right.minNode()
		// 删除节点的值和后继节点值交换
		avlNode.Data = min.Data
		// 因为使用了后继，肯定在右边，所以在右边子树中删除最小节点
		avlNode.Right = avlNode.Right.Remove(min.Data)

		// 如果是前驱可以这样写
		// max := avlNode.Left.maxNode
		// avlNode.Data = max.Data
		// avlNode.Lfet = avlNode.Left.Remove(max.Data)
	}
	return avlNode
}

// 遍历找到左子树中最小的节点(当前节点的最左子节点必定是最小，如果没有左子树则本节点最小)
func (avlNode *AVLNode) minNode() *AVLNode {
	for avlNode.Left != nil {
		avlNode = avlNode.Left
	}
	return avlNode
}

// 遍历找到右子树中最大的节点(当前节点的最右子节点必定是最大，如果没有右子树则本节点最大)
func (avlNode *AVLNode) maxNode() *AVLNode {
	for avlNode.Right != nil {
		avlNode = avlNode.Right
	}
	return avlNode
}

// 前序遍历

func (p *AVLNode) PreTraverse() {
	if p == nil {
		return
	}
	fmt.Printf("%d ", p.Data)
	if p.Left != nil {
		p.Left.PreTraverse()
	}
	if p.Right != nil {
		p.Right.PreTraverse()
	}
}

// 中序遍历

func (p *AVLNode) InTraverse() {
	if p == nil {
		return
	}
	if p.Left != nil {
		p.Left.InTraverse()
	}
	fmt.Printf("%d ", p.Data)
	if p.Right != nil {
		p.Right.InTraverse()
	}
}

// 后序遍历

func (p *AVLNode) PostTraverse() {
	if p == nil {
		return
	}
	if p.Left != nil {
		p.Left.PostTraverse()
	}
	if p.Right != nil {
		p.Right.PostTraverse()
	}
	fmt.Printf("%d ", p.Data)
}

func main() {
	var append_list = []int{2, 3, 5, 6, 10, 11, 17}

	var remove_list = []int{5, 1}
	treeNode := &AVLNode{
		Left:   nil,
		Right:  nil,
		Height: 0,
		Data:   1,
	}
	// 循环添加元素
	for i := 0; i < len(append_list); i++ {
		treeNode = treeNode.Insert(append_list[i])
	}
	fmt.Println("------循环插入数据--------")
	fmt.Println("------中序遍历--------")
	treeNode.InTraverse()
	fmt.Println("\n")
	// 循环删除元素
	for i := 0; i < len(remove_list); i++ {
		treeNode = treeNode.Remove(remove_list[i])
	}
	fmt.Println("------删除后--------")
	fmt.Println("------前序遍历--------")
	treeNode.PreTraverse()
	fmt.Println("\n")

	fmt.Println("------中序遍历--------")
	treeNode.InTraverse()
	fmt.Println("\n")

	fmt.Println("------后序遍历--------")
	treeNode.PostTraverse()
	fmt.Println("\n")
}
