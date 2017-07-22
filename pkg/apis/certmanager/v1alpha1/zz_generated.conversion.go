// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	certmanager "github.com/jetstack/cert-manager/pkg/apis/certmanager"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ACMECertificateConfig_To_certmanager_ACMECertificateConfig,
		Convert_certmanager_ACMECertificateConfig_To_v1alpha1_ACMECertificateConfig,
		Convert_v1alpha1_ACMECertificateDNS01Config_To_certmanager_ACMECertificateDNS01Config,
		Convert_certmanager_ACMECertificateDNS01Config_To_v1alpha1_ACMECertificateDNS01Config,
		Convert_v1alpha1_ACMECertificateDomainConfig_To_certmanager_ACMECertificateDomainConfig,
		Convert_certmanager_ACMECertificateDomainConfig_To_v1alpha1_ACMECertificateDomainConfig,
		Convert_v1alpha1_ACMECertificateHTTP01Config_To_certmanager_ACMECertificateHTTP01Config,
		Convert_certmanager_ACMECertificateHTTP01Config_To_v1alpha1_ACMECertificateHTTP01Config,
		Convert_v1alpha1_ACMEDomainAuthorization_To_certmanager_ACMEDomainAuthorization,
		Convert_certmanager_ACMEDomainAuthorization_To_v1alpha1_ACMEDomainAuthorization,
		Convert_v1alpha1_ACMEIssuer_To_certmanager_ACMEIssuer,
		Convert_certmanager_ACMEIssuer_To_v1alpha1_ACMEIssuer,
		Convert_v1alpha1_ACMEIssuerDNS01Config_To_certmanager_ACMEIssuerDNS01Config,
		Convert_certmanager_ACMEIssuerDNS01Config_To_v1alpha1_ACMEIssuerDNS01Config,
		Convert_v1alpha1_ACMEIssuerDNS01Provider_To_certmanager_ACMEIssuerDNS01Provider,
		Convert_certmanager_ACMEIssuerDNS01Provider_To_v1alpha1_ACMEIssuerDNS01Provider,
		Convert_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS_To_certmanager_ACMEIssuerDNS01ProviderCloudDNS,
		Convert_certmanager_ACMEIssuerDNS01ProviderCloudDNS_To_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS,
		Convert_v1alpha1_Certificate_To_certmanager_Certificate,
		Convert_certmanager_Certificate_To_v1alpha1_Certificate,
		Convert_v1alpha1_CertificateACMEStatus_To_certmanager_CertificateACMEStatus,
		Convert_certmanager_CertificateACMEStatus_To_v1alpha1_CertificateACMEStatus,
		Convert_v1alpha1_CertificateList_To_certmanager_CertificateList,
		Convert_certmanager_CertificateList_To_v1alpha1_CertificateList,
		Convert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec,
		Convert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec,
		Convert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus,
		Convert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus,
		Convert_v1alpha1_Issuer_To_certmanager_Issuer,
		Convert_certmanager_Issuer_To_v1alpha1_Issuer,
		Convert_v1alpha1_IssuerList_To_certmanager_IssuerList,
		Convert_certmanager_IssuerList_To_v1alpha1_IssuerList,
		Convert_v1alpha1_IssuerSpec_To_certmanager_IssuerSpec,
		Convert_certmanager_IssuerSpec_To_v1alpha1_IssuerSpec,
		Convert_v1alpha1_IssuerStatus_To_certmanager_IssuerStatus,
		Convert_certmanager_IssuerStatus_To_v1alpha1_IssuerStatus,
	)
}

func autoConvert_v1alpha1_ACMECertificateConfig_To_certmanager_ACMECertificateConfig(in *ACMECertificateConfig, out *certmanager.ACMECertificateConfig, s conversion.Scope) error {
	out.Config = *(*[]certmanager.ACMECertificateDomainConfig)(unsafe.Pointer(&in.Config))
	return nil
}

func Convert_v1alpha1_ACMECertificateConfig_To_certmanager_ACMECertificateConfig(in *ACMECertificateConfig, out *certmanager.ACMECertificateConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMECertificateConfig_To_certmanager_ACMECertificateConfig(in, out, s)
}

