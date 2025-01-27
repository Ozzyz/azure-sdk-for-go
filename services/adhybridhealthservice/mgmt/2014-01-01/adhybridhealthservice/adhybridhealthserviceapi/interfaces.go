package adhybridhealthserviceapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/adhybridhealthservice/mgmt/2014-01-01/adhybridhealthservice"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/gofrs/uuid"
)

// AddsServicesClientAPI contains the set of methods on the AddsServicesClient type.
type AddsServicesClientAPI interface {
	Add(ctx context.Context, service adhybridhealthservice.ServiceProperties) (result adhybridhealthservice.ServiceProperties, err error)
	Delete(ctx context.Context, serviceName string, confirm *bool) (result autorest.Response, err error)
	Get(ctx context.Context, serviceName string) (result adhybridhealthservice.ServiceProperties, err error)
	GetForestSummary(ctx context.Context, serviceName string) (result adhybridhealthservice.ForestSummary, err error)
	GetMetricMetadata(ctx context.Context, serviceName string, metricName string) (result adhybridhealthservice.MetricMetadata, err error)
	GetMetricMetadataForGroup(ctx context.Context, serviceName string, metricName string, groupName string, groupKey string, fromDate *date.Time, toDate *date.Time) (result adhybridhealthservice.MetricSets, err error)
	List(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesPage, err error)
	ListComplete(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesIterator, err error)
	ListMetricMetadata(ctx context.Context, serviceName string, filter string, perfCounter *bool) (result adhybridhealthservice.MetricMetadataListPage, err error)
	ListMetricMetadataComplete(ctx context.Context, serviceName string, filter string, perfCounter *bool) (result adhybridhealthservice.MetricMetadataListIterator, err error)
	ListMetricsAverage(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsPage, err error)
	ListMetricsAverageComplete(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsIterator, err error)
	ListMetricsSum(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsPage, err error)
	ListMetricsSumComplete(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsIterator, err error)
	ListPremiumServices(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesPage, err error)
	ListPremiumServicesComplete(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesIterator, err error)
	ListReplicationDetails(ctx context.Context, serviceName string, filter string, withDetails *bool) (result adhybridhealthservice.ReplicationDetailsList, err error)
	ListReplicationSummary(ctx context.Context, serviceName string, isGroupbySite bool, query string, filter string, takeCount *int32) (result adhybridhealthservice.ReplicationSummaryList, err error)
	ListServerAlerts(ctx context.Context, serviceMemberID uuid.UUID, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsPage, err error)
	ListServerAlertsComplete(ctx context.Context, serviceMemberID uuid.UUID, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsIterator, err error)
	Update(ctx context.Context, serviceName string, service adhybridhealthservice.ServiceProperties) (result adhybridhealthservice.ServiceProperties, err error)
}

var _ AddsServicesClientAPI = (*adhybridhealthservice.AddsServicesClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	ListAddsAlerts(ctx context.Context, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsPage, err error)
	ListAddsAlertsComplete(ctx context.Context, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsIterator, err error)
}

var _ AlertsClientAPI = (*adhybridhealthservice.AlertsClient)(nil)

// ConfigurationClientAPI contains the set of methods on the ConfigurationClient type.
type ConfigurationClientAPI interface {
	Add(ctx context.Context) (result adhybridhealthservice.Tenant, err error)
	Get(ctx context.Context) (result adhybridhealthservice.Tenant, err error)
	ListAddsConfigurations(ctx context.Context, serviceName string, grouping string) (result adhybridhealthservice.AddsConfigurationPage, err error)
	ListAddsConfigurationsComplete(ctx context.Context, serviceName string, grouping string) (result adhybridhealthservice.AddsConfigurationIterator, err error)
	Update(ctx context.Context, tenant adhybridhealthservice.Tenant) (result adhybridhealthservice.Tenant, err error)
}

var _ ConfigurationClientAPI = (*adhybridhealthservice.ConfigurationClient)(nil)

