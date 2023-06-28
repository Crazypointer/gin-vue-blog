package models

type FullTextModel struct {
	ID    string `json:"id" structs:"id"`       // es的id
	Key   string `json:"key"`                   // 文章关联的id
	Title string `json:"title" structs:"title"` // 文章标题
	Slug  string `json:"slug" structs:"slug"`   // 标题的跳转地址
	Body  string `json:"body" structs:"body"`   // 文章内容
}

func (FullTextModel) Index() string {
	return "full_text_index"
}

func (FullTextModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
	  "key": {
        "type": "keyword"
      },
      "title": { 
        "type": "text"
      },
      "slug": { 
        "type": "keyword"
      },
      "body": { 
        "type": "text"
      }
    }
  }
}
`
}
