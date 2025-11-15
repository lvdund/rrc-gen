package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FilterConfigCLI_r16 struct {
	filterCoefficientSRS_RSRP_r16 FilterCoefficient `madatory`
	filterCoefficientCLI_RSSI_r16 FilterCoefficient `madatory`
}

func (ie *FilterConfigCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.filterCoefficientSRS_RSRP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode filterCoefficientSRS_RSRP_r16", err)
	}
	if err = ie.filterCoefficientCLI_RSSI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode filterCoefficientCLI_RSSI_r16", err)
	}
	return nil
}

func (ie *FilterConfigCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.filterCoefficientSRS_RSRP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode filterCoefficientSRS_RSRP_r16", err)
	}
	if err = ie.filterCoefficientCLI_RSSI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode filterCoefficientCLI_RSSI_r16", err)
	}
	return nil
}
