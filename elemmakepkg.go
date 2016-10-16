package xsd

import (
	"fmt"
	"path"
	"strings"
	"unicode"

	"github.com/go-utils/ustr"

	xsdt "github.com/wherethebitsroam/go-xsd/types"
)

const (
	idPrefix = "XsdGoPkg"
)

func (a *All) makePkg(bag *PkgBag) {
	a.elemBase.beforeMakePkg(bag)
	a.hasElemsElement.makePkg(bag)
	a.elemBase.afterMakePkg(bag)
}

func (a *Annotation) makePkg(bag *PkgBag) {
	a.elemBase.beforeMakePkg(bag)
	a.hasElemsAppInfo.makePkg(bag)
	a.hasElemsDocumentation.makePkg(bag)
	a.elemBase.afterMakePkg(bag)
}

func (a *Any) makePkg(bag *PkgBag) {
	a.elemBase.beforeMakePkg(bag)
	a.elemBase.afterMakePkg(bag)
}

func (a *AnyAttribute) makePkg(bag *PkgBag) {
	a.elemBase.beforeMakePkg(bag)
	a.elemBase.afterMakePkg(bag)
}

func (a *AppInfo) makePkg(bag *PkgBag) {
	a.elemBase.beforeMakePkg(bag)
	a.elemBase.afterMakePkg(bag)
}

func (a *Attribute) makePkg(bag *PkgBag) {
	var safeName, typeName, tmp, key, defVal, impName string
	var defName = "Default"
	a.elemBase.beforeMakePkg(bag)
	if len(a.Form) == 0 {
		a.Form = bag.Schema.AttributeFormDefault
	}
	a.hasElemsSimpleType.makePkg(bag)
	if len(a.Ref) > 0 {
		key = bag.resolveQnameRef(a.Ref.String(), "", &impName)
		tmp = ustr.PrefixWithSep(impName, ".", idPrefix+"HasAttr_"+bag.safeName(a.Ref.String()[(strings.Index(a.Ref.String(), ":")+1):]))
		if bag.attRefImps[a], bag.attsKeys[a] = impName, key; len(bag.attsCache[key]) == 0 {
			bag.attsCache[key] = tmp
		}
	} else {
		safeName = bag.safeName(a.Name.String())
		if typeName = a.Type.String(); (len(typeName) == 0) && (len(a.SimpleTypes) > 0) {
			typeName = a.SimpleTypes[0].Name.String()
		} else {
			if len(typeName) == 0 {
				typeName = bag.xsdStringTypeRef()
			}
			typeName = bag.resolveQnameRef(typeName, "T", &impName)
		}
		if defVal = a.Default; len(defVal) == 0 {
			defName, defVal = "Fixed", a.Fixed
		}
		if a.Parent() == bag.Schema {
			key = safeName
		} else {
			key = safeName + "_" + bag.safeName(typeName) + "_" + bag.safeName(defVal)
		}
		if len(bag.attsCache[key]) == 0 {
			tmp = idPrefix + "HasAttr_" + key
			bag.attsKeys[a] = key
			bag.attsCache[key] = tmp
			var td = bag.addType(a, tmp, "", a.Annotation)
			td.addField(a, safeName, typeName, ustr.Ifs(len(bag.Schema.TargetNamespace) > 0, bag.Schema.TargetNamespace.String()+" ", "")+a.Name.String()+",attr", a.Annotation)
			if isPt := bag.isParseType(typeName); len(defVal) > 0 {
				doc := sfmt("Returns the %v value for %v -- "+ustr.Ifs(isPt, "%v", "%#v"), strings.ToLower(defName), safeName, defVal)
				if isPt {
					if PkgGen.ForceParseForDefaults {
						td.addMethod(nil, tmp, safeName+defName, typeName, sfmt("var x = new(%v); x.Set(%#v); return *x", typeName, defVal), doc)
					} else {
						td.addMethod(nil, tmp, safeName+defName, typeName, sfmt("return %v(%v)", typeName, defVal), doc)
					}
				} else {
					td.addMethod(nil, tmp, safeName+defName, typeName, sfmt("return %v(%#v)", typeName, defVal), doc)
				}
			}
		} else {
			bag.attsKeys[a] = key
		}
	}
	a.elemBase.afterMakePkg(bag)
}

