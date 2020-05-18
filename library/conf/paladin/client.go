package paladin

type Getter interface {
	Get(string) *Value
	GetAll() *Map
}

type Setter interface {
	Set(string) error
}

type Client interface {
	Getter
}
