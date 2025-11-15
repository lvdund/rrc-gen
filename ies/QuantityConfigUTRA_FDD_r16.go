package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfigUTRA_FDD_r16 struct {
	filterCoefficientRSCP_r16 FilterCoefficient `madatory`
	filterCoefficientEcNO_r16 FilterCoefficient `madatory`
}

func (ie *QuantityConfigUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.filterCoefficientRSCP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode filterCoefficientRSCP_r16", err)
	}
	if err = ie.filterCoefficientEcNO_r16.Encode(w); err != nil {
		return utils.WrapError("Encode filterCoefficientEcNO_r16", err)
	}
	return nil
}

func (ie *QuantityConfigUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.filterCoefficientRSCP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode filterCoefficientRSCP_r16", err)
	}
	if err = ie.filterCoefficientEcNO_r16.Decode(r); err != nil {
		return utils.WrapError("Decode filterCoefficientEcNO_r16", err)
	}
	return nil
}
