// Code generated by entc, DO NOT EDIT.

package auth

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ArifulProtik/sheba-api/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Sessionid applies equality check predicate on the "sessionid" field. It's identical to SessionidEQ.
func Sessionid(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionid), v))
	})
}

// IsBlocked applies equality check predicate on the "is_blocked" field. It's identical to IsBlockedEQ.
func IsBlocked(v bool) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsBlocked), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// SessionidEQ applies the EQ predicate on the "sessionid" field.
func SessionidEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionid), v))
	})
}

// SessionidNEQ applies the NEQ predicate on the "sessionid" field.
func SessionidNEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionid), v))
	})
}

// SessionidIn applies the In predicate on the "sessionid" field.
func SessionidIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSessionid), v...))
	})
}

// SessionidNotIn applies the NotIn predicate on the "sessionid" field.
func SessionidNotIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSessionid), v...))
	})
}

// SessionidGT applies the GT predicate on the "sessionid" field.
func SessionidGT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionid), v))
	})
}

// SessionidGTE applies the GTE predicate on the "sessionid" field.
func SessionidGTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionid), v))
	})
}

// SessionidLT applies the LT predicate on the "sessionid" field.
func SessionidLT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionid), v))
	})
}

// SessionidLTE applies the LTE predicate on the "sessionid" field.
func SessionidLTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionid), v))
	})
}

// IsBlockedEQ applies the EQ predicate on the "is_blocked" field.
func IsBlockedEQ(v bool) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsBlocked), v))
	})
}

// IsBlockedNEQ applies the NEQ predicate on the "is_blocked" field.
func IsBlockedNEQ(v bool) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsBlocked), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
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
func Not(p predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		p(s.Not())
	})
}
