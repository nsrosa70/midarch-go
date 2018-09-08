package proxy

type Proxy struct {
	Host string
	Port int
	ObjectId int
	TypeElem interface{}
}

//func Invoke(any interface{}, name string, args ... interface{}){
//
//	shared.Invoke(reflect.ValueOf(any).FieldByName("TypeElem"),name,args)
//}
