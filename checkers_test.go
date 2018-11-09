package checkers

import (
	"testing"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/linttest"
)

func TestCheckers(t *testing.T) { linttest.TestCheckers(t) }

func TestStableList(t *testing.T) {
	// Verify that new checker is not added without "experimental"
	// tag by accident. When stable checker is about to be added,
	// slice above should be modified to include new checker name.

	stableList := []string{
		"appendAssign",
		"appendCombine",
		"assignOp",
		"caseOrder",
		"dupArg",
		"dupBranchBody",
		"dupCase",
		"flagDeref",
		"ifElseChain",
		"rangeExprCopy",
		"rangeValCopy",
		"regexpMust",
		"singleCaseSwitch",
		"sloppyLen",
		"switchTrue",
		"typeSwitchVar",
		"underef",
		"unlambda",
		"unslice",
	}

	m := make(map[string]bool)
	for _, name := range stableList {
		m[name] = true
	}

	for _, info := range lintpack.GetCheckersInfo() {
		if info.HasTag("experimental") {
			continue
		}
		if !m[info.Name] {
			t.Errorf("%q checker misses `experimental` tag", info.Name)
		}
	}
}
