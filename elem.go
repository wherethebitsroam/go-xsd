package xsd

import (
	xsdt "github.com/wherethebitsroam/go-xsd/types"
)

type element interface {
	base() *elemBase
	init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake)
	Parent() element
}

type elemBase struct {
	atts         []beforeAfterMake
	parent, self element // self is the struct that embeds elemBase, rather than the elemBase pseudo-field
	xsdName      xsdt.NCName
	hasNameAttr  bool
}

func (me *elemBase) afterMakePkg(bag *PkgBag) {
	if !me.hasNameAttr {
		bag.Stacks.Name.Pop()
	}
	for _, a := range me.atts {
		a.afterMakePkg(bag)
	}
}

func (me *elemBase) beforeMakePkg(bag *PkgBag) {
	if !me.hasNameAttr {
		bag.Stacks.Name.Push(me.xsdName)
	}
	for _, a := range me.atts {
		a.beforeMakePkg(bag)
	}
}

func (me *elemBase) base() *elemBase { return me }

func (me *elemBase) init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake) {
	me.parent, me.self, me.xsdName, me.atts = parent, self, xsdName, atts
	for _, a := range atts {
		if _, me.hasNameAttr = a.(*hasAttrName); me.hasNameAttr {
			break
		}
	}
}

func (me *elemBase) longSafeName(bag *PkgBag) (ln string) {
	var els = []element{}
	for el := me.self; (el != nil) && (el != bag.Schema); el = el.Parent() {
		els = append(els, el)
	}
	for i := len(els) - 1; i >= 0; i-- {
		ln += bag.safeName(els[i].base().selfName().String())
	}
	return
}

func (me *elemBase) selfName() xsdt.NCName {
	if me.hasNameAttr {
		for _, at := range me.atts {
			if an, ok := at.(*hasAttrName); ok {
				return an.Name
			}
		}
	}
	return me.xsdName
}

func (me *elemBase) Parent() element { return me.parent }

type All struct {
	elemBase
	//	XMLName xml.Name `xml:"all"`
	hasAttrID
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemsElement
}

type Annotation struct {
	elemBase
	//	XMLName xml.Name `xml:"annotation"`
	hasElemsAppInfo
	hasElemsDocumentation
}

type Any struct {
	elemBase
	//	XMLName xml.Name `xml:"any"`
	hasAttrID
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrNamespace
	hasAttrProcessContents
	hasElemAnnotation
}

type AnyAttribute struct {
	elemBase
	//	XMLName xml.Name `xml:"anyAttribute"`
	hasAttrID
	hasAttrNamespace
	hasAttrProcessContents
	hasElemAnnotation
}

type AppInfo struct {
	elemBase
	//	XMLName xml.Name `xml:"appinfo"`
	hasAttrSource
	hasCdata
}

type Attribute struct {
	elemBase
	//	XMLName xml.Name `xml:"attribute"`
	hasAttrDefault
	hasAttrFixed
	hasAttrForm
	hasAttrID
	hasAttrName
	hasAttrRef
	hasAttrType
	hasAttrUse
	hasElemAnnotation
	hasElemsSimpleType
}

type AttributeGroup struct {
	elemBase
	//	XMLName xml.Name `xml:"attributeGroup"`
	hasAttrID
	hasAttrName
	hasAttrRef
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
}

type Choice struct {
	elemBase
	//	XMLName xml.Name `xml:"choice"`
	hasAttrID
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemsAny
	hasElemsChoice
	hasElemsElement
	hasElemsGroup
	hasElemsSequence
}

type ComplexContent struct {
	elemBase
	//	XMLName xml.Name `xml:"complexContent"`
	hasAttrID
	hasAttrMixed
	hasElemAnnotation
	hasElemExtensionComplexContent
	hasElemRestrictionComplexContent
}

type ComplexType struct {
	elemBase
	//	XMLName xml.Name `xml:"complexType"`
	hasAttrAbstract
	hasAttrBlock
	hasAttrFinal
	hasAttrID
	hasAttrMixed
	hasAttrName
	hasElemAll
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemChoice
	hasElemComplexContent
	hasElemGroup
	hasElemSequence
	hasElemSimpleContent
}

