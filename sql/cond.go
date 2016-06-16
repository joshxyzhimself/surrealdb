// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

func (p *Parser) parseCond() (mul []Expr, err error) {

	var tok Token
	var lit string

	// Remove the WHERE keyword
	if _, _, exi := p.mightBe(WHERE); !exi {
		return nil, nil
	}

	for {

		one := &BinaryExpression{}

		tok, lit, err = p.shouldBe(IDENT, ID, TIME, TRUE, FALSE, STRING, NUMBER, DOUBLE)
		if err != nil {
			return nil, &ParseError{Found: lit, Expected: []string{"field name"}}
		}

		one.LHS, err = declare(tok, lit)
		if err != nil {
			return nil, err
		}

		tok, lit, err = p.shouldBe(IN, EQ, NEQ, GT, LT, GTE, LTE, EQR, NER, SEQ, SNE)
		if err != nil {
			return nil, err
		}
		one.Op = lit

		tok, lit, err = p.shouldBe(IDENT, ID, NULL, EMPTY, NOW, DATE, TIME, TRUE, FALSE, STRING, REGION, NUMBER, DOUBLE, REGEX, JSON, ARRAY)
		if err != nil {
			return nil, &ParseError{Found: lit, Expected: []string{"field value"}}
		}

		one.RHS, err = declare(tok, lit)
		if err != nil {
			return nil, err
		}

		mul = append(mul, one)

		// Remove the WHERE keyword
		if _, _, exi := p.mightBe(AND, OR); !exi {
			break
		}

	}

	return mul, nil

}