// DimensionsClientAPI contains the set of methods on the DimensionsClient type.
type DimensionsClientAPI interface {
	ListAddsDimensions(ctx context.Context, serviceName string, dimension string) (result adhybridhealthservice.DimensionsPage, err error)
	ListAddsDimensionsComplete(ctx context.Context, serviceName string, dimension string) (result adhybridhealthservice.DimensionsIterator, err error)
}

var _ DimensionsClientAPI = (*adhybridhealthservice.DimensionsClient)(nil)

// AddsServiceMembersClientAPI contains the set of methods on the AddsServiceMembersClient type.
type AddsServiceMembersClientAPI interface {
	Delete(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, confirm *bool) (result autorest.Response, err error)
	Get(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result adhybridhealthservice.ServiceMember, err error)
	List(ctx context.Context, serviceName string, filter string) (result adhybridhealthservice.AddsServiceMembersPage, err error)
	ListComplete(ctx context.Context, serviceName string, filter string) (result adhybridhealthservice.AddsServiceMembersIterator, err error)
	ListCredentials(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, filter string) (result adhybridhealthservice.Credentials, err error)
}

var _ AddsServiceMembersClientAPI = (*adhybridhealthservice.AddsServiceMembersClient)(nil)

// AdDomainServiceMembersClientAPI contains the set of methods on the AdDomainServiceMembersClient type.
type AdDomainServiceMembersClientAPI interface {
	List(ctx context.Context, serviceName string, isGroupbySite bool, filter string, query string, takeCount *int32) (result adhybridhealthservice.AddsServiceMembersPage, err error)
	ListComplete(ctx context.Context, serviceName string, isGroupbySite bool, filter string, query string, takeCount *int32) (result adhybridhealthservice.AddsServiceMembersIterator, err error)
}

var _ AdDomainServiceMembersClientAPI = (*adhybridhealthservice.AdDomainServiceMembersClient)(nil)

// AddsServicesUserPreferenceClientAPI contains the set of methods on the AddsServicesUserPreferenceClient type.
type AddsServicesUserPreferenceClientAPI interface {
	Add(ctx context.Context, serviceName string, featureName string, setting adhybridhealthservice.UserPreference) (result autorest.Response, err error)
	Delete(ctx context.Context, serviceName string, featureName string) (result autorest.Response, err error)
	Get(ctx context.Context, serviceName string, featureName string) (result adhybridhealthservice.UserPreference, err error)
}

var _ AddsServicesUserPreferenceClientAPI = (*adhybridhealthservice.AddsServicesUserPreferenceClient)(nil)

// AddsServiceClientAPI contains the set of methods on the AddsServiceClient type.
type AddsServiceClientAPI interface {
	GetMetrics(ctx context.Context, serviceName string, metricName string, groupName string, groupKey string, fromDate *date.Time, toDate *date.Time) (result adhybridhealthservice.MetricSets, err error)
}

var _ AddsServiceClientAPI = (*adhybridhealthservice.AddsServiceClient)(nil)

// AddsServicesReplicationStatusClientAPI contains the set of methods on the AddsServicesReplicationStatusClient type.
type AddsServicesReplicationStatusClientAPI interface {
	Get(ctx context.Context, serviceName string) (result adhybridhealthservice.ReplicationStatus, err error)
}

var _ AddsServicesReplicationStatusClientAPI = (*adhybridhealthservice.AddsServicesReplicationStatusClient)(nil)

// AddsServicesServiceMembersClientAPI contains the set of methods on the AddsServicesServiceMembersClient type.
type AddsServicesServiceMembersClientAPI interface {
	Add(ctx context.Context, serviceName string, serviceMember adhybridhealthservice.ServiceMember) (result adhybridhealthservice.ServiceMember, err error)
	List(ctx context.Context, serviceName string, filter string, dimensionType string, dimensionSignature string) (result adhybridhealthservice.ServiceMembersPage, err error)
	ListComplete(ctx context.Context, serviceName string, filter string, dimensionType string, dimensionSignature string) (result adhybridhealthservice.ServiceMembersIterator, err error)
}

var _ AddsServicesServiceMembersClientAPI = (*adhybridhealthservice.AddsServicesServiceMembersClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result adhybridhealthservice.OperationListResponsePage, err error)
	ListComplete(ctx context.Context) (result adhybridhealthservice.OperationListResponseIterator, err error)
}