func (a *AttributeGroup) makePkg(bag *PkgBag) {
	var refName, refImp string
	a.elemBase.beforeMakePkg(bag)
	a.hasElemsAttribute.makePkg(bag)
	a.hasElemsAnyAttribute.makePkg(bag)
	a.hasElemsAttributeGroup.makePkg(bag)
	if len(a.Ref) > 0 {
		if len(bag.attGroups[a]) == 0 {
			refName = bag.resolveQnameRef(a.Ref.String(), "", &refImp)
			bag.attGroups[a] = idPrefix + "HasAtts_" + refName
			bag.attGroupRefImps[a] = refImp
		}
	} else {
		safeName := bag.safeName(a.Name.String())
		tmp := idPrefix + "HasAtts_" + safeName
		var td = bag.addType(a, tmp, "", a.Annotation)
		bag.attGroups[a] = tmp
		for _, ag := range a.AttributeGroups {
			if len(ag.Ref) == 0 {
				ag.Ref.Set(ag.Name.String())
			}
			if refName = bag.resolveQnameRef(ag.Ref.String(), "", &refImp); len(refImp) > 0 {
				td.addEmbed(ag, refImp+"."+idPrefix+"HasAtts_"+refName[(len(refImp)+1):], ag.Annotation)
			} else {
				td.addEmbed(ag, idPrefix+"HasAtts_"+refName, ag.Annotation)
			}
		}
		for _, att := range a.Attributes {
			if key := bag.attsKeys[att]; len(key) > 0 {
				td.addEmbed(att, bag.attsCache[key], att.Annotation)
			}
		}
	}
	a.elemBase.afterMakePkg(bag)
}

func (c *Choice) makePkg(bag *PkgBag) {
	c.elemBase.beforeMakePkg(bag)
	c.hasElemsAny.makePkg(bag)
	c.hasElemsChoice.makePkg(bag)
	c.hasElemsGroup.makePkg(bag)
	c.hasElemsSequence.makePkg(bag)
	c.hasElemsElement.makePkg(bag)
	c.elemBase.afterMakePkg(bag)
}

func (c *ComplexContent) makePkg(bag *PkgBag) {
	c.elemBase.beforeMakePkg(bag)
	c.hasElemExtensionComplexContent.makePkg(bag)
	c.hasElemRestrictionComplexContent.makePkg(bag)
	c.elemBase.afterMakePkg(bag)
}

