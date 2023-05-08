// Copyright 2023 The Infratographer Authors
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
//
// Code generated by entc, DO NOT EDIT.

package annotation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldUpdatedAt, v))
}

// MetadataID applies equality check predicate on the "metadata_id" field. It's identical to MetadataIDEQ.
func MetadataID(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldMetadataID, v))
}

// AnnotationNamespaceID applies equality check predicate on the "annotation_namespace_id" field. It's identical to AnnotationNamespaceIDEQ.
func AnnotationNamespaceID(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldAnnotationNamespaceID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldUpdatedAt, v))
}

// MetadataIDEQ applies the EQ predicate on the "metadata_id" field.
func MetadataIDEQ(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldMetadataID, v))
}

// MetadataIDNEQ applies the NEQ predicate on the "metadata_id" field.
func MetadataIDNEQ(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldMetadataID, v))
}

// MetadataIDIn applies the In predicate on the "metadata_id" field.
func MetadataIDIn(vs ...gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldMetadataID, vs...))
}

// MetadataIDNotIn applies the NotIn predicate on the "metadata_id" field.
func MetadataIDNotIn(vs ...gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldMetadataID, vs...))
}

// MetadataIDGT applies the GT predicate on the "metadata_id" field.
func MetadataIDGT(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldMetadataID, v))
}

// MetadataIDGTE applies the GTE predicate on the "metadata_id" field.
func MetadataIDGTE(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldMetadataID, v))
}

// MetadataIDLT applies the LT predicate on the "metadata_id" field.
func MetadataIDLT(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldMetadataID, v))
}

// MetadataIDLTE applies the LTE predicate on the "metadata_id" field.
func MetadataIDLTE(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldMetadataID, v))
}

// MetadataIDContains applies the Contains predicate on the "metadata_id" field.
func MetadataIDContains(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldContains(FieldMetadataID, vc))
}

// MetadataIDHasPrefix applies the HasPrefix predicate on the "metadata_id" field.
func MetadataIDHasPrefix(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldHasPrefix(FieldMetadataID, vc))
}

// MetadataIDHasSuffix applies the HasSuffix predicate on the "metadata_id" field.
func MetadataIDHasSuffix(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldHasSuffix(FieldMetadataID, vc))
}

// MetadataIDEqualFold applies the EqualFold predicate on the "metadata_id" field.
func MetadataIDEqualFold(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldEqualFold(FieldMetadataID, vc))
}

// MetadataIDContainsFold applies the ContainsFold predicate on the "metadata_id" field.
func MetadataIDContainsFold(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldContainsFold(FieldMetadataID, vc))
}

// AnnotationNamespaceIDEQ applies the EQ predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDEQ(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldAnnotationNamespaceID, v))
}

// AnnotationNamespaceIDNEQ applies the NEQ predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDNEQ(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldAnnotationNamespaceID, v))
}

// AnnotationNamespaceIDIn applies the In predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDIn(vs ...gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldAnnotationNamespaceID, vs...))
}

// AnnotationNamespaceIDNotIn applies the NotIn predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDNotIn(vs ...gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldAnnotationNamespaceID, vs...))
}

// AnnotationNamespaceIDGT applies the GT predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDGT(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldAnnotationNamespaceID, v))
}

// AnnotationNamespaceIDGTE applies the GTE predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDGTE(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldAnnotationNamespaceID, v))
}

// AnnotationNamespaceIDLT applies the LT predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDLT(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldAnnotationNamespaceID, v))
}

// AnnotationNamespaceIDLTE applies the LTE predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDLTE(v gidx.PrefixedID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldAnnotationNamespaceID, v))
}

// AnnotationNamespaceIDContains applies the Contains predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDContains(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldContains(FieldAnnotationNamespaceID, vc))
}

// AnnotationNamespaceIDHasPrefix applies the HasPrefix predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDHasPrefix(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldHasPrefix(FieldAnnotationNamespaceID, vc))
}

// AnnotationNamespaceIDHasSuffix applies the HasSuffix predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDHasSuffix(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldHasSuffix(FieldAnnotationNamespaceID, vc))
}

// AnnotationNamespaceIDEqualFold applies the EqualFold predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDEqualFold(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldEqualFold(FieldAnnotationNamespaceID, vc))
}

// AnnotationNamespaceIDContainsFold applies the ContainsFold predicate on the "annotation_namespace_id" field.
func AnnotationNamespaceIDContainsFold(v gidx.PrefixedID) predicate.Annotation {
	vc := string(v)
	return predicate.Annotation(sql.FieldContainsFold(FieldAnnotationNamespaceID, vc))
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.AnnotationNamespace) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := newNamespaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
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
func Not(p predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		p(s.Not())
	})
}