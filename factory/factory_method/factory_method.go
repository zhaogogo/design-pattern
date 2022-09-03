package factory_method

type IRouleConfigParse interface {
	Parse([]byte)
}

type jsonRuleConfigParse struct{}

func (j jsonRuleConfigParse) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParse struct{}

func (y yamlRuleConfigParse) Parse(data []byte) {
	panic("implement me")
}

//工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParse() IRouleConfigParse
}

type yamlRuleConfigParserFactory struct{}

func (y yamlRuleConfigParserFactory) CreateParse() IRouleConfigParse {
	//配置初始化
	return yamlRuleConfigParse{}
}

type jsonRuleConfigParserFactory struct{}

func (j jsonRuleConfigParserFactory) CreateParse() IRouleConfigParse {
	//配置初始化
	return jsonRuleConfigParse{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
