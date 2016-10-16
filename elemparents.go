package xsd

func (a *All) initElement(parent element) {
	a.elemBase.init(parent, a, "all", &a.hasAttrId, &a.hasAttrMaxOccurs, &a.hasAttrMinOccurs)
	a.hasElemAnnotation.initChildren(a)
	a.hasElemsElement.initChildren(a)
}

func (a *Annotation) initElement(parent element) {
	a.elemBase.init(parent, a, "annotation")
	a.hasElemsAppInfo.initChildren(a)
	a.hasElemsDocumentation.initChildren(a)
}

func (a *Any) initElement(parent element) {
	a.elemBase.init(parent, a, "any", &a.hasAttrId, &a.hasAttrNamespace, &a.hasAttrMaxOccurs, &a.hasAttrMinOccurs, &a.hasAttrProcessContents)
	a.hasElemAnnotation.initChildren(a)
}

func (a *AnyAttribute) initElement(parent element) {
	a.elemBase.init(parent, a, "anyAttribute", &a.hasAttrId, &a.hasAttrNamespace, &a.hasAttrProcessContents)
	a.hasElemAnnotation.initChildren(a)
}

func (a *AppInfo) initElement(parent element) {
	a.elemBase.init(parent, a, "appInfo", &a.hasAttrSource)
}

func (a *Attribute) initElement(parent element) {
	a.elemBase.init(parent, a, "attribute", &a.hasAttrDefault, &a.hasAttrFixed, &a.hasAttrForm, &a.hasAttrId, &a.hasAttrName, &a.hasAttrRef, &a.hasAttrType, &a.hasAttrUse)
	a.hasElemAnnotation.initChildren(a)
	a.hasElemsSimpleType.initChildren(a)
}

func (a *AttributeGroup) initElement(parent element) {
	a.elemBase.init(parent, a, "attributeGroup", &a.hasAttrId, &a.hasAttrName, &a.hasAttrRef)
	a.hasElemAnnotation.initChildren(a)
	a.hasElemsAttribute.initChildren(a)
	a.hasElemsAnyAttribute.initChildren(a)
	a.hasElemsAttributeGroup.initChildren(a)
}

func (c *Choice) initElement(parent element) {
	c.elemBase.init(parent, c, "choice", &c.hasAttrId, &c.hasAttrMaxOccurs, &c.hasAttrMinOccurs)
	c.hasElemAnnotation.initChildren(c)
	c.hasElemsAny.initChildren(c)
	c.hasElemsChoice.initChildren(c)
	c.hasElemsElement.initChildren(c)
	c.hasElemsGroup.initChildren(c)
	c.hasElemsSequence.initChildren(c)
}

func (c *ComplexContent) initElement(parent element) {
	c.elemBase.init(parent, c, "complexContent", &c.hasAttrId, &c.hasAttrMixed)
	c.hasElemAnnotation.initChildren(c)
	c.hasElemExtensionComplexContent.initChildren(c)
	c.hasElemRestrictionComplexContent.initChildren(c)
}

func (c *ComplexType) initElement(parent element) {
	c.elemBase.init(parent, c, "complexType", &c.hasAttrAbstract, &c.hasAttrBlock, &c.hasAttrFinal, &c.hasAttrId, &c.hasAttrMixed, &c.hasAttrName)
	c.hasElemAnnotation.initChildren(c)
	c.hasElemAll.initChildren(c)
	c.hasElemChoice.initChildren(c)
	c.hasElemsAttribute.initChildren(c)
	c.hasElemGroup.initChildren(c)
	c.hasElemSequence.initChildren(c)
	c.hasElemComplexContent.initChildren(c)
	c.hasElemSimpleContent.initChildren(c)
	c.hasElemsAnyAttribute.initChildren(c)
	c.hasElemsAttributeGroup.initChildren(c)
}

func (d *Documentation) initElement(parent element) {
	d.elemBase.init(parent, d, "documentation", &d.hasAttrLang, &d.hasAttrSource)
}

