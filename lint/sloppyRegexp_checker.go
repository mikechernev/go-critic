package lint

import (
	"go/ast"
	"regexp"
)

func init() {
	addChecker(&sloppyRegexpChecker{}, attrExperimental)
}

type sloppyRegexpChecker struct {
	checkerBase

	sloppyParamRE *regexp.Regexp
}

func (c *sloppyRegexpChecker) Init() {
	c.sloppyParamRE = regexp.MustCompile("^\\^\\w")
}

func (c *sloppyRegexpChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects regexp that can be rewritten with a simpler alternative"
	d.Before = `var re = regexp.MustCompile("^TODO")
re.MatchString(s)`
	d.After = `strings.HasPrefix(s, "TODO")`
}

func (c *sloppyRegexpChecker) VisitExpr(x ast.Expr) {
	call, ok := x.(*ast.CallExpr)
	if !ok || !c.IsRegexpCreator(call.Fun) {
		return
	}

	param := "^TODO" //astutil.Unparen(call.Args[0])
	if c.sloppyParamRE.MatchString(param) {
		c.warn(call, "strings.HasPrefix")
	}
}

func (c *sloppyRegexpChecker) IsRegexpCreator(call ast.Expr) bool {
	switch qualifiedName(call) {
	case "regexp.Compile", "regexp.CompilePOSIX", "regexp.MustCompile", "regexp.MustCompilePOSIX":
		return true
	default:
		return false
	}
}

func (c *sloppyRegexpChecker) warn(cause *ast.CallExpr, suggestion string) {
	c.ctx.Warn(cause, "regexp can be rewritten with %s", suggestion)
}