func (c *ComplexType) makePkg(bag *PkgBag) {
	var att *Attribute
	var attGroup *AttributeGroup
	var ctBaseType, ctValueType, typeSafeName string
	var allAtts = map[*Attribute]bool{}
	var allAttGroups = map[*AttributeGroup]bool{}
	var allElems = map[*Element]bool{}
	var allElemGroups = map[*Group]bool{}
	var elsDone, grsDone = map[string]bool{}, map[string]bool{}
	var allChoices, tmpChoices = []*Choice{}, []*Choice{c.Choice}
	var allSeqs, tmpSeqs = []*Sequence{}, []*Sequence{c.Sequence}
	var el *Element
	var elGr *Group
	var mixed = false
	c.elemBase.beforeMakePkg(bag)
	c.hasElemsAttribute.makePkg(bag)
	c.hasElemsAnyAttribute.makePkg(bag)
	c.hasElemsAttributeGroup.makePkg(bag)
	c.hasElemAll.makePkg(bag)
	c.hasElemChoice.makePkg(bag)
	c.hasElemGroup.makePkg(bag)
	c.hasElemSequence.makePkg(bag)
	c.hasElemComplexContent.makePkg(bag)
	c.hasElemSimpleContent.makePkg(bag)
	if len(c.Name) == 0 {
		c.Name = bag.AnonName(c.longSafeName(bag))
	}
	typeSafeName = bag.safeName(ustr.PrependIf(c.Name.String(), "T"))
	var td = bag.addType(c, typeSafeName, "", c.Annotation)
	for _, att = range c.Attributes {
		allAtts[att] = true
	}
	for _, attGroup = range c.AttributeGroups {
		allAttGroups[attGroup] = true
	}
	allChoices, allSeqs = Flattened(tmpChoices, tmpSeqs)
	if c.All != nil {
		for _, el = range c.All.Elements {
			allElems[el] = true
		}
	}
	if c.Group != nil {
		allElemGroups[c.Group] = true
	}
	if mixed = c.Mixed; c.ComplexContent != nil {
		mixed = mixed || c.ComplexContent.Mixed
		td.addAnnotations(c.ComplexContent.Annotation)
		if c.ComplexContent.ExtensionComplexContent != nil {
			td.addAnnotations(c.ComplexContent.ExtensionComplexContent.Annotation)
			if c.ComplexContent.ExtensionComplexContent.All != nil {
				for _, el = range c.ComplexContent.ExtensionComplexContent.All.Elements {
					allElems[el] = true
				}
			}
			for _, elGr = range c.ComplexContent.ExtensionComplexContent.Groups {
				allElemGroups[elGr] = true
			}
			tmpChoices, tmpSeqs = Flattened(c.ComplexContent.ExtensionComplexContent.Choices, c.ComplexContent.ExtensionComplexContent.Sequences)
			allChoices, allSeqs = append(allChoices, tmpChoices...), append(allSeqs, tmpSeqs...)
			for _, att = range c.ComplexContent.ExtensionComplexContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range c.ComplexContent.ExtensionComplexContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if len(c.ComplexContent.ExtensionComplexContent.Base) > 0 {
				ctBaseType = c.ComplexContent.ExtensionComplexContent.Base.String()
			}
		}
		if c.ComplexContent.RestrictionComplexContent != nil {
			td.addAnnotations(c.ComplexContent.RestrictionComplexContent.Annotation)
			if c.ComplexContent.RestrictionComplexContent.All != nil {
				for _, el = range c.ComplexContent.RestrictionComplexContent.All.Elements {
					allElems[el] = true
				}
			}
			tmpChoices, tmpSeqs = Flattened(c.ComplexContent.RestrictionComplexContent.Choices, c.ComplexContent.RestrictionComplexContent.Sequences)
			allChoices, allSeqs = append(allChoices, tmpChoices...), append(allSeqs, tmpSeqs...)
			for _, att = range c.ComplexContent.RestrictionComplexContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range c.ComplexContent.RestrictionComplexContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if len(c.ComplexContent.RestrictionComplexContent.Base) > 0 {
				ctBaseType = c.ComplexContent.RestrictionComplexContent.Base.String()
			}
		}
	}
	if c.SimpleContent != nil {
		td.addAnnotations(c.SimpleContent.Annotation)
		if c.SimpleContent.ExtensionSimpleContent != nil {
			if len(c.SimpleContent.ExtensionSimpleContent.Base) > 0 {
				ctBaseType = c.SimpleContent.ExtensionSimpleContent.Base.String()
			}
			td.addAnnotations(c.SimpleContent.ExtensionSimpleContent.Annotation)
			for _, att = range c.SimpleContent.ExtensionSimpleContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range c.SimpleContent.ExtensionSimpleContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if (len(ctValueType) == 0) && (len(c.SimpleContent.ExtensionSimpleContent.Base) > 0) {
				ctValueType = c.SimpleContent.ExtensionSimpleContent.Base.String()
			}
		}
		if c.SimpleContent.RestrictionSimpleContent != nil {
			if len(c.SimpleContent.RestrictionSimpleContent.Base) > 0 {
				ctBaseType = c.SimpleContent.RestrictionSimpleContent.Base.String()
			}
			td.addAnnotations(c.SimpleContent.RestrictionSimpleContent.Annotation)
			for _, att = range c.SimpleContent.RestrictionSimpleContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range c.SimpleContent.RestrictionSimpleContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if (len(ctValueType) == 0) && (len(c.SimpleContent.RestrictionSimpleContent.Base) > 0) {
				ctValueType = c.SimpleContent.RestrictionSimpleContent.Base.String()
			}
			if (len(ctValueType) == 0) && (len(c.SimpleContent.RestrictionSimpleContent.SimpleTypes) > 0) {
				ctValueType = c.SimpleContent.RestrictionSimpleContent.SimpleTypes[0].Name.String()
			}
			for _, enum := range c.SimpleContent.RestrictionSimpleContent.Enumerations {
				println("ENUMTODO!?! Whoever sees this message, please post an issue at github.com/metaleap/go-xsd with a link to the XSD..." + enum.selfName().String())
			}
		}
	}
	if ctBaseType = bag.resolveQnameRef(ctBaseType, "T", nil); len(ctBaseType) > 0 {
		td.addEmbed(nil, bag.safeName(ctBaseType))
	} else if ctValueType = bag.resolveQnameRef(ctValueType, "T", nil); len(ctValueType) > 0 {
		bag.simpleContentValueTypes[typeSafeName] = ctValueType
		td.addField(nil, idPrefix+"Value", ctValueType, ",chardata")
		chain := sfmt("me.%vValue", idPrefix)
		td.addMethod(nil, "*"+typeSafeName, sfmt("To%v", bag.safeName(ctValueType)), ctValueType, sfmt("return %v", chain), sfmt("Simply returns the value of its %vValue field.", idPrefix))
		ttn := ctValueType
		for ttd := bag.declTypes[ctValueType]; ttd != nil; ttd = bag.declTypes[ttn] {
			if ttd != nil {
				bag.declConvs[ttd.Name] = true
			}
			if ttn = ttd.Type; len(ttn) > 0 {
				chain += sfmt(".To%v()", bag.safeName(ttn))
				td.addMethod(nil, "*"+typeSafeName, sfmt("To%v", bag.safeName(ttn)), ttn, sfmt("return %v", chain), sfmt("Returns the value of its %vValue field as a %v (which %v is just aliasing).", idPrefix, ttn, ctValueType))
			} else {
				break
			}
		}
		if (!strings.HasPrefix(ctValueType, "xsdt.")) && (bag.declTypes[ctValueType] == nil) {
			println("NOTFOUND: " + ctValueType)
		}
	} else if mixed {
		td.addEmbed(nil, idPrefix+"HasCdata")
	}
	for elGr = range allElemGroups {
		subMakeElemGroup(bag, td, elGr, grsDone, anns(nil, c.ComplexContent)...)
	}
	for el = range allElems {
		subMakeElem(bag, td, el, elsDone, 1, anns(c.All, nil)...)
	}
	for _, ch := range allChoices {
		for _, el = range ch.Elements {
			subMakeElem(bag, td, el, elsDone, ch.hasAttrMaxOccurs.Value(), ch.Annotation)
		}
		for _, elGr = range ch.Groups {
			subMakeElemGroup(bag, td, elGr, grsDone, ch.Annotation)
		}
	}
	for _, seq := range allSeqs {
		for _, el = range seq.Elements {
			subMakeElem(bag, td, el, elsDone, seq.hasAttrMaxOccurs.Value(), seq.Annotation)
		}
		for _, elGr = range seq.Groups {
			subMakeElemGroup(bag, td, elGr, grsDone, seq.Annotation)
		}
	}
	for attGroup = range allAttGroups {
		td.addEmbed(attGroup, ustr.PrefixWithSep(bag.attGroupRefImps[attGroup], ".", bag.attGroups[attGroup][(strings.Index(bag.attGroups[attGroup], ".")+1):]), attGroup.Annotation)
	}

	for att = range allAtts {
		if key := bag.attsKeys[att]; len(key) > 0 {
			td.addEmbed(att, ustr.PrefixWithSep(bag.attRefImps[att], ".", bag.attsCache[key][(strings.Index(bag.attsCache[key], ".")+1):]), att.Annotation)
		}
	}
	c.elemBase.afterMakePkg(bag)
}

