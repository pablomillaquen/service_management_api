package workorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusValidation(t *testing.T) {
	assert.True(t, StatusPending.IsValid())
	assert.True(t, StatusAssigned.IsValid())
	assert.True(t, StatusInProgress.IsValid())
	assert.True(t, StatusWaitingParts.IsValid())
	assert.True(t, StatusCompleted.IsValid())
	assert.True(t, StatusCancelled.IsValid())
	assert.False(t, Status("invalid").IsValid())
}

func TestPriorityValidation(t *testing.T) {
	assert.True(t, PriorityLow.IsValid())
	assert.True(t, PriorityMedium.IsValid())
	assert.True(t, PriorityHigh.IsValid())
	assert.True(t, PriorityCritical.IsValid())
	assert.False(t, Priority("invalid").IsValid())
}

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from Status
		to   Status
		ok   bool
	}{
		{StatusPending, StatusAssigned, true},
		{StatusPending, StatusCancelled, true},
		{StatusAssigned, StatusInProgress, true},
		{StatusAssigned, StatusCancelled, true},
		{StatusInProgress, StatusWaitingParts, true},
		{StatusInProgress, StatusCompleted, true},
		{StatusInProgress, StatusCancelled, true},
		{StatusWaitingParts, StatusInProgress, true},
		{StatusWaitingParts, StatusCancelled, true},
		{StatusCompleted, StatusPending, false},
		{StatusCancelled, StatusAssigned, false},
		{StatusPending, StatusCompleted, false},
		{StatusPending, StatusInProgress, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.from)+"_to_"+string(tt.to), func(t *testing.T) {
			assert.Equal(t, tt.ok, IsValidTransition(tt.from, tt.to))
		})
	}
}