func autoConvert_certmanager_ACMECertificateConfig_To_v1alpha1_ACMECertificateConfig(in *certmanager.ACMECertificateConfig, out *ACMECertificateConfig, s conversion.Scope) error {
	if in.Config == nil {
		out.Config = make([]ACMECertificateDomainConfig, 0)
	} else {
		out.Config = *(*[]ACMECertificateDomainConfig)(unsafe.Pointer(&in.Config))
	}
	return nil
}

func Convert_certmanager_ACMECertificateConfig_To_v1alpha1_ACMECertificateConfig(in *certmanager.ACMECertificateConfig, out *ACMECertificateConfig, s conversion.Scope) error {
	return autoConvert_certmanager_ACMECertificateConfig_To_v1alpha1_ACMECertificateConfig(in, out, s)
}

func autoConvert_v1alpha1_ACMECertificateDNS01Config_To_certmanager_ACMECertificateDNS01Config(in *ACMECertificateDNS01Config, out *certmanager.ACMECertificateDNS01Config, s conversion.Scope) error {
	out.Provider = in.Provider
	return nil
}

func Convert_v1alpha1_ACMECertificateDNS01Config_To_certmanager_ACMECertificateDNS01Config(in *ACMECertificateDNS01Config, out *certmanager.ACMECertificateDNS01Config, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMECertificateDNS01Config_To_certmanager_ACMECertificateDNS01Config(in, out, s)
}

func autoConvert_certmanager_ACMECertificateDNS01Config_To_v1alpha1_ACMECertificateDNS01Config(in *certmanager.ACMECertificateDNS01Config, out *ACMECertificateDNS01Config, s conversion.Scope) error {
	out.Provider = in.Provider
	return nil
}

func Convert_certmanager_ACMECertificateDNS01Config_To_v1alpha1_ACMECertificateDNS01Config(in *certmanager.ACMECertificateDNS01Config, out *ACMECertificateDNS01Config, s conversion.Scope) error {
	return autoConvert_certmanager_ACMECertificateDNS01Config_To_v1alpha1_ACMECertificateDNS01Config(in, out, s)
}

func autoConvert_v1alpha1_ACMECertificateDomainConfig_To_certmanager_ACMECertificateDomainConfig(in *ACMECertificateDomainConfig, out *certmanager.ACMECertificateDomainConfig, s conversion.Scope) error {
	out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	out.HTTP01 = (*certmanager.ACMECertificateHTTP01Config)(unsafe.Pointer(in.HTTP01))
	out.DNS01 = (*certmanager.ACMECertificateDNS01Config)(unsafe.Pointer(in.DNS01))
	return nil
}

func Convert_v1alpha1_ACMECertificateDomainConfig_To_certmanager_ACMECertificateDomainConfig(in *ACMECertificateDomainConfig, out *certmanager.ACMECertificateDomainConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMECertificateDomainConfig_To_certmanager_ACMECertificateDomainConfig(in, out, s)
}

func autoConvert_certmanager_ACMECertificateDomainConfig_To_v1alpha1_ACMECertificateDomainConfig(in *certmanager.ACMECertificateDomainConfig, out *ACMECertificateDomainConfig, s conversion.Scope) error {
	if in.Domains == nil {
		out.Domains = make([]string, 0)
	} else {
		out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	}
	out.HTTP01 = (*ACMECertificateHTTP01Config)(unsafe.Pointer(in.HTTP01))
	out.DNS01 = (*ACMECertificateDNS01Config)(unsafe.Pointer(in.DNS01))
	return nil
}

func Convert_certmanager_ACMECertificateDomainConfig_To_v1alpha1_ACMECertificateDomainConfig(in *certmanager.ACMECertificateDomainConfig, out *ACMECertificateDomainConfig, s conversion.Scope) error {
	return autoConvert_certmanager_ACMECertificateDomainConfig_To_v1alpha1_ACMECertificateDomainConfig(in, out, s)
}

func autoConvert_v1alpha1_ACMECertificateHTTP01Config_To_certmanager_ACMECertificateHTTP01Config(in *ACMECertificateHTTP01Config, out *certmanager.ACMECertificateHTTP01Config, s conversion.Scope) error {
	out.Ingress = in.Ingress
	out.IngressClass = (*string)(unsafe.Pointer(in.IngressClass))
	return nil
}

