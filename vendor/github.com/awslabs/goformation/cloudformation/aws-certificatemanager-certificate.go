package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSCertificateManagerCertificate AWS CloudFormation Resource (AWS::CertificateManager::Certificate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html
type AWSCertificateManagerCertificate struct {

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainname
	DomainName string `json:"DomainName,omitempty"`

	// DomainValidationOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainvalidationoptions
	DomainValidationOptions []AWSCertificateManagerCertificate_DomainValidationOption `json:"DomainValidationOptions,omitempty"`

	// SubjectAlternativeNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-subjectalternativenames
	SubjectAlternativeNames []string `json:"SubjectAlternativeNames,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-tags
	Tags []Tag `json:"Tags,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCertificateManagerCertificate) AWSCloudFormationType() string {
	return "AWS::CertificateManager::Certificate"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSCertificateManagerCertificate) MarshalJSON() ([]byte, error) {
	type Properties AWSCertificateManagerCertificate
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSCertificateManagerCertificate) UnmarshalJSON(b []byte) error {
	type Properties AWSCertificateManagerCertificate
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSCertificateManagerCertificate(*res.Properties)
	}

	return nil
}

// GetAllAWSCertificateManagerCertificateResources retrieves all AWSCertificateManagerCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSCertificateManagerCertificateResources() map[string]AWSCertificateManagerCertificate {
	results := map[string]AWSCertificateManagerCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSCertificateManagerCertificate:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::CertificateManager::Certificate" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCertificateManagerCertificate
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSCertificateManagerCertificateWithName retrieves all AWSCertificateManagerCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCertificateManagerCertificateWithName(name string) (AWSCertificateManagerCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSCertificateManagerCertificate:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::CertificateManager::Certificate" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCertificateManagerCertificate
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSCertificateManagerCertificate{}, errors.New("resource not found")
}
