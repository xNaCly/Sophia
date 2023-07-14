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
	}
	return val
}

func printer(stmt Statement) {
	strOut, _ := Eval(stmt.Children)
	fmt.Printf("~ %v\n", strOut)
}

func sumOperator(operator int, children []Node) float64 {
	var err bool
	var errNode Value
	if len(children) == 0 {
		return 0
	}
	val := getValue(children[0])
	if !val.isFloat {
		errNode = val
		err = true
	}
	res := val.Float

	if !err {
		switch operator {
		case ADD:
			for _, c := range children[1:] {
				val := getValue(c)
				if !val.isFloat {
					err = true
					errNode = val
					break
				}
				res += val.Float
			}
		case SUB:
			for _, c := range children[1:] {
				val := getValue(c)
				if !val.isFloat {
					err = true
					errNode = val
					break
				}
				res -= val.Float
			}
		case DIV:
			for _, c := range children[1:] {
				val := getValue(c)
				if !val.isFloat {
					err = true
					errNode = val
					break
				}
				res /= val.Float
			}
		case MUL:
			for _, c := range children[1:] {
				val := getValue(c)
				if !val.isFloat {
					err = true
					errNode = val
					break
				}
				res *= val.Float
			}
		}
	}

	if err {
		log.Printf("cant use a non float: '%s' as a value in '%s'-Operation", errNode.String, TOKEN_NAME_MAP[operator])
		return 0
	}

	return res
}

func Eval(nodes []Node) ([]string, []float64) {
	resF := make([]float64, 0)
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
				out := sumOperator(node.GetToken().Type, stmt.Children)
				res = append(res, fmt.Sprint(out))
				resF = append(resF, out)
			}
		case *Float:
			v := getValue(node).Float
			resF = append(resF, v)
			res = append(res, fmt.Sprint(v))
		case *String:
			res = append(res, getValue(node).String)
		}
	}

	return res, resF
}
