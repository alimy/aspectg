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
	OnAdvisor(JoinPoint)
}

// AdvisorFunc used convert func(...) to Advisor
type AdvisorFunc func(JoinPoint)

// AdvisorFunc used convert func(...) to ProceedingAdvisor
type ProceedingAdvisorFunc func(ProceedingJoinPoint)

// ProceedingAdvisor define aspect code
type ProceedingAdvisor interface {
	OnProceedingAdvisor(ProceedingJoinPoint)
}

// Proceed invoke ProceedingJoinPoint's method
func (p *ProceedingJoinPoint) Proceed() {
	// TODO
}

// OnAdvisor implement Advisor interface
func (f AdvisorFunc) OnAdvisor(jp JoinPoint) {
	f(jp)
}

// OnProceedingAdvisor implement ProceedingAdvisor interface
func (f ProceedingAdvisorFunc) OnProceedingAdvisor(jp ProceedingJoinPoint) {
	f(jp)
}