func Convert_v1alpha1_ACMECertificateHTTP01Config_To_certmanager_ACMECertificateHTTP01Config(in *ACMECertificateHTTP01Config, out *certmanager.ACMECertificateHTTP01Config, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMECertificateHTTP01Config_To_certmanager_ACMECertificateHTTP01Config(in, out, s)
}

func autoConvert_certmanager_ACMECertificateHTTP01Config_To_v1alpha1_ACMECertificateHTTP01Config(in *certmanager.ACMECertificateHTTP01Config, out *ACMECertificateHTTP01Config, s conversion.Scope) error {
	out.Ingress = in.Ingress
	out.IngressClass = (*string)(unsafe.Pointer(in.IngressClass))
	return nil
}

func Convert_certmanager_ACMECertificateHTTP01Config_To_v1alpha1_ACMECertificateHTTP01Config(in *certmanager.ACMECertificateHTTP01Config, out *ACMECertificateHTTP01Config, s conversion.Scope) error {
	return autoConvert_certmanager_ACMECertificateHTTP01Config_To_v1alpha1_ACMECertificateHTTP01Config(in, out, s)
}

func autoConvert_v1alpha1_ACMEDomainAuthorization_To_certmanager_ACMEDomainAuthorization(in *ACMEDomainAuthorization, out *certmanager.ACMEDomainAuthorization, s conversion.Scope) error {
	out.Domain = in.Domain
	out.URI = in.URI
	return nil
}

func Convert_v1alpha1_ACMEDomainAuthorization_To_certmanager_ACMEDomainAuthorization(in *ACMEDomainAuthorization, out *certmanager.ACMEDomainAuthorization, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEDomainAuthorization_To_certmanager_ACMEDomainAuthorization(in, out, s)
}

func autoConvert_certmanager_ACMEDomainAuthorization_To_v1alpha1_ACMEDomainAuthorization(in *certmanager.ACMEDomainAuthorization, out *ACMEDomainAuthorization, s conversion.Scope) error {
	out.Domain = in.Domain
	out.URI = in.URI
	return nil
}

func Convert_certmanager_ACMEDomainAuthorization_To_v1alpha1_ACMEDomainAuthorization(in *certmanager.ACMEDomainAuthorization, out *ACMEDomainAuthorization, s conversion.Scope) error {
	return autoConvert_certmanager_ACMEDomainAuthorization_To_v1alpha1_ACMEDomainAuthorization(in, out, s)
}

func autoConvert_v1alpha1_ACMEIssuer_To_certmanager_ACMEIssuer(in *ACMEIssuer, out *certmanager.ACMEIssuer, s conversion.Scope) error {
	out.Email = in.Email
	out.Server = in.Server
	out.PrivateKey = in.PrivateKey
	out.URI = in.URI
	out.DNS01 = (*certmanager.ACMEIssuerDNS01Config)(unsafe.Pointer(in.DNS01))
	return nil
}

func Convert_v1alpha1_ACMEIssuer_To_certmanager_ACMEIssuer(in *ACMEIssuer, out *certmanager.ACMEIssuer, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEIssuer_To_certmanager_ACMEIssuer(in, out, s)
}

func autoConvert_certmanager_ACMEIssuer_To_v1alpha1_ACMEIssuer(in *certmanager.ACMEIssuer, out *ACMEIssuer, s conversion.Scope) error {
	out.Email = in.Email
	out.Server = in.Server
	out.PrivateKey = in.PrivateKey
	out.URI = in.URI
	out.DNS01 = (*ACMEIssuerDNS01Config)(unsafe.Pointer(in.DNS01))
	return nil
}

func Convert_certmanager_ACMEIssuer_To_v1alpha1_ACMEIssuer(in *certmanager.ACMEIssuer, out *ACMEIssuer, s conversion.Scope) error {
	return autoConvert_certmanager_ACMEIssuer_To_v1alpha1_ACMEIssuer(in, out, s)
}

func autoConvert_v1alpha1_ACMEIssuerDNS01Config_To_certmanager_ACMEIssuerDNS01Config(in *ACMEIssuerDNS01Config, out *certmanager.ACMEIssuerDNS01Config, s conversion.Scope) error {
	out.Providers = *(*[]certmanager.ACMEIssuerDNS01Provider)(unsafe.Pointer(&in.Providers))
	return nil
}

