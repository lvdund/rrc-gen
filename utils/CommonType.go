package utils

import (
	"github.com/lvdund/asn1go/uper"
)

const (
	ChoiceNothing = 0
	ChoiceRelease = 1
	ChoiceSetup   = 2
)

type BOOLEAN struct {
	Value bool
}

func NewBOOLEAN(v bool, ext bool) BOOLEAN {
	return BOOLEAN{
		Value: v,
	}
}

func (t *BOOLEAN) Encode(w *uper.UperWriter) (err error) {
	err = w.WriteBoolean(t.Value)
	return
}
func (t *BOOLEAN) Decode(r *uper.UperReader) (err error) {
	v, err := r.ReadBoolean()
	if err != nil {
		return err
	}
	t.Value = v
	return
}

type ENUMERATED struct {
	Value uper.Enumerated
	c     uper.Constraint
	ext   bool
}

func NewENUMERATED(v int64, c uper.Constraint, ext bool) ENUMERATED {
	return ENUMERATED{
		Value: uper.Enumerated(v),
		c:     c,
		ext:   ext,
	}
}
func (t *ENUMERATED) Encode(w *uper.UperWriter) (err error) {
	err = w.WriteEnumerate(uint64(t.Value), t.c, t.ext)
	return
}
func (t *ENUMERATED) Decode(r *uper.UperReader) (err error) {
	v, err := r.ReadEnumerate(t.c, t.ext)
	t.Value = uper.Enumerated(v)
	return
}

type BITSTRING struct {
	Value uper.BitString
	c     uper.Constraint
	ext   bool
}

func NewBITSTRING(v uper.BitString, c uper.Constraint, ext bool) BITSTRING {
	return BITSTRING{
		Value: uper.BitString{
			Bytes:   v.Bytes,
			NumBits: v.NumBits,
		},
		c:   c,
		ext: ext,
	}
}
func (t *BITSTRING) Encode(w *uper.UperWriter) (err error) {
	if t.c.Lb == t.c.Ub {
		t.Value.NumBits = uint64(t.c.Lb)
	} else if len(t.Value.Bytes)*8 < int(t.c.Lb) {
		t.Value.NumBits = uint64(t.c.Lb)
	}
	err = w.WriteBitString(t.Value.Bytes, uint(t.Value.NumBits), &t.c, t.ext)
	return
}
func (t *BITSTRING) Decode(r *uper.UperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&t.c, t.ext); err != nil {
		return
	}
	t.Value.Bytes = v
	t.Value.NumBits = uint64(n)
	return
}

type OCTETSTRING struct {
	Value uper.OctetString
	c     uper.Constraint
	ext   bool
}

func NewOCTETSTRING(v []byte, c uper.Constraint, ext bool) OCTETSTRING {
	return OCTETSTRING{
		Value: v,
		c:     c,
		ext:   ext,
	}
}
func (t *OCTETSTRING) Encode(w *uper.UperWriter) (err error) {
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		err = w.WriteOctetString(t.Value, nil, t.ext)
	} else {
		err = w.WriteOctetString(t.Value, &t.c, t.ext)
	}
	return
}
func (t *OCTETSTRING) Decode(r *uper.UperReader) (err error) {
	var v uper.OctetString
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		if v, err = r.ReadOctetString(nil, t.ext); err != nil {
			return
		}
	} else {
		if v, err = r.ReadOctetString(&t.c, t.ext); err != nil {
			return
		}
	}

	t.Value = v
	return
}

type INTEGER struct {
	Value uper.Integer
	c     uper.Constraint
	ext   bool
}

func NewINTEGER(v int64, c uper.Constraint, ext bool) INTEGER {
	return INTEGER{
		Value: uper.Integer(v),
		c:     c,
		ext:   ext,
	}
}
func (t *INTEGER) Encode(w *uper.UperWriter) (err error) {
	err = w.WriteInteger(int64(t.Value), &t.c, t.ext)
	return
}
func (t *INTEGER) Decode(r *uper.UperReader) (err error) {
	v, err := r.ReadInteger(&t.c, t.ext)
	t.Value = uper.Integer(v)
	return
}

type Sequence[T uper.IE] struct {
	Value []T
	c     uper.Constraint
	ext   bool
}

func NewSequence[T uper.IE](items []T, c uper.Constraint, ext bool) Sequence[T] {
	return Sequence[T]{
		Value: items,
		c:     c,
		ext:   ext,
	}
}

func (s *Sequence[T]) Encode(w *uper.UperWriter) (err error) {
	if err = uper.WriteSequenceOf[T](s.Value, w, &s.c, s.ext); err != nil {
		return
	}
	return
}
func (s *Sequence[T]) Decode(r *uper.UperReader, fn func() T) (err error) {
	var newItems []T
	newItems, err = uper.ReadSequenceOfEx(fn, r, &s.c, s.ext)
	if err != nil {
		return
	}
	s.Value = []T{}
	s.Value = append(s.Value, newItems...)
	return
}
