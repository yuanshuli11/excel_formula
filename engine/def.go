package engine

import "math"

const (
	RadianMode = iota
	AngleMode
)

type defS struct {
	argc int
	fun  func(expr ...ExprAST) float64
}

// enum "RadianMode", "AngleMode"
var TrigonometricMode = RadianMode

var defConst = map[string]float64{
	"pi": math.Pi,
}

var defFunc map[string]defS

func init() {
	defFunc = map[string]defS{
		"sin": {1, defSin},
		"cos": {1, defCos},
		"tan": {1, defTan},
		"cot": {1, defCot},
		"sec": {1, defSec},
		"csc": {1, defCsc},

		"abs":   {1, defAbs},
		"ceil":  {1, defCeil},
		"floor": {1, defFloor},
		"round": {1, defRound},
		"sqrt":  {1, defSqrt},
		"cbrt":  {1, defCbrt},

		"noerr": {1, defNoerr},

		"max": {2, defMax},
		"min": {2, defMin},
		"ifs": {-1, defIfs},
	}
}

// sin(pi/2) = 1
func defSin(expr ...ExprAST) float64 {
	return math.Sin(expr2Radian(expr[0]))
}

// sin(pi/2) = 1
func defIfs(expr ...ExprAST) float64 {
	length := len(expr)
	for i, v := range expr {
		if i+1 == length {
			return float64(0)
		}
		if ExprASTResult(v) == 1 {
			return ExprASTResult(expr[i+1])
		}
	}
	return float64(0)
}

// cos(0) = 1
func defCos(expr ...ExprAST) float64 {
	return math.Cos(expr2Radian(expr[0]))
}

// tan(pi/4) = 1
func defTan(expr ...ExprAST) float64 {
	return math.Tan(expr2Radian(expr[0]))
}

// cot(pi/4) = 1
func defCot(expr ...ExprAST) float64 {
	return 1 / defTan(expr...)
}

// sec(0) = 1
func defSec(expr ...ExprAST) float64 {
	return 1 / defCos(expr...)
}

// csc(pi/2) = 1
func defCsc(expr ...ExprAST) float64 {
	return 1 / defSin(expr...)
}

// abs(-2) = 2
func defAbs(expr ...ExprAST) float64 {
	return math.Abs(ExprASTResult(expr[0]))
}

// ceil(4.2) = ceil(4.8) = 5
func defCeil(expr ...ExprAST) float64 {
	return math.Ceil(ExprASTResult(expr[0]))
}

// floor(4.2) = floor(4.8) = 4
func defFloor(expr ...ExprAST) float64 {
	return math.Floor(ExprASTResult(expr[0]))
}

// round(4.2) = 4
// round(4.6) = 5
func defRound(expr ...ExprAST) float64 {
	return math.Round(ExprASTResult(expr[0]))
}

// sqrt(4) = 2
func defSqrt(expr ...ExprAST) float64 {
	return math.Sqrt(ExprASTResult(expr[0]))
}

// cbrt(27) = 3
func defCbrt(expr ...ExprAST) float64 {
	return math.Cbrt(ExprASTResult(expr[0]))
}

// max(2, 3) = 3
func defMax(expr ...ExprAST) float64 {
	return math.Max(p2(expr...))
}

// max(2, 3) = 2
func defMin(expr ...ExprAST) float64 {
	return math.Min(p2(expr...))
}

func p2(expr ...ExprAST) (float64, float64) {
	return ExprASTResult(expr[0]), ExprASTResult(expr[1])
}

// noerr(1/0) = 0
// noerr(2.5/(1-1)) = 0
func defNoerr(expr ...ExprAST) (r float64) {
	defer func() {
		if e := recover(); e != nil {
			r = 0
		}
	}()
	return ExprASTResult(expr[0])
}
