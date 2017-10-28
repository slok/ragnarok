package v1

import (
	"errors"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/slok/ragnarok/api"
	"github.com/slok/ragnarok/attack"
	"github.com/slok/ragnarok/log"
)

const (
	// FailureKind is the kind a failure.
	FailureKind = "failure"
	// FailureVersion is the version of the failure
	FailureVersion = "chaos/v1"
)

// FailureState is the state a failure can be.
type FailureState int

const (
	// UnknownFailureState is an unknown status.
	UnknownFailureState FailureState = iota
	// EnabledFailureState is when the failure should be making stuff.
	EnabledFailureState
	// ExecutingFailureState is when the failure its making stuff.
	ExecutingFailureState
	// RevertingFailureState is when the failure is being reverted.
	RevertingFailureState
	// DisabledFailureState is when the failure is should be not making stuff.
	DisabledFailureState
	// StaleFailureState is when the failure has go through alll the lifecycle and should be archived ((reverted already).
	StaleFailureState
	// ErroredFailureState is when the failure is not making stuff (due to an error).
	ErroredFailureState
	// ErroredRevertingFailureState is when the failure is not making stuff (due to an error reverting).
	ErroredRevertingFailureState
)

func (f FailureState) String() string {
	switch f {
	case EnabledFailureState:
		return "enabled"
	case ExecutingFailureState:
		return "executing"
	case RevertingFailureState:
		return "reverting"
	case DisabledFailureState:
		return "disabled"
	case StaleFailureState:
		return "stale"
	case ErroredFailureState:
		return "errored"
	case ErroredRevertingFailureState:
		return "erroredreverting"

	default:
		return "unknown"
	}
}

// AttackMap is a type that defines a list of map of attackers.
type AttackMap map[string]attack.Opts

// FailureMetadata has information about the object.
type FailureMetadata struct {
	ID     string // ID is the id of the Failure.
	NodeID string // NodeID is the id of the Node.
}

// FailureStatus has all the information of a failure to create an injection
type FailureStatus struct {
	CurrentState  FailureState // CurrentState is the state of the failure.
	ExpectedState FailureState // ExpectedState is the state the failure should be.
	Creation      time.Time    // Creation is when the failure injection was created.
	Executed      time.Time    // Executed is when the failure injectionwas executed.
	Finished      time.Time    // Finished is when the failure injection was reverted.
}

// FailureSpec is the specification that has the information to it can be created and applied.
type FailureSpec struct {
	// Timeout is
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// Attacks used an array so the no repeated elements of map limitation can be bypassed.
	Attacks []AttackMap `yaml:"attacks,omitempty"`
	// TODO: accuracy
}

// Failure is the way a failure is defined.
type Failure struct {
	// Metadta is additional data of a failure object.
	Metadata FailureMetadata `yaml:"metadata,omitempty"`
	// Spec has all the required data to create a Failure and use it.
	Spec FailureSpec `yaml:"spec,omitempty"`
	// Status is the current information and status of the Failure.
	Status FailureStatus `yaml:"status,omitempty"`
}

// GetObjectKind satisfies Object interface.
func (f *Failure) GetObjectKind() api.Kind {
	return FailureKind
}

// GetObjectVersion satisfies Object interface.
func (f *Failure) GetObjectVersion() api.Version {
	return FailureVersion
}

// ReadFailure reads a config yaml failure and returns a failure object.
func ReadFailure(data []byte) (Failure, error) {
	log.Debug("reading failure")
	d := &Failure{}
	err := yaml.Unmarshal(data, d)
	return *d, err
}

// ReadFailureSpec reads a config yaml failure spec and returns a failure Spec object.
func ReadFailureSpec(data []byte) (FailureSpec, error) {
	log.Debug("reading failure spec")
	s := &FailureSpec{}
	err := yaml.Unmarshal(data, s)
	return *s, err
}

// Render renders a yaml form a Failure object.
func (f *Failure) Render() ([]byte, error) {
	log.Debug("rendering failure")

	// Check if there are more than one elements on the maps of the list.
	for _, a := range f.Spec.Attacks {
		if len(a) != 1 {
			return nil, errors.New("each attack map of the attack list needs to be a single map")
		}
	}
	// Marshal to yaml
	return yaml.Marshal(f)
}

// Render renders a yaml form of a Failure Spec object.
func (f *FailureSpec) Render() ([]byte, error) {
	log.Debug("rendering failure spec")

	// Check if there are more than one elements on the maps of the list.
	for _, a := range f.Attacks {
		if len(a) != 1 {
			return nil, errors.New("each attack map of the attack list needs to be a single map")
		}
	}
	// Marshal to yaml
	return yaml.Marshal(f)
}

// UnmarshalYAML wraps yaml lib unmarshalling to have extra validations.
func (f *Failure) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Made to bypass the unmarshaling recursion.
	type plain Failure
	ff := &Failure{}
	if err := unmarshal((*plain)(ff)); err != nil {
		return err
	}

	// Check if there are more then one elements on the maps of the list.
	for _, a := range ff.Spec.Attacks {
		if len(a) != 1 {
			return errors.New("attacks format error, tip: check identantion and '-' indicator")
		}
	}

	*f = *ff
	return nil
}