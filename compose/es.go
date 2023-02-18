package compose

import (
	"fmt"
	"gopkg.in/olivere/elastic.v2"
)

var EsClient *elastic.Client

func InitEs() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://43.136.39.33:9200/"))
	if err != nil {
		panic("connect es base" + err.Error())
		return
	}
	fmt.Println("client", client)
	EsClient = client
}
