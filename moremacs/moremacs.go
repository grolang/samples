// Copyright 2016 The Groo Authors. All rights reserved.
// Use of this source code is governed by the same BSD-style
// license as Go that can be found in the LICENSE file.

package moremacs

import (
	"go/ast"
	"github.com/grolang/gro/macro"
)

type Struct struct{}

func (m Struct) Init(p macro.Parser) {
}

func (m Struct) Main(p macro.Parser) ast.Stmt {
	s:= &ast.EmptyStmt{Semicolon: p.Pos(), Implicit: true}
	p.Next()
	return s
}

