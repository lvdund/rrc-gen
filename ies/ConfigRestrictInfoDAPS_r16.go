package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoDAPS_r16 struct {
	powerCoordination_r16 *ConfigRestrictInfoDAPS_r16_powerCoordination_r16 `optional`
}

func (ie *ConfigRestrictInfoDAPS_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.powerCoordination_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.powerCoordination_r16 != nil {
		if err = ie.powerCoordination_r16.Encode(w); err != nil {
			return utils.WrapError("Encode powerCoordination_r16", err)
		}
	}
	return nil
}

func (ie *ConfigRestrictInfoDAPS_r16) Decode(r *uper.UperReader) error {
	var err error
	var powerCoordination_r16Present bool
	if powerCoordination_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if powerCoordination_r16Present {
		ie.powerCoordination_r16 = new(ConfigRestrictInfoDAPS_r16_powerCoordination_r16)
		if err = ie.powerCoordination_r16.Decode(r); err != nil {
			return utils.WrapError("Decode powerCoordination_r16", err)
		}
	}
	return nil
}
