package main

// source: https://leetcode.com/problems/design-underground-system/

type path struct {
	src  string
	time int
}

type UndergroundSystem struct {
	time    map[string]map[string][]int
	clients map[int]path
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		time:    make(map[string]map[string][]int),
		clients: make(map[int]path),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.clients[id] = path{
		src:  stationName,
		time: t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if this.time[this.clients[id].src] == nil {
		this.time[this.clients[id].src] = make(map[string][]int)

	}
	this.time[this.clients[id].src][stationName] = append(this.time[this.clients[id].src][stationName], t-this.clients[id].time)
	delete(this.clients, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	var sum, cnt int
	for _, i := range this.time[startStation][endStation] {
		sum += i
		cnt++
	}
	return float64(sum) / float64(cnt)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
