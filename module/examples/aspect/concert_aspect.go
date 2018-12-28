//+build AspectG

package aspect

import (
	. "github.com/alimy/aspectg/aspect"
)

type Audience struct {
	Performance       PointCut          `aspect:"execution(github.com/alimy/aspectg/module/examples/concert.Performance.Perform(...))"`
	ModulePerformance PointCut          `aspect:"execution(.../concert.Performance.Perform(...))"`
	WatchPerformance  ProceedingAdvisor `aspect:"ModulePerformance(),Around"`
	TakeSeats         Advisor           `aspect:"Performance(),Before"`
	Applause          Advisor           `aspect:"Performance(),AfterReturning"`
	DemandRefund      Advisor           `aspect:"Performance,AfterPanic"`
	GoAway            Advisor           `aspect:"execution(.../concert.Performance.Perform(...)),After"`
}

func (a *Audience) watchPerformance(jp ProceedingJoinPoint) {
	defer func() {
		if err := recover(); err != nil {
			// AfterPanic invoke ProceedingJoinPoint method
			// TODO
		}
	}()

	// Before invoke ProceedingJoinPoint method
	// TODO

	// invoke ProceedingJoinPoint method
	jp.Proceed()

	// After invoke ProceedingJoinPoint method
	// TODO
}

// TakeSeats's advisor
func (a *Audience) takeSeats(jp JoinPoint) {
	// TODO
}

// Applause's advisor
func (a *Audience) applause(jp JoinPoint) {
	// TODO
}

// DemandRefund's advisor
func (a *Audience) demandRefund(jp JoinPoint) {
	// TODO
}

// GoAway's advisor
func (a *Audience) goAway(jp JoinPoint) {
	// TODO
}

// Advisor make a aspect rules instance
func (a *Audience) Aspect() interface{} {
	a.WatchPerformance = ProceedingAdvisorFunc(a.watchPerformance)
	a.TakeSeats = AdvisorFunc(a.takeSeats)
	a.Applause = AdvisorFunc(a.applause)
	a.DemandRefund = AdvisorFunc(a.demandRefund)
	a.GoAway = AdvisorFunc(a.goAway)

	return a
}
