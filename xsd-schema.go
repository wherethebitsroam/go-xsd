package xsd

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"fmt"

	"github.com/go-utils/ufs"
	"github.com/go-utils/unet"
	"github.com/go-utils/ustr"
)

const (
	goPkgPrefix     = ""
	goPkgSuffix     = "_go"
	protSep         = "://"
	xsdNamespaceURI = "http://www.w3.org/2001/XMLSchema"
)

var (
	loadedSchemas = map[string]*Schema{}
)

// Schema ...
type Schema struct {
	elemBase
	XMLName            xml.Name          `xml:"schema"`
	XMLNamespacePrefix string            `xml:"-"`
	XMLNamespaces      map[string]string `xml:"-"`
	XMLIncludedSchemas []*Schema         `xml:"-"`
	XSDNamespacePrefix string            `xml:"-"`
	XSDParentSchema    *Schema           `xml:"-"`

	hasAttrAttributeFormDefault
	hasAttrBlockDefault
	hasAttrElementFormDefault
	hasAttrFinalDefault
	hasAttrLang
	hasAttrId
	hasAttrSchemaLocation
	hasAttrTargetNamespace
	hasAttrVersion
	hasElemAnnotation
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsComplexType
	hasElemsElement
	hasElemsGroup
	hasElemsInclude
	hasElemsImport
	hasElemsNotation
	hasElemsRedefine
	hasElemsSimpleType

	loadLocalPath, loadURI string
}

func (s *Schema) allSchemas(loadedSchemas map[string]bool) (schemas []*Schema) {
	schemas = append(schemas, s)
	loadedSchemas[s.loadURI] = true
	for _, ss := range s.XMLIncludedSchemas {
		if v, ok := loadedSchemas[ss.loadURI]; ok && v {
			continue
		}
		schemas = append(schemas, ss.allSchemas(loadedSchemas)...)
	}
	return
}

func (s *Schema) collectGlobals(bag *PkgBag, loadedSchemas map[string]bool) {
	loadedSchemas[s.loadURI] = true
	for _, att := range s.Attributes {
		bag.allAtts = append(bag.allAtts, att)
	}
	for _, agr := range s.AttributeGroups {
		bag.allAttGroups = append(bag.allAttGroups, agr)
	}
	for _, el := range s.Elements {
		bag.allElems = append(bag.allElems, el)
	}
	for _, egr := range s.Groups {
		bag.allElemGroups = append(bag.allElemGroups, egr)
	}
	for _, not := range s.Notations {
		bag.allNotations = append(bag.allNotations, not)
	}
	for _, ss := range s.XMLIncludedSchemas {
		if v, ok := loadedSchemas[ss.loadURI]; ok && v {
			continue
		}
		ss.collectGlobals(bag, loadedSchemas)
	}
}

func (s *Schema) globalComplexType(bag *PkgBag, name string, loadedSchemas map[string]bool) (ct *ComplexType) {
	var imp string
	for _, ct = range s.ComplexTypes {
		if bag.resolveQnameRef(ustr.PrefixWithSep(s.XMLNamespacePrefix, ":", ct.Name.String()), "T", &imp) == name {
			return
		}
	}
	loadedSchemas[s.loadURI] = true
	for _, ss := range s.XMLIncludedSchemas {
		if v, ok := loadedSchemas[ss.loadURI]; ok && v {
			//fmt.Printf("Ignoring processed schema: %s\n", ss.loadUri)
			continue
		}
		if ct = ss.globalComplexType(bag, name, loadedSchemas); ct != nil {
			return
		}
	}
	ct = nil
	return
}

func (s *Schema) globalElement(bag *PkgBag, name string) (el *Element) {
	var imp string
	if len(name) > 0 {
		var rname = bag.resolveQnameRef(name, "", &imp)
		for _, el = range s.Elements {
			if bag.resolveQnameRef(ustr.PrefixWithSep(s.XMLNamespacePrefix, ":", el.Name.String()), "", &imp) == rname {
				return
			}
		}
		for _, ss := range s.XMLIncludedSchemas {
			if el = ss.globalElement(bag, name); el != nil {
				return
			}
		}
	}
	el = nil
	return
}

func (s *Schema) globalSubstitutionElems(el *Element, loadedSchemas map[string]bool) (els []*Element) {
	var elName = el.Ref.String()
	if len(elName) == 0 {
		elName = el.Name.String()
	}
	for _, tle := range s.Elements {
		if (tle != el) && (len(tle.SubstitutionGroup) > 0) {
			if tle.SubstitutionGroup.String()[(strings.Index(tle.SubstitutionGroup.String(), ":")+1):] == elName {
				els = append(els, tle)
			}
		}
	}
	loadedSchemas[s.loadURI] = true
	for _, inc := range s.XMLIncludedSchemas {
		if v, ok := loadedSchemas[inc.loadURI]; ok && v {
			//fmt.Printf("Ignoring processed schema: %s\n", inc.loadUri)
			continue
		}
		els = append(els, inc.globalSubstitutionElems(el, loadedSchemas)...)
	}
	return
}