func (e *Element) initElement(parent element) {
	e.elemBase.init(parent, e, "element", &e.hasAttrAbstract, &e.hasAttrBlock, &e.hasAttrDefault, &e.hasAttrFinal, &e.hasAttrFixed, &e.hasAttrForm, &e.hasAttrId, &e.hasAttrName, &e.hasAttrNillable, &e.hasAttrRef, &e.hasAttrType, &e.hasAttrMaxOccurs, &e.hasAttrMinOccurs, &e.hasAttrSubstitutionGroup)
	e.hasElemAnnotation.initChildren(e)
	e.hasElemUnique.initChildren(e)
	e.hasElemsKey.initChildren(e)
	e.hasElemComplexType.initChildren(e)
	e.hasElemKeyRef.initChildren(e)
	e.hasElemsSimpleType.initChildren(e)
}

func (e *ExtensionComplexContent) initElement(parent element) {
	e.elemBase.init(parent, e, "extension", &e.hasAttrBase, &e.hasAttrId)
	e.hasElemAnnotation.initChildren(e)
	e.hasElemAll.initChildren(e)
	e.hasElemsAttribute.initChildren(e)
	e.hasElemsChoice.initChildren(e)
	e.hasElemsGroup.initChildren(e)
	e.hasElemsSequence.initChildren(e)
	e.hasElemsAnyAttribute.initChildren(e)
	e.hasElemsAttributeGroup.initChildren(e)
}

func (e *ExtensionSimpleContent) initElement(parent element) {
	e.elemBase.init(parent, e, "extension", &e.hasAttrBase, &e.hasAttrId)
	e.hasElemAnnotation.initChildren(e)
	e.hasElemsAttribute.initChildren(e)
	e.hasElemsAnyAttribute.initChildren(e)
	e.hasElemsAttributeGroup.initChildren(e)
}

func (f *Field) initElement(parent element) {
	f.elemBase.init(parent, f, "field", &f.hasAttrId, &f.hasAttrXpath)
	f.hasElemAnnotation.initChildren(f)
}

func (g *Group) initElement(parent element) {
	g.elemBase.init(parent, g, "group", &g.hasAttrId, &g.hasAttrName, &g.hasAttrRef, &g.hasAttrMaxOccurs, &g.hasAttrMinOccurs)
	g.hasElemAnnotation.initChildren(g)
	g.hasElemAll.initChildren(g)
	g.hasElemChoice.initChildren(g)
	g.hasElemSequence.initChildren(g)
}

func (i *Import) initElement(parent element) {
	i.elemBase.init(parent, i, "import", &i.hasAttrId, &i.hasAttrNamespace, &i.hasAttrSchemaLocation)
	i.hasElemAnnotation.initChildren(i)
}

func (k *Key) initElement(parent element) {
	k.elemBase.init(parent, k, "key", &k.hasAttrId, &k.hasAttrName)
	k.hasElemAnnotation.initChildren(k)
	k.hasElemField.initChildren(k)
	k.hasElemSelector.initChildren(k)
}

func (k *KeyRef) initElement(parent element) {
	k.elemBase.init(parent, k, "keyref", &k.hasAttrId, &k.hasAttrName, &k.hasAttrRefer)
	k.hasElemAnnotation.initChildren(k)
	k.hasElemField.initChildren(k)
	k.hasElemSelector.initChildren(k)
}

func (l *List) initElement(parent element) {
	l.elemBase.init(parent, l, "list", &l.hasAttrId, &l.hasAttrItemType)
	l.hasElemAnnotation.initChildren(l)
	l.hasElemsSimpleType.initChildren(l)
}

func (n *Notation) initElement(parent element) {
	n.elemBase.init(parent, n, "notation", &n.hasAttrId, &n.hasAttrName, &n.hasAttrPublic, &n.hasAttrSystem)
	n.hasElemAnnotation.initChildren(n)
}

