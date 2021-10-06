package main

func _(){
	p := Person{}
	var bed Bed
	bed = p
	bed.getup(1)
}

type Bed interface {
	sleep(int) string
	getup(int) string
}

type Person struct {

}

func (p Person) sleep(i int) string {
	return "implement me"
}

func (p Person) getup(i int) string {
	return "implement me"
}

