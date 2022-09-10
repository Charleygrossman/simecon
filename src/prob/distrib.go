package prob

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Distribution represents a probability distribution.
type Distribution interface {
	// Generate returns a random variable
	// that follows the distribution.
	Generate() float64
	// Indicate generates a random variable x and returns
	// an indicator random variable represented by
	// true = 1, false = 0, if x satisfies a success event.
	Indicate() bool
}

type DistribType = string

const (
	DistribExp  DistribType = "exponential"
	DistribNorm DistribType = "normal"
	DistribUni  DistribType = "uniform"
)

var DistribTypes = []DistribType{
	DistribExp,
	DistribNorm,
	DistribUni,
}

type DistribTypeError struct {
	Type string
}

func NewDistribTypeError(distribType string) *DistribTypeError {
	return &DistribTypeError{Type: distribType}
}

func (e DistribTypeError) Error() string {
	return fmt.Sprintf("unsupported distribution type: supported=%s got=%s", strings.Join(DistribTypes, ", "), e.Type)
}

type Exponential struct {
	Prob   float64
	Lambda float64
}

func NewExponential(prob, lambda float64) Exponential {
	return Exponential{
		Prob:   prob,
		Lambda: lambda,
	}
}

func (e Exponential) Generate() float64 {
	return rand.ExpFloat64() / e.Lambda
}

func (e Exponential) Indicate() bool {
	return e.Prob >= normalizedFloat64(e.Generate(), 0, math.MaxFloat64)
}

type Normal struct {
	Prob   float64
	Mean   float64
	StdDev float64
}

func NewNormal(prob, mean, stdDev float64) Normal {
	return Normal{
		Prob:   prob,
		Mean:   mean,
		StdDev: stdDev,
	}
}

func (n Normal) Generate() float64 {
	return rand.NormFloat64()*n.StdDev + n.Mean
}

func (n Normal) Indicate() bool {
	return n.Prob >= normalizedFloat64(n.Generate(), -math.MaxFloat64, math.MaxFloat64)
}

type Uniform struct {
	Prob float64
}

func NewUniform(prob float64) Uniform {
	return Uniform{Prob: prob}
}

func (u Uniform) Generate() float64 {
	return rand.Float64()
}

func (u Uniform) Indicate() bool {
	return u.Prob >= u.Generate()
}

func normalizedFloat64(v, min, max float64) float64 {
	return (v - min) / (max - min)
}
