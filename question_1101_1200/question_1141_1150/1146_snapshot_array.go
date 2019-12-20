package question_1141_1150

// 1146. 快照数组
// https://leetcode-cn.com/problems/snapshot-array
// Topics: 数组

type SnapshotArray struct {
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{}
}

func (this *SnapshotArray) Set(index int, val int) {

}

func (this *SnapshotArray) Snap() int {
	return 0
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
