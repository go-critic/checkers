package checkers

import (
	"go/ast"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "captLocal"
	info.Tags = []string{"style"}
	info.Summary = "Detects capitalized names for local variables"
	info.Before = `func f(IN int, OUT *int) (ERR error) {}`
	info.After = `func f(in int, out *int) (err error) {}`

	lintpack.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		c := &captLocalChecker{ctx: ctx}
		c.checkLocals = c.ctx.Params.Bool("checkLocals", true)
		return astwalk.WalkerForLocalDef(c, ctx.TypesInfo)
	})
}

type captLocalChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext

	upcaseNames map[string]bool
	checkLocals bool
}

func (c *captLocalChecker) VisitLocalDef(def astwalk.Name, _ ast.Expr) {
	if !c.checkLocals && def.Kind != astwalk.NameParam {
		return
	}
	if ast.IsExported(def.ID.Name) {
		c.warn(def.ID)
	}
}

func (c *captLocalChecker) warn(id ast.Node) {
	c.ctx.Warn(id, "`%s' should not be capitalized", id)
}
