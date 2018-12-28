package aspect

// JoinPoint a point that used in Advisor
type JoinPoint interface {
	Args() map[string]interface{} // return point argument slice
	// TODO
}

// ProceedingJoinPoint a point that  used in ProceedingAdvisor
type ProceedingJoinPoint interface {
	JoinPoint
	Proceed() // Proceed invoke ProceedingJoinPoint's method
	// TODO
}

// PointCut used in tag JoinPoint
type PointCut func()

// Aspect used in Build to obtain rules
type Aspect interface {
	Aspect() interface{}
}

// Advisor define aspect code
type Advisor interface {
	OnAdvisor(JoinPoint)
}

// ProceedingAdvisor define aspect code
type ProceedingAdvisor interface {
	OnProceedingAdvisor(ProceedingJoinPoint)
}

// AdvisorFunc used convert func(...) to Advisor
type AdvisorFunc func(JoinPoint)

// AdvisorFunc used convert func(...) to ProceedingAdvisor
type ProceedingAdvisorFunc func(ProceedingJoinPoint)

// OnAdvisor implement Advisor interface
func (f AdvisorFunc) OnAdvisor(jp JoinPoint) {
	f(jp)
}

// OnProceedingAdvisor implement ProceedingAdvisor interface
func (f ProceedingAdvisorFunc) OnProceedingAdvisor(jp ProceedingJoinPoint) {
	f(jp)
}