func (r *Redefine) initElement(parent element) {
	r.elemBase.init(parent, r, "redefine", &r.hasAttrId, &r.hasAttrSchemaLocation)
	r.hasElemAnnotation.initChildren(r)
	r.hasElemsGroup.initChildren(r)
	r.hasElemsAttributeGroup.initChildren(r)
	r.hasElemsComplexType.initChildren(r)
	r.hasElemsSimpleType.initChildren(r)
}

func (r *RestrictionComplexContent) initElement(parent element) {
	r.elemBase.init(parent, r, "restriction", &r.hasAttrBase, &r.hasAttrId)
	r.hasElemAnnotation.initChildren(r)
	r.hasElemAll.initChildren(r)
	r.hasElemsAttribute.initChildren(r)
	r.hasElemsChoice.initChildren(r)
	r.hasElemsSequence.initChildren(r)
	r.hasElemsAnyAttribute.initChildren(r)
	r.hasElemsAttributeGroup.initChildren(r)
}

func (r *RestrictionSimpleContent) initElement(parent element) {
	r.elemBase.init(parent, r, "restriction", &r.hasAttrBase, &r.hasAttrId)
	r.hasElemAnnotation.initChildren(r)
	r.hasElemLength.initChildren(r)
	r.hasElemPattern.initChildren(r)
	r.hasElemsAttribute.initChildren(r)
	r.hasElemsEnumeration.initChildren(r)
	r.hasElemFractionDigits.initChildren(r)
	r.hasElemMaxExclusive.initChildren(r)
	r.hasElemMaxInclusive.initChildren(r)
	r.hasElemMaxLength.initChildren(r)
	r.hasElemMinExclusive.initChildren(r)
	r.hasElemMinInclusive.initChildren(r)
	r.hasElemMinLength.initChildren(r)
	r.hasElemTotalDigits.initChildren(r)
	r.hasElemWhiteSpace.initChildren(r)
	r.hasElemsAnyAttribute.initChildren(r)
	r.hasElemsAttributeGroup.initChildren(r)
	r.hasElemsSimpleType.initChildren(r)
}

func (r *RestrictionSimpleEnumeration) initElement(parent element) {
	r.elemBase.init(parent, r, "enumeration", &r.hasAttrValue)
}

func (r *RestrictionSimpleFractionDigits) initElement(parent element) {
	r.elemBase.init(parent, r, "fractionDigits", &r.hasAttrValue)
}

func (r *RestrictionSimpleLength) initElement(parent element) {
	r.elemBase.init(parent, r, "length", &r.hasAttrValue)
}

func (r *RestrictionSimpleMaxExclusive) initElement(parent element) {
	r.elemBase.init(parent, r, "maxExclusive", &r.hasAttrValue)
}

func (r *RestrictionSimpleMaxInclusive) initElement(parent element) {
	r.elemBase.init(parent, r, "maxInclusive", &r.hasAttrValue)
}

func (r *RestrictionSimpleMaxLength) initElement(parent element) {
	r.elemBase.init(parent, r, "maxLength", &r.hasAttrValue)
}

func (r *RestrictionSimpleMinExclusive) initElement(parent element) {
	r.elemBase.init(parent, r, "minExclusive", &r.hasAttrValue)
}

func (r *RestrictionSimpleMinInclusive) initElement(parent element) {
	r.elemBase.init(parent, r, "minInclusive", &r.hasAttrValue)
}

func (r *RestrictionSimpleMinLength) initElement(parent element) {
	r.elemBase.init(parent, r, "minLength", &r.hasAttrValue)
}

func (r *RestrictionSimplePattern) initElement(parent element) {
	r.elemBase.init(parent, r, "pattern", &r.hasAttrValue)
}

func (r *RestrictionSimpleTotalDigits) initElement(parent element) {
	r.elemBase.init(parent, r, "totalDigits", &r.hasAttrValue)
}

