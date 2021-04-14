package main

import (
	"bufio"
	"fmt"
	"github.com/yuanshuli11/excel_formula/engine"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	engine.RegFunction("if", -1, func(expr ...engine.ExprAST) float64 {

		length := len(expr)
		if length < 2 {
			return 0
		}
		if engine.ExprASTResult(expr[0]) == 0 {
			if length == 2 {
				return 0
			}
			return engine.ExprASTResult(expr[2])
		} else {
			return engine.ExprASTResult(expr[1])
		}
	})

	engine.RegFunction("roundup", 2, func(expr ...engine.ExprAST) float64 {
		if engine.ExprASTResult(expr[1]) <= 0 {
			return math.Round(engine.ExprASTResult(expr[0]))
		}
		n := strings.Split(fmt.Sprintf("%v", engine.ExprASTResult(expr[1])), ".")

		value, _ := strconv.ParseFloat(fmt.Sprintf("%."+n[0]+"f", engine.ExprASTResult(expr[0])), 64)
		return value

	})
	engine.RegFunction("rounddown", 2, func(expr ...engine.ExprAST) float64 {
		//先多保留一位小数
		if engine.ExprASTResult(expr[1]) <= 0 {
			return math.Ceil(engine.ExprASTResult(expr[0]))
		}
		n := strings.Split(fmt.Sprintf("%v", engine.ExprASTResult(expr[1])+1), ".")
		vstr := fmt.Sprintf("%."+n[0]+"f", engine.ExprASTResult(expr[0]))
		value, _ := strconv.ParseFloat(vstr[:len(vstr)-1], 64)
		return value

	})
	engine.RegFunction("ceiling", 1, func(expr ...engine.ExprAST) float64 {
		return math.Ceil(engine.ExprASTResult(expr[0]))
	})
	engine.RegFunction("max", -1, func(expr ...engine.ExprAST) float64 {

		max := -math.MaxFloat64
		for _, v := range expr {
			if engine.ExprASTResult(v) > max {
				max = engine.ExprASTResult(v)
			}
		}
		return max

	})

	engine.RegFunction("and", -1, func(expr ...engine.ExprAST) float64 {
		for _, v := range expr {
			if engine.ExprASTResult(v) == 0 {
				return 0
			}
		}
		return 1

	})
	engine.RegFunction("or", -1, func(expr ...engine.ExprAST) float64 {
		for _, v := range expr {
			if engine.ExprASTResult(v) > 0 {
				return float64(1)
			}
		}
		return float64(0)
	})
}
func main() {
	//varkeys := []string{"ccccc", "aaa", "ccc"}
	//sort(varkeys)
	//
	//fmt.Println(varkeys)
	outValue, e := engine.ParseAndExec("IF(216=0,0,IF(216<=216,6,ROUNDUP(216/432,0)*12))")
	fmt.Println("====", outValue, e)
	//loop()
}

// input loop
func loop() {
	engine.RegFunction("double", 1, func(expr ...engine.ExprAST) float64 {
		return engine.ExprASTResult(expr[0]) * 2
	})
	for {
		fmt.Print("input /> ")
		f := bufio.NewReader(os.Stdin)
		s, err := f.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if s == "exit" || s == "quit" || s == "q" {
			fmt.Println("bye")
			break
		}
		start := time.Now()
		exec(s)
		cost := time.Since(start)
		fmt.Println("time: " + cost.String())
	}
}

// call engine
func exec(exp string) {
	// input text -> []token
	toks, err := engine.Parse(exp)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	// []token -> AST Tree
	ast := engine.NewAST(toks, exp)
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		return
	}
	// AST builder
	ar := ast.ParseExpression()
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		return
	}
	fmt.Printf("ExprAST: %+v\n", ar)
	// catch runtime errors
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("ERROR: ", e)
		}
	}()
	// AST traversal -> result
	r := engine.ExprASTResult(ar)
	fmt.Println("progressing ...\t", r)
	fmt.Printf("%s = %v\n", exp, r)
}
