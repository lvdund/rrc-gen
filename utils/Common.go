package utils

import (
	"bytes"
	"fmt"
	"io"
	"github.com/lvdund/asn1go/uper"
)

func encodeTransferMessage(ies []NgapMessageIE) (w []byte, err error) {
	var buf bytes.Buffer
	aw := uper.NewWriter(&buf)
	if err = aw.WriteBool(uper.Zero); err != nil {
		return
	}
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	if err = uper.WriteSequenceOf[NgapMessageIE](ies, aw, &uper.Constraint{
		Lb: 0,
		Ub: int64(uper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	err = aw.Close()
	w = buf.Bytes()
	return
}

func encodeMessage(w io.Writer, present uint8, procedureCode int64, criticality uper.Enumerated, ies []NgapMessageIE) (err error) {
	aw := uper.NewWriter(w)
	if err = aw.WriteBool(uper.Zero); err != nil {
		return
	}
	if err = aw.WriteChoice(uint64(present), 2, true); err != nil {
		return
	}
	pCode := ProcedureCode{
		Value: uper.Integer(procedureCode),
	}
	if err = pCode.Encode(aw); err != nil {
		return
	}
	cr := Criticality{
		Value: criticality,
	}
	if err = cr.Encode(aw); err != nil {
		return
	}
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	var buf bytes.Buffer
	cW := uper.NewWriter(&buf) //container writer
	cW.WriteBool(uper.Zero)
	if err = uper.WriteSequenceOf[NgapMessageIE](ies, cW, &uper.Constraint{
		Lb: 0,
		Ub: int64(uper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	if err = cW.Close(); err != nil { //finalize
		return
	}
	if err = aw.WriteOpenType(buf.Bytes()); err != nil {
		return
	}
	err = aw.Close()
	return
}

// represent an IE in Ngap messages
type NgapMessageIE struct {
	Id          ProtocolIEID //protocol IE identity
	Criticality Criticality
	Value       uper.UperMarshaller //open type
}

func (ie NgapMessageIE) Encode(w *uper.UperWriter) (err error) {
	//1. encode protocol Ie Id
	if err = ie.Id.Encode(w); err != nil {
		return
	}
	//2. encode criticality
	if err = ie.Criticality.Encode(w); err != nil {
		return
	}
	//3. encode NgapIE
	//encode IE into a byte array first
	var buf bytes.Buffer
	ieW := uper.NewWriter(&buf)
	if err = ie.Value.Encode(ieW); err != nil {
		return
	}
	ieW.Close()
	//then write the array as open type
	err = w.WriteOpenType(buf.Bytes())
	return
}

type ProcedureCode struct {
	Value uper.Integer `aper:"valueLB:0,valueUB:255"`
}

func (ie *ProcedureCode) Decode(r *uper.UperReader) error {
	if v, err := r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = uper.Integer(v)
	}
	return nil
}
func (ie *ProcedureCode) Encode(r *uper.UperWriter) (err error) {
	if err = r.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return err
	}
	return nil
}

type TriggeringMessage struct {
	Value uper.Enumerated `aper:"valueLB:0,valueUB:2"`
}

func (ie *TriggeringMessage) Decode(r *uper.UperReader) error {
	if v, err := r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = uper.Enumerated(v)
	}
	return nil
}
func (ie *TriggeringMessage) Encode(r *uper.UperWriter) (err error) {
	if err = r.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	}
	return nil
}

type Criticality struct {
	Value uper.Enumerated `aper:"valueLB:0,valueUB:2"`
}

func (ie *Criticality) Decode(r *uper.UperReader) error {
	if v, err := r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = uper.Enumerated(v)
	}
	return nil
}
func (ie *Criticality) Encode(r *uper.UperWriter) (err error) {
	if err = r.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	}
	return nil
}

type ProtocolIEID struct {
	Value uper.Integer `aper:"valueLB:0,valueUB:65535"`
}

func (ie *ProtocolIEID) Decode(r *uper.UperReader) error {
	if v, err := r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return err
	} else {
		ie.Value = uper.Integer(v)
	}
	return nil
}
func (ie *ProtocolIEID) Encode(r *uper.UperWriter) (err error) {
	if err = r.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return err
	}
	return nil
}

func msgErrors(err1, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	}
	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}
	return fmt.Errorf("%v: %v", err1, err2)
}

// NgapPdu - Present
const (
	NgapPresentNothing uint8 = iota
	NgapPduInitiatingMessage
	NgapPduSuccessfulOutcome
	NgapPduUnsuccessfulOutcome
)
