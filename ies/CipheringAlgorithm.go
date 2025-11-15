package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CipheringAlgorithm_Enum_nea0   uper.Enumerated = 0
	CipheringAlgorithm_Enum_nea1   uper.Enumerated = 1
	CipheringAlgorithm_Enum_nea2   uper.Enumerated = 2
	CipheringAlgorithm_Enum_nea3   uper.Enumerated = 3
	CipheringAlgorithm_Enum_spare4 uper.Enumerated = 4
	CipheringAlgorithm_Enum_spare3 uper.Enumerated = 5
	CipheringAlgorithm_Enum_spare2 uper.Enumerated = 6
	CipheringAlgorithm_Enum_spare1 uper.Enumerated = 7
)

type CipheringAlgorithm struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CipheringAlgorithm) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CipheringAlgorithm", err)
	}
	return nil
}

func (ie *CipheringAlgorithm) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CipheringAlgorithm", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
