package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1620 struct {
	multiDCI_MultiTRP_r16  *MultiDCI_MultiTRP_r16                                `optional`
	supportFDM_SchemeB_r16 *FeatureSetDownlinkPerCC_v1620_supportFDM_SchemeB_r16 `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1620) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.multiDCI_MultiTRP_r16 != nil, ie.supportFDM_SchemeB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.multiDCI_MultiTRP_r16 != nil {
		if err = ie.multiDCI_MultiTRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode multiDCI_MultiTRP_r16", err)
		}
	}
	if ie.supportFDM_SchemeB_r16 != nil {
		if err = ie.supportFDM_SchemeB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportFDM_SchemeB_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1620) Decode(r *uper.UperReader) error {
	var err error
	var multiDCI_MultiTRP_r16Present bool
	if multiDCI_MultiTRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportFDM_SchemeB_r16Present bool
	if supportFDM_SchemeB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if multiDCI_MultiTRP_r16Present {
		ie.multiDCI_MultiTRP_r16 = new(MultiDCI_MultiTRP_r16)
		if err = ie.multiDCI_MultiTRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode multiDCI_MultiTRP_r16", err)
		}
	}
	if supportFDM_SchemeB_r16Present {
		ie.supportFDM_SchemeB_r16 = new(FeatureSetDownlinkPerCC_v1620_supportFDM_SchemeB_r16)
		if err = ie.supportFDM_SchemeB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportFDM_SchemeB_r16", err)
		}
	}
	return nil
}
