package structs

type JSON string

type JSONEncodable interface {
	ToJSON() (JSON, error)
	FromJSON(JSON) error
}
