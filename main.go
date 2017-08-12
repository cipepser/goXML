package main

import (
	"encoding/xml"
	"fmt"
)

type ChiChidleyRoot314159 struct {
	ChiPeople *ChiPeople `xml:" people,omitempty" json:"people,omitempty"`
}

type ChiAge struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ChiMarried struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ChiName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ChiPeople struct {
	ChiPerson []*ChiPerson `xml:" person,omitempty" json:"person,omitempty"`
}

type ChiPerson struct {
	ChiAge     *ChiAge     `xml:" age,omitempty" json:"age,omitempty"`
	ChiMarried *ChiMarried `xml:" married,omitempty" json:"married,omitempty"`
	ChiName    *ChiName    `xml:" name,omitempty" json:"name,omitempty"`
}

func main() {
	x := []byte(`<people>

    <person>
      <name>bill</name>
      <age>37</age>
      <married>true</married>
    </person>

    <person>
      <name>sarah</name>
      <age>24</age>
      <married>false</married>
    </person>

  </people>`)

	p := &ChiPeople{}
	xml.Unmarshal(x, p)

	fmt.Println(p.ChiPerson[0].ChiAge)

}
