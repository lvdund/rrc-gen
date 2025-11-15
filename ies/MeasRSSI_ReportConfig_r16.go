package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasRSSI_ReportConfig_r16 struct {
	channelOccupancyThreshold_r16 *RSSI_Range_r16 `optional`
}

func (ie *MeasRSSI_ReportConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.channelOccupancyThreshold_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.channelOccupancyThreshold_r16 != nil {
		if err = ie.channelOccupancyThreshold_r16.Encode(w); err != nil {
			return utils.WrapError("Encode channelOccupancyThreshold_r16", err)
		}
	}
	return nil
}

func (ie *MeasRSSI_ReportConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var channelOccupancyThreshold_r16Present bool
	if channelOccupancyThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if channelOccupancyThreshold_r16Present {
		ie.channelOccupancyThreshold_r16 = new(RSSI_Range_r16)
		if err = ie.channelOccupancyThreshold_r16.Decode(r); err != nil {
			return utils.WrapError("Decode channelOccupancyThreshold_r16", err)
		}
	}
	return nil
}