// MakeGoPkgSrcFile ...
func (s *Schema) MakeGoPkgSrcFile() (goOutFilePath string, err error) {
	var goOutDirPath = filepath.Join(filepath.Dir(s.loadLocalPath), goPkgPrefix+filepath.Base(s.loadLocalPath)+goPkgSuffix)
	goOutFilePath = filepath.Join(goOutDirPath, path.Base(s.loadURI)+".go")
	var bag = newPkgBag(s)
	loadedSchemas := make(map[string]bool)
	for _, inc := range s.allSchemas(loadedSchemas) {
		bag.Schema = inc
		inc.makePkg(bag)
	}
	bag.Schema = s
	s.hasElemAnnotation.makePkg(bag)
	bag.appendFmt(true, "")
	s.makePkg(bag)
	if err = ufs.EnsureDirExists(filepath.Dir(goOutFilePath)); err == nil {
		err = ufs.WriteTextFile(goOutFilePath, bag.assembleSource())
	}
	return
}

func (s *Schema) onLoad(rootAtts []xml.Attr, loadURI, localPath string) (err error) {
	var tmpURL string
	var sd *Schema
	loadedSchemas[loadURI] = s
	s.loadLocalPath = localPath
	s.loadURI = loadURI
	s.XMLNamespaces = map[string]string{}
	for _, att := range rootAtts {
		if att.Name.Space == "xmlns" {
			s.XMLNamespaces[att.Name.Local] = att.Value
		} else if len(att.Name.Space) > 0 {

		} else if att.Name.Local == "xmlns" {
			s.XMLNamespaces[""] = att.Value
		}
	}
	for k, v := range s.XMLNamespaces {
		if v == xsdNamespaceURI {
			s.XSDNamespacePrefix = k
		} else if v == s.TargetNamespace.String() {
			s.XMLNamespacePrefix = k
		}
	}
	if len(s.XMLNamespaces["xml"]) == 0 {
		s.XMLNamespaces["xml"] = "http://www.w3.org/XML/1998/namespace"
	}
	s.XMLIncludedSchemas = []*Schema{}
	for _, inc := range s.Includes {
		if tmpURL = inc.SchemaLocation.String(); strings.Index(tmpURL, protSep) < 0 {
			tmpURL = path.Join(path.Dir(loadURI), tmpURL)
		}
		var ok bool
		var toLoadURI string
		if pos := strings.Index(tmpURL, protSep); pos >= 0 {
			toLoadURI = tmpURL[pos+len(protSep):]
		} else {
			toLoadURI = tmpURL
		}
		if sd, ok = loadedSchemas[toLoadURI]; !ok {
			if sd, err = LoadSchema(tmpURL, len(localPath) > 0); err != nil {
				return
			}
		}
		sd.XSDParentSchema = s
		s.XMLIncludedSchemas = append(s.XMLIncludedSchemas, sd)
	}
	s.initElement(nil)
	return
}

// RootSchema ...
func (s *Schema) RootSchema(pathSchemas []string) *Schema {
	if s.XSDParentSchema != nil {
		for _, sch := range pathSchemas {
			if s.XSDParentSchema.loadURI == sch {
				fmt.Printf("schema loop detected %+v - > %s!\n", pathSchemas, s.XSDParentSchema.loadURI)
				return s
			}
		}
		pathSchemas = append(pathSchemas, s.loadURI)
		return s.XSDParentSchema.RootSchema(pathSchemas)
	}
	return s
}

// ClearLoadedSchemasCache ...
func ClearLoadedSchemasCache() {
	loadedSchemas = map[string]*Schema{}
}

func loadSchema(r io.Reader, loadURI, localPath string) (sd *Schema, err error) {
	var data []byte
	var rootAtts []xml.Attr
	if data, err = ioutil.ReadAll(r); err == nil {
		var t xml.Token
		sd = new(Schema)
		for xd := xml.NewDecoder(bytes.NewReader(data)); err == nil; {
			if t, err = xd.Token(); err == nil {
				if startEl, ok := t.(xml.StartElement); ok {
					rootAtts = startEl.Attr
					break
				}
			}
		}
		if err = xml.Unmarshal(data, sd); err == nil {
			err = sd.onLoad(rootAtts, loadURI, localPath)
		}
		if err != nil {
			sd = nil
		}
	}
	return
}

func loadSchemaFile(filename string, loadURI string) (sd *Schema, err error) {
	var file *os.File
	if file, err = os.Open(filename); err == nil {
		defer file.Close()
		sd, err = loadSchema(file, loadURI, filename)
	}
	return
}

// LoadSchema ...
func LoadSchema(uri string, localCopy bool) (sd *Schema, err error) {
	var protocol, localPath string
	var rc io.ReadCloser

	if pos := strings.Index(uri, protSep); pos < 0 {
		protocol = "http" + protSep
	} else {
		protocol = uri[:pos+len(protSep)]
		uri = uri[pos+len(protSep):]
	}
	if localCopy {
		if localPath = filepath.Join(PkgGen.BaseCodePath, uri); !ufs.FileExists(localPath) {
			if err = ufs.EnsureDirExists(filepath.Dir(localPath)); err == nil {
				err = unet.DownloadFile(protocol+uri, localPath)
			}
		}
		if err == nil {
			if sd, err = loadSchemaFile(localPath, uri); sd != nil {
				sd.loadLocalPath = localPath
			}
		}
	} else if rc, err = unet.OpenRemoteFile(protocol + uri); err == nil {
		defer rc.Close()
		sd, err = loadSchema(rc, uri, "")
	}
	return
}