var _ OperationsClientAPI = (*adhybridhealthservice.OperationsClient)(nil)

// ReportsClientAPI contains the set of methods on the ReportsClient type.
type ReportsClientAPI interface {
	GetDevOps(ctx context.Context) (result adhybridhealthservice.Result, err error)
}

var _ ReportsClientAPI = (*adhybridhealthservice.ReportsClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	Add(ctx context.Context, service adhybridhealthservice.ServiceProperties) (result adhybridhealthservice.ServiceProperties, err error)
	AddAlertFeedback(ctx context.Context, serviceName string, alertFeedback adhybridhealthservice.AlertFeedback) (result adhybridhealthservice.AlertFeedback, err error)
	Delete(ctx context.Context, serviceName string, confirm *bool) (result autorest.Response, err error)
	Get(ctx context.Context, serviceName string) (result adhybridhealthservice.ServiceProperties, err error)
	GetFeatureAvailibility(ctx context.Context, serviceName string, featureName string) (result adhybridhealthservice.Result, err error)
	GetMetricMetadata(ctx context.Context, serviceName string, metricName string) (result adhybridhealthservice.MetricMetadata, err error)
	GetMetricMetadataForGroup(ctx context.Context, serviceName string, metricName string, groupName string, groupKey string, fromDate *date.Time, toDate *date.Time) (result adhybridhealthservice.MetricSets, err error)
	GetTenantWhitelisting(ctx context.Context, serviceName string, featureName string) (result adhybridhealthservice.Result, err error)
	List(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesPage, err error)
	ListComplete(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesIterator, err error)
	ListAlertFeedback(ctx context.Context, serviceName string, shortName string) (result adhybridhealthservice.AlertFeedbacks, err error)
	ListAlerts(ctx context.Context, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsPage, err error)
	ListAlertsComplete(ctx context.Context, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsIterator, err error)
	ListAllRiskyIPDownloadReport(ctx context.Context, serviceName string) (result adhybridhealthservice.RiskyIPBlobUris, err error)
	ListCurrentRiskyIPDownloadReport(ctx context.Context, serviceName string) (result adhybridhealthservice.RiskyIPBlobUris, err error)
	ListExportErrors(ctx context.Context, serviceName string) (result adhybridhealthservice.ErrorCounts, err error)
	ListExportErrorsV2(ctx context.Context, serviceName string, errorBucket string) (result adhybridhealthservice.MergedExportErrors, err error)
	ListExportStatus(ctx context.Context, serviceName string) (result adhybridhealthservice.ExportStatusesPage, err error)
	ListExportStatusComplete(ctx context.Context, serviceName string) (result adhybridhealthservice.ExportStatusesIterator, err error)
	ListMetricMetadata(ctx context.Context, serviceName string, filter string, perfCounter *bool) (result adhybridhealthservice.MetricMetadataListPage, err error)
	ListMetricMetadataComplete(ctx context.Context, serviceName string, filter string, perfCounter *bool) (result adhybridhealthservice.MetricMetadataListIterator, err error)
	ListMetricsAverage(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsPage, err error)
	ListMetricsAverageComplete(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsIterator, err error)
	ListMetricsSum(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsPage, err error)
	ListMetricsSumComplete(ctx context.Context, serviceName string, metricName string, groupName string) (result adhybridhealthservice.MetricsIterator, err error)
	ListMonitoringConfigurations(ctx context.Context, serviceName string) (result adhybridhealthservice.Items, err error)
	ListPremium(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesPage, err error)
	ListPremiumComplete(ctx context.Context, filter string, serviceType string, skipCount *int32, takeCount *int32) (result adhybridhealthservice.ServicesIterator, err error)
	ListUserBadPasswordReport(ctx context.Context, serviceName string, dataSource string) (result adhybridhealthservice.ErrorReportUsersEntries, err error)
	Update(ctx context.Context, serviceName string, service adhybridhealthservice.ServiceProperties) (result adhybridhealthservice.ServiceProperties, err error)
	UpdateMonitoringConfiguration(ctx context.Context, serviceName string, configurationSetting adhybridhealthservice.Item) (result autorest.Response, err error)
}

