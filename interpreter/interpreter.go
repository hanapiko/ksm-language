package interpreter

import (
	"fmt"
	"strconv"
	"strings"

	"ksm/parser"
)

const (
	blue    = "\033[34m"
	reset   = "\033[0m"
	magenta = "\033[35m"
	green   = "\033[32m"
	red     = "\033[31m"
)

type Interpreter struct {
	variables map[string]string
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		variables: make(map[string]string),
	}
}

func (i *Interpreter) Interpret(node *parser.Node) error {
	if node == nil {
		return nil
	}

	switch node.Type {
	case parser.NodeVarDecl:
		parts := strings.Split(node.Literal, "=")
		if len(parts) != 2 {
			return fmt.Errorf(blue+"invalid variable declaration: %s"+reset, node.Literal)
		}
		varName := strings.TrimSpace(parts[0])
		varValue := strings.TrimSpace(parts[1])
		i.variables[varName] = varValue
		fmt.Printf(blue+"Variable Declaration: %s = %s\n"+reset, varName, varValue)

	case parser.NodePrint:
		value := strings.TrimPrefix(node.Literal, "print ")
		fmt.Printf(magenta+"Print Statement: %s\n"+reset, i.evaluateExpression(value))

	case parser.NodeIf:
		condition := strings.TrimPrefix(node.Literal, "if ")
		if i.evaluateCondition(condition) {
			fmt.Println(green+"If Statement (True):"+reset, condition)
			for _, child := range node.Children {
				if err := i.Interpret(child); err != nil {
					return err
				}
			}
		} else {
			fmt.Println(red+"If Statement (False):"+reset, condition)
		}

	case parser.NodeOtherwise:
		fmt.Println("Otherwise Statement")
		for _, child := range node.Children {
			if err := i.Interpret(child); err != nil {
				return err
			}
		}

	case parser.NodeBlock:
		for _, child := range node.Children {
			if err := i.Interpret(child); err != nil {
				return err
			}
		}
	}

	return nil
}

func (i *Interpreter) evaluateExpression(expr string) string {
	if val, ok := i.variables[expr]; ok {
		return val
	}
	return expr
}

func (i *Interpreter) evaluateCondition(condition string) bool {
	parts := strings.Split(condition, " ")
	if len(parts) != 3 {
		return false
	}

	left := i.evaluateExpression(parts[0])
	operator := parts[1]
	right := i.evaluateExpression(parts[2])

	leftNum, leftErr := strconv.Atoi(left)
	rightNum, rightErr := strconv.Atoi(right)

	if leftErr == nil && rightErr == nil {
		switch operator {
		case ">":
			return leftNum > rightNum
		case "<":
			return leftNum < rightNum
		case "==":
			return leftNum == rightNum
		}
	}

	return false
}
