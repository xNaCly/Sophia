package core

// TODO: add string support

import (
	"fmt"
	"log"
)

type Value struct {
	isFloat bool
	Float   float64
	String  string
}

func getValue(node Node) Value {
	val := Value{}
	switch node.(type) {
	case *String:
		val.isFloat = false
		val.String = node.(*String).Token.Raw
	case *Float:
		val.isFloat = true
		val.Float = node.(*Float).Token.Float
	default:
		val.isFloat = true
		val.Float = sumOperator(node.GetToken().Type, node.(*Statement).Children)
		// log.Printf("err: unknown node of type '%T', can't get value", node)
	}
	return val
}

func printer(stmt Statement) {
	fmt.Println(Eval(stmt.Children))
}

func sumOperator(operator int, children []Node) float64 {

	val := getValue(children[0])
	if !val.isFloat {
		log.Printf("cant use a non float as a value")
		return 0
	}
	res := val.Float

	switch operator {
	case ADD:
		for _, c := range children[1:] {
			val := getValue(c)
			if !val.isFloat {
				log.Printf("cant use a non float as a value")
				return 0
			}
			res += val.Float
		}
	case SUB:
		for _, c := range children[1:] {
			val := getValue(c)
			if !val.isFloat {
				log.Printf("cant use a non float as a value")
				return 0
			}
			res -= val.Float
		}
	case DIV:
		for _, c := range children[1:] {
			val := getValue(c)
			if !val.isFloat {
				log.Printf("cant use a non float as a value")
				return 0
			}
			res /= val.Float
		}
	case MUL:
		for _, c := range children[1:] {
			val := getValue(c)
			if !val.isFloat {
				log.Printf("cant use a non float as a value")
				return 0
			}
			res *= val.Float
		}
	}

	return res
}

func Eval(nodes []Node) []string {
	res := make([]string, 0)
	for _, node := range nodes {
		switch node.(type) {
		case *Statement:
			stmt, ok := node.(*Statement)
			if !ok {
				log.Println("err: in eval, couldnt cast Node to Statement")
			}
			switch stmt.GetToken().Type {
			case PUTV:
				printer(*stmt)
			default:
				res = append(res, fmt.Sprint(sumOperator(node.GetToken().Type, stmt.Children)))
			}
		case *Float:
			res = append(res, fmt.Sprint(getValue(node).Float))
		case *String:
			res = append(res, getValue(node).String)
		}
	}
	return res
}
