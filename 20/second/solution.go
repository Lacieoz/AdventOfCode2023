package main

import (
	"fmt"
	"strings"
	"syscall"
	"time"
)

func main() {

	startTime := time.Now()

	var rows = strings.Split(input, "\n")
	mapModules := make(map[string]Module)
	for _, row := range rows {
		module := createModule(row)
		mapModules[module.getName()] = module
	}

	for key, value := range mapModules {
		destinations := value.getDestinations()
		for _, destination := range destinations {
			module := mapModules[destination]
			if mod, ok := module.(*Conjunction); ok {
				mod.addConnNode(key)
			}
		}
	}

	buttonModule := Button{"button"}
	mapModules[buttonModule.name] = &buttonModule

	index := 0

	mapLoops := make(map[string]int)
	findLoops := make(map[string]bool)
	loops := make(map[string]int)

	// RESOLUTION
	cond := true
	for cond {
		var pulseOps []PulseOp
		pulseOps = append(pulseOps, PulseOp{Low, "input", buttonModule.getName()})
		ind := 0
		for {
			if ind >= len(pulseOps) {
				break
			}
			pulseOp := pulseOps[ind]
			if pulseOp.to == "vd" && pulseOp.pulse == High {
				val, isOk := mapLoops[pulseOp.from]
				_, isOk2 := findLoops[pulseOp.from]
				if !isOk {
					mapLoops[pulseOp.from] = index
				} else if isOk && !isOk2 {
					findLoops[pulseOp.from] = true
					loops[pulseOp.from] = index - val
					if len(loops) == 4 {
						cond = false
						break
					}
				}
			}
			if pulseOp.pulse == Low && pulseOp.to == "rx" {
				cond = false
				break
			}
			currModule, isOk := mapModules[pulseOp.to]
			if isOk {
				resPulseOps := currModule.calcRes(pulseOp)
				if len(resPulseOps) > 0 {
					pulseOps = append(pulseOps, resPulseOps...)
				}
			}
			ind++
		}
		index++
	}

	var vals []int
	for _, val := range loops {
		vals = append(vals, val)
	}
	res := LCM(vals[0], vals[1], vals[2:]...)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Your function took %s\n", elapsedTime)

	fmt.Println(res)
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func createModule(row string) Module {
	row = strings.Replace(row, " ", "", -1)
	splitted := strings.Split(row, "->")

	to := splitted[1]
	destinations := strings.Split(to, ",")

	from := splitted[0]
	if string(from[0]) == "%" {
		return &Flipflop{from[1:],
			false,
			destinations}
	} else if string(from[0]) == "&" {
		var inputNodes []string
		return &Conjunction{from[1:],
			make(map[string]Pulse),
			destinations,
			inputNodes}
	} else if from == "broadcaster" {
		return &Broadcaster{from,
			destinations}
	}

	fmt.Println("NOTHING RECOGNIZED")
	syscall.Exit(1)
	return &Button{"button"}
}

type Module interface {
	calcRes(PulseOp) []PulseOp
	getName() string
	getDestinations() []string
}
type Flipflop struct {
	name         string
	onoff        bool
	destinations []string
}

func (t *Flipflop) calcRes(input PulseOp) []PulseOp {
	var res []PulseOp
	if input.pulse == Low {
		t.onoff = !t.onoff
		pulseVal := Low
		if t.onoff {
			pulseVal = High
		}
		for _, dest := range t.destinations {
			res = append(res, PulseOp{pulseVal, t.name, dest})
		}
	}
	return res
}
func (t Flipflop) getName() string {
	return t.name
}
func (t Flipflop) getDestinations() []string {
	return t.destinations
}

type Conjunction struct {
	name           string
	mapRecentPulse map[string]Pulse
	destinations   []string
	inputNodes     []string
}

func (t Conjunction) addConnNode(input string) {
	for _, node := range t.inputNodes {
		if node == input {
			return
		}
	}
	t.inputNodes = append(t.inputNodes, input)
	t.mapRecentPulse[input] = Low
}
func (t *Conjunction) calcRes(input PulseOp) []PulseOp {
	var res []PulseOp
	t.mapRecentPulse[input.from] = input.pulse
	for _, value := range t.mapRecentPulse {
		if value == Low {
			for _, dest := range t.destinations {
				res = append(res, PulseOp{High, t.name, dest})
			}
			return res
		}
	}
	for _, dest := range t.destinations {
		res = append(res, PulseOp{Low, t.name, dest})
	}
	return res
}
func (t Conjunction) getName() string {
	return t.name
}
func (t Conjunction) getDestinations() []string {
	return t.destinations
}

type Broadcaster struct {
	name         string
	destinations []string
}

func (t *Broadcaster) calcRes(input PulseOp) []PulseOp {
	var res []PulseOp
	for _, dest := range t.destinations {
		res = append(res, PulseOp{input.pulse, t.name, dest})
	}
	return res
}
func (t Broadcaster) getName() string {
	return t.name
}
func (t Broadcaster) getDestinations() []string {
	return t.destinations
}

type Button struct {
	name string
}

func (t *Button) calcRes(input PulseOp) []PulseOp {
	var res []PulseOp
	res = append(res, PulseOp{Low, t.name, "broadcaster"})
	return res
}
func (t Button) getName() string {
	return t.name
}
func (t Button) getDestinations() []string {
	var destinations []string
	destinations = append(destinations, "broadcaster")
	return destinations
}

type Pulse int

const (
	Low Pulse = iota
	High
)

type PulseOp struct {
	pulse Pulse
	from  string
	to    string
}

const times = 1000
const input = `%jb -> fz
%xz -> ck, bg
%xm -> qt, cs
%df -> hc, lq
%mt -> sx
%fr -> ks, hc
%tn -> pf
%gt -> pp, kb
%jn -> ck, nz
%td -> kz
&rd -> vd
%pp -> gv, kb
&qt -> jb, vx, bt, gh, td, gb
%ms -> xz
%vx -> fp
%rb -> ck, mt
%nz -> hh
%fp -> rp, qt
%gd -> gc
%gv -> kb
%nl -> cc, hc
%cs -> qt
%kz -> jb, qt
%vg -> fr, hc
%zq -> qt, xm
%pv -> ps
&bt -> vd
%ps -> kb, rf
%hh -> ck, ms
broadcaster -> gn, gb, rb, df
%gh -> td
%rf -> kb, nm
%rp -> qt, gh
%gc -> kb, pv
%gb -> vx, qt
%rq -> ck, ts
%nm -> gt
%gn -> kb, tn
&ck -> nz, fv, rb, sx, ms, mt
&fv -> vd
%cc -> vg
%bg -> ck, rq
&hc -> qr, ch, df, dj, cc, rd
%qr -> dj
%gq -> hc, ch
&pr -> vd
%ks -> lc, hc
%dj -> nl
%fz -> qt, zq
%lq -> gq, hc
&kb -> pv, pr, tn, nm, pf, gn, gd
%ts -> ck
%lc -> hc
%jl -> ck, jn
%sx -> jl
%pf -> gd
&vd -> rx
%ch -> qr`

const result = 228300182686739
