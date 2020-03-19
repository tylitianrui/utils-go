package map2struct

type MapParser interface {
    Parse() error

}


type maper struct {
    originMap  map[string]interface{}

}

// 解析
func (m *maper) Parse() error {
    panic("implement me")
}

func NewMaper() *maper {
    return &maper{
    }
}


// 暴力加载data
func (m *maper)ForceLoad(data  map[string]interface{}) *maper {
    return &maper{
        originMap:data,
    }
}



// 加载data
func (m *maper)Load(data  map[string]interface{}) (*maper,bool) {
    if m.originMap !=nil{
        return m,false
    }
    return &maper{
        originMap:data,
    },true
}

