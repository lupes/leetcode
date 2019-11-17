package question_221_230

import "testing"

func Test_countNodes(t *testing.T) {
	root1 := &TreeNode{Val: 1}
	//root2 := &TreeNode{Val: 1,
	//	Left:  &TreeNode{Val: 21},
	//	Right: &TreeNode{Val: 22},
	//}
	//root3 := &TreeNode{Val: 1,
	//	Left: &TreeNode{Val: 21,
	//		Left:  &TreeNode{Val: 31},
	//		Right: &TreeNode{Val: 32},
	//	},
	//	Right: &TreeNode{Val: 22,
	//		Left:  &TreeNode{Val: 33},
	//		Right: &TreeNode{Val: 34},
	//	},
	//}
	//root5 := &TreeNode{Val: 1,
	//	Left: &TreeNode{Val: 21,
	//		Left: &TreeNode{Val: 31,
	//			Left: &TreeNode{Val: 41,
	//				Left:  &TreeNode{Val: 501},
	//				Right: &TreeNode{Val: 502},
	//			},
	//			Right: &TreeNode{Val: 42,
	//				Left:  &TreeNode{Val: 503},
	//				Right: &TreeNode{Val: 504},
	//			},
	//		},
	//		Right: &TreeNode{Val: 32,
	//			Left: &TreeNode{Val: 43,
	//				//Left:  &TreeNode{Val: 505},
	//				//Right: &TreeNode{Val: 506},
	//			},
	//			Right: &TreeNode{Val: 44,
	//				//Left:  &TreeNode{Val: 507},
	//				//Right: &TreeNode{Val: 508},
	//			},
	//		},
	//	},
	//	Right: &TreeNode{Val: 22,
	//		Left: &TreeNode{Val: 33,
	//			Left: &TreeNode{Val: 45,
	//				//Left:  &TreeNode{Val: 509},
	//				//Right: &TreeNode{Val: 510},
	//			},
	//			Right: &TreeNode{Val: 46,
	//				//Left:  &TreeNode{Val: 511},
	//				//Right: &TreeNode{Val: 512},
	//			},
	//		},
	//		Right: &TreeNode{Val: 34,
	//			Left: &TreeNode{Val: 47,
	//				//Left:  &TreeNode{Val: 513},
	//				//Right: &TreeNode{Val: 514},
	//			},
	//			Right: &TreeNode{Val: 48,
	//				//Left:  &TreeNode{Val: 515},
	//				//Right: &TreeNode{Val: 516},
	//			},
	//		},
	//	},
	//}
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		//{"test", nil, 0},
		{"test", root1, 1},
		//{"test", root2, 3},
		//{"test", root3, 7},
		//{"test", root5, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
