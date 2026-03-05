package solutions

import "sort"

type Packet struct {
	source, destination, timestamp int
}

type Router struct {
	set         map[Packet]bool
	d           map[int][]Packet
	packets     []Packet
	memoryLimit int
}

func RouterConstructor(memoryLimit int) Router {
	return Router{
		set:         make(map[Packet]bool),
		d:           make(map[int][]Packet),
		packets:     []Packet{},
		memoryLimit: memoryLimit,
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := Packet{source, destination, timestamp}
	if _, ok := this.set[packet]; ok {
		return false
	}
	if len(this.packets)+1 > this.memoryLimit {
		this.ForwardPacket()
	}
	this.set[packet] = true
	destPackets := this.d[packet.destination]
	idx := sort.Search(len(destPackets), func(i int) bool { return destPackets[i].timestamp > packet.timestamp })
	destPackets = append(destPackets, Packet{})
	copy(destPackets[idx+1:], destPackets[idx:])
	destPackets[idx] = packet
	this.d[packet.destination] = destPackets
	this.packets = append(this.packets, packet)
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.packets) == 0 {
		return []int{}
	}
	first := this.packets[0]
	delete(this.set, first)
	destPackets := this.d[first.destination]
	idx := sort.Search(len(destPackets), func(i int) bool { return destPackets[i].timestamp >= first.timestamp })
	if idx < len(destPackets) && destPackets[idx] == first {
		destPackets = append(destPackets[:idx], destPackets[idx+1:]...)
	}
	if len(destPackets) == 0 {
		delete(this.d, first.destination)
	} else {
		this.d[first.destination] = destPackets
	}
	this.packets = this.packets[1:]
	return []int{first.source, first.destination, first.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	packets := this.d[destination]
	left := sort.Search(len(packets), func(i int) bool { return packets[i].timestamp >= startTime })
	right := sort.Search(len(packets), func(i int) bool { return packets[i].timestamp > endTime })
	return right - left
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
