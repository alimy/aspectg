package aspect

// JoinPoint a point that used in Advisor
type JoinPoint struct {
	// TODO
}

// ProceedingJoinPoint a point that  used in ProceedingAdvisor
type ProceedingJoinPoint struct {
	// TODO
}

// PointCut used in tag JoinPoint
type PointCut func()

// Advisor define aspect code
type Advisor interface {
	OnAdvisor(point JoinPoint, args ...interface{})
}

// ProceedingAdvisor define aspect code
type ProceedingAdvisor interface {
	OnProceedingAdvisor(point ProceedingJoinPoint, args ...interface{})
}

// AdvisorFunc used convert func(...) to Advisor
type AdvisorFunc func(point JoinPoint, args ...interface{})

// AdvisorFunc used convert func(...) to ProceedingAdvisor
type ProceedingAdvisorFunc func(point ProceedingJoinPoint, args ...interface{})

// Proceed invoke ProceedingJoinPoint's method
func (p *ProceedingJoinPoint) Proceed() {
	// TODO
}

// OnAdvisor implement Advisor interface
func (f AdvisorFunc) OnAdvisor(point JoinPoint, args ...interface{}) {
	f(point, args...)
}

// OnProceedingAdvisor implement ProceedingAdvisor interface
func (f ProceedingAdvisorFunc) OnProceedingAdvisor(pint ProceedingJoinPoint, args ...interface{}) {
	f(pint, args...)
}
