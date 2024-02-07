package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i learngo-pockets/templates/internal/handlers.habitsClient -o ./mocks/habits_client_mock.go -n HabitsClientMock

import (
	"context"
	habit "learngo-pockets/templates/internal/habits"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// HabitsClientMock implements handlers.habitsClient
type HabitsClientMock struct {
	t minimock.Tester

	funcListHabits          func(ctx context.Context) (ha1 []habit.Habit, err error)
	inspectFuncListHabits   func(ctx context.Context)
	afterListHabitsCounter  uint64
	beforeListHabitsCounter uint64
	ListHabitsMock          mHabitsClientMockListHabits
}

// NewHabitsClientMock returns a mock for handlers.habitsClient
func NewHabitsClientMock(t minimock.Tester) *HabitsClientMock {
	m := &HabitsClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ListHabitsMock = mHabitsClientMockListHabits{mock: m}
	m.ListHabitsMock.callArgs = []*HabitsClientMockListHabitsParams{}

	return m
}

type mHabitsClientMockListHabits struct {
	mock               *HabitsClientMock
	defaultExpectation *HabitsClientMockListHabitsExpectation
	expectations       []*HabitsClientMockListHabitsExpectation

	callArgs []*HabitsClientMockListHabitsParams
	mutex    sync.RWMutex
}

// HabitsClientMockListHabitsExpectation specifies expectation struct of the habitsClient.ListHabits
type HabitsClientMockListHabitsExpectation struct {
	mock    *HabitsClientMock
	params  *HabitsClientMockListHabitsParams
	results *HabitsClientMockListHabitsResults
	Counter uint64
}

// HabitsClientMockListHabitsParams contains parameters of the habitsClient.ListHabits
type HabitsClientMockListHabitsParams struct {
	ctx context.Context
}

// HabitsClientMockListHabitsResults contains results of the habitsClient.ListHabits
type HabitsClientMockListHabitsResults struct {
	ha1 []habit.Habit
	err error
}

// Expect sets up expected params for habitsClient.ListHabits
func (mmListHabits *mHabitsClientMockListHabits) Expect(ctx context.Context) *mHabitsClientMockListHabits {
	if mmListHabits.mock.funcListHabits != nil {
		mmListHabits.mock.t.Fatalf("HabitsClientMock.ListHabits mock is already set by Set")
	}

	if mmListHabits.defaultExpectation == nil {
		mmListHabits.defaultExpectation = &HabitsClientMockListHabitsExpectation{}
	}

	mmListHabits.defaultExpectation.params = &HabitsClientMockListHabitsParams{ctx}
	for _, e := range mmListHabits.expectations {
		if minimock.Equal(e.params, mmListHabits.defaultExpectation.params) {
			mmListHabits.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmListHabits.defaultExpectation.params)
		}
	}

	return mmListHabits
}

// Inspect accepts an inspector function that has same arguments as the habitsClient.ListHabits
func (mmListHabits *mHabitsClientMockListHabits) Inspect(f func(ctx context.Context)) *mHabitsClientMockListHabits {
	if mmListHabits.mock.inspectFuncListHabits != nil {
		mmListHabits.mock.t.Fatalf("Inspect function is already set for HabitsClientMock.ListHabits")
	}

	mmListHabits.mock.inspectFuncListHabits = f

	return mmListHabits
}

// Return sets up results that will be returned by habitsClient.ListHabits
func (mmListHabits *mHabitsClientMockListHabits) Return(ha1 []habit.Habit, err error) *HabitsClientMock {
	if mmListHabits.mock.funcListHabits != nil {
		mmListHabits.mock.t.Fatalf("HabitsClientMock.ListHabits mock is already set by Set")
	}

	if mmListHabits.defaultExpectation == nil {
		mmListHabits.defaultExpectation = &HabitsClientMockListHabitsExpectation{mock: mmListHabits.mock}
	}
	mmListHabits.defaultExpectation.results = &HabitsClientMockListHabitsResults{ha1, err}
	return mmListHabits.mock
}

// Set uses given function f to mock the habitsClient.ListHabits method
func (mmListHabits *mHabitsClientMockListHabits) Set(f func(ctx context.Context) (ha1 []habit.Habit, err error)) *HabitsClientMock {
	if mmListHabits.defaultExpectation != nil {
		mmListHabits.mock.t.Fatalf("Default expectation is already set for the habitsClient.ListHabits method")
	}

	if len(mmListHabits.expectations) > 0 {
		mmListHabits.mock.t.Fatalf("Some expectations are already set for the habitsClient.ListHabits method")
	}

	mmListHabits.mock.funcListHabits = f
	return mmListHabits.mock
}