func (d *Documentation) makePkg(bag *PkgBag) {
	d.elemBase.beforeMakePkg(bag)
	if len(d.CDATA) > 0 {
		var s, ln string
		for _, ln = range ustr.Split(d.CDATA, "\n") {
			if s = strings.Trim(ln, " \t\r\n"); len(s) > 0 {
				bag.appendFmt(false, "//\t%s", s)
			}
		}
	}
	d.elemBase.afterMakePkg(bag)
}

func (e *Element) makePkg(bag *PkgBag) {
	var (
		safeName, typeName, valueType, tmp, key, defVal, impName string
		subEl                                                    *Element
	)
	asterisk, defName, doc := "", "Default", ""
	e.elemBase.beforeMakePkg(bag)
	if len(e.Form) == 0 {
		e.Form = bag.Schema.ElementFormDefault
	}
	e.hasElemsSimpleType.makePkg(bag)
	e.hasElemComplexType.makePkg(bag)
	if len(e.Ref) > 0 {
		key = bag.resolveQnameRef(e.Ref.String(), "", &impName)
		for pref, cache := range map[string]map[string]string{"HasElem_": bag.elemsCacheOnce, "HasElems_": bag.elemsCacheMult} {
			tmp = ustr.PrefixWithSep(impName, ".", idPrefix+pref+bag.safeName(e.Ref.String()[(strings.Index(e.Ref.String(), ":")+1):]))
			if bag.elemRefImps[e], bag.elemKeys[e] = impName, key; len(cache[key]) == 0 {
				cache[key] = tmp
			}
		}
	} else {
		safeName = bag.safeName(e.Name.String())
		if typeName = e.Type.String(); (len(typeName) == 0) && ((e.ComplexType != nil) || (len(e.SimpleTypes) > 0)) {
			if e.ComplexType != nil {
				asterisk, typeName = "*", e.ComplexType.Name.String()
			} else {
				typeName = e.SimpleTypes[0].Name.String()
			}
		} else {
			if len(typeName) == 0 {
				typeName = bag.xsdStringTypeRef()
			}
			loadedSchemas := make(map[string]bool)
			if typeName = bag.resolveQnameRef(typeName, "T", &impName); bag.Schema.RootSchema([]string{bag.Schema.loadURI}).globalComplexType(bag, typeName, loadedSchemas) != nil {
				asterisk = "*"
			}
		}
		if defVal = e.Default; len(defVal) == 0 {
			defName, defVal = "Fixed", e.Fixed
		}
		if e.Parent() == bag.Schema {
			key = safeName
		} else {
			key = bag.safeName(bag.Stacks.FullName()) + "_" + safeName + "_" + bag.safeName(typeName) + "_" + bag.safeName(defVal)
		}
		if valueType = bag.simpleContentValueTypes[typeName]; len(valueType) == 0 {
			valueType = typeName
		}
		isPt := bag.isParseType(valueType)
		if _, isChoice := e.Parent().(*Choice); isChoice && isPt {
			asterisk = "*"
		}
		for pref, cache := range map[string]map[string]string{"HasElem_": bag.elemsCacheOnce, "HasElems_": bag.elemsCacheMult} {
			if tmp = idPrefix + pref + key; !bag.elemsWritten[tmp] {
				bag.elemsWritten[tmp], bag.elemKeys[e] = true, key
				cache[key] = tmp
				var td = bag.addType(e, tmp, "", e.Annotation)
				td.addField(e, ustr.Ifs(pref == "HasElems_", pluralize(safeName), safeName), ustr.Ifs(pref == "HasElems_", "[]"+asterisk+typeName, asterisk+typeName), ustr.Ifs(len(bag.Schema.TargetNamespace) > 0, bag.Schema.TargetNamespace.String()+" ", "")+e.Name.String(), e.Annotation)
				if e.parent == bag.Schema {
					loadedSchemas := make(map[string]bool)
					for _, subEl = range bag.Schema.RootSchema([]string{bag.Schema.loadURI}).globalSubstitutionElems(e, loadedSchemas) {
						td.addEmbed(subEl, idPrefix+pref+bag.safeName(subEl.Name.String()), subEl.Annotation)
					}
				}
				if len(defVal) > 0 {
					doc = sfmt("Returns the %v value for %v -- "+ustr.Ifs(isPt, "%v", "%#v"), strings.ToLower(defName), safeName, defVal)
					if isPt {
						if PkgGen.ForceParseForDefaults {
							td.addMethod(nil, tmp, safeName+defName, valueType, sfmt("var x = new(%v); x.Set(%#v); return *x", valueType, defVal), doc)
						} else {
							td.addMethod(nil, tmp, safeName+defName, valueType, sfmt("return %v(%v)", valueType, defVal), doc)
						}
					} else {
						td.addMethod(nil, tmp, safeName+defName, valueType, sfmt("return %v(%#v)", valueType, defVal), doc)
					}
				}
			}
		}
	}
	e.elemBase.afterMakePkg(bag)
}