func Convert_v1alpha1_ACMEIssuerDNS01Config_To_certmanager_ACMEIssuerDNS01Config(in *ACMEIssuerDNS01Config, out *certmanager.ACMEIssuerDNS01Config, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEIssuerDNS01Config_To_certmanager_ACMEIssuerDNS01Config(in, out, s)
}

func autoConvert_certmanager_ACMEIssuerDNS01Config_To_v1alpha1_ACMEIssuerDNS01Config(in *certmanager.ACMEIssuerDNS01Config, out *ACMEIssuerDNS01Config, s conversion.Scope) error {
	if in.Providers == nil {
		out.Providers = make([]ACMEIssuerDNS01Provider, 0)
	} else {
		out.Providers = *(*[]ACMEIssuerDNS01Provider)(unsafe.Pointer(&in.Providers))
	}
	return nil
}

func Convert_certmanager_ACMEIssuerDNS01Config_To_v1alpha1_ACMEIssuerDNS01Config(in *certmanager.ACMEIssuerDNS01Config, out *ACMEIssuerDNS01Config, s conversion.Scope) error {
	return autoConvert_certmanager_ACMEIssuerDNS01Config_To_v1alpha1_ACMEIssuerDNS01Config(in, out, s)
}

func autoConvert_v1alpha1_ACMEIssuerDNS01Provider_To_certmanager_ACMEIssuerDNS01Provider(in *ACMEIssuerDNS01Provider, out *certmanager.ACMEIssuerDNS01Provider, s conversion.Scope) error {
	out.Name = in.Name
	out.CloudDNS = (*certmanager.ACMEIssuerDNS01ProviderCloudDNS)(unsafe.Pointer(in.CloudDNS))
	return nil
}

func Convert_v1alpha1_ACMEIssuerDNS01Provider_To_certmanager_ACMEIssuerDNS01Provider(in *ACMEIssuerDNS01Provider, out *certmanager.ACMEIssuerDNS01Provider, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEIssuerDNS01Provider_To_certmanager_ACMEIssuerDNS01Provider(in, out, s)
}

func autoConvert_certmanager_ACMEIssuerDNS01Provider_To_v1alpha1_ACMEIssuerDNS01Provider(in *certmanager.ACMEIssuerDNS01Provider, out *ACMEIssuerDNS01Provider, s conversion.Scope) error {
	out.Name = in.Name
	out.CloudDNS = (*ACMEIssuerDNS01ProviderCloudDNS)(unsafe.Pointer(in.CloudDNS))
	return nil
}

func Convert_certmanager_ACMEIssuerDNS01Provider_To_v1alpha1_ACMEIssuerDNS01Provider(in *certmanager.ACMEIssuerDNS01Provider, out *ACMEIssuerDNS01Provider, s conversion.Scope) error {
	return autoConvert_certmanager_ACMEIssuerDNS01Provider_To_v1alpha1_ACMEIssuerDNS01Provider(in, out, s)
}

func autoConvert_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS_To_certmanager_ACMEIssuerDNS01ProviderCloudDNS(in *ACMEIssuerDNS01ProviderCloudDNS, out *certmanager.ACMEIssuerDNS01ProviderCloudDNS, s conversion.Scope) error {
	out.ServiceAccount = in.ServiceAccount
	out.Project = in.Project
	return nil
}

func Convert_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS_To_certmanager_ACMEIssuerDNS01ProviderCloudDNS(in *ACMEIssuerDNS01ProviderCloudDNS, out *certmanager.ACMEIssuerDNS01ProviderCloudDNS, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS_To_certmanager_ACMEIssuerDNS01ProviderCloudDNS(in, out, s)
}

func autoConvert_certmanager_ACMEIssuerDNS01ProviderCloudDNS_To_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS(in *certmanager.ACMEIssuerDNS01ProviderCloudDNS, out *ACMEIssuerDNS01ProviderCloudDNS, s conversion.Scope) error {
	out.ServiceAccount = in.ServiceAccount
	out.Project = in.Project
	return nil
}

