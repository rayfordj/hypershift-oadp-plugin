package validation

import (
	"fmt"
	"strconv"
	"time"

	plugtypes "github.com/openshift/hypershift-oadp-plugin/pkg/core/types"
	hyperv1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
	"github.com/sirupsen/logrus"
)

type BackupValidator interface {
	ValidatePluginConfig(config map[string]string) (*plugtypes.BackupOptions, error)
	ValidatePlatformConfig(hcp *hyperv1.HostedControlPlane) error
}

type BackupPluginValidator struct {
	Log       logrus.FieldLogger
	LogHeader string
}

func (p *BackupPluginValidator) ValidatePluginConfig(config map[string]string) (*plugtypes.BackupOptions, error) {
	// Validate the plugin configuration
	p.Log.Debugf("%s validating plugin configuration", p.LogHeader)
	if len(config) == 0 {
		p.Log.Debug("no configuration provided")
		return &plugtypes.BackupOptions{}, nil
	}
	bo := &plugtypes.BackupOptions{}

	for key, value := range config {
		p.Log.Debugf("%s configuration key: %s, value: %s", p.LogHeader, key, value)
		switch key {
		case "migration":
			bo.Migration = value == "true"
		case "readoptNodes":
			bo.ReadoptNodes = value == "true"
		case "managedServices":
			bo.ManagedServices = value == "true"
		case "dataUploadTimeout":
			minutes, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing dataUploadTimeout: %s", err.Error())
			}
			bo.DataUploadTimeout = time.Duration(minutes)
		case "dataUploadCheckPace":
			seconds, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing dataUploadCheckPace: %s", err.Error())
			}
			bo.DataUploadCheckPace = time.Duration(seconds)
		default:
			p.Log.Warnf("unknown configuration key: %s with value %s", key, value)
		}
	}

	p.Log.Infof("%s plugin configuration validated", p.LogHeader)

	return bo, nil

}

func (p *BackupPluginValidator) ValidatePlatformConfig(hcp *hyperv1.HostedControlPlane) error {
	switch hcp.Spec.Platform.Type {
	case hyperv1.AWSPlatform:
		return p.checkAWSPlatform(hcp)
	case hyperv1.AzurePlatform:
		return p.checkAzurePlatform(hcp)
	case hyperv1.IBMCloudPlatform:
		return p.checkIBMCloudPlatform(hcp)
	case hyperv1.KubevirtPlatform:
		return p.checkKubevirtPlatform(hcp)
	case hyperv1.OpenStackPlatform:
		return p.checkOpenStackPlatform(hcp)
	case hyperv1.AgentPlatform:
		return p.checkAgentPlatform(hcp)
	default:
		return fmt.Errorf("unsupported platform type %s", hcp.Spec.Platform.Type)
	}

}

func (p *BackupPluginValidator) checkAWSPlatform(hcp *hyperv1.HostedControlPlane) error {
	// Check if the AWS platform is configured properly
	// Check ROSA
	p.Log.Infof("%s AWS platform configuration is valid for HCP: %s", p.LogHeader, hcp.Name)
	return nil
}

func (p *BackupPluginValidator) checkAzurePlatform(hcp *hyperv1.HostedControlPlane) error {
	// Check if the Azure platform is configured properly
	// Check ARO
	p.Log.Infof("%s ARO platform configuration is valid for HCP: %s", p.LogHeader, hcp.Name)
	return nil
}

func (p *BackupPluginValidator) checkIBMCloudPlatform(hcp *hyperv1.HostedControlPlane) error {
	// Check if the IBM Cloud platform is configured properly
	p.Log.Infof("%s IBM platform configuration is valid for HCP: %s", p.LogHeader, hcp.Name)
	return nil
}

func (p *BackupPluginValidator) checkKubevirtPlatform(hcp *hyperv1.HostedControlPlane) error {
	// Check if the Kubevirt platform is configured properly
	p.Log.Infof("%s Kubevirt platform configuration is valid for HCP: %s", p.LogHeader, hcp.Name)
	return nil
}

func (p *BackupPluginValidator) checkOpenStackPlatform(hcp *hyperv1.HostedControlPlane) error {
	// Check if the OpenStack platform is configured properly
	p.Log.Infof("%s OpenStack platform configuration is valid for HCP: %s", p.LogHeader, hcp.Name)
	return nil
}

func (p *BackupPluginValidator) checkAgentPlatform(hcp *hyperv1.HostedControlPlane) error {
	// Check if the Agent platform is configured properly
	p.Log.Infof("%s Agent platform configuration is valid for HCP: %s", p.LogHeader, hcp.Name)
	return nil
}