func (e *ExtensionComplexContent) makePkg(bag *PkgBag) {
	e.elemBase.beforeMakePkg(bag)
	e.hasElemsAttribute.makePkg(bag)
	e.hasElemsAnyAttribute.makePkg(bag)
	e.hasElemsAttributeGroup.makePkg(bag)
	e.hasElemAll.makePkg(bag)
	e.hasElemsChoice.makePkg(bag)
	e.hasElemsGroup.makePkg(bag)
	e.hasElemsSequence.makePkg(bag)
	e.elemBase.afterMakePkg(bag)
}

func (e *ExtensionSimpleContent) makePkg(bag *PkgBag) {
	e.elemBase.beforeMakePkg(bag)
	e.hasElemsAttribute.makePkg(bag)
	e.hasElemsAnyAttribute.makePkg(bag)
	e.hasElemsAttributeGroup.makePkg(bag)
	e.elemBase.afterMakePkg(bag)
}

func (f *Field) makePkg(bag *PkgBag) {
	f.elemBase.beforeMakePkg(bag)
	f.elemBase.afterMakePkg(bag)
}

func (g *Group) makePkg(bag *PkgBag) {
	var refName, refImp string
	var choices = []*Choice{g.Choice}
	var seqs = []*Sequence{g.Sequence}
	var el *Element
	var gr *Group
	var elsDone, grsDone = map[string]bool{}, map[string]bool{}
	g.elemBase.beforeMakePkg(bag)
	g.hasElemAll.makePkg(bag)
	g.hasElemChoice.makePkg(bag)
	g.hasElemSequence.makePkg(bag)
	if len(g.Ref) > 0 {
		if len(bag.elemGroups[g]) == 0 {
			refName = bag.resolveQnameRef(g.Ref.String(), "", &refImp)
			bag.elemGroups[g] = idPrefix + "HasGroup_" + refName
			bag.elemGroupRefImps[g] = refImp
		}
	} else {
		g.Ref.Set(g.Name.String())
		safeName := bag.safeName(g.Name.String())
		tmp := idPrefix + "HasGroup_" + safeName
		bag.elemGroups[g] = tmp
		var td = bag.addType(g, tmp, "", g.Annotation)
		choices, seqs = Flattened(choices, seqs)
		if g.All != nil {
			for _, el = range g.All.Elements {
				subMakeElem(bag, td, el, elsDone, 1, g.All.Annotation)
			}
		}
		for _, ch := range choices {
			for _, el = range ch.Elements {
				subMakeElem(bag, td, el, elsDone, ch.hasAttrMaxOccurs.Value(), ch.Annotation)
			}
			for _, gr = range ch.Groups {
				subMakeElemGroup(bag, td, gr, grsDone, ch.Annotation)
			}
		}
		for _, seq := range seqs {
			for _, el = range seq.Elements {
				subMakeElem(bag, td, el, elsDone, seq.hasAttrMaxOccurs.Value(), seq.Annotation)
			}
			for _, gr = range seq.Groups {
				subMakeElemGroup(bag, td, gr, grsDone, seq.Annotation)
			}
		}
	}

	g.elemBase.afterMakePkg(bag)
}

