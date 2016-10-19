package main

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestTargetEvaluate(t *testing.T) {
	target := &Target{ni: 5, no: 5, name: "echo"}
	z := []float64{0.0, 0.25, 0.5, 0.75, 1.0}
	u, _ := target.evaluate(z)
	assert.Equal(u, z, t)
}
