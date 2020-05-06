package main

type stormtrooper struct {
	member string
	weapon string
	armor  string
}

var troop []stormtrooper

func enroll(troopers ...stormtrooper) {
	for _, trooper := range troopers {
		if !trooper.isRenegade() {
			troop = append(troop, trooper)
		}
	}
}

func (trooper stormtrooper) isRenegade() bool {
	if trooper.member == "FN-2187" {
		return true
	}
	return false
}

func main() {
	clone1 := stormtrooper{
		member: "TK-14057",
		weapon: "E-11 blaster rifle",
		armor:  "plastoid",
	}
	enroll(clone1, )
}