func (i *Import) makePkg(bag *PkgBag) {
	i.elemBase.beforeMakePkg(bag)
	var impName, impPath string
	var pos int
	i.hasElemAnnotation.makePkg(bag)
	for k, v := range bag.Schema.XMLNamespaces {
		if v == i.Namespace {
			impName = safeIdentifier(k)
			break
		}
	}
	if len(impName) > 0 {
		if pos, impPath = strings.Index(i.SchemaLocation.String(), protSep), i.SchemaLocation.String(); pos > 0 {
			impPath = impPath[pos+len(protSep):]
		} else {
			impPath = path.Join(path.Dir(bag.Schema.loadURI), impPath)
		}
		impPath = path.Join(path.Dir(impPath), goPkgPrefix+path.Base(impPath)+goPkgSuffix)
		bag.imports[impName] = path.Join(PkgGen.BasePath, impPath)
	}
	i.elemBase.afterMakePkg(bag)
}

func (k *Key) makePkg(bag *PkgBag) {
	k.elemBase.beforeMakePkg(bag)
	k.hasElemField.makePkg(bag)
	k.hasElemSelector.makePkg(bag)
	k.elemBase.afterMakePkg(bag)
}

func (k *KeyRef) makePkg(bag *PkgBag) {
	k.elemBase.beforeMakePkg(bag)
	k.hasElemField.makePkg(bag)
	k.hasElemSelector.makePkg(bag)
	k.elemBase.afterMakePkg(bag)
}

func (l *List) makePkg(bag *PkgBag) {
	var safeName string
	l.elemBase.beforeMakePkg(bag)
	l.hasElemsSimpleType.makePkg(bag)
	rtr := bag.resolveQnameRef(l.ItemType.String(), "T", nil)
	if len(rtr) == 0 {
		rtr = l.SimpleTypes[0].Name.String()
	}
	st := bag.Stacks.CurSimpleType()
	safeName = bag.safeName(ustr.PrependIf(st.Name.String(), "T"))
	body, doc := "", sfmt("%v declares a String containing a whitespace-separated list of %v values. This Values() method creates and returns a slice of all elements in that list", safeName, rtr)
	body = sfmt("svals := %v.ListValues(string(me)); list = make([]%v, len(svals)); for i, s := range svals { list[i].Set(s) }; return", bag.impName, rtr)
	bag.ctd.addMethod(l, safeName, "Values", sfmt("(list []%v)", rtr), body, doc+".", l.Annotation)
	for baseType := bag.simpleBaseTypes[rtr]; len(baseType) > 0; baseType = bag.simpleBaseTypes[baseType] {
		body = sfmt("svals := %v.ListValues(string(me)); list = make([]%v, len(svals)); for i, s := range svals { list[i].Set(s) }; return", bag.impName, baseType)
		bag.ctd.addMethod(l, safeName, "Values"+bag.safeName(baseType), sfmt("(list []%v)", baseType), body, sfmt("%s, typed as %s.", doc, baseType), l.Annotation)
	}
	l.elemBase.afterMakePkg(bag)
}

func (n *Notation) makePkg(bag *PkgBag) {
	n.elemBase.beforeMakePkg(bag)
	n.hasElemAnnotation.makePkg(bag)
	bag.appendFmt(false, "%vNotations.Add(%#v, %#v, %#v, %#v)", idPrefix, n.Id, n.Name, n.Public, n.System)
	n.elemBase.afterMakePkg(bag)
}

func (r *Redefine) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.hasElemsSimpleType.makePkg(bag)
	r.hasElemsAttributeGroup.makePkg(bag)
	r.hasElemsGroup.makePkg(bag)
	r.hasElemsComplexType.makePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionComplexContent) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.hasElemsAttribute.makePkg(bag)
	r.hasElemsAnyAttribute.makePkg(bag)
	r.hasElemsAttributeGroup.makePkg(bag)
	r.hasElemAll.makePkg(bag)
	r.hasElemsChoice.makePkg(bag)
	r.hasElemsSequence.makePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleContent) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.hasElemsSimpleType.makePkg(bag)
	r.hasElemsAttribute.makePkg(bag)
	r.hasElemsAnyAttribute.makePkg(bag)
	r.hasElemsAttributeGroup.makePkg(bag)
	r.hasElemLength.makePkg(bag)
	r.hasElemPattern.makePkg(bag)
	r.hasElemsEnumeration.makePkg(bag)
	r.hasElemFractionDigits.makePkg(bag)
	r.hasElemMaxExclusive.makePkg(bag)
	r.hasElemMaxInclusive.makePkg(bag)
	r.hasElemMaxLength.makePkg(bag)
	r.hasElemMinExclusive.makePkg(bag)
	r.hasElemMinInclusive.makePkg(bag)
	r.hasElemMinLength.makePkg(bag)
	r.hasElemTotalDigits.makePkg(bag)
	r.hasElemWhiteSpace.makePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleEnumeration) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	safeName := bag.safeName(ustr.PrependIf(bag.Stacks.CurSimpleType().Name.String(), "T"))
	var doc = sfmt("Returns true if the value of this enumerated %v is %#v.", safeName, r.Value)
	bag.ctd.addMethod(r, safeName, "Is"+bag.safeName(r.Value), "bool", sfmt("return me.String() == %#v", r.Value), doc)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleFractionDigits) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleLength) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleMaxExclusive) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleMaxInclusive) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleMaxLength) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleMinExclusive) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleMinInclusive) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleMinLength) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimplePattern) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleTotalDigits) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleType) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.hasElemsSimpleType.makePkg(bag)
	r.hasElemLength.makePkg(bag)
	r.hasElemPattern.makePkg(bag)
	r.hasElemsEnumeration.makePkg(bag)
	r.hasElemFractionDigits.makePkg(bag)
	r.hasElemMaxExclusive.makePkg(bag)
	r.hasElemMaxInclusive.makePkg(bag)
	r.hasElemMaxLength.makePkg(bag)
	r.hasElemMinExclusive.makePkg(bag)
	r.hasElemMinInclusive.makePkg(bag)
	r.hasElemMinLength.makePkg(bag)
	r.hasElemTotalDigits.makePkg(bag)
	r.hasElemWhiteSpace.makePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (r *RestrictionSimpleWhiteSpace) makePkg(bag *PkgBag) {
	r.elemBase.beforeMakePkg(bag)
	r.elemBase.afterMakePkg(bag)
}

