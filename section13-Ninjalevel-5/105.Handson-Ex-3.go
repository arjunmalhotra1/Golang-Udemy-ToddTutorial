package main

import (
	"fmt"
)

type veh struct{
	doors int
	color string
}

type truck struct{
	veh
	fourwheel bool
}

type sedan struct{
	//vehicle veh // This is another way too. He did is using simply "veh"
	//VIMT To not this will not make the field doors part of the sedan
	//but only "veh" will make the field doors part of sedan"
	veh
	luxury bool
}

func main() {
	t := truck {
		veh : veh{
			doors : 2,
			color :"white",
		},
		fourwheel : true,
	}


	x := sedan {
		//vehicle : veh{
		veh:veh{
			doors : 4,
			color : "orange",
		},
		luxury : true,
	}

	fmt.Println(t)
	fmt.Println(x)
	fmt.Println(t.doors)
	//fmt.Println(x.vehicle.doors) // If you do vehicle veh then this wil print the doors.
	fmt.Println(x.doors)

}