type Documentation struct {
	elemBase
	//	XMLName xml.Name `xml:"documentation"`
	hasAttrLang
	hasAttrSource
	hasCdata
}

type Element struct {
	elemBase
	//	XMLName xml.Name `xml:"element"`
	hasAttrAbstract
	hasAttrBlock
	hasAttrDefault
	hasAttrFinal
	hasAttrFixed
	hasAttrForm
	hasAttrID
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrName
	hasAttrNillable
	hasAttrRef
	hasAttrSubstitutionGroup
	hasAttrType
	hasElemAnnotation
	hasElemComplexType
	hasElemsKey
	hasElemKeyRef
	hasElemsSimpleType
	hasElemUnique
}

type ExtensionComplexContent struct {
	elemBase
	//	XMLName xml.Name `xml:"extension"`
	hasAttrBase
	hasAttrID
	hasElemAll
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsChoice
	hasElemsGroup
	hasElemsSequence
}

type ExtensionSimpleContent struct {
	elemBase
	//	XMLName xml.Name `xml:"extension"`
	hasAttrBase
	hasAttrID
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
}

type Field struct {
	elemBase
	//	XMLName xml.Name `xml:"field"`
	hasAttrID
	hasAttrXpath
	hasElemAnnotation
}

type Group struct {
	elemBase
	//	XMLName xml.Name `xml:"group"`
	hasAttrID
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrName
	hasAttrRef
	hasElemAll
	hasElemAnnotation
	hasElemChoice
	hasElemSequence
}

type Include struct {
	elemBase
	//	XMLName xml.Name `xml:"include"`
	hasAttrID
	hasAttrSchemaLocation
	hasElemAnnotation
}

type Import struct {
	elemBase
	//	XMLName xml.Name `xml:"import"`
	hasAttrID
	hasAttrNamespace
	hasAttrSchemaLocation
	hasElemAnnotation
}

type Key struct {
	elemBase
	//	XMLName xml.Name `xml:"key"`
	hasAttrID
	hasAttrName
	hasElemAnnotation
	hasElemField
	hasElemSelector
}

type KeyRef struct {
	elemBase
	//	XMLName xml.Name `xml:"keyref"`
	hasAttrID
	hasAttrName
	hasAttrRefer
	hasElemAnnotation
	hasElemField
	hasElemSelector
}

type List struct {
	elemBase
	//	XMLName xml.Name `xml:"list"`
	hasAttrID
	hasAttrItemType
	hasElemAnnotation
	hasElemsSimpleType
}

type Notation struct {
	elemBase
	//	XMLName xml.Name `xml:"notation"`
	hasAttrID
	hasAttrName
	hasAttrPublic
	hasAttrSystem
	hasElemAnnotation
}

type Redefine struct {
	elemBase
	//	XMLName xml.Name `xml:"redefine"`
	hasAttrID
	hasAttrSchemaLocation
	hasElemAnnotation
	hasElemsAttributeGroup
	hasElemsComplexType
	hasElemsGroup
	hasElemsSimpleType
}

type RestrictionComplexContent struct {
	elemBase
	//	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrID
	hasElemAll
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsChoice
	hasElemsSequence
}

type RestrictionSimpleContent struct {
	elemBase
	//	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrID
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsEnumeration
	hasElemFractionDigits
	hasElemLength
	hasElemMaxExclusive
	hasElemMaxInclusive
	hasElemMaxLength
	hasElemMinExclusive
	hasElemMinInclusive
	hasElemMinLength
	hasElemPattern
	hasElemsSimpleType
	hasElemTotalDigits
	hasElemWhiteSpace
}

type RestrictionSimpleEnumeration struct {
	elemBase
	//	XMLName xml.Name `xml:"enumeration"`
	hasAttrValue
}

