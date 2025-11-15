package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfigCommon_v1700 struct {
	initialUplinkBWP_RedCap_r17 *BWP_UplinkCommon `optional`
}

func (ie *UplinkConfigCommon_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.initialUplinkBWP_RedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.initialUplinkBWP_RedCap_r17 != nil {
		if err = ie.initialUplinkBWP_RedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode initialUplinkBWP_RedCap_r17", err)
		}
	}
	return nil
}

func (ie *UplinkConfigCommon_v1700) Decode(r *uper.UperReader) error {
	var err error
	var initialUplinkBWP_RedCap_r17Present bool
	if initialUplinkBWP_RedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if initialUplinkBWP_RedCap_r17Present {
		ie.initialUplinkBWP_RedCap_r17 = new(BWP_UplinkCommon)
		if err = ie.initialUplinkBWP_RedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode initialUplinkBWP_RedCap_r17", err)
		}
	}
	return nil
}
