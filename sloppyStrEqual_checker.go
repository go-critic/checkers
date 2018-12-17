package checkers

import (
	"go/ast"
	"go/token"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/typep"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "sloppyStrEqual"
	info.Tags = []string{"performance"}
	info.Summary = "Detects unoptimal string equal."
	info.Before = `strings.ToLower(x) == y`
	info.After = `strings.EqualFold(x, y)`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForExpr(&sloppyStrEqualChecker{ctx: ctx})
	})
}

type sloppyStrEqualChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext
}

func (c *sloppyStrEqualChecker) VisitExpr(e ast.Expr) {
	expr, ok := e.(*ast.BinaryExpr)
	if !ok {
		return
	}

	if expr.Op != token.EQL && expr.Op != token.NEQ {
		return
	}

	call := astcast.ToCallExpr(expr.X)
	if qualifiedName(call.Fun) != "strings.ToLower" && qualifiedName(call.Fun) != "strings.ToUpper" {
		return
	}

	stringConv := astcast.ToCallExpr(call.Args[0])
	// if qualifiedName(stringConv.Fun) != "string" {
	// 	return
	// }

	x := stringConv.Args[0]
	y := call.Args[1]
	if typep.SideEffectFree(c.ctx.TypesInfo, x) && typep.SideEffectFree(c.ctx.TypesInfo, y) {
		c.warn(e, x, y)
	}
}

func (c *sloppyStrEqualChecker) warn(cause ast.Node, x, y ast.Expr) {
	c.ctx.Warn(cause, "consider replacing %s with strings.EqualFold(%s, %s)", cause, x, y)
}