// When sets expectation for the habitsClient.ListHabits which will trigger the result defined by the following
// Then helper
func (mmListHabits *mHabitsClientMockListHabits) When(ctx context.Context) *HabitsClientMockListHabitsExpectation {
	if mmListHabits.mock.funcListHabits != nil {
		mmListHabits.mock.t.Fatalf("HabitsClientMock.ListHabits mock is already set by Set")
	}

	expectation := &HabitsClientMockListHabitsExpectation{
		mock:   mmListHabits.mock,
		params: &HabitsClientMockListHabitsParams{ctx},
	}
	mmListHabits.expectations = append(mmListHabits.expectations, expectation)
	return expectation
}

// Then sets up habitsClient.ListHabits return parameters for the expectation previously defined by the When method
func (e *HabitsClientMockListHabitsExpectation) Then(ha1 []habit.Habit, err error) *HabitsClientMock {
	e.results = &HabitsClientMockListHabitsResults{ha1, err}
	return e.mock
}

// ListHabits implements handlers.habitsClient
func (mmListHabits *HabitsClientMock) ListHabits(ctx context.Context) (ha1 []habit.Habit, err error) {
	mm_atomic.AddUint64(&mmListHabits.beforeListHabitsCounter, 1)
	defer mm_atomic.AddUint64(&mmListHabits.afterListHabitsCounter, 1)

	if mmListHabits.inspectFuncListHabits != nil {
		mmListHabits.inspectFuncListHabits(ctx)
	}

	mm_params := &HabitsClientMockListHabitsParams{ctx}

	// Record call args
	mmListHabits.ListHabitsMock.mutex.Lock()
	mmListHabits.ListHabitsMock.callArgs = append(mmListHabits.ListHabitsMock.callArgs, mm_params)
	mmListHabits.ListHabitsMock.mutex.Unlock()

	for _, e := range mmListHabits.ListHabitsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ha1, e.results.err
		}
	}

	if mmListHabits.ListHabitsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmListHabits.ListHabitsMock.defaultExpectation.Counter, 1)
		mm_want := mmListHabits.ListHabitsMock.defaultExpectation.params
		mm_got := HabitsClientMockListHabitsParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmListHabits.t.Errorf("HabitsClientMock.ListHabits got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmListHabits.ListHabitsMock.defaultExpectation.results
		if mm_results == nil {
			mmListHabits.t.Fatal("No results are set for the HabitsClientMock.ListHabits")
		}
		return (*mm_results).ha1, (*mm_results).err
	}
	if mmListHabits.funcListHabits != nil {
		return mmListHabits.funcListHabits(ctx)
	}
	mmListHabits.t.Fatalf("Unexpected call to HabitsClientMock.ListHabits. %v", ctx)
	return
}

// ListHabitsAfterCounter returns a count of finished HabitsClientMock.ListHabits invocations
func (mmListHabits *HabitsClientMock) ListHabitsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmListHabits.afterListHabitsCounter)
}

// ListHabitsBeforeCounter returns a count of HabitsClientMock.ListHabits invocations
func (mmListHabits *HabitsClientMock) ListHabitsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmListHabits.beforeListHabitsCounter)
}

// Calls returns a list of arguments used in each call to HabitsClientMock.ListHabits.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmListHabits *mHabitsClientMockListHabits) Calls() []*HabitsClientMockListHabitsParams {
	mmListHabits.mutex.RLock()

	argCopy := make([]*HabitsClientMockListHabitsParams, len(mmListHabits.callArgs))
	copy(argCopy, mmListHabits.callArgs)

	mmListHabits.mutex.RUnlock()

	return argCopy
}

// MinimockListHabitsDone returns true if the count of the ListHabits invocations corresponds
// the number of defined expectations
func (m *HabitsClientMock) MinimockListHabitsDone() bool {
	for _, e := range m.ListHabitsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListHabitsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListHabitsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcListHabits != nil && mm_atomic.LoadUint64(&m.afterListHabitsCounter) < 1 {
		return false
	}
	return true
}

// MinimockListHabitsInspect logs each unmet expectation
func (m *HabitsClientMock) MinimockListHabitsInspect() {
	for _, e := range m.ListHabitsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to HabitsClientMock.ListHabits with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListHabitsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListHabitsCounter) < 1 {
		if m.ListHabitsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to HabitsClientMock.ListHabits")
		} else {
			m.t.Errorf("Expected call to HabitsClientMock.ListHabits with params: %#v", *m.ListHabitsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcListHabits != nil && mm_atomic.LoadUint64(&m.afterListHabitsCounter) < 1 {
		m.t.Error("Expected call to HabitsClientMock.ListHabits")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *HabitsClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockListHabitsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *HabitsClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *HabitsClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockListHabitsDone()
}
