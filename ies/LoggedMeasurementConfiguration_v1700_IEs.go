package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LoggedMeasurementConfiguration_v1700_IEs struct {
	sigLoggedMeasType_r17   *LoggedMeasurementConfiguration_v1700_IEs_sigLoggedMeasType_r17   `optional`
	earlyMeasIndication_r17 *LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17 `optional`
	areaConfiguration_v1700 *AreaConfiguration_v1700                                          `optional`
	nonCriticalExtension    interface{}                                                       `optional`
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sigLoggedMeasType_r17 != nil, ie.earlyMeasIndication_r17 != nil, ie.areaConfiguration_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sigLoggedMeasType_r17 != nil {
		if err = ie.sigLoggedMeasType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sigLoggedMeasType_r17", err)
		}
	}
	if ie.earlyMeasIndication_r17 != nil {
		if err = ie.earlyMeasIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode earlyMeasIndication_r17", err)
		}
	}
	if ie.areaConfiguration_v1700 != nil {
		if err = ie.areaConfiguration_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode areaConfiguration_v1700", err)
		}
	}
	return nil
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sigLoggedMeasType_r17Present bool
	if sigLoggedMeasType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var earlyMeasIndication_r17Present bool
	if earlyMeasIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var areaConfiguration_v1700Present bool
	if areaConfiguration_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sigLoggedMeasType_r17Present {
		ie.sigLoggedMeasType_r17 = new(LoggedMeasurementConfiguration_v1700_IEs_sigLoggedMeasType_r17)
		if err = ie.sigLoggedMeasType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sigLoggedMeasType_r17", err)
		}
	}
	if earlyMeasIndication_r17Present {
		ie.earlyMeasIndication_r17 = new(LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17)
		if err = ie.earlyMeasIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode earlyMeasIndication_r17", err)
		}
	}
	if areaConfiguration_v1700Present {
		ie.areaConfiguration_v1700 = new(AreaConfiguration_v1700)
		if err = ie.areaConfiguration_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode areaConfiguration_v1700", err)
		}
	}
	return nil
}
