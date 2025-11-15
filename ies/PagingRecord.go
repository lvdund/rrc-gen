package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PagingRecord struct {
	ue_Identity PagingUE_Identity        `madatory`
	accessType  *PagingRecord_accessType `optional`
}

func (ie *PagingRecord) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.accessType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ue_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode ue_Identity", err)
	}
	if ie.accessType != nil {
		if err = ie.accessType.Encode(w); err != nil {
			return utils.WrapError("Encode accessType", err)
		}
	}
	return nil
}

func (ie *PagingRecord) Decode(r *uper.UperReader) error {
	var err error
	var accessTypePresent bool
	if accessTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ue_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode ue_Identity", err)
	}
	if accessTypePresent {
		ie.accessType = new(PagingRecord_accessType)
		if err = ie.accessType.Decode(r); err != nil {
			return utils.WrapError("Decode accessType", err)
		}
	}
	return nil
}