func (s *Schema) makePkg(bag *PkgBag) {
	s.elemBase.beforeMakePkg(bag)
	s.hasElemsImport.makePkg(bag)
	s.hasElemsSimpleType.makePkg(bag)
	s.hasElemsAttribute.makePkg(bag)
	s.hasElemsAttributeGroup.makePkg(bag)
	s.hasElemsComplexType.makePkg(bag)
	s.hasElemsElement.makePkg(bag)
	s.hasElemsGroup.makePkg(bag)
	s.hasElemsRedefine.makePkg(bag)
	s.elemBase.afterMakePkg(bag)
}

func (s *Selector) makePkg(bag *PkgBag) {
	s.elemBase.beforeMakePkg(bag)
	s.hasElemAnnotation.makePkg(bag)
	s.elemBase.afterMakePkg(bag)
}

func (s *Sequence) makePkg(bag *PkgBag) {
	s.elemBase.beforeMakePkg(bag)
	s.hasElemsAny.makePkg(bag)
	s.hasElemsChoice.makePkg(bag)
	s.hasElemsGroup.makePkg(bag)
	s.hasElemsSequence.makePkg(bag)
	s.hasElemsElement.makePkg(bag)
	s.elemBase.afterMakePkg(bag)
}

func (s *SimpleContent) makePkg(bag *PkgBag) {
	s.elemBase.beforeMakePkg(bag)
	s.hasElemExtensionSimpleContent.makePkg(bag)
	s.hasElemRestrictionSimpleContent.makePkg(bag)
	s.elemBase.afterMakePkg(bag)
}

func (s *SimpleType) makePkg(bag *PkgBag) {
	var typeName = s.Name
	var baseType, safeName = "", ""
	var resolve = true
	var isPt bool
	if len(typeName) == 0 {
		typeName = bag.AnonName(s.longSafeName(bag))
		s.Name = typeName
	} else {
		s.Name = typeName
	}
	typeName = xsdt.NCName(ustr.PrependIf(typeName.String(), "T"))
	s.elemBase.beforeMakePkg(bag)
	bag.Stacks.SimpleType.Push(s)
	safeName = bag.safeName(typeName.String())
	if s.RestrictionSimpleType != nil {
		if baseType = s.RestrictionSimpleType.Base.String(); (len(baseType) == 0) && (len(s.RestrictionSimpleType.SimpleTypes) > 0) {
			resolve, baseType = false, s.RestrictionSimpleType.SimpleTypes[0].Name.String()
		}
	}
	if len(baseType) == 0 {
		baseType = bag.xsdStringTypeRef()
	}
	if resolve {
		baseType = bag.resolveQnameRef(baseType, "T", nil)
	}
	bag.simpleBaseTypes[safeName] = baseType
	if isPt = bag.isParseType(baseType); isPt {
		bag.parseTypes[safeName] = true
	}
	var td = bag.addType(s, safeName, baseType, s.Annotation)
	var doc string
	if isPt {
		doc = sfmt("Since %v is a non-string scalar type (either boolean or numeric), sets the current value obtained from parsing the specified string.", safeName)
	} else {
		doc = sfmt("Since %v is just a simple String type, this merely sets the current value from the specified string.", safeName)
	}
	td.addMethod(nil, "*"+safeName, "Set (s string)", "", sfmt("(*%v)(me).Set(s)", baseType), doc)
	if isPt {
		doc = sfmt("Returns a string representation of this %v's current non-string scalar value.", safeName)
	} else {
		doc = sfmt("Since %v is just a simple String type, this merely returns the current string value.", safeName)
	}
	td.addMethod(nil, safeName, "String", "string", sfmt("return %v(me).String()", baseType), doc)
	doc = sfmt("This convenience method just performs a simple type conversion to %v's alias type %v.", safeName, baseType)
	td.addMethod(nil, safeName, "To"+bag.safeName(baseType), baseType, sfmt("return %v(me)", baseType), doc)
	s.hasElemRestrictionSimpleType.makePkg(bag)
	s.hasElemList.makePkg(bag)
	s.hasElemUnion.makePkg(bag)
	bag.Stacks.SimpleType.Pop()
	s.elemBase.afterMakePkg(bag)
}