var _ ServicesClientAPI = (*adhybridhealthservice.ServicesClient)(nil)

// ServiceClientAPI contains the set of methods on the ServiceClient type.
type ServiceClientAPI interface {
	GetMetrics(ctx context.Context, serviceName string, metricName string, groupName string, groupKey string, fromDate *date.Time, toDate *date.Time) (result adhybridhealthservice.MetricSets, err error)
}

var _ ServiceClientAPI = (*adhybridhealthservice.ServiceClient)(nil)

// ServiceMembersClientAPI contains the set of methods on the ServiceMembersClient type.
type ServiceMembersClientAPI interface {
	Add(ctx context.Context, serviceName string, serviceMember adhybridhealthservice.ServiceMember) (result adhybridhealthservice.ServiceMember, err error)
	Delete(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, confirm *bool) (result autorest.Response, err error)
	DeleteData(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result adhybridhealthservice.ServiceMember, err error)
	GetConnectorMetadata(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, metricName string) (result adhybridhealthservice.ConnectorMetadata, err error)
	GetMetrics(ctx context.Context, serviceName string, metricName string, groupName string, serviceMemberID uuid.UUID, groupKey string, fromDate *date.Time, toDate *date.Time) (result adhybridhealthservice.MetricSets, err error)
	GetServiceConfiguration(ctx context.Context, serviceName string, serviceMemberID string) (result adhybridhealthservice.ServiceConfiguration, err error)
	List(ctx context.Context, serviceName string, filter string, dimensionType string, dimensionSignature string) (result adhybridhealthservice.ServiceMembersPage, err error)
	ListComplete(ctx context.Context, serviceName string, filter string, dimensionType string, dimensionSignature string) (result adhybridhealthservice.ServiceMembersIterator, err error)
	ListAlerts(ctx context.Context, serviceMemberID uuid.UUID, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsPage, err error)
	ListAlertsComplete(ctx context.Context, serviceMemberID uuid.UUID, serviceName string, filter string, state string, from *date.Time, toParameter *date.Time) (result adhybridhealthservice.AlertsIterator, err error)
	ListConnectors(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result adhybridhealthservice.Connectors, err error)
	ListCredentials(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, filter string) (result adhybridhealthservice.Credentials, err error)
	ListDataFreshness(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result adhybridhealthservice.DataFreshnessDetails, err error)
	ListExportStatus(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result adhybridhealthservice.ExportStatusesPage, err error)
	ListExportStatusComplete(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result adhybridhealthservice.ExportStatusesIterator, err error)
	ListGlobalConfiguration(ctx context.Context, serviceName string, serviceMemberID string) (result adhybridhealthservice.GlobalConfigurations, err error)
}

var _ ServiceMembersClientAPI = (*adhybridhealthservice.ServiceMembersClient)(nil)

// ListClientAPI contains the set of methods on the ListClient type.
type ListClientAPI interface {
	IPAddressAggregatesByService(ctx context.Context, serviceName string, skiptoken string) (result adhybridhealthservice.IPAddressAggregatesPage, err error)
	IPAddressAggregatesByServiceComplete(ctx context.Context, serviceName string, skiptoken string) (result adhybridhealthservice.IPAddressAggregatesIterator, err error)
	IPAddressAggregateSettings(ctx context.Context, serviceName string) (result adhybridhealthservice.IPAddressAggregateSetting, err error)
}

var _ ListClientAPI = (*adhybridhealthservice.ListClient)(nil)

// UpdateClientAPI contains the set of methods on the UpdateClient type.
type UpdateClientAPI interface {
	IPAddressAggregateSettings(ctx context.Context, serviceName string, IPAddressAggregateSetting adhybridhealthservice.IPAddressAggregateSetting) (result adhybridhealthservice.IPAddressAggregateSetting, err error)
}

var _ UpdateClientAPI = (*adhybridhealthservice.UpdateClient)(nil)
