package solutions

import (
	"container/heap"
)

type TaskManager struct {
	h     *TaskHeap
	tasks map[int]*Task
}

type Task struct {
	taskId   int
	userId   int
	priority int
	index    int
}

type TaskHeap []*Task

func (t TaskHeap) Len() int { return len(t) }
func (t TaskHeap) Less(i, j int) bool {
	if t[i].priority != t[j].priority {
		return t[i].priority > t[j].priority
	}
	return t[i].taskId > t[j].taskId
}

func (t TaskHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
	t[i].index = i
	t[j].index = j
}

func (t *TaskHeap) Push(x any) {
	task := x.(*Task)
	task.index = len(*t)
	*t = append(*t, task)
}

func (t *TaskHeap) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*t = old[:n-1]
	return x
}

func TaskManagerConstructor(tasks [][]int) TaskManager {
	tl := &TaskHeap{}
	m := make(map[int]*Task)
	for _, task := range tasks {
		t := &Task{
			userId:   task[0],
			taskId:   task[1],
			priority: task[2],
		}
		m[task[1]] = t
		heap.Push(tl, t)
	}

	return TaskManager{
		tasks: m,
		h:     tl,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	t := &Task{
		userId:   userId,
		taskId:   taskId,
		priority: priority,
	}
	this.tasks[taskId] = t
	heap.Push(this.h, t)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	t := this.tasks[taskId]
	t.priority = newPriority
	heap.Fix(this.h, t.index)
}

func (this *TaskManager) Rmv(taskId int) {
	t := this.tasks[taskId]
	heap.Remove(this.h, t.index)
	delete(this.tasks, taskId)
}

func (this *TaskManager) ExecTop() int {
	if this.h.Len() > 0 {
		t := heap.Pop(this.h).(*Task)
		delete(this.tasks, t.taskId)
		return t.userId
	}
	return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
