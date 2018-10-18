package checkers

import (
	"go/ast"
	"go/types"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "unslice"
	info.Summary = "Detects slice expressions that can be simplified to sliced expression itself"
	info.Before = `
f(s[:])               // s is string
copy(b[:], values...) // b is []byte`
	info.After = `
f(s)
copy(b, values...)`
	lintpack.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForExpr(&unsliceChecker{ctx: ctx})
	})
}

type unsliceChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext
}

func (c *unsliceChecker) VisitExpr(expr ast.Expr) {
	if expr, ok := expr.(*ast.SliceExpr); ok {
		// No need to worry about 3-index slicing,
		// because it's only permitted if expr.High is not nil.
		if expr.Low != nil || expr.High != nil {
			return
		}
		switch c.ctx.TypesInfo.TypeOf(expr.X).(type) {
		case *types.Slice, *types.Basic:
			// Basic kind catches strings, Slice cathes everything else.
			c.warn(expr)
		}
	}
}

func (c *unsliceChecker) warn(expr *ast.SliceExpr) {
	c.ctx.Warn(expr, "could simplify %s to %s", expr, expr.X)
}
