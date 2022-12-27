package abstract_factory

import (
	div "design-pattern-go/abstract_factory/div"
	f "design-pattern-go/abstract_factory/factory"
	list "design-pattern-go/abstract_factory/list"
	"fmt"
)

func GetFactory(factoryName string) (f.FactoryInterface, error) {
	if factoryName == "list" {
		return list.NewListFactory(), nil
	} else if factoryName == "div" {
		return div.NewDivFactory(), nil
	} else {
		return nil, fmt.Errorf("specified factory does not exist")
	}
}
