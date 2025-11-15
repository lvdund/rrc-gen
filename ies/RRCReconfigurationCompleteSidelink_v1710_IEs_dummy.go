package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfigurationCompleteSidelink_v1710_IEs_dummy_Enum_true uper.Enumerated = 0
)

type RRCReconfigurationCompleteSidelink_v1710_IEs_dummy struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RRCReconfigurationCompleteSidelink_v1710_IEs_dummy) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RRCReconfigurationCompleteSidelink_v1710_IEs_dummy", err)
	}
	return nil
}

func (ie *RRCReconfigurationCompleteSidelink_v1710_IEs_dummy) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RRCReconfigurationCompleteSidelink_v1710_IEs_dummy", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