type RestrictionSimpleFractionDigits struct {
	elemBase
	//	XMLName xml.Name `xml:"fractionDigits"`
	hasAttrValue
}

type RestrictionSimpleLength struct {
	elemBase
	//	XMLName xml.Name `xml:"length"`
	hasAttrValue
}

type RestrictionSimpleMaxExclusive struct {
	elemBase
	//	XMLName xml.Name `xml:"maxExclusive"`
	hasAttrValue
}

type RestrictionSimpleMaxInclusive struct {
	elemBase
	//	XMLName xml.Name `xml:"maxInclusive"`
	hasAttrValue
}

type RestrictionSimpleMaxLength struct {
	elemBase
	//	XMLName xml.Name `xml:"maxLength"`
	hasAttrValue
}

type RestrictionSimpleMinExclusive struct {
	elemBase
	//	XMLName xml.Name `xml:"minExclusive"`
	hasAttrValue
}

type RestrictionSimpleMinInclusive struct {
	elemBase
	//	XMLName xml.Name `xml:"minInclusive"`
	hasAttrValue
}

type RestrictionSimpleMinLength struct {
	elemBase
	//	XMLName xml.Name `xml:"minLength"`
	hasAttrValue
}

type RestrictionSimplePattern struct {
	elemBase
	//	XMLName xml.Name `xml:"pattern"`
	hasAttrValue
}

type RestrictionSimpleTotalDigits struct {
	elemBase
	//	XMLName xml.Name `xml:"totalDigits"`
	hasAttrValue
}

type RestrictionSimpleType struct {
	elemBase
	//	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrID
	hasElemAnnotation
	hasElemsEnumeration
	hasElemFractionDigits
	hasElemLength
	hasElemMaxExclusive
	hasElemMaxInclusive
	hasElemMaxLength
	hasElemMinExclusive
	hasElemMinInclusive
	hasElemMinLength
	hasElemPattern
	hasElemsSimpleType
	hasElemTotalDigits
	hasElemWhiteSpace
}

type RestrictionSimpleWhiteSpace struct {
	elemBase
	//	XMLName xml.Name `xml:"whiteSpace"`
	hasAttrValue
}

type Selector struct {
	elemBase
	//	XMLName xml.Name `xml:"selector"`
	hasAttrID
	hasAttrXpath
	hasElemAnnotation
}

type Sequence struct {
	elemBase
	//	XMLName xml.Name `xml:"sequence"`
	hasAttrID
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemsAny
	hasElemsChoice
	hasElemsElement
	hasElemsGroup
	hasElemsSequence
}

type SimpleContent struct {
	elemBase
	//	XMLName xml.Name `xml:"simpleContent"`
	hasAttrID
	hasElemAnnotation
	hasElemExtensionSimpleContent
	hasElemRestrictionSimpleContent
}

type SimpleType struct {
	elemBase
	//	XMLName xml.Name `xml:"simpleType"`
	hasAttrFinal
	hasAttrID
	hasAttrName
	hasElemAnnotation
	hasElemList
	hasElemRestrictionSimpleType
	hasElemUnion
}

type Union struct {
	elemBase
	//	XMLName xml.Name `xml:"union"`
	hasAttrID
	hasAttrMemberTypes
	hasElemAnnotation
	hasElemsSimpleType
}

type Unique struct {
	elemBase
	//	XMLName xml.Name `xml:"unique"`
	hasAttrID
	hasAttrName
	hasElemAnnotation
	hasElemField
	hasElemSelector
}

func Flattened(choices []*Choice, seqs []*Sequence) (allChoices []*Choice, allSeqs []*Sequence) {
	var tmpChoices []*Choice
	var tmpSeqs []*Sequence
	for _, ch := range choices {
		if ch != nil {
			allChoices = append(allChoices, ch)
			tmpChoices, tmpSeqs = Flattened(ch.Choices, ch.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
	}
	for _, seq := range seqs {
		if seq != nil {
			allSeqs = append(allSeqs, seq)
			tmpChoices, tmpSeqs = Flattened(seq.Choices, seq.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
	}
	return
}
