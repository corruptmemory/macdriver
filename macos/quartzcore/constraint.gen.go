// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Constraint] class.
var ConstraintClass = _ConstraintClass{objc.GetClass("CAConstraint")}

type _ConstraintClass struct {
	objc.Class
}

// An interface definition for the [Constraint] class.
type IConstraint interface {
	objc.IObject
	Scale() float64
	Attribute() ConstraintAttribute
	SourceName() string
	Offset() float64
	SourceAttribute() ConstraintAttribute
}

// A representation of a single layout constraint between two layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint?language=objc
type Constraint struct {
	objc.Object
}

func ConstraintFrom(ptr unsafe.Pointer) Constraint {
	return Constraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ConstraintClass) ConstraintWithAttributeRelativeToAttribute(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute) Constraint {
	rv := objc.Call[Constraint](cc, objc.Sel("constraintWithAttribute:relativeTo:attribute:"), attr, srcId, srcAttr)
	return rv
}

// Creates and returns an CAConstraint object with the specified parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/1521924-constraintwithattribute?language=objc
func Constraint_ConstraintWithAttributeRelativeToAttribute(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute) Constraint {
	return ConstraintClass.ConstraintWithAttributeRelativeToAttribute(attr, srcId, srcAttr)
}

func (c_ Constraint) InitWithAttributeRelativeToAttributeScaleOffset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, m float64, c float64) Constraint {
	rv := objc.Call[Constraint](c_, objc.Sel("initWithAttribute:relativeTo:attribute:scale:offset:"), attr, srcId, srcAttr, m, c)
	return rv
}

// Returns an CAConstraint object with the specified parameters. Designated initializer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/1522213-initwithattribute?language=objc
func Constraint_InitWithAttributeRelativeToAttributeScaleOffset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, m float64, c float64) Constraint {
	return ConstraintClass.Alloc().InitWithAttributeRelativeToAttributeScaleOffset(attr, srcId, srcAttr, m, c)
}

func (cc _ConstraintClass) Alloc() Constraint {
	rv := objc.Call[Constraint](cc, objc.Sel("alloc"))
	return rv
}

func Constraint_Alloc() Constraint {
	return ConstraintClass.Alloc()
}

func (cc _ConstraintClass) New() Constraint {
	rv := objc.Call[Constraint](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewConstraint() Constraint {
	return ConstraintClass.New()
}

func (c_ Constraint) Init() Constraint {
	rv := objc.Call[Constraint](c_, objc.Sel("init"))
	return rv
}

// Scale factor of the constraint attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/1521911-scale?language=objc
func (c_ Constraint) Scale() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scale"))
	return rv
}

// The attribute the constraint affects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/1522186-attribute?language=objc
func (c_ Constraint) Attribute() ConstraintAttribute {
	rv := objc.Call[ConstraintAttribute](c_, objc.Sel("attribute"))
	return rv
}

// Name of the layer that the constraint is calculated relative to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/1522224-sourcename?language=objc
func (c_ Constraint) SourceName() string {
	rv := objc.Call[string](c_, objc.Sel("sourceName"))
	return rv
}

// Offset value of the constraint attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/1522142-offset?language=objc
func (c_ Constraint) Offset() float64 {
	rv := objc.Call[float64](c_, objc.Sel("offset"))
	return rv
}

// The constraint attribute of the layer the receiver is calculated relative to [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/1522385-sourceattribute?language=objc
func (c_ Constraint) SourceAttribute() ConstraintAttribute {
	rv := objc.Call[ConstraintAttribute](c_, objc.Sel("sourceAttribute"))
	return rv
}