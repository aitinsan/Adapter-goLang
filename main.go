package main

type someone struct{

}
func (p *someone) insertRussianPS(socket powerSocket){
	socket.InsertRussianConductor()
}
type powerSocket interface{
	InsertRussianConductor()
}
type nokia struct{

}

func (n *nokia) InsertRussianConductor() {

}


type iphone struct{

}
func (i *iphone) insertAmericanConductor(){

}
type iphoneAdapter struct{
	applePhone *iphone
}

func (a *iphoneAdapter) InsertRussianConductor() {
	a.applePhone.insertAmericanConductor()
}



func main(){
	someone :=&someone{}
	nokia:=&nokia{}
	someone.insertRussianPS(nokia)
	applePhone:=&iphone{}
	appleAdapter:=&iphoneAdapter{
		applePhone: applePhone,
	}
	someone.insertRussianPS(appleAdapter)
}