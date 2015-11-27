// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type XmlExpr struct {
	Xpr       Expr          `json:"xpr"`
	Op        XmlExprOp     `json:"op"`         /* xml function ID */
	Name      *string       `json:"name"`       /* name in xml(NAME foo ...) syntaxes */
	NamedArgs []Node        `json:"named_args"` /* non-XML expressions for xml_attributes */
	ArgNames  []Node        `json:"arg_names"`  /* parallel list of Value strings */
	Args      []Node        `json:"args"`       /* list of expressions */
	Xmloption XmlOptionType `json:"xmloption"`  /* DOCUMENT or CONTENT */
	Type      Oid           `json:"type"`       /* target type/typmod for XMLSERIALIZE */
	Typmod    int32         `json:"typmod"`
	Location  int           `json:"location"` /* token location, or -1 if unknown */
}

func (node XmlExpr) MarshalJSON() ([]byte, error) {
	type XmlExprMarshalAlias XmlExpr
	return json.Marshal(map[string]interface{}{
		"XMLEXPR": (*XmlExprMarshalAlias)(&node),
	})
}

func (node *XmlExpr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["op"] != nil {
		err = json.Unmarshal(fields["op"], &node.Op)
		if err != nil {
			return
		}
	}

	if fields["name"] != nil {
		err = json.Unmarshal(fields["name"], &node.Name)
		if err != nil {
			return
		}
	}

	if fields["named_args"] != nil {
		node.NamedArgs, err = UnmarshalNodeArrayJSON(fields["named_args"])
		if err != nil {
			return
		}
	}

	if fields["arg_names"] != nil {
		node.ArgNames, err = UnmarshalNodeArrayJSON(fields["arg_names"])
		if err != nil {
			return
		}
	}

	if fields["args"] != nil {
		node.Args, err = UnmarshalNodeArrayJSON(fields["args"])
		if err != nil {
			return
		}
	}

	if fields["xmloption"] != nil {
		err = json.Unmarshal(fields["xmloption"], &node.Xmloption)
		if err != nil {
			return
		}
	}

	if fields["type"] != nil {
		err = json.Unmarshal(fields["type"], &node.Type)
		if err != nil {
			return
		}
	}

	if fields["typmod"] != nil {
		err = json.Unmarshal(fields["typmod"], &node.Typmod)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}