func (u *Union) makePkg(bag *PkgBag) {
	var memberTypes []string
	var rtn, rtnSafeName, safeName string
	u.elemBase.beforeMakePkg(bag)
	u.hasElemsSimpleType.makePkg(bag)
	memberTypes = ustr.Split(u.MemberTypes, " ")
	for _, st := range u.SimpleTypes {
		memberTypes = append(memberTypes, st.Name.String())
	}
	for _, mt := range memberTypes {
		rtn = bag.resolveQnameRef(mt, "T", nil)
		safeName, rtnSafeName = bag.safeName(ustr.PrependIf(bag.Stacks.CurSimpleType().Name.String(), "T")), bag.safeName(rtn)
		bag.ctd.addMethod(u, safeName, "To"+rtnSafeName, rtn, sfmt(ustr.Ifs(bag.isParseType(rtn), "var x = new(%v); x.Set(me.String()); return *x", "return %v(me)"), rtn), sfmt("%v is an XSD union-type of several types. This is a simple type conversion to %v, but keep in mind the actual value may or may not be a valid %v value.", safeName, rtnSafeName, rtnSafeName), u.Annotation)
	}
	u.elemBase.afterMakePkg(bag)
}

func (u *Unique) makePkg(bag *PkgBag) {
	u.elemBase.beforeMakePkg(bag)
	u.hasElemField.makePkg(bag)
	u.hasElemSelector.makePkg(bag)
	u.elemBase.afterMakePkg(bag)
}

func anns(a *All, cc *ComplexContent) (anns []*Annotation) {
	if (a != nil) && (a.Annotation != nil) {
		anns = append(anns, a.Annotation)
	}
	if cc != nil {
		if cc.Annotation != nil {
			anns = append(anns, cc.Annotation)
		}
		if ecc := cc.ExtensionComplexContent; (ecc != nil) && (ecc.Annotation != nil) {
			anns = append(anns, ecc.Annotation)
		}
	}
	return
}

func pluralize(s string) string {
	for _, psp := range PkgGen.PluralizeSpecialPrefixes {
		if strings.HasPrefix(s, psp) {
			return ustr.Pluralize(s[len(psp):] + s[:len(psp)])
		}
	}
	return ustr.Pluralize(s)
}

func sfmt(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}

// For any rune, return a rune that is a valid in an identifier
func coerceToIdentifierRune(ch rune) rune {
	if !unicode.IsLetter(ch) && !unicode.IsNumber(ch) {
		return '_'
	}
	return ch
}

// Take any string and convert it to a valid identifier
// Appends an underscore if the first rune is a number
func safeIdentifier(s string) string {
	s = strings.Map(coerceToIdentifierRune, s)
	if unicode.IsNumber([]rune(s)[0]) {
		s = fmt.Sprint("_", s)
	}
	return s
}

func subMakeElem(bag *PkgBag, td *declType, el *Element, done map[string]bool, parentMaxOccurs xsdt.Long, anns ...*Annotation) {
	var elCache map[string]string
	anns = append(anns, el.Annotation)
	if refName := bag.elemKeys[el]; (len(refName) > 0) && (!done[refName]) {
		if done[refName], elCache = true, ustr.Ifm((parentMaxOccurs == 1) && (el.hasAttrMaxOccurs.Value() == 1), bag.elemsCacheOnce, bag.elemsCacheMult); !strings.HasPrefix(elCache[refName], bag.impName+"."+idPrefix) {
			td.addEmbed(el, elCache[refName], anns...)
		}
	}
}

func subMakeElemGroup(bag *PkgBag, td *declType, gr *Group, done map[string]bool, anns ...*Annotation) {
	var refImp string
	anns = append(anns, gr.Annotation)
	if refName := bag.resolveQnameRef(gr.Ref.String(), "", &refImp); !done[refName] {
		if done[refName] = true; len(refImp) > 0 {
			if !strings.HasPrefix(refName, bag.impName+"."+idPrefix) {
				td.addEmbed(gr, refImp+"."+idPrefix+"HasGroup_"+refName[(len(refImp)+1):], anns...)
			}
		} else {
			td.addEmbed(gr, idPrefix+"HasGroup_"+refName, anns...)
		}
	}
}