func (r *RestrictionSimpleType) initElement(parent element) {
	r.elemBase.init(parent, r, "restriction", &r.hasAttrBase, &r.hasAttrId)
	r.hasElemAnnotation.initChildren(r)
	r.hasElemLength.initChildren(r)
	r.hasElemPattern.initChildren(r)
	r.hasElemsEnumeration.initChildren(r)
	r.hasElemFractionDigits.initChildren(r)
	r.hasElemMaxExclusive.initChildren(r)
	r.hasElemMaxInclusive.initChildren(r)
	r.hasElemMaxLength.initChildren(r)
	r.hasElemMinExclusive.initChildren(r)
	r.hasElemMinInclusive.initChildren(r)
	r.hasElemMinLength.initChildren(r)
	r.hasElemTotalDigits.initChildren(r)
	r.hasElemWhiteSpace.initChildren(r)
	r.hasElemsSimpleType.initChildren(r)
}

func (r *RestrictionSimpleWhiteSpace) initElement(parent element) {
	r.elemBase.init(parent, r, "whiteSpace", &r.hasAttrValue)
}

func (s *Schema) initElement(parent element) {
	s.elemBase.init(parent, s, "schema", &s.hasAttrId, &s.hasAttrLang, &s.hasAttrVersion, &s.hasAttrBlockDefault, &s.hasAttrFinalDefault, &s.hasAttrSchemaLocation, &s.hasAttrTargetNamespace, &s.hasAttrAttributeFormDefault, &s.hasAttrElementFormDefault)
	s.hasElemAnnotation.initChildren(s)
	s.hasElemsAttribute.initChildren(s)
	s.hasElemsElement.initChildren(s)
	s.hasElemsGroup.initChildren(s)
	s.hasElemsImport.initChildren(s)
	s.hasElemsNotation.initChildren(s)
	s.hasElemsRedefine.initChildren(s)
	s.hasElemsAttributeGroup.initChildren(s)
	s.hasElemsComplexType.initChildren(s)
	s.hasElemsSimpleType.initChildren(s)
}

func (s *Selector) initElement(parent element) {
	s.elemBase.init(parent, s, "selector", &s.hasAttrId, &s.hasAttrXpath)
	s.hasElemAnnotation.initChildren(s)
}

func (s *Sequence) initElement(parent element) {
	s.elemBase.init(parent, s, "sequence", &s.hasAttrId, &s.hasAttrMaxOccurs, &s.hasAttrMinOccurs)
	s.hasElemAnnotation.initChildren(s)
	s.hasElemsAny.initChildren(s)
	s.hasElemsChoice.initChildren(s)
	s.hasElemsElement.initChildren(s)
	s.hasElemsGroup.initChildren(s)
	s.hasElemsSequence.initChildren(s)
}

func (s *SimpleContent) initElement(parent element) {
	s.elemBase.init(parent, s, "simpleContent", &s.hasAttrId)
	s.hasElemAnnotation.initChildren(s)
	s.hasElemExtensionSimpleContent.initChildren(s)
	s.hasElemRestrictionSimpleContent.initChildren(s)
}

func (s *SimpleType) initElement(parent element) {
	s.elemBase.init(parent, s, "simpleType", &s.hasAttrFinal, &s.hasAttrId, &s.hasAttrName)
	s.hasElemAnnotation.initChildren(s)
	s.hasElemRestrictionSimpleType.initChildren(s)
	s.hasElemList.initChildren(s)
	s.hasElemUnion.initChildren(s)
}

func (u *Union) initElement(parent element) {
	u.elemBase.init(parent, u, "union", &u.hasAttrId, &u.hasAttrMemberTypes)
	u.hasElemAnnotation.initChildren(u)
	u.hasElemsSimpleType.initChildren(u)
}

func (u *Unique) initElement(parent element) {
	u.elemBase.init(parent, u, "unique", &u.hasAttrId, &u.hasAttrName)
	u.hasElemAnnotation.initChildren(u)
	u.hasElemField.initChildren(u)
	u.hasElemSelector.initChildren(u)
}
