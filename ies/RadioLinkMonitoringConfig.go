package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RadioLinkMonitoringConfig struct {
	failureDetectionResourcesToAddModList  []RadioLinkMonitoringRS                                `lb:1,ub:maxNrofFailureDetectionResources,optional`
	failureDetectionResourcesToReleaseList []RadioLinkMonitoringRS_Id                             `lb:1,ub:maxNrofFailureDetectionResources,optional`
	beamFailureInstanceMaxCount            *RadioLinkMonitoringConfig_beamFailureInstanceMaxCount `optional`
	beamFailureDetectionTimer              *RadioLinkMonitoringConfig_beamFailureDetectionTimer   `optional`
	beamfailure_r17                        *BeamFailureDetection_r17                              `optional,ext-1`
}

func (ie *RadioLinkMonitoringConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.beamfailure_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.failureDetectionResourcesToAddModList) > 0, len(ie.failureDetectionResourcesToReleaseList) > 0, ie.beamFailureInstanceMaxCount != nil, ie.beamFailureDetectionTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.failureDetectionResourcesToAddModList) > 0 {
		tmp_failureDetectionResourcesToAddModList := utils.NewSequence[*RadioLinkMonitoringRS]([]*RadioLinkMonitoringRS{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		for _, i := range ie.failureDetectionResourcesToAddModList {
			tmp_failureDetectionResourcesToAddModList.Value = append(tmp_failureDetectionResourcesToAddModList.Value, &i)
		}
		if err = tmp_failureDetectionResourcesToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode failureDetectionResourcesToAddModList", err)
		}
	}
	if len(ie.failureDetectionResourcesToReleaseList) > 0 {
		tmp_failureDetectionResourcesToReleaseList := utils.NewSequence[*RadioLinkMonitoringRS_Id]([]*RadioLinkMonitoringRS_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		for _, i := range ie.failureDetectionResourcesToReleaseList {
			tmp_failureDetectionResourcesToReleaseList.Value = append(tmp_failureDetectionResourcesToReleaseList.Value, &i)
		}
		if err = tmp_failureDetectionResourcesToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode failureDetectionResourcesToReleaseList", err)
		}
	}
	if ie.beamFailureInstanceMaxCount != nil {
		if err = ie.beamFailureInstanceMaxCount.Encode(w); err != nil {
			return utils.WrapError("Encode beamFailureInstanceMaxCount", err)
		}
	}
	if ie.beamFailureDetectionTimer != nil {
		if err = ie.beamFailureDetectionTimer.Encode(w); err != nil {
			return utils.WrapError("Encode beamFailureDetectionTimer", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.beamfailure_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RadioLinkMonitoringConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.beamfailure_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode beamfailure_r17 optional
			if ie.beamfailure_r17 != nil {
				if err = ie.beamfailure_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamfailure_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *RadioLinkMonitoringConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var failureDetectionResourcesToAddModListPresent bool
	if failureDetectionResourcesToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var failureDetectionResourcesToReleaseListPresent bool
	if failureDetectionResourcesToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var beamFailureInstanceMaxCountPresent bool
	if beamFailureInstanceMaxCountPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var beamFailureDetectionTimerPresent bool
	if beamFailureDetectionTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if failureDetectionResourcesToAddModListPresent {
		tmp_failureDetectionResourcesToAddModList := utils.NewSequence[*RadioLinkMonitoringRS]([]*RadioLinkMonitoringRS{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		fn_failureDetectionResourcesToAddModList := func() *RadioLinkMonitoringRS {
			return new(RadioLinkMonitoringRS)
		}
		if err = tmp_failureDetectionResourcesToAddModList.Decode(r, fn_failureDetectionResourcesToAddModList); err != nil {
			return utils.WrapError("Decode failureDetectionResourcesToAddModList", err)
		}
		ie.failureDetectionResourcesToAddModList = []RadioLinkMonitoringRS{}
		for _, i := range tmp_failureDetectionResourcesToAddModList.Value {
			ie.failureDetectionResourcesToAddModList = append(ie.failureDetectionResourcesToAddModList, *i)
		}
	}
	if failureDetectionResourcesToReleaseListPresent {
		tmp_failureDetectionResourcesToReleaseList := utils.NewSequence[*RadioLinkMonitoringRS_Id]([]*RadioLinkMonitoringRS_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofFailureDetectionResources}, false)
		fn_failureDetectionResourcesToReleaseList := func() *RadioLinkMonitoringRS_Id {
			return new(RadioLinkMonitoringRS_Id)
		}
		if err = tmp_failureDetectionResourcesToReleaseList.Decode(r, fn_failureDetectionResourcesToReleaseList); err != nil {
			return utils.WrapError("Decode failureDetectionResourcesToReleaseList", err)
		}
		ie.failureDetectionResourcesToReleaseList = []RadioLinkMonitoringRS_Id{}
		for _, i := range tmp_failureDetectionResourcesToReleaseList.Value {
			ie.failureDetectionResourcesToReleaseList = append(ie.failureDetectionResourcesToReleaseList, *i)
		}
	}
	if beamFailureInstanceMaxCountPresent {
		ie.beamFailureInstanceMaxCount = new(RadioLinkMonitoringConfig_beamFailureInstanceMaxCount)
		if err = ie.beamFailureInstanceMaxCount.Decode(r); err != nil {
			return utils.WrapError("Decode beamFailureInstanceMaxCount", err)
		}
	}
	if beamFailureDetectionTimerPresent {
		ie.beamFailureDetectionTimer = new(RadioLinkMonitoringConfig_beamFailureDetectionTimer)
		if err = ie.beamFailureDetectionTimer.Decode(r); err != nil {
			return utils.WrapError("Decode beamFailureDetectionTimer", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			beamfailure_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode beamfailure_r17 optional
			if beamfailure_r17Present {
				ie.beamfailure_r17 = new(BeamFailureDetection_r17)
				if err = ie.beamfailure_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamfailure_r17", err)
				}
			}
		}
	}
	return nil
}
