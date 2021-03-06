// Code generated by entc, DO NOT EDIT.

package personal

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Personalname applies equality check predicate on the "Personalname" field. It's identical to PersonalnameEQ.
func Personalname(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPersonalname), v))
	})
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Password applies equality check predicate on the "Password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PersonalnameEQ applies the EQ predicate on the "Personalname" field.
func PersonalnameEQ(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPersonalname), v))
	})
}

// PersonalnameNEQ applies the NEQ predicate on the "Personalname" field.
func PersonalnameNEQ(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPersonalname), v))
	})
}

// PersonalnameIn applies the In predicate on the "Personalname" field.
func PersonalnameIn(vs ...string) predicate.Personal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Personal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPersonalname), v...))
	})
}

// PersonalnameNotIn applies the NotIn predicate on the "Personalname" field.
func PersonalnameNotIn(vs ...string) predicate.Personal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Personal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPersonalname), v...))
	})
}

// PersonalnameGT applies the GT predicate on the "Personalname" field.
func PersonalnameGT(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPersonalname), v))
	})
}

// PersonalnameGTE applies the GTE predicate on the "Personalname" field.
func PersonalnameGTE(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPersonalname), v))
	})
}

// PersonalnameLT applies the LT predicate on the "Personalname" field.
func PersonalnameLT(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPersonalname), v))
	})
}

// PersonalnameLTE applies the LTE predicate on the "Personalname" field.
func PersonalnameLTE(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPersonalname), v))
	})
}

// PersonalnameContains applies the Contains predicate on the "Personalname" field.
func PersonalnameContains(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPersonalname), v))
	})
}

// PersonalnameHasPrefix applies the HasPrefix predicate on the "Personalname" field.
func PersonalnameHasPrefix(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPersonalname), v))
	})
}

// PersonalnameHasSuffix applies the HasSuffix predicate on the "Personalname" field.
func PersonalnameHasSuffix(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPersonalname), v))
	})
}

// PersonalnameEqualFold applies the EqualFold predicate on the "Personalname" field.
func PersonalnameEqualFold(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPersonalname), v))
	})
}

// PersonalnameContainsFold applies the ContainsFold predicate on the "Personalname" field.
func PersonalnameContainsFold(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPersonalname), v))
	})
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.Personal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Personal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.Personal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Personal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PasswordEQ applies the EQ predicate on the "Password" field.
func PasswordEQ(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "Password" field.
func PasswordNEQ(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "Password" field.
func PasswordIn(vs ...string) predicate.Personal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Personal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "Password" field.
func PasswordNotIn(vs ...string) predicate.Personal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Personal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "Password" field.
func PasswordGT(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "Password" field.
func PasswordGTE(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "Password" field.
func PasswordLT(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "Password" field.
func PasswordLTE(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "Password" field.
func PasswordContains(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "Password" field.
func PasswordHasPrefix(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "Password" field.
func PasswordHasSuffix(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "Password" field.
func PasswordEqualFold(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "Password" field.
func PasswordContainsFold(v string) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTitle applies the HasEdge predicate on the "title" edge.
func HasTitle() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTitleWith applies the HasEdge predicate on the "title" edge with a given conditions (other predicates).
func HasTitleWith(preds ...predicate.Title) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGender applies the HasEdge predicate on the "gender" edge.
func HasGender() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenderWith applies the HasEdge predicate on the "gender" edge with a given conditions (other predicates).
func HasGenderWith(preds ...predicate.Gender) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFix applies the HasEdge predicate on the "fix" edge.
func HasFix() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FixTable, FixColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixWith applies the HasEdge predicate on the "fix" edge with a given conditions (other predicates).
func HasFixWith(preds ...predicate.Fix) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FixTable, FixColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonal applies the HasEdge predicate on the "personal" edge.
func HasPersonal() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonalWith applies the HasEdge predicate on the "personal" edge with a given conditions (other predicates).
func HasPersonalWith(preds ...predicate.Adminrepair) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReceipt applies the HasEdge predicate on the "receipt" edge.
func HasReceipt() predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReceiptTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiptTable, ReceiptColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiptWith applies the HasEdge predicate on the "receipt" edge with a given conditions (other predicates).
func HasReceiptWith(preds ...predicate.Receipt) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
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
func And(predicates ...predicate.Personal) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Personal) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
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
func Not(p predicate.Personal) predicate.Personal {
	return predicate.Personal(func(s *sql.Selector) {
		p(s.Not())
	})
}
