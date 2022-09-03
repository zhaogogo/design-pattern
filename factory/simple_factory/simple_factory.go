package simple_factory

type IRuleConfigParse interface {
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

func NewIRuleConfigParse(t string) IRuleConfigParse {
	switch t {
	case "json":
		return jsonRuleConfigParse{}
	case "yaml":
		return yamlRuleConfigParse{}
	}
	return nil
}
