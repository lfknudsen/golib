package structs

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func (v *Version) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: v.String(),
	}, nil
}

func (v *Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *Version) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(v.Major, start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(v.Minor, start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(v.Patch, start)
	if err != nil {
		return err
	}
	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	return err
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) Bytes() []byte {
	return []byte{v.Major, v.Minor, v.Patch}
}

func FromString(s string) (*Version, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid version string")
	}
	var major uint64
	if parts[0] == "v" {
		var e error
		major, e = strconv.ParseUint(parts[0][1:], 10, 8)
		if e != nil {
			return nil, e
		}
	} else {
		var e error
		major, e = strconv.ParseUint(parts[0], 10, 8)
		if e != nil {
			return nil, e
		}
	}
	minor, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return nil, err
	}
	patch, err := strconv.ParseUint(parts[2], 10, 8)
	if err != nil {
		return nil, err
	}
	return &Version{
		Major: uint8(major),
		Minor: uint8(minor),
		Patch: uint8(patch),
	}, nil
}

func FromBytes(b [3]byte) (*Version, error) {
	return &Version{
		Major: b[0],
		Minor: b[1],
		Patch: b[2],
	}, nil
}

func DecompileVersion(r io.Reader) (*Version, error) {
	buffer := [3]byte{}
	n, err := r.Read(buffer[:])
	if err != nil {
		return nil, err
	}
	if n != 3 {
		return nil, errors.New("invalid version string")
	}
	return FromBytes(buffer)
}
