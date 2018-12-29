package aspect

// Build generate AOP code base on rules
func Build(aspects ...Aspect) {
	for _, a := range aspects {
		rule := a.AspectOf()
		if rule != nil {
			buildWith(rule)
		}
	}
	// TODO
}

func buildWith(rule interface{}) {
	// TODO
}
