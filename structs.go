package json2docs

type Format struct {
	Header     []HeaderField
	BodyHeader []BodyHeaderField
	BodyFields []BodyField
	Summary    []SummaryField
}

type HeaderField struct {
	Line int
	Text string
}

type BodyHeaderField struct {
	Index int
	Text  string
	Width int
}

type BodyField struct {
	Index int    `json:"Index"`
	Field string `json:"Field"`
	Width int    `json:"Width"`
}

type SummaryField struct {
	Index int
	Text  string
	Width int
}
