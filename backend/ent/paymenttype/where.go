// Code generated by entc, DO NOT EDIT.

package paymenttype

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Typename applies equality check predicate on the "Typename" field. It's identical to TypenameEQ.
func Typename(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypename), v))
	})
}

// TypenameEQ applies the EQ predicate on the "Typename" field.
func TypenameEQ(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypename), v))
	})
}

// TypenameNEQ applies the NEQ predicate on the "Typename" field.
func TypenameNEQ(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypename), v))
	})
}

// TypenameIn applies the In predicate on the "Typename" field.
func TypenameIn(vs ...string) predicate.PaymentType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTypename), v...))
	})
}

// TypenameNotIn applies the NotIn predicate on the "Typename" field.
func TypenameNotIn(vs ...string) predicate.PaymentType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTypename), v...))
	})
}

// TypenameGT applies the GT predicate on the "Typename" field.
func TypenameGT(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypename), v))
	})
}

// TypenameGTE applies the GTE predicate on the "Typename" field.
func TypenameGTE(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypename), v))
	})
}

// TypenameLT applies the LT predicate on the "Typename" field.
func TypenameLT(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypename), v))
	})
}

// TypenameLTE applies the LTE predicate on the "Typename" field.
func TypenameLTE(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypename), v))
	})
}

// TypenameContains applies the Contains predicate on the "Typename" field.
func TypenameContains(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTypename), v))
	})
}

// TypenameHasPrefix applies the HasPrefix predicate on the "Typename" field.
func TypenameHasPrefix(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTypename), v))
	})
}

// TypenameHasSuffix applies the HasSuffix predicate on the "Typename" field.
func TypenameHasSuffix(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTypename), v))
	})
}

// TypenameEqualFold applies the EqualFold predicate on the "Typename" field.
func TypenameEqualFold(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTypename), v))
	})
}

// TypenameContainsFold applies the ContainsFold predicate on the "Typename" field.
func TypenameContainsFold(v string) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTypename), v))
	})
}

// HasReceipt applies the HasEdge predicate on the "receipt" edge.
func HasReceipt() predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiptTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiptTable, ReceiptColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiptWith applies the HasEdge predicate on the "receipt" edge with a given conditions (other predicates).
func HasReceiptWith(preds ...predicate.Receipt) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiptInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiptTable, ReceiptColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.PaymentType) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.PaymentType) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PaymentType) predicate.PaymentType {
	return predicate.PaymentType(func(s *sql.Selector) {
		p(s.Not())
	})
}