func Convert_certmanager_ACMEIssuerDNS01ProviderCloudDNS_To_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS(in *certmanager.ACMEIssuerDNS01ProviderCloudDNS, out *ACMEIssuerDNS01ProviderCloudDNS, s conversion.Scope) error {
	return autoConvert_certmanager_ACMEIssuerDNS01ProviderCloudDNS_To_v1alpha1_ACMEIssuerDNS01ProviderCloudDNS(in, out, s)
}

func autoConvert_v1alpha1_Certificate_To_certmanager_Certificate(in *Certificate, out *certmanager.Certificate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_Certificate_To_certmanager_Certificate(in *Certificate, out *certmanager.Certificate, s conversion.Scope) error {
	return autoConvert_v1alpha1_Certificate_To_certmanager_Certificate(in, out, s)
}

func autoConvert_certmanager_Certificate_To_v1alpha1_Certificate(in *certmanager.Certificate, out *Certificate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_certmanager_Certificate_To_v1alpha1_Certificate(in *certmanager.Certificate, out *Certificate, s conversion.Scope) error {
	return autoConvert_certmanager_Certificate_To_v1alpha1_Certificate(in, out, s)
}

func autoConvert_v1alpha1_CertificateACMEStatus_To_certmanager_CertificateACMEStatus(in *CertificateACMEStatus, out *certmanager.CertificateACMEStatus, s conversion.Scope) error {
	out.Authorizations = *(*[]certmanager.ACMEDomainAuthorization)(unsafe.Pointer(&in.Authorizations))
	return nil
}

func Convert_v1alpha1_CertificateACMEStatus_To_certmanager_CertificateACMEStatus(in *CertificateACMEStatus, out *certmanager.CertificateACMEStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateACMEStatus_To_certmanager_CertificateACMEStatus(in, out, s)
}

func autoConvert_certmanager_CertificateACMEStatus_To_v1alpha1_CertificateACMEStatus(in *certmanager.CertificateACMEStatus, out *CertificateACMEStatus, s conversion.Scope) error {
	if in.Authorizations == nil {
		out.Authorizations = make([]ACMEDomainAuthorization, 0)
	} else {
		out.Authorizations = *(*[]ACMEDomainAuthorization)(unsafe.Pointer(&in.Authorizations))
	}
	return nil
}

func Convert_certmanager_CertificateACMEStatus_To_v1alpha1_CertificateACMEStatus(in *certmanager.CertificateACMEStatus, out *CertificateACMEStatus, s conversion.Scope) error {
	return autoConvert_certmanager_CertificateACMEStatus_To_v1alpha1_CertificateACMEStatus(in, out, s)
}

func autoConvert_v1alpha1_CertificateList_To_certmanager_CertificateList(in *CertificateList, out *certmanager.CertificateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]certmanager.Certificate)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_CertificateList_To_certmanager_CertificateList(in *CertificateList, out *certmanager.CertificateList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateList_To_certmanager_CertificateList(in, out, s)
}

func autoConvert_certmanager_CertificateList_To_v1alpha1_CertificateList(in *certmanager.CertificateList, out *CertificateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]Certificate, 0)
	} else {
		out.Items = *(*[]Certificate)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_certmanager_CertificateList_To_v1alpha1_CertificateList(in *certmanager.CertificateList, out *CertificateList, s conversion.Scope) error {
	return autoConvert_certmanager_CertificateList_To_v1alpha1_CertificateList(in, out, s)
}

func autoConvert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(in *CertificateSpec, out *certmanager.CertificateSpec, s conversion.Scope) error {
	out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	out.SecretName = in.SecretName
	out.Issuer = in.Issuer
	out.ACME = (*certmanager.ACMECertificateConfig)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(in *CertificateSpec, out *certmanager.CertificateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(in, out, s)
}

func autoConvert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(in *certmanager.CertificateSpec, out *CertificateSpec, s conversion.Scope) error {
	if in.Domains == nil {
		out.Domains = make([]string, 0)
	} else {
		out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	}
	out.SecretName = in.SecretName
	out.Issuer = in.Issuer
	out.ACME = (*ACMECertificateConfig)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(in *certmanager.CertificateSpec, out *CertificateSpec, s conversion.Scope) error {
	return autoConvert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(in, out, s)
}

func autoConvert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(in *CertificateStatus, out *certmanager.CertificateStatus, s conversion.Scope) error {
	out.ACME = (*certmanager.CertificateACMEStatus)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(in *CertificateStatus, out *certmanager.CertificateStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(in, out, s)
}

func autoConvert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(in *certmanager.CertificateStatus, out *CertificateStatus, s conversion.Scope) error {
	out.ACME = (*CertificateACMEStatus)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(in *certmanager.CertificateStatus, out *CertificateStatus, s conversion.Scope) error {
	return autoConvert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(in, out, s)
}

func autoConvert_v1alpha1_Issuer_To_certmanager_Issuer(in *Issuer, out *certmanager.Issuer, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_IssuerSpec_To_certmanager_IssuerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_IssuerStatus_To_certmanager_IssuerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_Issuer_To_certmanager_Issuer(in *Issuer, out *certmanager.Issuer, s conversion.Scope) error {
	return autoConvert_v1alpha1_Issuer_To_certmanager_Issuer(in, out, s)
}

func autoConvert_certmanager_Issuer_To_v1alpha1_Issuer(in *certmanager.Issuer, out *Issuer, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_certmanager_IssuerSpec_To_v1alpha1_IssuerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_certmanager_IssuerStatus_To_v1alpha1_IssuerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_certmanager_Issuer_To_v1alpha1_Issuer(in *certmanager.Issuer, out *Issuer, s conversion.Scope) error {
	return autoConvert_certmanager_Issuer_To_v1alpha1_Issuer(in, out, s)
}

func autoConvert_v1alpha1_IssuerList_To_certmanager_IssuerList(in *IssuerList, out *certmanager.IssuerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]certmanager.Issuer)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_IssuerList_To_certmanager_IssuerList(in *IssuerList, out *certmanager.IssuerList, s conversion.Scope) error {
	return autoConvert_v1alpha1_IssuerList_To_certmanager_IssuerList(in, out, s)
}

func autoConvert_certmanager_IssuerList_To_v1alpha1_IssuerList(in *certmanager.IssuerList, out *IssuerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]Issuer, 0)
	} else {
		out.Items = *(*[]Issuer)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_certmanager_IssuerList_To_v1alpha1_IssuerList(in *certmanager.IssuerList, out *IssuerList, s conversion.Scope) error {
	return autoConvert_certmanager_IssuerList_To_v1alpha1_IssuerList(in, out, s)
}

func autoConvert_v1alpha1_IssuerSpec_To_certmanager_IssuerSpec(in *IssuerSpec, out *certmanager.IssuerSpec, s conversion.Scope) error {
	out.ACME = (*certmanager.ACMEIssuer)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_v1alpha1_IssuerSpec_To_certmanager_IssuerSpec(in *IssuerSpec, out *certmanager.IssuerSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_IssuerSpec_To_certmanager_IssuerSpec(in, out, s)
}

func autoConvert_certmanager_IssuerSpec_To_v1alpha1_IssuerSpec(in *certmanager.IssuerSpec, out *IssuerSpec, s conversion.Scope) error {
	out.ACME = (*ACMEIssuer)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_certmanager_IssuerSpec_To_v1alpha1_IssuerSpec(in *certmanager.IssuerSpec, out *IssuerSpec, s conversion.Scope) error {
	return autoConvert_certmanager_IssuerSpec_To_v1alpha1_IssuerSpec(in, out, s)
}

func autoConvert_v1alpha1_IssuerStatus_To_certmanager_IssuerStatus(in *IssuerStatus, out *certmanager.IssuerStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	return nil
}

func Convert_v1alpha1_IssuerStatus_To_certmanager_IssuerStatus(in *IssuerStatus, out *certmanager.IssuerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_IssuerStatus_To_certmanager_IssuerStatus(in, out, s)
}

func autoConvert_certmanager_IssuerStatus_To_v1alpha1_IssuerStatus(in *certmanager.IssuerStatus, out *IssuerStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	return nil
}

func Convert_certmanager_IssuerStatus_To_v1alpha1_IssuerStatus(in *certmanager.IssuerStatus, out *IssuerStatus, s conversion.Scope) error {
	return autoConvert_certmanager_IssuerStatus_To_v1alpha1_IssuerStatus(in, out, s)
}
